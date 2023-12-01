package main

import "fmt"

func main() {
    fmt.Println("Should answer [1,2,3]: ", quick_sort([]int{3,2,1}))
    fmt.Println("Should answer [1,1,2,2,3,3]: ", quick_sort([]int{3,2,3,1,1,2}))
    
    fmt.Println("Should answer 0: ", binary_search([]int{1,2,3,4}, 1))
    fmt.Println("Should answer 1: ", binary_search([]int{1,2,3,4}, 2))
    fmt.Println("Should answer -1: ", binary_search([]int{1,2,3,4}, 10))
    
    input := []int{3,1,2,5,4}
    fmt.Println("Should answer true: ", single_pass(input, 4))
    fmt.Println("Should answer false: ", single_pass(input, 100))
    fmt.Println("Should answer false: ", single_pass(input, 10))
}

func single_pass(input []int, k int) bool {
    sorted_input := quick_sort(input)
    return binary_search(sorted_input, k) >= 0
}

func quick_sort(array []int) []int {
    var low, same, high []int

    if len(array) == 0 {
        return array
    }

    pivot := array[0]

    for _, item := range array {
        if item < pivot {
            low = append(low, item)
        } else if item == pivot {
            same = append(same, item)
        } else if item > pivot {
            high = append(high, item)
        }
    }

    return append(append(quick_sort(low), same...), quick_sort(high)...)
}

func binary_search(array []int, k int) int {
    var low, mid, high int
    high = len(array) - 1

    for low <= high {
        mid = (high + low) / 2

        if array[mid] < k {
            low = mid + 1
        } else if array[mid] > k {
            high = mid - 1
        } else {
            return mid
        }
    }

    return -1
}
