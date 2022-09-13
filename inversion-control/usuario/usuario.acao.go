package usuario

import (
	"comunicacao"
	"fmt"
)

func CadastrarUsuario(nome string, comunicar comunicacao.EnviarComunicacao) *Usuario {
	usuario := &Usuario{nome}

	comunicar(fmt.Sprintf("Usuário %s cadastrado com sucesso!", usuario.nome))

	return usuario
}
