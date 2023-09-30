# Go-Canvas-Todos

Go-Canvas-Todos is a Go-based utility that helps you synchronize your Canvas LMS todos with a task calendar. This tool allows you to automatically add, update, and manage your tasks based on assignments in Canvas.

## Prerequisites

Before you begin using this tool, make sure you have the following:

- Go Programming Language (installed and configured)
- Access to Canvas LMS
- Access to a WebDAV server (for calendar synchronization)

## Configuration

To configure the Go-Canvas-Todos tool, create a `config.json` file based on the provided `config-example.json`. Here is what each configuration parameter means:

- `canvas_host`: The hostname of your Canvas instance (e.g., `canvas.example.com`).
- `canvas_key`: Your Canvas API key. Replace `<YOUR CANVAS KEY>` with your actual API key.
- `dav_url`: The URL of your WebDAV server (e.g., `https://dav.example.com`).
- `dav_username`: Your WebDAV username. Replace `<YOUR DAV USER>` with your actual username.
- `dav_password`: Your WebDAV password. Replace `<YOUR DAV PASSWORD>` with your actual password.
- `task_calendar`: The name or identifier of the calendar where tasks will be added.
- `disable_tls`: Set this to `true` to disable TLS certificate verification when making HTTP requests to Canvas. Only use this option in a trusted environment, as it may pose security risks.

## Usage

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/go-canvas-todos.git
   cd go-canvas-todos
   ```
   
2. Build the Go application:
  ```bash
  go build
  ```

3. Run the application
  ```bash
  ./go-canvas-todos
  ```

4. The application will fetch todos from Canvas and synchronize them with your specified task calendar.