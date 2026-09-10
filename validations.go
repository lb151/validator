package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"slices"
	"time"
)

// Required returns a validation that ensures the value is
// not the zero value for its type.
//
// For strings, this means non-empty. For numbers, this means
// non-zero. For pointers, this means non-nil.
//
// An optional custom error message can be provided as the
// parameter.
//
// Example:
//
//	validator.Val("").Validate(validator.Required[string]())  // fails
//	validator.Val("John").Validate(validator.Required[string]())  // passes
func Required[T comparable](errsMsg ...string) func(Value[T]) error {
	return func(v Value[T]) error {
		var zero T
		if v.value == zero {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return errors.New(v.name + " is required")
		}

		return nil
	}
}

// Max returns a validation that ensures the value does
// not exceed the given maximum.
//
// Works with all numeric types defined by the Ordered
// constraint.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val(100).Validate(validator.Max(100))
func Max[T Ordered](max T, errsMsg ...string) func(Value[T]) error {
	return func(v Value[T]) error {
		if v.value > max {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return fmt.Errorf("%s cannot be larger than %v", v.name, max)
		}

		return nil
	}
}

// Min returns a validation that ensures the value is
// at least the given minimum.
//
// Works with all numeric types defined by the Ordered
// constraint.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val(5).Validate(validator.Min(1))
func Min[T Ordered](min T, errsMsg ...string) func(Value[T]) error {
	return func(v Value[T]) error {
		if v.value < min {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return fmt.Errorf("%s cannot be smaller than %v", v.name, min)
		}

		return nil
	}
}

// MaxLengthString returns a validation that ensures the
// length of a string does not exceed the given maximum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val("username").Validate(validator.MaxLengthString(20))
func MaxLengthString(max int, errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if len([]rune(v.value)) > max {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return fmt.Errorf("%s's length cannot be larger than %v", v.name, max)
		}

		return nil
	}
}

// MaxLengthSlice returns a validation that ensures the
// length of a slice does not exceed the given maximum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val([]int{1}).Validate(validator.MaxLengthSlice(2))
func MaxLengthSlice[T any](max int, errsMsg ...string) func(Value[[]T]) error {
	return func(v Value[[]T]) error {
		if len(v.value) > max {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return fmt.Errorf("%s's length cannot be larger than %v", v.name, max)
		}

		return nil
	}
}

// MaxLengthMap returns a validation that ensures the
// length of a map does not exceed the given maximum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val(map[string]int{"no": 1}).Validate(validator.MaxLengthMap(2))
func MaxLengthMap[K comparable, V any](max int, errsMsg ...string) func(Value[map[K]V]) error {
	return func(v Value[map[K]V]) error {
		if len(v.value) > max {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return fmt.Errorf("%s's length cannot be larger than %v", v.name, max)
		}

		return nil
	}
}

// MinLengthString returns a validation that ensures the
// length of a string is at least the given minimum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val("username").Validate(validator.MinLengthString(5))
func MinLengthString(min int, errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if len([]rune(v.value)) < min {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return fmt.Errorf("%s's length cannot be smaller than %v", v.name, min)
		}

		return nil
	}
}

// MinLengthSlice returns a validation that ensures the
// length of a slice is at least the given minimum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val([]int{1}).Validate(validator.MinLengthSlice(1))
func MinLengthSlice[T any](min int, errsMsg ...string) func(Value[[]T]) error {
	return func(v Value[[]T]) error {
		if len(v.value) < min {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return fmt.Errorf("%s's length cannot be smaller than %v", v.name, min)
		}

		return nil
	}
}

// MinLengthMap returns a validation that ensures the
// length of a map is at least the given minimum.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val(map[string]int{"no": 1}).Validate(validator.MinLengthMap(1))
func MinLengthMap[K comparable, V any](min int, errsMsg ...string) func(Value[map[K]V]) error {
	return func(v Value[map[K]V]) error {
		if len(v.value) < min {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return fmt.Errorf("%s's length cannot be smaller than %v", v.name, min)
		}

		return nil
	}
}

// Email returns a validation that ensures the value
// is a valid email address.
//
// It uses a practical, internationally-aware pattern
// that catches common errors, while remaining permissive.
//
// For true validation, send a confirmation email.
//
// An optional custom error message can be provided as the
// parameter.
//
// Example:
//
//	validator.Val("user@example.com").Validate(validator.Email())
func Email(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		flag := true
		if !rxEmail.MatchString(v.value) {
			flag = false
		}

		if flag {
			_, err := mail.ParseAddress(v.value)
			if err != nil {
				flag = false
			}
		}

		if !flag {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return errors.New(v.name + " must be in correct email format")
		}

		return nil
	}
}

// OneOf returns a validation that ensures the value matches
// one of the provided allowed values.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val("pending").Validate(validator.OneOf([]string{"pending", "approved", "rejected"}))
func OneOf[T comparable](values []T, errsMsg ...string) func(Value[T]) error {
	return func(v Value[T]) error {
		if !slices.Contains(values, v.value) {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return fmt.Errorf("%s must be one of: %v", v.name, values)
		}

		return nil
	}
}

// NotIn returns a validation that ensures the value
// does not match any of the provided forbidden values.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val("john").Validate(validator.NotIn([]string{"admin", "root", "system"}))
func NotIn[T comparable](values []T, errsMsg ...string) func(Value[T]) error {
	return func(v Value[T]) error {
		if slices.Contains(values, v.value) {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return fmt.Errorf("%s cannot be one of: %v", v.name, values)
		}

		return nil
	}
}

// IdNo returns a validation that ensures the value
// does not match any of the provided forbidden values.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val("140101198001010016").Validate(validator.IdNo("is not valid ID card format")

func IdNo(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		flag := true
		if !rxIdNo.MatchString(v.value) {
			flag = false

		}

		if flag {
			pt, err := time.Parse("20060102", v.value[6:14])
			if err != nil || pt.After(time.Now()) {
				flag = false
			}
		}

		if flag {
			weights := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
			checkMap := map[int]byte{0: '1', 1: '0', 2: 'X', 3: '9', 4: '8', 5: '7', 6: '6', 7: '5', 8: '4', 9: '3', 10: '2'}

			sum := 0
			for i := range 17 {
				sum += int(v.value[i]-'0') * weights[i]
			}
			want := checkMap[sum%11]
			last := v.value[17]
			if last == 'x' {
				last = 'X'
			}

			if last != want {
				flag = false
			}

		}

		if !flag {
			// Return custom error message, if provided
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return errors.New(v.name + " must be in the valid ID card format")
		}

		return nil
	}
}

// Example:
//
//	validator.Val(`1502900957`).Validate(validator.Mobile("is not a mobile number"))

func Mobile(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if rxMobile.MatchString(v.value) {
			return nil
		}

		if len(errsMsg) > 0 && errsMsg[0] != "" {
			return errors.New(errsMsg[0])
		}

		return errors.New(v.name + " must be a mobile number")
	}
}

// Example:
//
//	validator.Val(`86-1502900957`).Validate(validator.Mobile("is not a mobile number"))

func MobileWithCode(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if rxMobileCode.MatchString(v.value) {
			return nil
		}

		if len(errsMsg) > 0 && errsMsg[0] != "" {
			return errors.New(errsMsg[0])
		}

		return errors.New(v.name + " must be a mobile number")
	}
}

// TimeYm returns a validation that ensures the value
// does not match any of the provided forbidden values.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val("2026-09").Validate(validator.TimeYm("Inconsistent time format"))

func TimeYm(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		_, err := time.Parse(timeLayoutYm, v.value)
		if err != nil {
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return errors.New(v.name + " not a valid date format")
		}

		return nil
	}

}

// TimeYmd returns a validation that ensures the value
// does not match any of the provided forbidden values.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val("2026-09-01").Validate(validator.TimeYmd("Inconsistent time format"))

func TimeYmd(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		_, err := time.Parse(timeLayoutYmd, v.value)
		if err != nil {
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return errors.New(v.name + " not a valid date format")
		}

		return nil
	}

}

// TimeYmdHis returns a validation that ensures the value
// does not match any of the provided forbidden values.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val("2026-09-01 10:10:01").Validate(validator.TimeYmdHis("Inconsistent time format"))

func TimeYmdHis(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		_, err := time.Parse(timeLayoutYmdHis, v.value)
		if err != nil {
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return errors.New(v.name + " not a valid date format")
		}

		return nil
	}

}

// URL returns a validation that ensures the value
// does not match any of the provided forbidden values.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val("https://www.baidu.com").Validate(validator.URL("is not a valid URL"))

func URL(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		flag := true
		urlLen := len([]rune(v.value))
		if v.value == "" || urlLen >= maxURLRuneCount || urlLen <= minURLRuneCount {
			flag = false
		}

		if flag {
			u, err := url.Parse(v.value)
			if err != nil || u.Scheme == "" || u.Host == "" {
				flag = false
			}
		}

		if !flag {
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return errors.New(v.name + " is not a valid URL")
		}

		return nil
	}
}

// IP returns a validation that ensures the value
// does not match any of the provided forbidden values.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val("127.0.0.1").Validate(validator.IP("is not a valid IP"))

func IP(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		addr := net.ParseIP(v.value)
		if addr == nil {
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return errors.New(v.name + " is not a valid IP")
		}

		return nil
	}
}

// JSON returns a validation that ensures the value
// does not match any of the provided forbidden values.
//
// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val(`{"name": 123}`).Validate(validator.JSON("is not a valid JSON"))

func JSON(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if json.Valid([]byte(v.value)) {
			return nil
		}

		if len(errsMsg) > 0 && errsMsg[0] != "" {
			return errors.New(errsMsg[0])
		}

		return errors.New(v.name + " is not a valid JSON")
	}
}

// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val(`abC`).Validate(validator.Alpha("is not a valid alpha"))

func Alpha(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if rxAlpha.MatchString(v.value) {
			return nil
		}

		if len(errsMsg) > 0 && errsMsg[0] != "" {
			return errors.New(errsMsg[0])
		}

		return errors.New(v.name + " is not a valid alpha")
	}
}

// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val(`Ac09`).Validate(validator.Alphanumeric("is not a valid alphanumeric"))

func Alphanumeric(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if rxAlphanumeric.MatchString(v.value) {
			return nil
		}

		if len(errsMsg) > 0 && errsMsg[0] != "" {
			return errors.New(errsMsg[0])
		}

		return errors.New(v.name + " is not a valid alphanumeric")
	}
}

// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val(-1).Validate(validator.Numeric("cannot be a numeric"))

func Numeric[T IntOrdered](errsMsg ...string) func(Value[T]) error {
	return func(v Value[T]) error {
		var zero T
		if v.value < zero {
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return errors.New(v.name + " cannot be a numeric")
		}

		return nil
	}
}

// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val(100).Validate(validator.IncID("must be greater than zero"))

func IncID[T IntOrdered](errsMsg ...string) func(Value[T]) error {
	return func(v Value[T]) error {
		var zero T
		if v.value <= zero {
			if len(errsMsg) > 0 && errsMsg[0] != "" {
				return errors.New(errsMsg[0])
			}

			return errors.New(v.name + " must be greater than zero")
		}

		return nil
	}
}

// An optional custom error message can be provided as the
// last parameter.
//
// Example:
//
//	validator.Val(`1,10,20`).Validate(validator.IDs("is not a valid comma-separated ID string"))

func IDs(errsMsg ...string) func(Value[string]) error {
	return func(v Value[string]) error {
		if rxIds.MatchString(v.value) {
			return nil
		}

		if len(errsMsg) > 0 && errsMsg[0] != "" {
			return errors.New(errsMsg[0])
		}

		return errors.New(v.name + " is not a valid comma-separated ID string")
	}
}
