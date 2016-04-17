package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ParseMonth(t *testing.T) {
	assert := assert.New(t)

	m, y, err := ParseMonth("4.2016")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(m, time.April)
	assert.Equal(y, 2016)

	m, y, err = ParseMonth("12.2013")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(m, time.December)
	assert.Equal(y, 2013)

	m, y, err = ParseMonth("01.2013")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(m, time.January)
	assert.Equal(y, 2013)
}

func Test_ParseMonth_WrongFormat(t *testing.T) {
	_, _, err := ParseMonth("1")
	if err == nil {
		t.Fatal("Expect error")
	}
}

func Test_FormatDate(t *testing.T) {
	d := time.Date(2016, time.April, 17, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, "So, 17. April 2016", FormatDate(d))
}
