package main

import (
	"fmt"
	"time"

	"github.com/ArchWizard56/go-webdav"
	"github.com/ArchWizard56/go-webdav/caldav"
	"github.com/emersion/go-ical"
	"github.com/google/uuid"
)

type Task struct {
	ID          string
	Title       string
	Category    string
	Description string
	DueDate     time.Time
}

func GetTasks() []Task {
	config := LoadConfig("config.json")
	c := webdav.HTTPClientWithBasicAuth(nil, config.DavUsername, config.DavPassword)
	caldavClient, err := caldav.NewClient(c, config.DavURL)
	if err != nil {
		panic(err)
	}
	compRequest := caldav.CalendarCompRequest{
		Name: "VTODO",
	}
	compFilter := caldav.CompFilter{
		Name: "VCALENDAR",
	}
	query := caldav.CalendarQuery{CompRequest: compRequest, CompFilter: compFilter}

	calendarObj, err := caldavClient.QueryCalendar("NDc", &query)
	if err != nil {
		panic(err)
	}
	tasks := []Task{}
	for i := range calendarObj {
		if calendarObj[i].Data.Name != "VCALENDAR" {
			continue
		}
		child := calendarObj[i].Data.Children[0]
		if child.Name != "VTODO" {
			continue
		}
		task := Task{}
		if child.Props.Get("SUMMARY") != nil {
			task.Title = child.Props.Get("SUMMARY").Value
		}
		if child.Props.Get("CATEGORIES") != nil {
			task.Category = child.Props.Get("CATEGORIES").Value
		}
		task.ID = child.Props.Get("UID").Value
		if child.Props.Get("DUE") != nil {
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
		tasks = append(tasks, task)
	}
	return tasks
}

func UpdateToDo(id string, title string, calendar string, category string, description string, due time.Time) {
	config := LoadConfig("config.json")
	c := webdav.HTTPClientWithBasicAuth(nil, config.DavUsername, config.DavPassword)
	caldavClient, err := caldav.NewClient(c, config.DavURL)
	if err != nil {
		panic(err)
	}

	vcal := ical.NewCalendar()
	vcal.Props.Add(&ical.Prop{Name: "PRODID", Value: "Open-Xchange"})
	vcal.Props.Add(&ical.Prop{Name: "VERSION", Value: "2.0"})
	vcal.Name = "VCALENDAR"

	todo := ical.NewCalendar()
	todo.Name = "VTODO"

	todo.Props.Add(&ical.Prop{Name: "SUMMARY", Value: title})
	todo.Props.Add(&ical.Prop{Name: "DTSTAMP", Value: time.Now().UTC().Format("20060102T150405Z")})
	if !due.IsZero() {
		todo.Props.Add(&ical.Prop{Name: "DUE", Value: due.UTC().Format("20060102T150405Z")})
	}
	todo.Props.Add(&ical.Prop{Name: "CATEGORIES", Value: category})
	todo.Props.Add(&ical.Prop{Name: "DESCRIPTION", Value: description})
	todo.Props.Add(&ical.Prop{Name: "UID", Value: id})

	vcal.Children = append(vcal.Children, todo.Component)
	_, err = caldavClient.PutCalendarObject(fmt.Sprintf("%s/%s.ics", calendar, id), vcal)
	if err != nil {
		panic(err)
	}
}
func AddTodo(title string, calendar string, category string, description string, due time.Time) {
	config := LoadConfig("config.json")
	c := webdav.HTTPClientWithBasicAuth(nil, config.DavUsername, config.DavPassword)
	caldavClient, err := caldav.NewClient(c, config.DavURL)
	if err != nil {
		panic(err)
	}

	vcal := ical.NewCalendar()
	vcal.Props.Add(&ical.Prop{Name: "PRODID", Value: "Open-Xchange"})
	vcal.Props.Add(&ical.Prop{Name: "VERSION", Value: "2.0"})
	vcal.Name = "VCALENDAR"

	todo := ical.NewCalendar()
	todo.Name = "VTODO"

	id := uuid.New().String()

	todo.Props.Add(&ical.Prop{Name: "SUMMARY", Value: title})
	todo.Props.Add(&ical.Prop{Name: "DTSTAMP", Value: time.Now().UTC().Format("20060102T150405Z")})
	if !due.IsZero() {
		todo.Props.Add(&ical.Prop{Name: "DUE", Value: due.UTC().Format("20060102T150405Z")})
	}
	todo.Props.Add(&ical.Prop{Name: "CATEGORIES", Value: category})
	todo.Props.Add(&ical.Prop{Name: "DESCRIPTION", Value: description})
	todo.Props.Add(&ical.Prop{Name: "UID", Value: id})

	vcal.Children = append(vcal.Children, todo.Component)
	_, err = caldavClient.PutCalendarObject(fmt.Sprintf("%s/%s.ics", calendar, id), vcal)
	if err != nil {
		panic(err)
	}
}
