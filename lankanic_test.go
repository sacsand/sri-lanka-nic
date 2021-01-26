package srilankanic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// var result1 = Result{
// 	isValidated: true,
// 	input:       "931260723v",
// 	inputFormat: "old",
// 	newFormat:   "199312600723",
// 	oldFormat:   "931260723v",
// 	nicLength:   10,
// 	gender:      "Male",
// 	character:   "--",
// 	birthday:    "05/05/1993",
// }

// var result2 = Result{
// 	isValidated: true,
// 	input:       "200184300068",
// 	year:        2001,
// 	inputFormat: "new",
// 	newFormat:   "200184300068",
// 	oldFormat:   "--",
// 	nicLength:   12,
// 	gender:      "Female",
// 	character:   "--",
// 	birthday:    "12/08/2001",
// }

// func TestINFO(t *testing.T) {
// 	assert := assert.New(t)

// 	res := Info("200184300068")
// 	assert.Equal(res, result2)

// 	fmt.Println()
// }

var result1 = Result{
	isValidated: true,
	input:       "199312600723",
	year:        1993,
	inputFormat: "new",
	newFormat:   "199312600723",
	oldFormat:   "931260723v",
	nicLength:   12,
	gender:      "Male",
	character:   "--",
	birthday:    "05/05/1993",
}

func TestINFOOld(t *testing.T) {
	assert := assert.New(t)

	res := Info("199312600723")
	assert.Equal(res, result1)

	fmt.Println()
}
