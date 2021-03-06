package apiControllerV1

import (
	"antriin/src/business/event"
	"time"
)

type EventRequest struct {
	Name     string            `json:"name" validate:"required"`
	Schedule []ScheduleRequest `json:"schedule,omitempty"`
}

type ScheduleRequest struct {
	Location     LocationRequest `json:"location"`
	Date         time.Time       `json:"date"`
	Start        time.Time       `json:"start"`
	End          time.Time       `json:"end"`
	Capacity     uint            `json:"capacity"`
	Progress     uint            `json:"progress"`
	AttendeeType string          `json:"attendeeType"`
}

type LocationRequest struct {
	Address     string  `json:"address"`
	District    string  `json:"district"`
	SubDistrict string  `json:"subDistrict"`
	City        string  `json:"city"`
	Province    string  `json:"province"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

func RequestToSchedule(request []ScheduleRequest) []event.Schedule {
	var scheduleData []event.Schedule
	for _, e := range request {
		scheduleData = append(scheduleData, event.Schedule{
			Location: event.Location{
				Address:     e.Location.Address,
				District:    e.Location.District,
				SubDistrict: e.Location.SubDistrict,
				City:        e.Location.City,
				Province:    e.Location.Province,
				Longitude:   e.Location.Longitude,
				Latitude:    e.Location.Latitude,
			},
			Date:         e.Date,
			Start:        e.Start,
			End:          e.End,
			Capacity:     e.Capacity,
			Progress:     e.Progress,
			AttendeeType: e.AttendeeType,
		})
	}
	return scheduleData
}

func RequestToEvent(data *EventRequest) event.Event {
	scheduleData := RequestToSchedule(data.Schedule)
	return event.Event{
		Name:     data.Name,
		Schedule: scheduleData,
	}
}

func NewEventRequest(data event.Event) EventRequest {
	var scheduleData []ScheduleRequest
	for _, e := range data.Schedule {
		scheduleData = append(scheduleData, ScheduleRequest{
			Location: LocationRequest{
				Address:     e.Location.Address,
				District:    e.Location.District,
				SubDistrict: e.Location.SubDistrict,
				City:        e.Location.City,
				Province:    e.Location.Province,
				Longitude:   e.Location.Longitude,
				Latitude:    e.Location.Latitude,
			},
			Date:         e.Date,
			Start:        e.Start,
			End:          e.End,
			Capacity:     e.Capacity,
			Progress:     e.Progress,
			AttendeeType: e.AttendeeType,
		})
	}
	return EventRequest{
		Name:     data.Name,
		Schedule: scheduleData,
	}
}
