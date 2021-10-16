package main

import (
	"github.com/Owicca/chan/infra"

	_ "github.com/Owicca/chan/routes/middleware"
	_ "github.com/Owicca/chan/routes/frontend"
	_ "github.com/Owicca/chan/routes/backend"
)

func main() {
	infra.S.Run()
}