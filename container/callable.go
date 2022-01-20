package container

import (
	"errors"
	"github.com/goal-web/contracts"
	"reflect"
)

var (
	FuncTypeError = errors.New("参数必须是一个函数")
)

type magicalFunc struct {
	in        int
	out       int
	value     reflect.Value
	arguments []reflect.Type
	returns   []reflect.Type
}

func NewMagicalFunc(fn interface{}) contracts.MagicalFunc {
	var (
		argValue = reflect.ValueOf(fn)
		argType  = reflect.TypeOf(fn)
	)

	if argValue.Kind() != reflect.Func {
		panic(FuncTypeError)
	}

	var (
		arguments    = make([]reflect.Type, 0)
		returns      = make([]reflect.Type, 0)
		argumentsLen = argType.NumIn()
		returnsLen   = argType.NumOut()
	)

	for argIndex := 0; argIndex < argumentsLen; argIndex++ {
		arguments = append(arguments, argType.In(argIndex))
	}

	for outIndex := 0; outIndex < returnsLen; outIndex++ {
		returns = append(returns, argType.Out(outIndex))
	}

	return &magicalFunc{
		in:        argumentsLen,
		out:       returnsLen,
		value:     argValue,
		arguments: arguments,
		returns:   returns,
	}
}

func (this *magicalFunc) Call(in []reflect.Value) []reflect.Value {
	return this.value.Call(in)
}

func (this *magicalFunc) Arguments() []reflect.Type {
	return this.arguments
}

func (this *magicalFunc) Returns() []reflect.Type {
	return this.returns
}

func (this *magicalFunc) NumOut() int {
	return this.out
}

func (this *magicalFunc) NumIn() int {
	return this.in
}
