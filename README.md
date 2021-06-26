# go-validator

*  validator package implements struct validation wherein user can provide tags to validate struct fields.
*  multiple validations for the same fields can also be done.


**Installation:**

`go get github.com/sainag9/go-validator`


**Example struct:**

```
type Details struct {
	UUID     string  `json:"uuid" validate:"uuid"`
	Name     string  `json:"name" validate:"required"`
	Email    string  `json:"email" validate:"required, email"`
	FilePath string  `json:"filepath" validate:"fileExists"`
	Age      float64 `json:"age" validate:"required"`
	WebSite  string  `json:"website" validate:"required, url"`
	Address  struct {
		HNO    string `json:"hno" validate:"required"`
		Street struct {
			Name string `json:"name" validate:"required"`
		} `json:"street"`
	} `json:"address"`
}
```



create a new validator and pass the struct for validation.


```
vd := validator.NewValidator()
vErr := vd.ValidateStruct(struct_to_be_validated)
```

if there are any validation errors validator returns a list of errors in the below format.

`type ValidationError struct {
	FieldName    string
	ErrorMessage string
}`

**input struct:**

```
dtls := Details{
		Name:     "name",
		Email:    "k.sai@gmail.com",
		FilePath: "somepath",
		UUID:     "the uuid",
		WebSite: "http://localhost:8080/api",
	}
	dtls.Address.HNO = "2-3-83/1"
```


**output for above details:** 

`[{.uuid the uuid is not a valid uuid} {.filepath somepath does not exists} {.website http is not valid url} {.address.street.name  cannot be empty}]`

**Available Tags:**

*  required
*  regex-{regular_expression}
*  email
*  uuid (v4 regex is used for validation)
*  fileExists (validates if file is present)
*  url
*  omitempty (not validated if empty value is provided, if present proceeds for further validation)
*  