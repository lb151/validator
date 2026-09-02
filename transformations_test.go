package validator_test

import (
	"fmt"
	"testing"

	"github.com/lb151/validator"
)

func TestUppercase(t *testing.T) {
	t.Run("string to upper case", func(t *testing.T) {
		v := validator.Val("hello").Transform(validator.Uppercase())
		if v.Value() != "HELLO" {
			t.Errorf("Expected transformation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestLowercase(t *testing.T) {
	t.Run("string to lower case", func(t *testing.T) {
		v := validator.Val("HELLO").Transform(validator.Lowercase())
		if v.Value() != "hello" {
			t.Errorf("Expected transformation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestTrimSpace(t *testing.T) {
	t.Run("string with trimmed space", func(t *testing.T) {
		v := validator.Val(" hello ").Transform(validator.TrimSpace())
		if v.Value() != "hello" {
			t.Errorf("Expected transformation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestCapitalise(t *testing.T) {
	t.Run("string with capital first letter", func(t *testing.T) {
		v := validator.Val("bobby").Transform(validator.Capitalise())
		if v.Value() != "Bobby" {
			t.Errorf("Expected transformation to pass, got errors: %v", v.Errors())
		}
	})
}

func TestMultipleTransformations(t *testing.T) {
	t.Run("all transformations pass", func(t *testing.T) {
		v := validator.Val(" hello ").Transform(
			validator.TrimSpace(),
			validator.Uppercase(),
		)

		if !v.IsValid() {
			t.Errorf("Expected all transformations to pass, got errors: %v", v.Errors())
		}
		if v.Value() != "HELLO" {
			t.Errorf("Expected transformed value, got: %s", v.Value())
		}
	})
}

func TestTransformWithError(t *testing.T) {
	t.Run("transformation error is collected", func(t *testing.T) {
		// Custom transformation that returns an error
		v := validator.Val("hello").Transform(func(v validator.Value[string]) (string, error) {
			return "", fmt.Errorf("transformation failed")
		})

		if v.IsValid() {
			t.Error("Expected transformation to fail")
		}

		if len(v.Errors()) != 1 {
			t.Errorf("Expected 1 error, got %d", len(v.Errors()))
		}

		// Verify value remains unchanged when transformation fails
		if v.Value() != "hello" {
			t.Errorf("Expected value to remain 'hello', got %q", v.Value())
		}
	})
}
