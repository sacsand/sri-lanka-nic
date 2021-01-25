package srilankanic

import (
	"fmt"
	"strconv"
)

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
		fmt.Println("NICX::" + nic)
		c := nic[0:4] //susbstring like js
		year, _ := strconv.Atoi(c)

		data.isValidated = true
		data.input = nic
		data.inputFormat = "new"
		data.year = year
		data.nicLength = 12
		data.birthday = "12/08/2001"
		data.character = "--"
		// data.dayNo = "v"
		data.gender = "Female"
		data.newFormat = nic
		data.oldFormat = "--" // defaualt

		fmt.Println(data)

		if nic[0:2] == "19" {

			st10 := nic[1:12]

			// st10 =
			b := byte(0)
			news := replaceAtIndex(st10, b, 5)
			fmt.Println(news)
			news = news + "v"
			fmt.Println(news)
			data.oldFormat = news
		}

	}

	if nicLength == 10 {

		c := nic[0:2] //susbstring like js
		year, _ := strconv.Atoi(c)

		data.isValidated = true
		data.input = nic
		data.inputFormat = "new"
		data.year = year
		data.nicLength = 12
		data.birthday = "05/05/1993"
		data.character = "v"
		data.gender = "Male"

		fmt.Println(data)

		st10 := nic[2:12]

		// news := replaceAtIndex(st10, b, 5)
		s := []rune(st10)
		res := delChar(s, 5)

		data.oldFormat = string(res)

	}
	return data
}
