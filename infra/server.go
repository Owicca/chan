package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/Owicca/chan/models/acl"

	"go.uber.org/zap"
)

var S *Server

// To be ran on server closing
var (
	undo              func()
	loggerSync        func() error
	_, filename, _, _ = runtime.Caller(0)
)

// Get config, db, logger
// set up settings and create Server
func init() {
	cfg, conn, logger := Setup(path.Join(path.Dir(filename), "../config.json"))
	loggerSync = logger.Sync
	undo = zap.ReplaceGlobals(logger)

	S = NewServer(
		cfg,
		conn,
		NewTemplate(),
	)
}

type Server struct {
	mux.Router
	Config   Config
	Conn     *gorm.DB
	Template *Template
}

func NewServer(
	cfg Config,
	conn *gorm.DB,
	tmpl *Template,
) *Server {
	if S == nil {
		S = &Server{
			Config:   cfg,
			Router:   *mux.NewRouter(),
			Conn:     conn,
			Template: tmpl,
		}
	}

	return S
}

// Get hostname and port.
func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%s", s.Config.HttpHost, s.Config.HttpPort)
}

// Server JSON response.
func (s *Server) JSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		json.NewEncoder(w).Encode(data)
		return nil
	}
	return fmt.Errorf("No data to return")
}

// Serve a media file.
func (s *Server) MEDIA(w http.ResponseWriter, status int, media []byte, mediaType string) {
	w.Header().Set("Content-Type", mediaType)
	w.Header().Set("Cache-Control", "max-age=31536000")
	w.WriteHeader(status)
	w.Write(media)
}

// Server a HTML response.
func (s *Server) HTML(w http.ResponseWriter, status int, htmlView string, data map[string]any) error {
	return s.Template.Render(w, status, htmlView, data)
}

func (s *Server) Redirect(w http.ResponseWriter, r *http.Request, status int, dst string) {
	http.Redirect(w, r, dst, status)
}

func (s *Server) GenerateUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s/", s.Addr(), strings.Trim(endpoint, "/"))
}

func (s *Server) Run() {
	acl.Run(s.Conn)

	addr := s.Addr()
	msg := fmt.Sprintf("Running at %s", addr)
	zap.L().Info(msg, zap.Int64("timestamp", time.Now().Unix()))

	httpServer := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				s.ServeHTTP(w, r)
			}),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	s.ShutdownOnInterrupt(httpServer)
}

func (s *Server) ShutdownOnInterrupt(srv *http.Server) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			msg := fmt.Sprintf("Shutting down error (%s)", err)
			zap.L().Info(msg, zap.Int64("timestamp", time.Now().Unix()))
		}
		zap.L().Info("Close everything!", zap.Int64("timestamp", time.Now().Unix()))
		defer loggerSync()
		defer undo()
		// s.Conn.Close()
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		zap.L().Info("Could not listen and serve!", zap.Int64("timestamp", time.Now().Unix()))
	}

	<-idleConnsClosed
}
