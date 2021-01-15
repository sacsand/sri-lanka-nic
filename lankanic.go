package srilankanic

import (
	"fmt"
	"strconv"
	"strings"
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

func replaceAtIndex(input string, replacement byte, index int) string {
	return strings.Join([]string{input[:index], string(replacement), input[index+1:]}, "")
}

// Info ..
func Info(nic string) Result {

	var data Result

	nicLength := len(nic)

	if nicLength == 12 {

		c := nic[0:4] //susbstring like js
		year, _ := strconv.Atoi(c)

		fmt.Println(year)

		data.isValidated = true
		data.input = nic
		data.inputFormat = "new"
		data.year = year
		data.nicLength = 12
		data.birthday = "05/05/1993"
		data.character = "v"
		data.dayNo = "v"
		data.gender = "Male"

		fmt.Println(data)

		//if nic[0:2] == "19" {

		st10 := nic[2:12]
		fmt.Println(st10)
		// st10 =
		b := byte(0)
		news := replaceAtIndex(st10, b, 5)
		fmt.Println(news)
		news = news + "v"
		fmt.Println(news)
		data.oldFormat = news

	}

	if nicLength == 10 {

		c := nic[0:2] //susbstring like js
		year, _ := strconv.Atoi(c)

		fmt.Println(year)

		data.isValidated = true
		data.input = nic
		data.inputFormat = "old"
		data.year = year
		data.nicLength = 12
		data.birthday = "05/05/1993"
		data.character = "v"
		data.gender = "Male"

		fmt.Println(data)

		//if nic[0:2] == "19" {

		st10 := nic[2:12]
		fmt.Println(st10)
		// st10 =
		b := byte(0)
		news := replaceAtIndex(st10, b, 5)
		fmt.Println(news)
		news = news + "v"
		fmt.Println(news)
		data.oldFormat = news

	}
	return data
}
