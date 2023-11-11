package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestLoadWeatherData tests the loadWeatherData function
func TestLoadWeatherData(t *testing.T) {
	t.Run("weather data fetch error", func(t *testing.T) {
		mockURL := "https://example.com/mockdata.csv"

		_, err := loadWeatherData(mockURL)
		require.Error(t, err)
	})
}

// TestFilterByDate tests the filterByDate function
func TestFilterByDate(t *testing.T) {
	// Create a test dataset
	testData := []WeatherData{
		{Date: "2023-11-01"},
		{Date: "2023-11-02"},
		{Date: "2023-11-03"},
	}

	testParams := map[string][]string{
		"date": []string{"2023-11-02"},
	}

	result := filterByDate(testData[1], testParams["date"])

	if !result {
		t.Error("Filtering by date failed")
	}

}

// TestFilterByWeather tests the filterByWeather function
func TestFilterByWeather(t *testing.T) {
	t.Run("filtrating by weather no error", func(t *testing.T) {
		// Create a test dataset
		testData := []WeatherData{
			{Weather: "rain"},
			{Weather: "cloudy"},
			{Weather: "sunny"},
		}

		testParams := map[string][]string{
			"weather": []string{"cloudy"},
		}

		result := filterByWeather(testData[1], testParams["weather"])
		require.Equal(t, result, true)

	})

	t.Run("filtrating by weather error", func(t *testing.T) {
		// Create a test dataset
		testData := []WeatherData{
			{Weather: "rain"},
			{Weather: "cloudy"},
			{Weather: "sunny"},
		}

		testParams := map[string][]string{
			"weather": []string{"shinny"},
		}

		result := filterByWeather(testData[1], testParams["weather"])
		require.Equal(t, result, false)
	})

}
