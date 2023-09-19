package bubble_sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	list := []int{9, 4, 8, 1, 4, 5, 33}
	BubbleSort(list)
	fmt.Println(list)
	assert.Equal(t, list, []int{1, 4, 4, 5, 8, 9, 33})
}
