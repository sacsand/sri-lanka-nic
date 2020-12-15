package lankanic

import "testing"

var result2 = Result{
	isValidated: true,
	input:       "200184300068",
	format:      "new",
	newFormat:   "200184300068",
	oldFormat:   "--",
	nicLength:   12,
	gender:      "Female",
	character:   "--",
	birthday:    "12/08/2001",
}

func TestValidation(t *testing.T) {
	lankanic.Validate("sdsdsds")
}
