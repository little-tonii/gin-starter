package middleware

import (
	"errors"
	"fmt"
	"gin-starter/internal/shared/constant"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindingValidator[T any]() gin.HandlerFunc {
	var t T
	if tType := reflect.TypeOf(t); tType.Kind() != reflect.Struct {
		panic("BindingValidator chỉ chấp nhận kiểu struct")
	}
	return func(c *gin.Context) {
		var requestData T
		if error := c.ShouldBind(&requestData); error != nil {
			validationErrors, parseResult := error.(validator.ValidationErrors)
			if parseResult {
				for _, error := range validationErrors {
					switch error.Tag() {
					case "required":
						c.Error(errors.New(fmt.Sprintf("Vui lòng cung cấp %s", error.Field())))
					case "email":
						c.Error(errors.New(fmt.Sprintf("Địa chỉ email %v không hợp lệ", error.Value())))
					case "min":
						c.Error(errors.New(fmt.Sprintf("%s phải có ít nhất %s ký tự", error.Field(), error.Param())))
					case "max":
						c.Error(errors.New(fmt.Sprintf("%s không được vượt quá %s ký tự", error.Field(), error.Param())))
					case "len":
						c.Error(errors.New(fmt.Sprintf("%s phải có độ dài chính xác là %s ký tự", error.Field(), error.Param())))
					case "gt":
						c.Error(errors.New(fmt.Sprintf("%s phải lớn hơn %s", error.Field(), error.Param())))
					case "gte":
						c.Error(errors.New(fmt.Sprintf("%s phải lớn hơn hoặc bằng %s", error.Field(), error.Param())))
					case "lt":
						c.Error(errors.New(fmt.Sprintf("%s phải nhỏ hơn %s", error.Field(), error.Param())))
					case "lte":
						c.Error(errors.New(fmt.Sprintf("%s phải nhỏ hơn hoặc bằng %s", error.Field(), error.Param())))
					case "alphanum":
						c.Error(errors.New(fmt.Sprintf("%s chỉ được chứa các ký tự chữ và số", error.Field())))
					case "url":
						c.Error(errors.New(fmt.Sprintf("%s phải là một URL hợp lệ", error.Field())))
					case "uuid":
						c.Error(errors.New(fmt.Sprintf("%s phải là một UUID hợp lệ", error.Field())))
					case "ip":
						c.Error(errors.New(fmt.Sprintf("%s phải là một địa chỉ IP hợp lệ", error.Field())))
					case "ipv4":
						c.Error(errors.New(fmt.Sprintf("%s phải là một địa chỉ IPv4 hợp lệ", error.Field())))
					case "ipv6":
						c.Error(errors.New(fmt.Sprintf("%s phải là một địa chỉ IPv6 hợp lệ", error.Field())))
					case "numeric":
						c.Error(errors.New(fmt.Sprintf("%s phải là một số", error.Field())))
					case "contains":
						c.Error(errors.New(fmt.Sprintf("%s phải chứa chuỗi con '%s'", error.Field(), error.Param())))
					case "startswith":
						c.Error(errors.New(fmt.Sprintf("%s phải bắt đầu bằng '%s'", error.Field(), error.Param())))
					case "endswith":
						c.Error(errors.New(fmt.Sprintf("%s phải kết thúc bằng '%s'", error.Field(), error.Param())))
					default:
						c.Error(errors.New("Lỗi không không xác định khi validation"))
					}
				}
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			c.Error(errors.New("Vui lòng cung cấp đầy đủ thông tin"))
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Set(constant.ContextKey.REQUEST_DATA, &requestData)
		c.Next()
	}
}
