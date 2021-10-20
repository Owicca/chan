package main

import (
	"github.com/Owicca/chan/infra"

	_ "github.com/Owicca/chan/routes/middleware"
	_ "github.com/Owicca/chan/routes/backend"
	_ "github.com/Owicca/chan/routes/frontend"
)

func main() {
	infra.S.Run()
}