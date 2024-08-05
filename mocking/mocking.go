package mocking

import (
	"fmt"
	"io"
)


func CountDown(writer io.Writer){
	fmt.Fprintf(writer,"3")
}