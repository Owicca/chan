package main

import (
	// "log"

	"github.com/Owicca/chan/infra"
	"github.com/Owicca/chan/routes"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	cfg, conn, logger := infra.Setup("./config.json")
	router := mux.NewRouter()
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	srv := infra.NewServer(
		cfg,
		conn,
		router,
		infra.NewTemplate(),
	)

	routes.LoadRoutes(&srv)

	srv.Run()
}