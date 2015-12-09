package main

import "time"
import "log"

type Basicstruct struct {
	id          int
	date        time.Time
	status      string
	hiddenField string
}

func main() {

	currentTime := time.Now()
	bs := Basicstruct{id: 10, date: currentTime, status: "FooBar"}

	if bs.ID() != 10 {
		log.Panicf("BasicStruct.ID() returned: %d not %d", bs.ID(), 10)
	}
	if !currentTime.Equal(bs.Date()) {
		log.Panicf("BasicStruct.Date() returned: %v, not %v", bs.Date(), currentTime)
	}
	if bs.Status() != "FooBar" {
		log.Panicf("BasicStruct.Status() returned: %s, not %s", bs.Status(), "FooBar")
	}

}
