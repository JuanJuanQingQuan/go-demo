package main

import (
	"errors"
	"fmt"
)

func returnError() error {
	var err error
	defer func() {
		err = nil
	}()
	err = errors.New("test error")
	return err
}

func returnError2() (err error) {
	defer func() {
		err = nil
	}()

	err = errors.New("test error")
	return err
}

func returnErrorPtr() *error {
	var err error
	defer func() {
		err = nil
	}()

	err = errors.New("test error")
	return &err
}

func ReturnDeferErrorMain() {
	// test error
	fmt.Printf("return error by return: %v\n", returnError())
	// nil
	fmt.Printf("return error by err return: %v\n", returnError2())
	// nil
	fmt.Printf("return error by ptr: %v\n", *returnErrorPtr())
}
