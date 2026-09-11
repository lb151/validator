package validator

type FieldCollector struct {
	run func() error
}

// Field declares a Value[T] along with several validation rules as a field to be executed (without applying any rules).
//
// Linked with the unified entry point Val(): The value and name are passed in wrapped by Val,
// and the Transform already applied on the Value will take effect; if the Value already has an error
// (such as a failed Transform), this field will fail directly with the first existing error.
//
// Multiple rules within a field are short-circuited in sequence: If the first rule fails, the process is terminated and no further rules are executed.
// Example:
//
//	err := validator.ValidateFast(
//	    validator.Field(validator.Val(u.Name, "name"),
//	        validator.Required[string](), validator.MinLengthString(3)),
//	    validator.Field(validator.Val(u.Email, "email"),
//	        validator.Required[string](), validator.Email()),
//	    validator.Field(validator.Val(u.Age, "age"), validator.Min(18)),
//
// )
//
//		if err ! = nil {
//		   return err // This returns the first error that occurred, and subsequent fields are not processed.
//	 }
//
// The same Value can also first go through the Collector's full mode, or first undergo Transformation and then perform short-circuit verification:
//
//	email := validator.Val(input.Email, "email").Transform(validator.TrimSpace())
//	err := validator.ValidateFast(
//	    validator.Field(email, validator.Email()),
//	    // ...
//	)
func Field[T any](v Value[T], rules ...func(Value[T]) error) FieldCollector {
	return FieldCollector{
		run: func() error {
			if len(v.errs) > 0 {
				return v.errs[0] // transmitting errors that exist (such as failed Transform)
			}

			for _, fn := range rules {
				if err := fn(v); err != nil {
					return err
				}
			}

			return nil
		},
	}
}

// ValidateFast: Perform a one-time verification of all fields:
//
// Execute in the order of declaration. As soon as the first failed field is encountered, break immediately.
// Return the first error of that field; if all pass, return nil.
//
// The unexecuted fields (and the associated rules) are not executed at all.
//
//Note: This method only returns error messages, not the Value[T] structure data.

func ValidateFast(fields ...FieldCollector) error {
	for _, f := range fields {
		if err := f.run(); err != nil {
			return err
		}
	}
	return nil
}
