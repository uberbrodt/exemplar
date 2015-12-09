package main

import "time"
import "log"
import "reflect"

type Basicstruct struct {
	id           int
	date         time.Time
	status       string
	hiddenField  string
	privateField string `propertizer:"private"`
}

func main() {

	privateField := "shouldnotread"

	currentTime := time.Now()
	bs := &Basicstruct{id: 10, date: currentTime, status: "FooBar", privateField: privateField}

	if bs.ID() != 10 {
		log.Panicf("BasicStruct.ID() returned: %d not %d", bs.ID(), 10)
	}
	if !currentTime.Equal(bs.Date()) {
		log.Panicf("BasicStruct.Date() returned: %v, not %v", bs.Date(), currentTime)
	}
	if bs.Status() != "FooBar" {
		log.Panicf("BasicStruct.Status() returned: %s, not %s", bs.Status(), "FooBar")
	}

	typ := reflect.ValueOf(bs)
	if val := typ.MethodByName("PrivateField"); val.IsValid() {
		log.Panicf("BasicStruct.PrivateField() created. 'private' struct tag ignored")
	}

}
