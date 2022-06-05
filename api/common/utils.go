package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()

	return res
}

func NewValidatorError(err error) CommonError {
	res := CommonError{}
	errs := err.(validator.ValidationErrors)

	for _, v := range errs {
		// can translate each error one at a time.
		//fmt.Println("gg",v.NameNamespace)
		if v.Param() != "" {
			res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag)
		}

	}
	return res
}

// https://github.com/gothinkster/golang-gin-realworld-example-app/blob/e4174a9d6450e9e993d9b89a0dace58b7994c045/common/utils.go#L80
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())

	return c.ShouldBindWith(obj, b)
}
