package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reg_exp = regexp.MustCompile(`[^0-9]+`)

func IsValidPhonenumber(phoneNumber string) (string, bool) {
	if string(phoneNumber[0:1]) == "+" && string(phoneNumber[1:2]) != "1" {
		return phoneNumber, false
	}

	phoneNumber = reg_exp.ReplaceAllString(phoneNumber, "")

	switch {
	case len(phoneNumber) < 10:
		return phoneNumber, false

	case len(phoneNumber) == 10:
		for i, dig := range strings.Split(phoneNumber, "") {
			num, _ := strconv.Atoi(dig)

			if (i == 0 || i == 3) && num < 2 {
				return phoneNumber, false
			}
		}
	case len(phoneNumber) == 11:
		if phoneNumber[0:1] != "1" {
			return phoneNumber, false
		} else {
			return IsValidPhonenumber(phoneNumber[1:])
		}

	case len(phoneNumber) > 11:
		return phoneNumber, false
	}

	return phoneNumber, true
}

func Number(phoneNumber string) (string, error) {
	phoneNumber = reg_exp.ReplaceAllString(phoneNumber, "")
	phoneNumber, isvalid := IsValidPhonenumber(phoneNumber)

	if !isvalid {
		return phoneNumber, errors.New("Invalid phonenumber")
	}

	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	//panic("Please implement the AreaCode function")
	number, is_valid := IsValidPhonenumber(phoneNumber)
	if is_valid {
		return number[0:3], nil
	} else {
		return phoneNumber, errors.New("invalid phonenumber " + number)
	}
}

func Format(phoneNumber string) (string, error) {
	//panic("Please implement the Format function")
	number, is_valid := IsValidPhonenumber(phoneNumber)
	if is_valid {
		return fmt.Sprintf("(%v) %v-%v", number[:3], number[3:6], number[6:]), nil
	} else {
		return phoneNumber, errors.New("invalid phonenumber")
	}
}
