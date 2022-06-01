package service

import (
	"fmt"
	"github.com/ricardojonathanromero/observer-pattern/internal/domain/models"
	"github.com/ricardojonathanromero/observer-pattern/internal/port"
)

type NotifierImpl struct {
	observers map[port.Observer]struct{}
}

var _ port.Notifier = (*NotifierImpl)(nil)

func (n *NotifierImpl) Register(o port.Observer) {
	n.observers[o] = struct{}{}
}

func (n *NotifierImpl) Unregister(o port.Observer) {
	delete(n.observers, o)
}

func (n *NotifierImpl) Notify(e models.Event) {
	for o := range n.observers {
		o.OnNotify(e)
	}
}

func (n *NotifierImpl) NotifyAsync(e models.Event) {
	go func(ev models.Event) {
		fmt.Println(len(n.observers))
		for o := range n.observers {
			go o.OnNotify(ev)
		}
	}(e)
}

func NewNotifier(observers map[port.Observer]struct{}) port.Notifier {
	return &NotifierImpl{observers: observers}
}
