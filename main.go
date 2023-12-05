package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-daq/smbus"
)

var tempThresholds = []int{30, 40, 50, 65}  // in degrees Celsius
var fanSpeeds = []int{0, 25, 50, 100}  // as a percentage of max speed


func getAverageTemperature(_path string) (float64, error) {
	var totalTemp int64
	var count int64

	// Get all thermal_zone* directories
	matches, err := filepath.Glob(_path + "/thermal_zone*")
	if err != nil {
		return 0, err
	}

	for _, match := range matches {
		// Read the temp file in each thermal_zone directory
		data, err := ioutil.ReadFile(match + "/temp")
		if err != nil {
			// Skip this directory if the temp file cannot be read
			continue
		}
		temp, err := strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64)
		if err != nil {
			return 0, err
		}
		totalTemp += temp
		count++
	}

	if count == 0 {
		return 0, fmt.Errorf("no temperature data found")
	}
	return float64(totalTemp) / float64(count), nil
}

func getFanSpeed(_temp float64) int {
	for i, threshold := range tempThresholds {
		if _temp < float64(threshold) {
			return fanSpeeds[i]
		}
	}
	return fanSpeeds[len(fanSpeeds)-1]
}


func main() {
	bus, err := smbus.Open(1, 0x1a) // replace 0x1a with your device address
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for {
		avgTemp, err := getAverageTemperature("/sys/class/thermal")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Average temperature: %.2f°C\n", avgTemp/1000)

		fanSpeed := getFanSpeed(avgTemp / 1000)
		fmt.Printf("Setting fan speed to %d%%\n", fanSpeed)

		err = bus.WriteReg(0x1a, 0x00,byte(fanSpeed)) // replace with your register address
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		time.Sleep(5 * time.Second)
	}
}

