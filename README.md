# Observer pattern

This pattern is used mostly in event-driven architectures.

## Description
In this example, whe use observer patter to create multiple observers and trigger all data that the handler receive and send the data to all observers.

Into the main function, you can see that two observers were created.

```
observers := make(map[port.Observer]struct{}, 0)
notifier := service.NewNotifier(observers)

// build observers
observerOne := service.NewObserver(1)
observerTwo := service.NewObserver(2)
...
```

those observers are registered eventually in the notifier.
```
notifier.Register(observerOne)
notifier.Register(observerTwo)
```

Finally, we use the function `Notify` of the notifier to send data and triggered an event.
The data is sending to every observer subscribed to the notifier.
