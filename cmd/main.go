package main

import (
	"fmt"
	"github.com/ricardojonathanromero/observer-pattern/internal/domain/models"
	"github.com/ricardojonathanromero/observer-pattern/internal/port"
	"github.com/ricardojonathanromero/observer-pattern/internal/service"
	"time"
)

func main() {
	// initialize observer map
	observers := make(map[port.Observer]struct{}, 0)
	// call to notifier constructor
	notifier := service.NewNotifier(observers)

	// build observers
	observerOne := service.NewObserver(1)
	observerTwo := service.NewObserver(2)

	// register new observers (1 & 2)
	notifier.Register(observerOne)
	notifier.Register(observerTwo)

	fmt.Println("----------- init -----------")

	// sending data to all observers
	notifier.Notify(models.Event{Data: 1})
	notifier.Notify(models.Event{Data: 101})
	notifier.Notify(models.Event{Data: 9999})

	// unregister observer 2
	notifier.Unregister(observerTwo)

	fmt.Println("----------------------------")
	// sending data to all observers
	notifier.Notify(models.Event{Data: 20})
	notifier.Notify(models.Event{Data: 99})
	notifier.Notify(models.Event{Data: 19283})

	// cleaning all observers
	notifier.Unregister(observerOne)

	fmt.Println("----------------------------")
	fmt.Println("building hundred observers")

	for i := 1; i <= 90000000; i++ {
		observer := service.NewObserver(i)
		notifier.Register(observer)
	}
	fmt.Println("sending data sync")
	init := time.Now()
	notifier.Notify(models.Event{Data: 20})
	notifier.Notify(models.Event{Data: 99})
	finish := time.Now()
	fmt.Println("finish after ", finish.Sub(init).Seconds())

	fmt.Println("----------------------------")
	fmt.Println("sending data async")
	initAsync := time.Now()
	notifier.NotifyAsync(models.Event{Data: 20})
	notifier.NotifyAsync(models.Event{Data: 99})
	finishAsync := time.Now()
	fmt.Println("finish after ", finishAsync.Sub(initAsync).Seconds())

	time.Sleep(5 * time.Second)
}
