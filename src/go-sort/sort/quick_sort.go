package sort
import "fmt"

func qsortInts(unsortedArray []int) []int {
	if len(unsortedArray) < 2 { return unsortedArray }

	left, right := 0, len(unsortedArray) - 1

	// Pick pivot point in the middle
	pivotIndex := left + right / 2

	// Swap the pivot with the item at left index for partitioning process
	fmt.Printf("\nSwapping left [%d] & pivot [%d]\n", pivotIndex, left)
	unsortedArray[pivotIndex], unsortedArray[left] = unsortedArray[left], unsortedArray[pivotIndex]
	fmt.Printf("%v\n\n", unsortedArray)
	// Iterate over each index in array
	for partitionIndex := range unsortedArray {
		// If partition index value is less than right index value, swap left and partition index
		fmt.Printf("%v partitionIndex [%d], left index [%d]\n", unsortedArray, partitionIndex, left)
		if unsortedArray[partitionIndex] < unsortedArray[right] {
			unsortedArray[partitionIndex], unsortedArray[left] = unsortedArray[left], unsortedArray[partitionIndex]
			left++
		}
	}

	// Swap the pivot with the last smaller element
	fmt.Printf("\n\nSwapping left [%d] & right [%d]\n", left, right)
	unsortedArray[left], unsortedArray[right] = unsortedArray[right], unsortedArray[left]
	fmt.Printf("%v\n\n", unsortedArray)

	// Recursively sort with left slice
	qsortInts(unsortedArray[:left])
	// Recursively sort with right slice
	qsortInts(unsortedArray[left + 1:])
	return unsortedArray
}