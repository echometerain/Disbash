package utils

import (
	"fmt"
)

func Eh(t interface{}, err error) interface{} { // error handler
	if err != nil {
		fmt.Print(err)
	}
	return t
}
