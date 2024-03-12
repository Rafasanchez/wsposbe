package functions

import (
	"time"

	"gorm.io/datatypes"
)

func Today() datatypes.Date {
	currentDate := datatypes.Date(time.Now())

	return currentDate
}
