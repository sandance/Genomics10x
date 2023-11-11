package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// WeatherData represents the structure of the weather data
type WeatherData struct {
	Date          string  `json:"date"`
	Precipitation float64 `json:"precipitation"`
	TempMax       float64 `json:"temp_max"`
	TempMin       float64 `json:"temp_min"`
	Wind          float64 `json:"wind"`
	Weather       string  `json:"weather"`
}

// parseFloat converts a string to a float64
func parseFloat(str string) float64 {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return val
}

// loadWeatherData reads weather data from a CSV file specified by the given URL
func loadWeatherData(url string) ([]WeatherData, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	reader := csv.NewReader(response.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var weatherData []WeatherData

	for _, record := range records {
		weather := WeatherData{
			Date:          record[0],
			Precipitation: parseFloat(record[1]),
			TempMax:       parseFloat(record[2]),
			TempMin:       parseFloat(record[3]),
			Wind:          parseFloat(record[4]),
			Weather:       record[5],
		}

		weatherData = append(weatherData, weather)
	}

	return weatherData, nil
}

// filterWeatherData filters weather data based on query parameters
func filterWeatherData(data []WeatherData, queryParams map[string][]string) []WeatherData {
	var filteredData []WeatherData

	for _, entry := range data {
		if filterByDate(entry, queryParams["date"]) && filterByWeather(entry, queryParams["weather"]) {
			filteredData = append(filteredData, entry)
		}
	}

	return filteredData
}

// filterByDate checks if the weather data entry matches the provided date parameter
func filterByDate(entry WeatherData, dateParams []string) bool {
	if len(dateParams) == 0 {
		return true
	}

	for _, dateParam := range dateParams {
		if entry.Date == dateParam {
			return true
		}
	}

	return false
}

// filterByWeather checks if the weather data entry matches the provided weather parameter
func filterByWeather(entry WeatherData, weatherParam []string) bool {

	if len(weatherParam) == 0 {
		return true
	}

	for _, dateParam := range weatherParam {
		if entry.Weather == dateParam {
			return true
		}
	}
	return false
}

func main() {
	// Load weather data from URL
	weatherData, err := loadWeatherData("https://raw.githubusercontent.com/vega/vega/main/docs/data/seattle-weather.csv")
	if err != nil {
		log.Fatal("Error loading weather data:", err)
	}

	// Define HTTP endpoint and handler
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()

		filteredData := filterWeatherData(weatherData, queryParams)

		limit, err := strconv.Atoi(queryParams.Get("limit"))
		if err == nil && limit > 0 && limit <= len(filteredData) {
			filteredData = filteredData[:limit]
		}

		responseJSON, err := json.Marshal(filteredData)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJSON)
	})

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
