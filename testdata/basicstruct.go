package main

import "time"
import "log"

type Basicstruct struct {
	ID     int
	Date   time.Time
	Status string
}

func main() {

	currentTime := time.Now()
	bs := Basicstruct{ID: 10, Date: currentTime, Status: "FooBar"}

	if bs.GetID() != 10 {
		log.Panicf("BasicStruct.GetID() returned: %d not %d", bs.GetID(), 10)
	}
	if !currentTime.Equal(bs.GetDate()) {
		log.Panicf("BasicStruct.GetDate() returned: %v, not %v", bs.GetDate(), currentTime)
	}
	if bs.GetStatus() != "FooBar" {
		log.Panicf("BasicStruct.GetStatus() returned: %s, not %s", bs.GetStatus(), "FooBar")
	}

}
