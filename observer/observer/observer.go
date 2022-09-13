package observable

import "fmt"

type Observable struct {
	maxListeners  int
	subscriptions map[*Subscription]bool
}

type Subscription struct {
	observable *Observable
	call       *func(any)
}

func (s *Subscription) Unsubscribe() {
	delete(s.observable.subscriptions, s)
}

func (o *Observable) Subscribe(listener func(any)) *Subscription {
	if o.maxListeners > 0 && len(o.subscriptions) >= o.maxListeners {
		fmt.Println("Atenção: Muitos observadores")
	}

	var subscription = &Subscription{
		observable: o,
		call:       &listener,
	}

	o.subscriptions[subscription] = true

	return subscription
}

func (o *Observable) Publish(data any) {
	for subscription := range o.subscriptions {
		(*subscription.call)(data)
	}
}

func Create() Observable {
	return Observable{
		subscriptions: make(map[*Subscription]bool),
		maxListeners:  0,
	}
}

func CreateWithMax(max int) Observable {
	return Observable{
		subscriptions: make(map[*Subscription]bool),
		maxListeners:  max,
	}
}
