
# WeatherApp


## Description

This application is a Go-based server that fetches and filters weather data from a CSV file and provides endpoints to query the data.
The project includes a Makefile and Docker setup to build, test, and run the application.

## Prerequisites

- Make sure you have Docker and Docker Compose installed.

## Makefile Commands
Build the Application

To build the application, use the following command:

```make build```

Run the Application

```make run ```

Clean the image

```make clean```

## Usage
- Ensure Docker is running on your machine.
- Use make build to build the Docker image.
- Run the application with `make run`.
- Test app containerapptest will continuously hit weatherapp every 5 seconds
