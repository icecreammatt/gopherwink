package scheduler

import (
	"fmt"
	"github.com/jonaz/astrotime"
	"os/exec"
	"time"
)

const LATITUDE = float64(38.547133)
const LONGITUDE = float64(-122.816380)

func isSunset() bool {
	t := astrotime.NextSunset(time.Now(), LATITUDE, LONGITUDE)
	tzname, _ := t.Zone()
	fmt.Println(tzname)
	fmt.Printf("The next sunrise is %d:%02d %s on %d/%d/%d.\n", t.Hour(), t.Minute(), tzname, t.Month(), t.Day(), t.Year())
	if t.Hour() == time.Now().Hour() && t.Minute() == time.Now().Minute() {
		return true
	}
	return false
}

func isSunrise() bool {
	t := astrotime.NextSunrise(time.Now(), LATITUDE, LONGITUDE)
	tzname, _ := t.Zone()
	fmt.Println(tzname)
	fmt.Printf("The next sunrise is %d:%02d %s on %d/%d/%d.\n", t.Hour(), t.Minute(), tzname, t.Month(), t.Day(), t.Year())
	if t.Hour() == time.Now().Hour() && t.Minute() == time.Now().Minute() {
		return true
	}
	return false
}

func Start(intervalInSeconds int) {
	ticker := time.NewTicker(time.Millisecond * time.Duration(intervalInSeconds))
	go func() {
		for t := range ticker.C {
			if isSunset() {
				fmt.Println("Sunset", t)
				args := []string{"-u", "-m" + "1", "-t1", "-v" + "ON"}
				cmd := exec.Command("/usr/sbin/aprontest", args...)
				cmd.Run()
			} else if isSunrise() {
				fmt.Println("Sunrise", t)
				args := []string{"-u", "-m" + "1", "-t1", "-v" + "OFF"}
				cmd := exec.Command("/usr/sbin/aprontest", args...)
				cmd.Run()
			}
		}
	}()
}
