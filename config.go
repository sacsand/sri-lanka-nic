package srilankanic

type Month struct {
	month string
	days  int
}

type Result struct {
	isValidated bool
	year        int
	input       string
	inputFormat string
	newFormat   string
	oldFormat   string
	nicLength   int
	gender      string
	character   string
	birthday    string
}

type BasicFetch struct {
	isValidated bool
	year        int
	input       string
	inputFormat string
	newFormat   string
	oldFormat   string
	nicLength   int
	gender      string
	character   string
	birthday    string
}
