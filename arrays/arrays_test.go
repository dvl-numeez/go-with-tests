package arrays

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T){
	got:=SumAll([]int{1,2},[]int{0,9})
	want:=[]int{3,9}
	if  ! reflect.DeepEqual(got,want){
		t.Errorf("got : %v want : %v",got,want)
	}
}

func TestSumAllTails(t *testing.T){
	got:=SumAllTail([]int{1,2},[]int{0,9})
	expected:=[]int{2,9}
	if  ! reflect.DeepEqual(got,expected){
		t.Errorf("got : %v want : %v",got,expected)
	}

	t.Run("entering empty slices also",func(t *testing.T){
		got:=SumAllTail([]int{},[]int{0,9})
	expected:=[]int{0,9}
	if  ! reflect.DeepEqual(got,expected){
		t.Errorf("got : %v want : %v",got,expected)
	}
	})

}