package core

import "go.uber.org/dig"

var container = dig.New()

func ProvideForDI(constructor interface{}, opts ...dig.ProvideOption) error {
	return container.Provide(constructor, opts...)
}

func DI(function interface{}, opts ...dig.InvokeOption) error {
	return container.Invoke(function, opts...)
}
