package main

import (
	"comunicacao"
	"usuario"
)

func main() {
	usuario.CadastrarUsuario("João", comunicacao.ViaVarios(
		comunicacao.ViaEmail,
		// comunicacao.ViaSms,
		comunicacao.ViaZapZap,
	))
}
