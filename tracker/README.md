## Tracker

Tracker is a service that is providing methods for tracking events. Event is specified by its ID and contains all the tracked periods. Every period has start and stop time. With Tracker service such events can be created, stopped and eventually started again. Service allows you to total all the periods when given event was active. 

## Code Example

The code below shows operations you can do with events:

    import "flowdock.eu/flowup/services/tracker"

    tracker := tracker.NewTracker()

    event := tracker.StartNew() // Event created and started

    tracker.Stop(event)    // Event stopped (paused)

    tracker.Start(event)   // Event started again 

If event includes some periods all periods can be sumed:

    sum := tracker.Total(event)

## Tests

All tests can be executed by running: 

    goconvey 

in the folder with test file, alternatively you can use command: 

    go test

## License