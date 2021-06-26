package validator

import (
	"github.com/sainag9/go-validator/models"
	"strings"
)

const (
	Required    = "required"
	Regex       = "regex"
	Email       = "email"
	UUID        = "uuid"
	FileExists  = "fileExists"
	EmailRegex  = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])"
	UUIDV4Regex = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	JSONTag     = "json"
	ValidateTag = "validate"
	URL         = "url"
	OmitEmpty   = "omitempty"
)

// formErrorMessage forms the error message based on inputs provided
func formErrorMessage(vType, fName string, value interface{}) string {

	switch vType {
	case Required:
		return strings.Trim(fName, ".") + " cannot be empty"
	case URL:
		return value.(string) + " is not valid url"
	}
	return ""
}

// getErrorMessage creates ValidationError
func getErrorMessage(fieldName, message string) models.ValidationError {
	return models.ValidationError{
		FieldName:    fieldName,
		ErrorMessage: message,
	}
}
