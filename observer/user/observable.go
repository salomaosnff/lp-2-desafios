package user

import (
	"observable"
)

func CreateObserver() *observable.Observable {
	var obs = observable.Create()

	obs.Subscribe(EnviarSms)
	obs.Subscribe(EnviarWhatsapp)
	obs.Subscribe(EnviarNotificacao)
	obs.Subscribe(ConectarAoServidorDoNatan)

	return &obs
}
