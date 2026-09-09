package validator_test

import (
	"testing"

	"github.com/lb151/validator"
)

func TestRequired(t *testing.T) {
	t.Run("empty string fails", func(t *testing.T) {
		v := validator.Val("").Validate(validator.Required[string]())
		if v.IsValid() {
			t.Error("Expected validation to fail for empty string")
		}
	})

	t.Run("non-empty string passes", func(t *testing.T) {
		v := validator.Val("hello").Validate(validator.Required[string]())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("zero int fails", func(t *testing.T) {
		v := validator.Val(0).Validate(validator.Required[int]())
		if v.IsValid() {
			t.Error("Expected validation to fail for zero int")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "Custom required error"
		v := validator.Val("").Validate(validator.Required[string](customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("below min fails", func(t *testing.T) {
		v := validator.Val(5).Validate(validator.Min(10))
		if v.IsValid() {
			t.Error("Expected validation to fail for value below min")
		}
	})

	t.Run("at min passes", func(t *testing.T) {
		v := validator.Val(10).Validate(validator.Min(10))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("above min passes", func(t *testing.T) {
		v := validator.Val(15).Validate(validator.Min(10))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "Value too small"
		v := validator.Val(5).Validate(validator.Min(10, customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("above max fails", func(t *testing.T) {
		v := validator.Val(15).Validate(validator.Max(10))
		if v.IsValid() {
			t.Error("Expected validation to fail for value above max")
		}
	})

	t.Run("at max passes", func(t *testing.T) {
		v := validator.Val(10).Validate(validator.Max(10))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "Value too large"
		v := validator.Val(15).Validate(validator.Max(10, customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestMinLengthString(t *testing.T) {
	t.Run("below min length fails", func(t *testing.T) {
		v := validator.Val("ab").Validate(validator.MinLengthString(5))
		if v.IsValid() {
			t.Error("Expected validation to fail for string below min length")
		}
	})

	t.Run("at min length passes", func(t *testing.T) {
		v := validator.Val("hello").Validate(validator.MinLengthString(5))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "String too short"
		v := validator.Val("ab").Validate(validator.MinLengthString(5, customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestMaxLengthString(t *testing.T) {
	t.Run("above max length fails", func(t *testing.T) {
		v := validator.Val("hello").Validate(validator.MaxLengthString(3))
		if v.IsValid() {
			t.Error("Expected validation to fail for string above max length")
		}
	})

	t.Run("at max length passes", func(t *testing.T) {
		v := validator.Val("abc").Validate(validator.MaxLengthString(3))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "String too long"
		v := validator.Val("hello").Validate(validator.MaxLengthString(3, customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestMinLengthSlice(t *testing.T) {
	t.Run("below min length fails", func(t *testing.T) {
		v := validator.Val([]int{1}).Validate(validator.MinLengthSlice[int](2))
		if v.IsValid() {
			t.Error("Expected validation to fail for slice below min length")
		}
	})

	t.Run("at min length passes", func(t *testing.T) {
		v := validator.Val([]int{1, 2}).Validate(validator.MinLengthSlice[int](2))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "Slice too short"
		v := validator.Val([]int{1}).Validate(validator.MinLengthSlice[int](2, customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestMaxLengthSlice(t *testing.T) {
	t.Run("above max length fails", func(t *testing.T) {
		v := validator.Val([]int{1, 2, 3}).Validate(validator.MaxLengthSlice[int](2))
		if v.IsValid() {
			t.Error("Expected validation to fail for slice above max length")
		}
	})

	t.Run("at max length passes", func(t *testing.T) {
		v := validator.Val([]int{1, 2}).Validate(validator.MaxLengthSlice[int](2))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "Slice too long"
		v := validator.Val([]int{1, 2, 3}).Validate(validator.MaxLengthSlice[int](2, customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestMinLengthMap(t *testing.T) {
	t.Run("below min length fails", func(t *testing.T) {
		v := validator.Val(map[string]int{"a": 1}).Validate(validator.MinLengthMap[string, int](2))
		if v.IsValid() {
			t.Error("Expected validation to fail for map below min length")
		}
	})

	t.Run("at min length passes", func(t *testing.T) {
		v := validator.Val(map[string]int{"a": 1, "b": 2}).Validate(validator.MinLengthMap[string, int](2))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "Map too small"
		v := validator.Val(map[string]int{"a": 1}).Validate(validator.MinLengthMap[string, int](2, customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestMaxLengthMap(t *testing.T) {
	t.Run("above max length fails", func(t *testing.T) {
		v := validator.Val(map[string]int{"a": 1, "b": 2, "c": 3}).Validate(validator.MaxLengthMap[string, int](2))
		if v.IsValid() {
			t.Error("Expected validation to fail for map above max length")
		}
	})

	t.Run("at max length passes", func(t *testing.T) {
		v := validator.Val(map[string]int{"a": 1, "b": 2}).Validate(validator.MaxLengthMap[string, int](2))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "Map too large"
		v := validator.Val(map[string]int{"a": 1, "b": 2, "c": 3}).Validate(validator.MaxLengthMap[string, int](2, customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestEmail(t *testing.T) {
	t.Run("valid email passes", func(t *testing.T) {
		v := validator.Val("test@example.com").Validate(validator.Email())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid email fails", func(t *testing.T) {
		v := validator.Val("not-an-email").Validate(validator.Email())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid email")
		}
	})

	t.Run("email with unicode passes", func(t *testing.T) {
		v := validator.Val("tëst@example.com").Validate(validator.Email())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass for unicode email, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "Invalid email address"
		v := validator.Val("not-an-email").Validate(validator.Email(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestOneOf(t *testing.T) {
	t.Run("valid one of passes", func(t *testing.T) {
		v := validator.Val("delivered").Validate(validator.OneOf([]string{"shipped", "delivered"}))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid one of fails", func(t *testing.T) {
		v := validator.Val("delivered").Validate(validator.OneOf([]string{"shipped", "returned"}))
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid enum")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "Value must be one of provided"
		v := validator.Val("delivered").Validate(validator.OneOf([]string{"shipped", "returned"}, customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestForbidden(t *testing.T) {
	t.Run("valid not in passes", func(t *testing.T) {
		v := validator.Val("delivered").Validate(validator.NotIn([]string{"shipped", "returned"}))
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid not in fails", func(t *testing.T) {
		v := validator.Val("delivered").Validate(validator.NotIn([]string{"shipped", "delivered"}))
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid enum")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "Value cannot be one of provided"
		v := validator.Val("delivered").Validate(validator.NotIn([]string{"shipped", "delivered"}, customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestMultipleValidations(t *testing.T) {
	t.Run("accumulates multiple errors", func(t *testing.T) {
		v := validator.Val("ab").Validate(
			validator.Required[string](),
			validator.MinLengthString(5),
			validator.MaxLengthString(1),
		)

		if len(v.Errors()) != 2 {
			t.Errorf("Expected 2 errors (min and max), got %d: %v", len(v.Errors()), v.Errors())
		}
	})

	t.Run("all validations pass", func(t *testing.T) {
		v := validator.Val("hello").Validate(
			validator.Required[string](),
			validator.MinLengthString(3),
			validator.MaxLengthString(10),
		)

		if !v.IsValid() {
			t.Errorf("Expected all validations to pass, got errors: %v", v.Errors())
		}
	})
}

func TestIdNo(t *testing.T) {
	t.Run("valid id card passes", func(t *testing.T) {
		v := validator.Val("140101198001010016").Validate(validator.IdNo())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid id card fails", func(t *testing.T) {
		v := validator.Val("140101800101001").Validate(validator.IdNo())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid id card")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid id card"
		v := validator.Val("140101800101001").Validate(validator.IdNo(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestTimeYm(t *testing.T) {
	t.Run("valid time format ym passes", func(t *testing.T) {
		v := validator.Val("2026-09").Validate(validator.TimeYm())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid time format ym fails", func(t *testing.T) {
		v := validator.Val("2026").Validate(validator.TimeYm())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid ym")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid time format ym"
		v := validator.Val("2026").Validate(validator.TimeYm(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestTimeYmd(t *testing.T) {
	t.Run("valid time format ymd passes", func(t *testing.T) {
		v := validator.Val("2026-09-01").Validate(validator.TimeYmd())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid time format ymd fails", func(t *testing.T) {
		v := validator.Val("2026").Validate(validator.TimeYmd())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid ymd")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid time format ymd"
		v := validator.Val("2026").Validate(validator.TimeYmd(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestTimeYmdHis(t *testing.T) {
	t.Run("valid time format ymdHis passes", func(t *testing.T) {
		v := validator.Val("2026-09-01 10:00:00").Validate(validator.TimeYmdHis())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid time format ymdHis fails", func(t *testing.T) {
		v := validator.Val("2026").Validate(validator.TimeYmdHis())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid ymdHis")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid time format ymdHis"
		v := validator.Val("2026").Validate(validator.TimeYmdHis(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestURL(t *testing.T) {
	t.Run("valid url passes", func(t *testing.T) {
		v := validator.Val("http://www.test.com").Validate(validator.URL())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid url fails", func(t *testing.T) {
		v := validator.Val("test.com").Validate(validator.URL())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid url")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid url link"
		v := validator.Val("test.com").Validate(validator.URL(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestIP(t *testing.T) {
	t.Run("valid ip passes", func(t *testing.T) {
		v := validator.Val("192.168.0.1").Validate(validator.IP())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid ip fails", func(t *testing.T) {
		v := validator.Val("192.168.0").Validate(validator.IP())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid ip")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid ip address"
		v := validator.Val("192.168.0").Validate(validator.IP(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestJSON(t *testing.T) {
	t.Run("valid json passes", func(t *testing.T) {
		v := validator.Val(`{"name":"jom"}`).Validate(validator.JSON())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid json fails", func(t *testing.T) {
		v := validator.Val("abc").Validate(validator.JSON())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid json")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid json address"
		v := validator.Val("abc").Validate(validator.JSON(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestAlpha(t *testing.T) {
	t.Run("valid alpha passes", func(t *testing.T) {
		v := validator.Val(`Abc`).Validate(validator.Alpha())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid alpha fails", func(t *testing.T) {
		v := validator.Val("123").Validate(validator.Alpha())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid alpha")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid alpha"
		v := validator.Val("123").Validate(validator.Alpha(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestAlphanumeric(t *testing.T) {
	t.Run("valid alpha numeric passes", func(t *testing.T) {
		v := validator.Val(`Abc`).Validate(validator.Alphanumeric())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid alpha numeric fails", func(t *testing.T) {
		v := validator.Val("123").Validate(validator.Alphanumeric())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid alpha numeric")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid alpha numeric"
		v := validator.Val("123").Validate(validator.Alphanumeric(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestNumeric(t *testing.T) {
	t.Run("valid numeric passes", func(t *testing.T) {
		v := validator.Val[int](0).Validate(validator.Numeric[int]())
		if !v.IsValid() {
			t.Errorf("Expected validation to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("invalid numeric fails", func(t *testing.T) {
		v := validator.Val[int](-10).Validate(validator.Numeric[int]())
		if v.IsValid() {
			t.Error("Expected validation to fail for invalid numeric")
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid numeric"
		v := validator.Val[int](-10).Validate(validator.Numeric[int](customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})
}

func TestIncID(t *testing.T) {
	t.Run("incremental ID validations error", func(t *testing.T) {
		v := validator.Val[int](0).Validate(validator.IncID[int]())

		if v.IsValid() {
			t.Errorf("Expected error, got %v", v.Errors())
		}
	})

	t.Run("ids validations pass", func(t *testing.T) {
		v := validator.Val[int](10).Validate(
			validator.IncID[int](),
		)

		if !v.IsValid() {
			t.Errorf("Expected validations to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "invalid auto-incremented ID"
		v := validator.Val[int](-10).Validate(validator.IncID[int](customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})

}

func TestIds(t *testing.T) {
	t.Run("ids validations error", func(t *testing.T) {
		v := validator.Val("1,2,a").Validate(
			validator.IDs(),
		)

		if v.IsValid() {
			t.Errorf("Expected error, got %v", v.Errors())
		}
	})

	t.Run("ids validations pass", func(t *testing.T) {
		v := validator.Val("1,2,3").Validate(
			validator.IDs(),
		)

		if !v.IsValid() {
			t.Errorf("Expected validations to pass, got errors: %v", v.Errors())
		}
	})

	t.Run("custom error message", func(t *testing.T) {
		customMsg := "valid comma-separated ID string"
		v := validator.Val("1,2,a").Validate(validator.IDs(customMsg))
		if v.IsValid() {
			t.Error("Expected validation to fail")
		}
		if v.Errors()[0].Error() != customMsg {
			t.Errorf("Expected %q, got %q", customMsg, v.Errors()[0].Error())
		}
	})

}
