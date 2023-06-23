package random

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"tabungan-api/utils/cast"
)

// GenerateAccountNo generates a random account no
func GenerateAccountNo() string {
	return strconv.Itoa(GenerateNumber(1000000000, 9999999999)) + strconv.Itoa(GenerateNumber(1000000000, 9999999999))
}

// GenerateBirthDate generates a random birth date
func GenerateBirthDate() string {
	// get current time
	now := time.Now()

	// set minimum year to be 18 years ago
	minYear := now.Year() - 18

	// set minimum and maximum date
	minDate := time.Date(minYear, now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	maxDate := time.Date(now.Year()-100, now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	if maxDate.Before(minDate) {
		minDate, maxDate = maxDate, minDate
	}

	// generate a random date between minimum and maximum date
	randomDate := minDate.Add(time.Duration(r.Int63n(maxDate.Unix()-minDate.Unix())) * time.Second)

	// format the date as "DD-MM-YYYY"
	formattedDate := randomDate.Format("02-01-2006")

	return formattedDate
}

// GenerateCardNo generates a random card no
func GenerateCardNo() string {
	return strconv.Itoa(GenerateNumber(1000000000000000, 9999999999999999))
}

// GenerateEmail generates a random email address
func GenerateEmail() string {
	usernameLength := GenerateNumber(5, 10)
	domainLength := GenerateNumber(5, 10)

	username := GenerateAlphanumericString(usernameLength)
	domain := GenerateAlphabetString(domainLength)

	emailDomain := GenerateFromSet(cast.StringsliceToAnyslice(strings.Split(emailExtensions, " ")))

	return fmt.Sprintf("%s@%s.%s", username, domain, emailDomain)
}

// GenerateGender generates a random gender
func GenerateGender() string {
	isMale := GenerateBool()

	if isMale {
		return "L"
	} else {
		return "P"
	}
}

// GenerateMaritalStatus generates a random marital status
func GenerateMaritalStatus() string {
	maritalStatus := GenerateFromSet([]interface{}{"single", "married", "divorced", "widowed", "separated"}).(string)

	switch maritalStatus {
	case "single":
		return "1"
	case "married":
		return "2"
	case "divorced":
		return "3"
	case "widowed":
		return "4"
	default:
		return "5"
	}
}

// GenerateOtp generates a random otp
func GenerateOtp() string {
	// return strconv.Itoa(GenerateNumber(100000, 999999))

	return "123456"
}

// GeneratePhoneNo generates a random phone no
func GeneratePhoneNo() string {
	return "08" + strconv.Itoa(GenerateNumber(100000, 999999)) + strconv.Itoa(GenerateNumber(100000, 999999))
}

// GeneratePostalCode generates a random postal code
func GeneratePostalCode() string {
	return strconv.Itoa(GenerateNumber(1000000000, 9999999999))
}

// GenerateRtRw generates a random rt or rw
func GenerateRtRw() string {
	return fmt.Sprintf("%03d", GenerateNumber(0, 999))
}
