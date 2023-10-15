package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

// Config is a struct that represents the configuration data loaded from "config.json".
type Config struct {
	CanvasHost   string `json:"canvas_host"`   // Host URL for the Canvas API, minus the https
	CanvasKey    string `json:"canvas_key"`    // API key for Canvas authentication
	DavURL       string `json:"dav_url"`       // Full URL for the dav server, including the https
	DavUsername  string `json:"dav_username"`  // Username for DAV server authentication
	DavPassword  string `json:"dav_password"`  // Password for DAV server authentication
	TaskCalendar string `json:"task_calendar"` // Calendar to add the final tasks to
	DisableTLS   bool   `json:"disable_tls"`   // Only for using a debug proxy - indicates whether TLS should be disabled
}

// LoadConfig loads the configuration data from the specified file path.
// It opens and reads the JSON file, unmarshals it into a Config struct,
// and returns a pointer to the loaded Config.
func LoadConfig() *Config {
	var config_path string
	if len(os.Args) == 2 {
		config_path = os.Args[1]
		fileinfo, err := os.Stat(config_path)
		if errors.Is(err, os.ErrNotExist) {
			panic(errors.New("provided config file does not exist"))
		} else if err != nil {
			panic(err)
		}
		if fileinfo.IsDir() {
			panic(errors.New("cannot load config from directory"))
		}
	} else {
		config_path = "config.json"
		fileinfo, err := os.Stat(config_path)
		if errors.Is(err, os.ErrNotExist) {
			panic(errors.New("provided config file does not exist"))
		} else if err != nil {
			panic(err)
		}
		if fileinfo.IsDir() {
			panic(errors.New("cannot load config from directory"))
		}
	}

	config_file, err := os.Open(config_path)
	if err != nil {
		// If there's an error opening the file, panic and terminate the program
		panic(err)
	}
	defer config_file.Close() // Ensure the file is closed when the function exits

	// Read the entire contents of the configuration file
	byteValue, _ := io.ReadAll(config_file)

	// Initialize an empty Config struct
	config := &Config{}

	// Unmarshal the JSON data into the Config struct
	err = json.Unmarshal(byteValue, config)
	if err != nil {
		// If there's an error unmarshaling the JSON data, panic and terminate the program
		panic(err)
	}

	// Return a pointer to the loaded Config
	return config
}
