package reflection


func walk(x interface{},fn func(input string)){
	fn("I still can't believe people are not using Go")
}