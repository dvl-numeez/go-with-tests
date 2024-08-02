package arrays

import "testing"

func TestSum(t *testing.T){
	nums:=[5]int{1,2,3,4,5}
	actual:=Sum(nums)
	expected:=15
	if actual!=expected{
		t.Errorf("Actual : %d Expected : %d",actual,expected)
	}
}