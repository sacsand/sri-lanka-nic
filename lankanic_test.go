package srilankanic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var result2 = Result{
	isValidated: true,
	input:       "200184300068",
	inputFormat: "new",
	newFormat:   "200184300068",
	oldFormat:   "--",
	nicLength:   12,
	gender:      "Female",
	character:   "--",
	birthday:    "12/08/2001",
}

func TestINFO(t *testing.T) {
	assert := assert.New(t)

	res := Info("200184300068")
	assert.Equal(res, result2)

	fmt.Println()
}
