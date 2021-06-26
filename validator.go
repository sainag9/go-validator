package validator

import (
	"fmt"
	"github.com/sainag9/go-validator/models"
	"reflect"
	"strings"
)

type Validator interface {
	ValidateStruct(s interface{}) []models.ValidationError
}

type Validate struct {
}

func NewValidator() Validator {
	return Validate{}
}

// ValidateStruct validates struct fields based on validate tag data
func (v Validate) ValidateStruct(s interface{}) []models.ValidationError {
	t := reflect.TypeOf(s)
	vl := reflect.ValueOf(s)

	var sMetaData []models.StructMetaData
	// initial(root) tag value is empty
	metaData := getStructInfo(t, vl, sMetaData, "")
	//fmt.Println(metaData)
	var vErr []models.ValidationError
	for _, m := range metaData {
		tagArray := strings.Split(m.ValidationType, ",")
		er := validate(tagArray, m)
		if len(er) != 0 {
			vErr = append(vErr, er...)
		}
	}

	return vErr
}

// getStructInfo uses reflect to get struct tags and form metadata
func getStructInfo(t reflect.Type, v reflect.Value, sArray []models.StructMetaData, tag string) []models.StructMetaData {

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i)
		//fmt.Println("&&&&&&&&&&&&", tag, val, field)
		if field.Type.Kind() == reflect.Struct {
			//tag = tag + "." + field.Tag.Get(JSONTag)
			sArray = getStructInfo(field.Type, val, sArray, tag+"."+field.Tag.Get(JSONTag))
			fmt.Println("***********", sArray, tag)
		} else if field.Type.Kind() == reflect.Array || field.Type.Kind() == reflect.Slice {
			for j := 0; j < val.Len(); j++ {
				vs := val.Convert(field.Type)
				ts := vs.Index(j).Type()
				//tag = tag + "." + field.Tag.Get(JSONTag)
				//fmt.Println("$$$$$$", ts, vs, tag)
				sArray = getStructInfo(ts, vs.Index(j), sArray, tag+"."+field.Tag.Get(JSONTag))
			}
		} else {
			column := field.Tag.Get(ValidateTag)
			jTag := field.Tag.Get(JSONTag)
			row := val.Interface()
			sArray = append(sArray, models.StructMetaData{
				Value:          row,
				JSONTag:        tag + "." + jTag,
				ValidationType: column,
				FieldName:      jTag,
			})
		}
	}
	//fmt.Println("%%%%%%%%%%")
	return sArray
}

func validate(tags []string, md models.StructMetaData) []models.ValidationError {

	var vErr []models.ValidationError

tagLoop:
	for _, v := range tags {
		v = strings.Trim(v, " ")
		switch {
		case v == Required:
			if !isRequired(md.Value) {
				vErr = append(vErr, getErrorMessage(md.JSONTag, formErrorMessage(v, md.FieldName, md.Value)))
			}
		case strings.Contains(v, Regex):
			if err := validateWithRegex(md.ValidationType, md.Value.(string)); err != nil {
				vErr = append(vErr, getErrorMessage(md.JSONTag, err.Error()))
			}
		case v == Email:
			if !isValidEmail(md.Value.(string)) {
				vErr = append(vErr, getErrorMessage(md.JSONTag, md.Value.(string)+InvalidEmail))
			}
		case v == UUID:
			if !isUUID(md.Value.(string)) {
				vErr = append(vErr, getErrorMessage(md.JSONTag, md.Value.(string)+InvalidUUID))
			}
		case v == FileExists:
			if !doesFileExists(md.Value.(string)) {
				vErr = append(vErr, getErrorMessage(md.JSONTag, md.Value.(string)+FileDoesNotExists))
			}
		case v == URL:
			if !isURL(md.Value.(string)) {
				vErr = append(vErr, getErrorMessage(md.JSONTag, formErrorMessage(v, md.FieldName, md.Value)))
			}
		case v == OmitEmpty:
			if isEmpty(md.Value) {
				break tagLoop
			}
		default:

		}
	}
	return vErr
}
