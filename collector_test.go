package validator_test

import (
	"testing"

	"github.com/lb151/validator"
)

func TestCollector(t *testing.T) {
	t.Run("collects errors from multiple values", func(t *testing.T) {
		c := validator.NewCollector()

		name := validator.Val("").Validate(validator.Required[string]()).Collect(c)
		age := validator.Val(15).Validate(validator.Min(18)).Collect(c)

		if c.IsValid() {
			t.Error("Collector should have errors")
		}

		if len(c.Errors()) != 2 {
			t.Errorf("Expected 2 errors, got %d: %v", len(c.Errors()), c.Errors())
		}

		// Verify collected values are returned correctly
		if name != "" {
			t.Errorf("Expected empty string, got %q", name)
		}
		if age != 15 {
			t.Errorf("Expected 15, got %d", age)
		}
	})

	t.Run("collector with no errors", func(t *testing.T) {
		c := validator.NewCollector()

		name := validator.Val(" John ").Validate(validator.Required[string]()).Transform(validator.TrimSpace()).Collect(c)
		age := validator.Val(25).Validate(validator.Min(18)).Collect(c)

		if !c.IsValid() {
			t.Errorf("Collector should be valid, got errors: %v", c.Errors())
		}

		if len(c.Errors()) != 0 {
			t.Errorf("Expected 0 errors, got %d", len(c.Errors()))
		}

		// Verify collected values
		if name != "John" {
			t.Errorf("Expected 'John', got %q", name)
		}
		if age != 25 {
			t.Errorf("Expected 25, got %d", age)
		}
	})
}
