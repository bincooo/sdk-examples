package model

import (
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	*A
}

// @Ioc(name="model.A", proxy="px.Echo")
func NewA() *A {
	return &A{}
}

// @Ioc(alias="model.B", qualifier="[0]:model.A")
func NewB(a *A) *B {
	return &B{a}
}

func (a A) Echo(name string) error {
	fmt.Printf("%d, A.Echo(%s)\n", a.Num, name)
	return nil
}

func (B) Echo() {
	fmt.Println("B.Echo()")
}
