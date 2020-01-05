package main

import (
	"errors"
	"log"
	"time"
)

func main() {
	currentYear := time.Now()
	endate := currentYear.Format("2006-01-02")
	begin := currentYear.AddDate(-1, 0, 0)
	prev := begin.AddDate(0, 0, 1).Format("2006-01-02")

	log.Println(prev)
	log.Println(begin)
	log.Println(endate)

	y, m, d := time.Now().Date()
	log.Println(y)
	log.Println(m)
	log.Println(d)

	DateFrom := time.Date(y, m, d, 23, 59, 59, 999999999, time.Local).Format("02.01.2006 15:04:05")
	log.Println(DateFrom)

	fd := time.Duration(-int(time.Now().YearDay())+1) * 24 * time.Hour
	fday := time.Now().Add(fd)
	log.Println(fday)

}

// sudo ./app nohup &

// DateValidation ...
func DateValidation(date string) (err error) {
	// date format is "DD.MM.YYYY HH:MM:SS"
	var layout = "02.01.2006"
	_, err = time.Parse(layout, date)
	if err != nil {
		err = errors.New("Формат даты должень быть: " + layout)
	}
	return
}
