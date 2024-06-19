package tools

import (
	"log"
	"time"
)

const defaultTimezone = "America/Bogota"

func ConfigureTimeZone() {
	var err error
	time.Local, err = time.LoadLocation(defaultTimezone)

	if err != nil {
		log.Fatal("Error loading timezone: " + err.Error())
		//panic("Error loading timezone: " + err.Error())
	}
}
