package validator_test

import (
	"testing"

	"github.com/lb151/validator"
)

func TestValidateFast(t *testing.T) {
	t.Run("all validations pass", func(t *testing.T) {
		email := validator.Val("1200857@qq.com ").Transform(validator.TrimSpace())

		err := validator.ValidateFast(validator.Field(email, validator.Required[string]("mail is not empty"), validator.Email("incorrect email format")),
			validator.Field(validator.Val("jom", "name"), validator.Alpha("")),
		)
		if err != nil {
			t.Errorf("Expected validation to pass, got errors: %v", err)
		}

	})

	t.Run("the first field verification fail", func(t *testing.T) {
		err := validator.ValidateFast(
			validator.Field(validator.Val("jom120", "name"), validator.Alpha("is not a valid alpha")),
			validator.Field(validator.Val("1200857@qq.com", "mail").Transform(validator.TrimSpace()), validator.Email("incorrect email format")),
			validator.Field(validator.Val[int](20, "age"), validator.Max[int](30, "the maximum value cannot exceed 30")),
		)
		if err == nil {
			t.Errorf("the first field for expected verification")
		}
	})

	t.Run("the last field verification fail", func(t *testing.T) {
		err := validator.ValidateFast(
			validator.Field(validator.Val("1200857@qq.com", "mail").Transform(validator.TrimSpace()), validator.Email("incorrect email format")),
			validator.Field(validator.Val("jom123", "name"), validator.Alphanumeric("is not a valid alpha number")),
			validator.Field(validator.Val[int](20, "age"), validator.Min[int](30, "the minimum value must not be lower than 30")),
		)
		if err == nil {
			t.Errorf("the first field for expected verification")
		}
	})
}
