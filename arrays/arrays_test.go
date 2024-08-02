package arrays

import "testing"

func TestSum(t *testing.T){
	t.Run("Collection of 5 numbers",func(t *testing.T){
		nums:=[]int{1,2,3,4,5}
	actual:=Sum(nums)
	expected:=15
	if actual!=expected{
		t.Errorf("Actual : %d Expected : %d",actual,expected)
	}
	})
	t.Run("Collection of number of any size",func(t *testing.T){
		nums:=[]int{1,2,3}
		actual:=Sum(nums)
		expected:=6
		if actual!=expected{
			t.Errorf("Actual : %d Expected : %d",actual,expected)
		}
		})
	
	
}