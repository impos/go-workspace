package wspace

import "github.com/pkg/errors"

type FooBar struct {
	Foo int
	Bar int
}

func FooBarFn(foo, bar int) FooBar {
	return FooBar{
		Foo: foo,
		Bar: bar,
	}
}

func Error(msg string) error {
	return errors.New(msg)
}
