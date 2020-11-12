package module02

import "sort"

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	/* 1. Worst approach
	for i,_ := range list {
		i=i
		for j:=0;j<len(list)-1; j++ {
			if(list[j] > list[j+1]) {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}

	}
	2. Better avoids rechecking items in sorted position
	for i,_ := range list {
		for j:=0; j<len(list)-1-i;j++ {
			if (list[j] > list[j+1]) {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
    */
	// 3. Best version early exit when list sorted
	for i := 0; i < len(list); i++ {
		swapped := false
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

}

// BubbleSortString is a bubble sort for string slices. Notice how it is EXACTLY
// the same as our int version aside from the input?
func BubbleSortString(list []string) {
	for i := 0; i < len(list); i++ {
		swapped := false
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {

	less := func(a, b Person) bool {
		if a.Age != b.Age {
			return a.Age < b.Age
		}
		if a.LastName != b.LastName {
			return a.LastName < b.LastName
		}
		return a.FirstName < b.FirstName
	}
	// Option 1 - write this by hand
	for sweepNum := 0; sweepNum < len(people); sweepNum++ {
		swapped := false
		for i := 0; i < len(people)-1-sweepNum; i++ {
			if less(people[i+1], people[i]) {
				people[i], people[i+1] = people[i+1], people[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	// Option 2 - implement the sort.Interface and use our sort.Interface impl.
	// See person.go for a bit more info/code showing how this People type works.
	// BubbleSort(People(people))
}

// BubbleSort uses the standard library's sort.Interface to sort
// This one is likely going to be a little more challenging to implement.
func BubbleSort(list sort.Interface) {
	for s := 0; s < list.Len(); s++ {
		swapped := false
		for i := 0; i < list.Len()-1-s; i++ {
			if list.Less(i+1, i) {
				list.Swap(i, i+1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
