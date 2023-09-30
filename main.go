package main

import (
	"crypto/tls"
	"net/http"
)

func main() {
	// Load configuration from a JSON file
	config := LoadConfig("config.json")

	// If TLS is disabled in the configuration, skip certificate verification
	if config.DisableTLS {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// Retrieve a list of Canvas todos
	todos := GetCanvasTodos()

	// Retrieve a list of tasks
	tasks := GetTasks()

	// Iterate through Canvas todos and synchronize them with tasks
	for todoNum := range todos {
		todoFound := false
		todo := todos[todoNum]
		assignment := todo.Assignment
		todoCourse := GetCourse(assignment.CourseID)
		dueTime := assignment.DueAt
		title := assignment.Name
		category := todoCourse.CourseCode
		description := todo.HTMLURL

		// Iterate through tasks to find a matching task
		for taskNum := range tasks {
			task := tasks[taskNum]

			// Check if the titles match
			if title == task.Title && todoCourse.CourseCode == task.Category {
				// Todo found, update it if due dates don't match
				todoFound = true
				if task.DueDate != dueTime {
					UpdateToDo(task.ID, title, config.TaskCalendar, category, description, dueTime)
				}
				break
			}
		}

		// If todo is not found in tasks, add it
		if !todoFound {
			AddTodo(title, config.TaskCalendar, category, description, dueTime)
		}
	}
}
