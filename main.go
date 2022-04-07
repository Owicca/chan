package main

import (
	"github.com/Owicca/chan/infra"

	_ "github.com/Owicca/chan/routes/backend"
	_ "github.com/Owicca/chan/routes/frontend"
	_ "github.com/Owicca/chan/routes/middleware"
)

func main() {
	infra.S.Run()
}
