package validator

import "regexp"

const (
	timeLayoutYm     = "2006-01"
	timeLayoutYmd    = "2006-01-02"
	timeLayoutYmdHis = "2006-01-02 15:04"

	maxURLRuneCount = 2083
	minURLRuneCount = 3

	emailExp        = `^(?:"(?:[^"]|\\")*"|[\p{L}\p{N}\p{M}._%+-]+)@[\p{L}\p{N}\p{M}.-]+\.[\p{L}\p{M}]{2,}$`
	idNoExp         = `(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$)`
	IdsExp          = `^[1-9]\d*(,[1-9]\d*)*$`
	alphaExp        = "^[a-zA-Z]+$"
	alphanumericExp = "^[a-zA-Z0-9]+$"
)

var (
	rxEmail        = regexp.MustCompile(emailExp)
	rxIdNo         = regexp.MustCompile(idNoExp)
	rxIds          = regexp.MustCompile(IdsExp)
	rxAlpha        = regexp.MustCompile(alphaExp)
	rxAlphanumeric = regexp.MustCompile(alphanumericExp)
)

// Ordered is a constraint that permits all numeric types
// that support comparison operations (<, >, <=, >=).
type (
	Ordered interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
	}
	IntOrdered interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
	}
)
