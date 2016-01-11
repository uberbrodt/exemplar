package main

import "time"

//import "log"
//import "reflect"
import "github.com/traitify/common-go/unixtime"

type Basicstruct struct {
	id           int             `db:"id"`
	date         time.Time       `db:"date"`
	prevDate     time.Time       `db:"prev_date"`
	status       string          `db:"status" whatever:"foo"`
	hiddenField  string          `db:"hidden_field"`
	privateField string          `db:"private_field" propertizer:"private" exclude_dao:"true"`
	modifiedAt   unixtime.Millis `db:"modified_at"`
	NeedsInsert  bool
}

//func main() {
//
//	privateField := "shouldnotread"
//
//	currentTime := time.Now()
//	bs := &Basicstruct{id: 10, date: currentTime, status: "FooBar", privateField: privateField}
//
//	if bs.ID() != 10 {
//		log.Panicf("BasicStruct.ID() returned: %d not %d", bs.ID(), 10)
//	}
//	if !currentTime.Equal(bs.Date()) {
//		log.Panicf("BasicStruct.Date() returned: %v, not %v", bs.Date(), currentTime)
//	}
//	if bs.Status() != "FooBar" {
//		log.Panicf("BasicStruct.Status() returned: %s, not %s", bs.Status(), "FooBar")
//	}
//
//	typ := reflect.ValueOf(bs)
//	if val := typ.MethodByName("PrivateField"); val.IsValid() {
//		log.Panicf("BasicStruct.PrivateField() created. 'private' struct tag ignored")
//	}
//
//}
