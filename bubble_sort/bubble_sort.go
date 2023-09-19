package bubble_sort

import "fmt"

func BubbleSort(list []int) {
	swapCount := len(list) - 1
	for i := 0; i < swapCount; i++ {
		for j := 0; j < swapCount-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
		fmt.Println(list)
	}

}
