package common

import (
	"github.com/kataras/iris/v12/x/errors"
	"reflect"
)

// 类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	// todo

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}
