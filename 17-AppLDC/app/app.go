package app

import (
	"github.com/urfave/cli/v2"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de comando"
	app.Usage = "Busca de IPs e Nomes de Servidor na internet"
	return app
}
