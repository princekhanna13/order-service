package rest

import (
	"reflect"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

const (
	dash       = "-"
	blank      = ""
	jsonVal    = "json"
	commaSplit = ","
)

// Validator - interface to expose functions to validate structs which use validator package
type Validator interface {
	ValidateStruct(object interface{}, logger *logrus.Entry) error
}

// DefaultValidator - implementation of interface Validator
type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

// lazyInit - using sync package's once functionality, lazy initialize the validator
func (v *DefaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.RegisterTagNameFunc(tagNameFunction)
	})
}

// tagNameFunction - Registering custom tag name function
func tagNameFunction(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get(jsonVal), commaSplit, 2)[0]
	if name == dash {
		return blank
	}
	return name
}

//ValidateStruct - validate the given structure and return validator errors
func (v *DefaultValidator) ValidateStruct(object interface{}, logger *logrus.Entry) error {
	v.lazyInit()
	return v.validate.Struct(object)
}
