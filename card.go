package main

import "fmt"

type Card interface {
	comparable
	fmt.Stringer
	Name() string
}
