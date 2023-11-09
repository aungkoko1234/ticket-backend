package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func getErrorMsg(fe validator.FieldError) string {
    switch fe.Tag() {
        case "required":
            return "This field is required"
		case "email" :
			 return  "This field must be email"	
		case "min" : 
		     return "This field must be longer than " + fe.Param() + " characters"	
		case "max" : 
		     return "This field must not be longer than " + fe.Param() + " characters" 
        case "lte":
            return "Should be less than " + fe.Param()
        case "gte":
            return "Should be greater than " + fe.Param()
    }
    return "Unknown error"
}

func ParseError(err error) interface {} {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make(map[string]string)
		for _, fe := range ve {
			out[fe.Field()] = getErrorMsg(fe)
		}
		return out
	}
	return nil
} 