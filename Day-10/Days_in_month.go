package main

import "fmt"

func checkLeap(year int) bool {
	flag := false
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				flag = true
			} else {
				flag = false
			}
		} else {
			flag = true
		}
	} else {
		flag = false
	}
	return flag
}

func daysInMonth(year, month int) {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	leap := checkLeap(year)
	if leap {
		days[1] = 29
	}
	fmt.Printf("The number of days in month #%v of year %v is %v\n", month, year, days[month-1])
}

func main() {
	var year, month int
	fmt.Printf("Enter year : ")
	fmt.Scanf("%d", &year)
	fmt.Printf("Enter month : ")
	fmt.Scanf("%d", &month)
	daysInMonth(year, month)
}
