package srilankanic

import "strings"

type Month struct {
	month string
	days  int
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

type Result struct {
	isValidated bool
	year        int
	input       string
	inputFormat string
	newFormat   string
	oldFormat   string
	dayNo       string
	nicLength   int
	gender      string
	character   string
	birthday    string
}

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

func replaceAtIndex(input string, replacement byte, index int) string {
	return strings.Join([]string{input[:index], string(replacement), input[index+1:]}, "")
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}
