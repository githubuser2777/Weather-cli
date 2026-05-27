package display

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"githubuser2777/Weather-cli/internal/weather"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorBold   = "\033[1m"
)

const boxWidth = 48 // Expanded slightly for forecasts

// PrintError formats and prints error messages to Stderr.
func PrintError(err error) {
	fmt.Fprintf(os.Stderr, "%s%s✖ Error:%s %v\n", ColorBold, ColorRed, ColorReset, err)
}

// formatTemp colorizes the temperature based on its value (Celsius).
// For Fahrenheit, it roughly maps equivalent ranges.
func formatTemp(temp float64, unit string) (string, string) {
	raw := fmt.Sprintf("%.1f°%s", temp, unit)
	
	celsiusTemp := temp
	if unit == "F" {
		celsiusTemp = (temp - 32) * 5 / 9
	}

	if celsiusTemp <= 15 {
		return raw, ColorBlue + raw + ColorReset
	} else if celsiusTemp <= 27 {
		return raw, ColorGreen + raw + ColorReset
	}
	return raw, ColorRed + raw + ColorReset
}

// formatHumidity colorizes the humidity.
func formatHumidity(hum int) (string, string) {
	raw := fmt.Sprintf("%d%%", hum)
	return raw, ColorCyan + raw + ColorReset
}

// formatConditions colorizes the conditions.
func formatConditions(cond string, icon string) (string, string) {
	raw := fmt.Sprintf("%s %s", icon, cond)
	return raw, ColorYellow + raw + ColorReset
}

// printRow prints a padded row for the ASCII widget.
func printRow(label, rawValue, coloredValue string) {
	contentLen := utf8.RuneCountInString(label) + utf8.RuneCountInString(rawValue)
	padding := boxWidth - contentLen
	if padding < 0 {
		padding = 0
	}
	fmt.Printf("│ %s%s%s │\n", label, coloredValue, strings.Repeat(" ", padding))
}

// RenderWeather prints the weather data and forecast inside a beautiful ASCII widget.
func RenderWeather(locationName string, data weather.WeatherData) {
	fmt.Println()
	fmt.Printf("╭%s╮\n", strings.Repeat("─", boxWidth+2))

	// Header row
	headerPrefix := " Weather for: "
	headerRawVal := locationName
	printRow(headerPrefix, headerRawVal, ColorBold+headerRawVal+ColorReset)

	fmt.Printf("├%s┤\n", strings.Repeat("─", boxWidth+2))

	// Data rows
	rawT, colT := formatTemp(data.Temperature, data.Unit)
	printRow(" Temperature: ", rawT, colT)

	rawH, colH := formatHumidity(data.Humidity)
	printRow(" Humidity:    ", rawH, colH)

	rawC, colC := formatConditions(data.Conditions, data.Icon)
	printRow(" Conditions:  ", rawC, colC)

	// Forecast rows
	if len(data.Forecast) > 0 {
		fmt.Printf("├%s┤\n", strings.Repeat("─", boxWidth+2))
		printRow(" Forecast:    ", "", "")
		for _, f := range data.Forecast {
			// e.g., "2023-10-02" -> "10-02"
			shortDate := f.Date
			if len(shortDate) >= 5 {
				shortDate = shortDate[5:]
			}

			rawMin, colMin := formatTemp(f.MinTemp, data.Unit)
			rawMax, colMax := formatTemp(f.MaxTemp, data.Unit)

			rawF := fmt.Sprintf("%s  L:%s H:%s %s %s", shortDate, rawMin, rawMax, f.Icon, f.Conditions)
			colF := fmt.Sprintf("%s%s%s  L:%s H:%s %s%s %s%s", ColorCyan, shortDate, ColorReset, colMin, colMax, ColorYellow, f.Icon, f.Conditions, ColorReset)

			printRow("  ", rawF, colF)
		}
	}

	fmt.Printf("╰%s╯\n", strings.Repeat("─", boxWidth+2))
	fmt.Println()
}
