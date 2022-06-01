package service_test

import (
	"github.com/ricardojonathanromero/observer-pattern/internal/domain/models"
	"github.com/ricardojonathanromero/observer-pattern/internal/port"
	"github.com/ricardojonathanromero/observer-pattern/internal/service"
	"testing"
	"time"
)

func TestObserverImpl_OnNotify(t *testing.T) {
	notifier := service.NewNotifier(make(map[port.Observer]struct{}, 0))

	observer := service.NewObserver(1)
	notifier.Register(observer)
	notifier.Notify(models.Event{Data: 10})
	notifier.NotifyAsync(models.Event{Data: 1})
	// sleep to exec async function before unregister observer
	time.Sleep(1 * time.Second)
	notifier.Unregister(observer)
}
