package comunicacao

func ViaVarios(canais ...EnviarComunicacao) EnviarComunicacao {
	return func(mensagem string) {
		for _, canal := range canais {
			canal(mensagem)
		}
	}
}
