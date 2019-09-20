package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/api/calendar/v3"
)

func main() {
	event := &calendar.Event{
		Summary:     "Google I/O 2015",
		Location:    "800 Howard St., San Francisco, CA 94103",
		Description: "A chance to hear more about Google's developer products.",
		Start: &calendar.EventDateTime{
			DateTime: "2015-05-28T09:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: "2015-05-28T17:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
		Attendees: []*calendar.EventAttendee{
			&calendar.EventAttendee{Email: "lpage@example.com"},
			&calendar.EventAttendee{Email: "sbrin@example.com"},
		},
	}

	ctx := context.Background()
	srv, err := calendar.NewService(ctx)

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
}
