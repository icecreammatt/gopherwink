package scheduler

import (
	"fmt"
	"github.com/jonaz/astrotime"
	"os/exec"
	"strconv"
	"time"
)

func isSunset(latitude float64, longitude float64) bool {
	t := astrotime.NextSunset(time.Now(), latitude, longitude)
	// tzname, _ := t.Zone()
	// fmt.Println(tzname)
	// fmt.Printf("The next sunrise is %d:%02d %s on %d/%d/%d.\n", t.Hour(), t.Minute(), tzname, t.Month(), t.Day(), t.Year())
	if t.Hour() == time.Now().Hour() && t.Minute() == time.Now().Minute() {
		return true
	}
	return false
}

func isSunrise(latitude float64, longitude float64) bool {
	t := astrotime.NextSunrise(time.Now(), latitude, longitude)
	tzname, _ := t.Zone()
	// fmt.Println(tzname)
	// fmt.Printf("The next sunrise is %d:%02d %s on %d/%d/%d.\n", t.Hour()-6, t.Minute()-25, tzname, t.Month(), t.Day(), t.Year())
	if t.Hour() == time.Now().Hour() && t.Minute() == time.Now().Minute() {
		fmt.Printf("The sunrise is %d:%02d %s on %d/%d/%d.\n", t.Hour(), t.Minute(), tzname, t.Month(), t.Day(), t.Year())
		return true
	}
	return false
}

func Start(intervalInSeconds int, latitude float64, longitude float64, autoNightLights []int64) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(intervalInSeconds))
	go func() {
		// Hard code lights to change for now
		for t := range ticker.C {
			var state string
			var changeLights bool
			if isSunset(latitude, longitude) {
				fmt.Println("Sunset", t)
				state = "ON"
				changeLights = true
			} else if isSunrise(latitude, longitude) {
				fmt.Println("Sunrise", t)
				state = "OFF"
				changeLights = true
			}
			if changeLights == true {
				fmt.Println("Updating lights")
				for _, light := range autoNightLights {
					args := []string{"-u", "-m" + strconv.FormatInt(light, 10), "-t1", "-v" + state}
					cmd := exec.Command("/usr/sbin/aprontest", args...)
					cmd.Run()
				}
			}
		}
	}()
}
