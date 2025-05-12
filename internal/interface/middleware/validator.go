package middleware

import (
	"errors"
	"fmt"
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
	return func(context *gin.Context) {
		var requestData T
		if error := context.ShouldBind(&requestData); error != nil {
			validationErrors, parseResult := error.(validator.ValidationErrors)
			if parseResult {
				for _, error := range validationErrors {
					switch error.Tag() {
					case "required":
						context.Error(errors.New(fmt.Sprintf("Vui lòng cung cấp %s", error.Field())))
					case "email":
						context.Error(errors.New(fmt.Sprintf("Địa chỉ email %v không hợp lệ", error.Value())))
					case "min":
						context.Error(errors.New(fmt.Sprintf("%s phải có ít nhất %s ký tự", error.Field(), error.Param())))
					case "max":
						context.Error(errors.New(fmt.Sprintf("%s không được vượt quá %s ký tự", error.Field(), error.Param())))
					case "len":
						context.Error(errors.New(fmt.Sprintf("%s phải có độ dài chính xác là %s ký tự", error.Field(), error.Param())))
					case "gt":
						context.Error(errors.New(fmt.Sprintf("%s phải lớn hơn %s", error.Field(), error.Param())))
					case "gte":
						context.Error(errors.New(fmt.Sprintf("%s phải lớn hơn hoặc bằng %s", error.Field(), error.Param())))
					case "lt":
						context.Error(errors.New(fmt.Sprintf("%s phải nhỏ hơn %s", error.Field(), error.Param())))
					case "lte":
						context.Error(errors.New(fmt.Sprintf("%s phải nhỏ hơn hoặc bằng %s", error.Field(), error.Param())))
					case "alphanum":
						context.Error(errors.New(fmt.Sprintf("%s chỉ được chứa các ký tự chữ và số", error.Field())))
					case "url":
						context.Error(errors.New(fmt.Sprintf("%s phải là một URL hợp lệ", error.Field())))
					case "uuid":
						context.Error(errors.New(fmt.Sprintf("%s phải là một UUID hợp lệ", error.Field())))
					case "ip":
						context.Error(errors.New(fmt.Sprintf("%s phải là một địa chỉ IP hợp lệ", error.Field())))
					case "ipv4":
						context.Error(errors.New(fmt.Sprintf("%s phải là một địa chỉ IPv4 hợp lệ", error.Field())))
					case "ipv6":
						context.Error(errors.New(fmt.Sprintf("%s phải là một địa chỉ IPv6 hợp lệ", error.Field())))
					case "numeric":
						context.Error(errors.New(fmt.Sprintf("%s phải là một số", error.Field())))
					case "contains":
						context.Error(errors.New(fmt.Sprintf("%s phải chứa chuỗi con '%s'", error.Field(), error.Param())))
					case "startswith":
						context.Error(errors.New(fmt.Sprintf("%s phải bắt đầu bằng '%s'", error.Field(), error.Param())))
					case "endswith":
						context.Error(errors.New(fmt.Sprintf("%s phải kết thúc bằng '%s'", error.Field(), error.Param())))
					case "vn_phone":
						context.Error(errors.New(fmt.Sprintf("%s không phải là số điện thoại hợp lệ", error.Value())))
					default:
						context.Error(errors.New("Lỗi không không xác định khi validation"))
					}
				}
				context.AbortWithStatus(http.StatusBadRequest)
				return
			}
			context.Error(errors.New("Vui lòng cung cấp đầy đủ thông tin"))
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		context.Set("request_data", &requestData)
		context.Next()
	}
}
