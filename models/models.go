package models

type StructMetaData struct {
	Value          interface{}
	JSONTag        string
	ValidationType string
	FieldName      string
}

type ValidationError struct {
	FieldName    string
	ErrorMessage string
}
