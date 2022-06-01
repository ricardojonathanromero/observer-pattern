package service

import (
	"fmt"
	"github.com/ricardojonathanromero/observer-pattern/internal/domain/models"
	"github.com/ricardojonathanromero/observer-pattern/internal/port"
)

type ObserverImpl struct {
	id int
}

var _ port.Observer = (*ObserverImpl)(nil)

func (o *ObserverImpl) OnNotify(e models.Event) {
	fmt.Printf("observer %d received event %d\n", o.id, e.Data)
}

func NewObserver(id int) port.Observer {
	return &ObserverImpl{id: id}
}
