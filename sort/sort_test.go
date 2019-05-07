package sort

import (
	"reflect"
	"testing"
)

// TestQuickSort - test quick sort functionality
func TestQuickSort(t *testing.T) {
	// input array
	input := []string{"5", "1", "3", "2", "4"}
	// expected output
	expected := []string{"1", "2", "3", "4", "5"}
	// performing quicksort on input array
	QuickSort(input, 0, len(input)-1)
	// compare input array after quicksort with expected output
	if reflect.DeepEqual(input, expected) == false {
		// when input array after sorting is not equal to expected output
		// throw error
		t.Errorf("%+v and %+v are not equal", input, expected)
	}
}

// TestMerge - test merge function
func TestMerge(t *testing.T) {
	// input arrays
	arr := []string{"1", "2", "3"}
	arr1 := []string{"4", "5"}
	// expected output
	expected := []string{"1", "2", "3", "4", "5"}
	// merge two input arrays
	actual := Merge(arr, arr1)
	// compare expected and actual output
	if reflect.DeepEqual(expected, actual) == false {
		// when expected and actual output are not same
		// throw error
		t.Errorf("expected output:%+v \nGot output:%+v \nMerge function failed", expected, actual)
	}
}
