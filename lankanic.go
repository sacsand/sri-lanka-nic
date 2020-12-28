package srilankanic

import (
	"fmt"
	"strconv"
)

var monthList = []Month{
	{
		month: "January",
		days:  31,
	},
	{
		month: "February",
		days:  29,
	},
	{
		month: "March",
		days:  31,
	},
	{
		month: "April",
		days:  30,
	},
	{
		month: "May",
		days:  31,
	},
	{
		month: "June",
		days:  30,
	},
	{
		month: "July",
		days:  31,
	},
	{
		month: "August",
		days:  31,
	},
	{
		month: "Septhember",
		days:  30,
	},
	{
		month: "October",
		days:  31,
	},
	{
		month: "November",
		days:  30,
	},
	{
		month: "December",
		days:  31,
	},
}

// func Validate(nationalIdNumber string) bool {
// 	return false
// }

// func Swap(nati onalIdNumber string) string {
// 	return "dsfsd"
// }

// Info ..
func Info(nic string) Result {

	var data Result

	nicLength := len(nic)

	if nicLength == 12 {

		data.input = nic
		data.inputFormat = "new"
		c := nic[0:4]

		year, _ := strconv.Atoi(c)

		data.year = year
		data.nicLength = 12
		data.character = ""

		fmt.Println(data)

	}

	// type Result struct {
	// 	isValidated bool
	// 	year        int
	// 	input       string
	// 	inputFormat string
	// 	newFormat   string
	// 	oldFormat   string
	// 	nicLength   int
	// 	gender      string
	// 	character   string
	// 	birthday    string
	// }

	// details := Result{
	// 	isValidated: true,
	// 	input:       "200184300068",
	// 	format:      "new",
	// 	newFormat:   "200184300068",
	// 	oldFormat:   "--",
	// 	nicLength:   12,
	// 	gender:      "Female",
	// 	character:   "--",
	// 	birthday:    "12/08/2001",
	// }
	return data
}
