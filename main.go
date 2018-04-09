package main

func Swap(list []int) {
	tmp := list[0]
	list[0] = list[1]
	list[1] = tmp
}

func MergeSort(list []int) {
	if len(list) == 1 {
		return
	}

	if len(list) == 2 {
		if list[0] > list[1] {
			Swap(list)
		}
		return
	}

	var (
		splitLen  = len(list) / 2
		leftSide  = list[:splitLen]
		rightSide = list[splitLen:]
	)

	MergeSort(leftSide)
	MergeSort(rightSide)

	merged := make([]int, len(list))
	for i, j, k := 0, 0, 0; k < len(list); k++ {
		if j >= len(rightSide) {
			merged[k] = leftSide[i]
			i++
		} else if i >= len(leftSide) || leftSide[i] >= rightSide[j] {
			merged[k] = rightSide[j]
			j++
		} else {
			merged[k] = leftSide[i]
			i++
		}
	}
	for i := range list {
		list[i] = merged[i]
	}
}

func main() {
}
