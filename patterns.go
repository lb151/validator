package validator

import "regexp"

const (
	timeLayoutYm     = "2006-01"
	timeLayoutYmd    = "2006-01-02"
	timeLayoutYmdHis = "2006-01-02 15:04:05"

	maxURLRuneCount = 2083
	minURLRuneCount = 3

	emailExp        = "^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+)*@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$"
	idNoExp         = `^[1-9]\d{5}(18|19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])\d{3}[0-9Xx]$`
	mobileExp       = `^1[3-9]\d{9}$`
	mobileCodeExp   = `^[1-9]\d{0,3}-\d{7,11}$`
	IdsExp          = `^[1-9]\d*(,[1-9]\d*)*$`
	alphaExp        = `^[a-zA-Z]+$`
	alphanumericExp = `^[a-zA-Z0-9]+$`
	verifyCodeExp   = `^\d{6}$`
)

var (
	rxEmail        = regexp.MustCompile(emailExp)
	rxIdNo         = regexp.MustCompile(idNoExp)
	rxMobile       = regexp.MustCompile(mobileExp)
	rxMobileCode   = regexp.MustCompile(mobileCodeExp)
	rxIds          = regexp.MustCompile(IdsExp)
	rxAlpha        = regexp.MustCompile(alphaExp)
	rxAlphanumeric = regexp.MustCompile(alphanumericExp)
	rxVerifyCode   = regexp.MustCompile(verifyCodeExp)
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
