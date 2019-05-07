package sort

// QuickSort - implements sorting using quick sort algorithm
func QuickSort(arr []string, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func partition(arr []string, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Merge - merges left and right slice into newly created slice
func Merge(left, right []string) []string {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]string, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
