package arrays






func Sum(numbers []int)int{
	sum:=0
	for _,num:= range numbers{
		sum+=num
	}
	return sum
}

func SumAll(numbersToSum ...[]int)[]int{
	var sum []int
	for _,numbers:= range numbersToSum{
		sum = append(sum, Sum(numbers))
	}
	return sum
}
func SumAllTail(numbersToSum ...[]int)[]int{
	var sum []int
	for _,numbers:= range numbersToSum{
		if len(numbers)==0{
			sum=append(sum, 0)
		}else{
			tail:=numbers[1:]
		sum = append(sum, Sum(tail))
		}
		
	}
	return sum
}