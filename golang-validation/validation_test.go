package validation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariabel(t *testing.T) {
	validate := validator.New()

	user := "eko"

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariabel(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmpassword := "rahasia"

	err := validate.VarWithValue(password, confirmpassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}
