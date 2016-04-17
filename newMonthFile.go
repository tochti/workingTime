package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func NewMonthFile(month string, tplPath string) {
	m, y, err := ParseMonth(month)
	if err != nil {
		log.Fatal(err)
	}

	daysN := DaysIn(m, y)
	daysInMonth := []string{}
	for n := 1; n <= daysN; n++ {
		tmp := time.Date(y, m, 0+n, 0, 0, 0, 0, time.UTC)
		f := FormatDate(tmp)
		daysInMonth = append(daysInMonth, f)
	}
	t := template.New("month")
	t = template.Must(t.ParseFiles(tplPath))

	outputFile := fmt.Sprintf("%v.%v.yml", m, y)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}

	ctx := struct {
		Days []string
	}{
		Days: daysInMonth,
	}

	err = t.Execute(f, ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func DaysIn(m time.Month, y int) int {
	return time.Date(y, m, 0, 0, 0, 0, 0, time.UTC).Day()
}

func ParseMonth(month string) (time.Month, int, error) {
	tmp := strings.Split(month, ".")
	if len(tmp) != 2 {
		return time.January, -1, fmt.Errorf("Wrong month format")
	}

	i, err := strconv.Atoi(tmp[0])
	if err != nil {
		return time.January, -1, err
	}

	m, err := Month(i)
	if err != nil {
		return time.January, -1, err
	}

	y, err := strconv.Atoi(tmp[1])
	if err != nil {
		return time.January, -1, err
	}

	return m, y, err
}

func Month(m int) (time.Month, error) {
	m = m - 1
	if m > 11 && m < 0 {
		return time.January, fmt.Errorf("%v out of range", m)
	}

	return []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}[m], nil
}

func FormatDate(d time.Time) string {
	s := fmt.Sprintf("%v, %v. %v %v",
		WeekdayShortGerman(d.Weekday()),
		d.Day(), d.Month(), d.Year())

	return s
}

func WeekdayShortGerman(d time.Weekday) string {
	return map[time.Weekday]string{
		time.Sunday:    "So",
		time.Monday:    "Mo",
		time.Tuesday:   "Di",
		time.Wednesday: "Mi",
		time.Thursday:  "Do",
		time.Friday:    "Fr",
		time.Saturday:  "Sa",
	}[d]
}
