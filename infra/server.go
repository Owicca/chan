package infra

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"
	"os/signal"
	"context"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/Owicca/chan/models/acl"

	"go.uber.org/zap"
)

type Server struct {
	Config Config
	Conn   *gorm.DB
	Router *mux.Router
	Template *Template
}

func NewServer(
	cfg Config,
	conn *gorm.DB,
	r *mux.Router,
	tmpl *Template,
) Server {
	return Server{
		Config: cfg,
		Conn:   conn,
		Router: r,
		Template: tmpl,
	}
}

// Merge hostname and port
func (s Server) GetAddr() string {
	return fmt.Sprintf("%s:%s", s.Config.HttpHost, s.Config.HttpPort)
}

func (s Server) RegisterRoute(
	path string,
	handler func(w http.ResponseWriter, r *http.Request),
	methods []string,
	name string,
) {
	s.Router.HandleFunc(path, handler).Methods(methods...).Name(name)
}

func (s Server) RegisterSubRoute(
	router *mux.Router,
	path string,
	handler func(w http.ResponseWriter, r *http.Request),
	methods []string,
	name string,
) {
	router.HandleFunc(path, handler).Methods(methods...).Name(name)
}

func (s Server) RegisterPathPrefix(
	path string,
	handler http.Handler,
	methods []string,
	name string,
) {
	s.Router.PathPrefix(path).Handler(handler).Methods(methods...).Name(name)
}

// Server JSON response
func (s Server) JSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		json.NewEncoder(w).Encode(data)
		return nil
	}
	return fmt.Errorf("No data to return")
}

// Serve a media file
func (s Server) MEDIA(w http.ResponseWriter, status int, media []byte, mediaType string) {
	w.Header().Set("Content-Type", mediaType)
	w.Header().Set("Cache-Control", "max-age=31536000")
	w.WriteHeader(status)
	w.Write(media)
}

// Server a HTML response
func (s Server) Render(w http.ResponseWriter, status int, htmlView string, data interface{}) error {
	return s.Template.Render(w, status, htmlView, data)
}

func (s Server) Redirect(w http.ResponseWriter, r *http.Request, status int, dst string) {
	http.Redirect(w, r, dst, status)
}

func (s Server) Run() {
	acl.Run(s.Conn)

	addr := s.GetAddr()
	msg := fmt.Sprintf("Running at %s", addr)
	zap.L().Info(msg, zap.Int64("timestamp", time.Now().Unix()))

	httpServer := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				s.Router.ServeHTTP(w, r)
			}),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	s.ShutdownOnInterrupt(httpServer)
}

func (s Server) ShutdownOnInterrupt(srv *http.Server) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<- sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			msg := fmt.Sprintf("Shutting down error (%s)", err)
			zap.L().Info(msg, zap.Int64("timestamp", time.Now().Unix()))
		}
		zap.L().Info("Close everything!", zap.Int64("timestamp", time.Now().Unix()))
		// s.Conn.Close()
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		zap.L().Info("Could not listen and serve!", zap.Int64("timestamp", time.Now().Unix()))
	}

	<- idleConnsClosed
}