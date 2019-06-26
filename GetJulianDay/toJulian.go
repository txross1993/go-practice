package main

import (
	"fmt"
	"time"
)

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func getJDN(year, month, day int) int {
	yearStart := Date(year, 1, 1)
	inputDate := Date(year, month, day)

	days := inputDate.Sub(yearStart).Hours() / 24

	return int(days) + 1

}

func main() {

	// May 22, 2012 => 143
	fmt.Println(getJDN(2012, 5, 22))

	// Jan 1, 2001 => 1
	fmt.Println(getJDN(2001, 1, 1))

	// Dec 31, 2015 => 365
	fmt.Println(getJDN(2015, 12, 31))

	// Dec 31, 2016 => 366
	fmt.Println(getJDN(2016, 12, 31))

}
