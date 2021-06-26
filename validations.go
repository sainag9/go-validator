package validator

import (
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func isRequired(value interface{}) bool {

	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return value.(string) != ""
	default:
		return true
	}
	return true
}

func isEmpty(value interface{}) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return value.(string) == ""
	}
	return true
}

func isUUID(id string) bool {
	isuUID, _ := regexp.MatchString(UUIDV4Regex, id)
	return isuUID
}

func doesFileExists(fPath string) bool {
	_, err := os.Stat(fPath)
	return err == nil
}

func isValidEmail(email string) bool {
	isEmail, _ := regexp.MatchString(EmailRegex, email)
	return isEmail
}

func validateWithRegex(reg string, value string) error {

	rArr := strings.Split(reg, "-")
	rExp := ""
	if len(rArr) > 1 {
		rExp = rArr[1]
	}
	_, err := regexp.MatchString(rExp, value)
	return err

}

func isURL(uri string) bool {
	u, err := url.Parse(uri)
	return err == nil && u.Scheme != "" && u.Host != ""
}
