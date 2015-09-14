package scheduler

import (
	"fmt"
	"github.com/jonaz/astrotime"
	"os/exec"
	"strconv"
	"time"
)

const LATITUDE = float64(38.547133)
const LONGITUDE = float64(-122.816380)

func isSunset() bool {
	t := astrotime.NextSunset(time.Now(), LATITUDE, LONGITUDE)
	// tzname, _ := t.Zone()
	// fmt.Println(tzname)
	// fmt.Printf("The next sunrise is %d:%02d %s on %d/%d/%d.\n", t.Hour(), t.Minute(), tzname, t.Month(), t.Day(), t.Year())
	if t.Hour() == time.Now().Hour() && t.Minute() == time.Now().Minute() {
		return true
	}
	return false
}

func isSunrise() bool {
	t := astrotime.NextSunrise(time.Now(), LATITUDE, LONGITUDE)
	tzname, _ := t.Zone()
	// fmt.Println(tzname)
	// fmt.Printf("The next sunrise is %d:%02d %s on %d/%d/%d.\n", t.Hour()-6, t.Minute()-25, tzname, t.Month(), t.Day(), t.Year())
	if t.Hour() == time.Now().Hour() && t.Minute() == time.Now().Minute() {
		fmt.Printf("The sunrise is %d:%02d %s on %d/%d/%d.\n", t.Hour(), t.Minute(), tzname, t.Month(), t.Day(), t.Year())
		return true
	}
	return false
}

func Start(intervalInSeconds int) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(intervalInSeconds))
	go func() {
		// Hard code lights to change for now
		lights := [2]int{1, 8}
		for t := range ticker.C {
			var state string
			var changeLights bool
			if isSunset() {
				fmt.Println("Sunset", t)
				state = "ON"
				changeLights = true
			} else if isSunrise() {
				fmt.Println("Sunrise", t)
				state = "OFF"
				changeLights = true
			}
			if changeLights == true {
				fmt.Println("Updating lights")
				for _, light := range lights {
					args := []string{"-u", "-m" + strconv.Itoa(light), "-t1", "-v" + state}
					cmd := exec.Command("/usr/sbin/aprontest", args...)
					cmd.Run()
				}
			}
		}
	}()
}
