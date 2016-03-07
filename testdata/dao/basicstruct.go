package main

import "time"

//import "log"
//import "reflect"
import "github.com/traitify/common-go/unixtime"

type Basicstruct struct {
	id            int             `db:"id"`
	date          time.Time       `db:"date"`
	prevDate      time.Time       `db:"prev_date"`
	status        string          `db:"status" whatever:"foo"`
	hiddenField   string          `db:"hidden_field"`
	privateField  []*string       `db:"private_field" propertizer:"private" exclude_dao:"true"`
	modifiedAt    unixtime.Millis `db:"modified_at"`
	unmarkedField string
	NeedsInsert   bool `propertizer:"private"`
}

func main() {

}
