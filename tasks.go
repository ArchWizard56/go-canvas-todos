package main

import (
	"fmt"
	"time"

	"github.com/ArchWizard56/go-webdav"
	"github.com/ArchWizard56/go-webdav/caldav"
	"github.com/emersion/go-ical"
	"github.com/google/uuid"
)

// Task represents a task with its properties.
type Task struct {
	ID          string
	Title       string
	Category    string
	Description string
	DueDate     time.Time
}

// GetTasks retrieves tasks from a WebDAV calendar and returns them as a slice of Task.
func GetTasks() []Task {
	// Load configuration from config.json
	config := LoadConfig()

	// Create an HTTP client with basic authentication for WebDAV
	c := webdav.HTTPClientWithBasicAuth(nil, config.DavUsername, config.DavPassword)
	caldavClient, err := caldav.NewClient(c, config.DavURL)
	if err != nil {
		panic(err)
	}

	// Define a calendar component request and filter
	compRequest := caldav.CalendarCompRequest{
		Name: "VTODO",
	}
	compFilter := caldav.CompFilter{
		Name: "VCALENDAR",
	}
	query := caldav.CalendarQuery{CompRequest: compRequest, CompFilter: compFilter}

	// Query the calendar for tasks
	calendarObj, err := caldavClient.QueryCalendar("NDc", &query)
	if err != nil {
		panic(err)
	}

	// Initialize a slice to store tasks
	tasks := []Task{}

	// Iterate through calendar objects to extract tasks
	for i := range calendarObj {
		if calendarObj[i].Data.Name != "VCALENDAR" {
			continue
		}
		child := calendarObj[i].Data.Children[0]
		if child.Name != "VTODO" {
			continue
		}

		task := Task{}

		// Extract task properties
		if child.Props.Get("SUMMARY") != nil {
			task.Title = child.Props.Get("SUMMARY").Value
		}
		if child.Props.Get("CATEGORIES") != nil {
			task.Category = child.Props.Get("CATEGORIES").Value
		}
		task.ID = child.Props.Get("UID").Value
		if child.Props.Get("DUE") != nil {
			// Parse due date with two possible formats
			task.DueDate, err = time.Parse("20060102T150405Z", child.Props.Get("DUE").Value)
			if err != nil {
				task.DueDate, err = time.Parse("20060102", child.Props.Get("DUE").Value)
				if err != nil {
					panic(err)
				}
			}
		}
		if child.Props.Get("DESCRIPTION") != nil {
			task.Description = child.Props.Get("DESCRIPTION").Value
		}

		// Append the task to the slice
		tasks = append(tasks, task)
	}

	return tasks
}

// UpdateToDo updates a task in the WebDAV calendar.
func UpdateToDo(id string, title string, calendar string, category string, description string, due time.Time) {
	// Load configuration from config.json
	config := LoadConfig()

	// Create an HTTP client with basic authentication for WebDAV
	c := webdav.HTTPClientWithBasicAuth(nil, config.DavUsername, config.DavPassword)
	caldavClient, err := caldav.NewClient(c, config.DavURL)
	if err != nil {
		panic(err)
	}

	// Create a new iCalendar
	vcal := ical.NewCalendar()
	vcal.Props.Add(&ical.Prop{Name: "PRODID", Value: "Open-Xchange"})
	vcal.Props.Add(&ical.Prop{Name: "VERSION", Value: "2.0"})
	vcal.Name = "VCALENDAR"

	// Create a new VTODO task
	todo := ical.NewCalendar()
	todo.Name = "VTODO"
	todo.Props.Add(&ical.Prop{Name: "SUMMARY", Value: title})
	todo.Props.Add(&ical.Prop{Name: "DTSTAMP", Value: time.Now().UTC().Format("20060102T150405Z")})

	// Set DUE property if due date is not zero
	if !due.IsZero() {
		todo.Props.Add(&ical.Prop{Name: "DUE", Value: due.UTC().Format("20060102T150405Z")})
	}
	todo.Props.Add(&ical.Prop{Name: "CATEGORIES", Value: category})
	todo.Props.Add(&ical.Prop{Name: "DESCRIPTION", Value: description})
	todo.Props.Add(&ical.Prop{Name: "UID", Value: id})

	// Add the VTODO task to the VCALENDAR
	vcal.Children = append(vcal.Children, todo.Component)

	// Put the updated calendar object in WebDAV
	_, err = caldavClient.PutCalendarObject(fmt.Sprintf("%s/%s.ics", calendar, id), vcal)
	if err != nil {
		panic(err)
	}
}

// AddTodo adds a new task to the WebDAV calendar.
func AddTodo(title string, calendar string, category string, description string, due time.Time) {
	// Load configuration from config.json
	config := LoadConfig()

	// Create an HTTP client with basic authentication for WebDAV
	c := webdav.HTTPClientWithBasicAuth(nil, config.DavUsername, config.DavPassword)
	caldavClient, err := caldav.NewClient(c, config.DavURL)
	if err != nil {
		panic(err)
	}

	// Create a new iCalendar
	vcal := ical.NewCalendar()
	vcal.Props.Add(&ical.Prop{Name: "PRODID", Value: "Open-Xchange"})
	vcal.Props.Add(&ical.Prop{Name: "VERSION", Value: "2.0"})
	vcal.Name = "VCALENDAR"

	// Create a new VTODO task
	todo := ical.NewCalendar()
	todo.Name = "VTODO"
	id := uuid.New().String()
	todo.Props.Add(&ical.Prop{Name: "SUMMARY", Value: title})
	todo.Props.Add(&ical.Prop{Name: "DTSTAMP", Value: time.Now().UTC().Format("20060102T150405Z")})

	// Set DUE property if due date is not zero
	if !due.IsZero() {
		todo.Props.Add(&ical.Prop{Name: "DUE", Value: due.UTC().Format("20060102T150405Z")})
	}
	todo.Props.Add(&ical.Prop{Name: "CATEGORIES", Value: category})
	todo.Props.Add(&ical.Prop{Name: "DESCRIPTION", Value: description})
	todo.Props.Add(&ical.Prop{Name: "UID", Value: id})

	// Add the VTODO task to the VCALENDAR
	vcal.Children = append(vcal.Children, todo.Component)

	// Put the new calendar object in WebDAV
	_, err = caldavClient.PutCalendarObject(fmt.Sprintf("%s/%s.ics", calendar, id), vcal)
	if err != nil {
		panic(err)
	}
}
