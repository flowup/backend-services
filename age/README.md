## Age
Age service is a service that provides you one simple method CalculateAge that is calculating how old the event is. This service is very simple and shows the style how you can write similiar services. 

## Code Example
This code example shows whole functionality of this Age service:

    event := time.Date(2013, 8, 2, 0, 0, 0, 0, time.UTC)
    now := time.Now()

    service := NewAgeService()
    period := service.CalculateAge(event, now)
