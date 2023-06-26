package main

import (
	"github.com/Owicca/chan/internal/infra"

	_ "github.com/Owicca/chan/internal/routes/backend"
	_ "github.com/Owicca/chan/internal/routes/frontend"
	_ "github.com/Owicca/chan/internal/routes/middleware"
)

func main() {
	infra.S.Run()
}
