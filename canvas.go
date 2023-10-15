package main

import (
	"github.com/ArchWizard56/go-canvas"
)

// GetCanvasTodos retrieves a list of TODO items from the Canvas LMS API.
// It loads configuration data from "config.json" and initializes the Canvas API client.
// The retrieved TODO items are returned as a slice of canvas.TODO values.
func GetCanvasTodos() []canvas.TODO {
	// Load configuration data from "config.json"
	config := LoadConfig()

	// Initialize the Canvas API client with the API key and host URL from the configuration
	canvas := canvas.WithHost(config.CanvasKey, config.CanvasHost)

	// Fetch the list of TODO items from the Canvas API
	todos, err := canvas.Todos()

	// Check for errors during the API call
	if err != nil {
		// If an error occurs, panic and terminate the program, displaying the error message
		panic(err)
	}

	// Return the retrieved TODO items
	return todos
}

func GetCourse(id int) *canvas.Course {
	// Load configuration data from "config.json"
	config := LoadConfig()

	// Initialize the Canvas API client with the API key and host URL from the configuration
	canvas := canvas.WithHost(config.CanvasKey, config.CanvasHost)

	course, err := canvas.GetCourse(id)
	if err != nil {
		panic(err)
	}
	return course
}
