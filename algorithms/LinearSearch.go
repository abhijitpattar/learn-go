package algorithms

import (
	"fmt"
	"time"
)

func LinearSearch(intArr []int, SearchKey int) int {
	retIndex := -1 // return -1 if no value is found

	for i := 0; i < len(intArr); i++ {
		if intArr[i] == SearchKey {
			return i
		}
	}
	return retIndex
}

func LinearSearchRecusive(intArr []int, leftIndex int, rightIndex int, SearchKey int) int {
	if leftIndex > rightIndex {
		return -1
	}
	if (intArr[leftIndex]) == SearchKey {
		return leftIndex
	}
	return LinearSearchRecusive(intArr, leftIndex+1, rightIndex, SearchKey)
}

func LinearSearch_2(intArr []int, SearchKey int) int {
	retIndex := -1 // return -1 if no value is found

	for i, n := 0, len(intArr)-1; i < n; i++ {
		if intArr[i] == SearchKey {
			return i
		}
		if intArr[n] == SearchKey {
			return n
		}
		n--
	}
	return retIndex
}

func LinearSearchRecusive_2(intArr []int, leftIndex int, rightIndex int, SearchKey int) int {
	if leftIndex > rightIndex {
		return -1
	}
	if (intArr[leftIndex]) == SearchKey {
		return leftIndex
	}
	if (intArr[rightIndex-1]) == SearchKey {
		return rightIndex - 1
	}
	return LinearSearchRecusive(intArr, leftIndex+1, rightIndex-1, SearchKey)
}

func CallLinearSearch() {
	aArr := []int{5, 3, 7, 9, 4, 8, 2, 89, 0}
	bArr := []int{554, 987, 578654, 68355, 654, 871, 687, 545687, 3216878, 5, 875, 754, 8754, 74, 87, 655}
	cArr := []int{554, 987, 54, 9871, 654, 87, 64, 21, 1, 032, 0, 65, 3, 66}

	now := time.Now().Nanosecond()
	fmt.Println(LinearSearch(aArr, 5))      // 0
	fmt.Println(LinearSearch(aArr, 2))      // 6
	fmt.Println(LinearSearch(bArr, 5))      // 9
	fmt.Println(LinearSearch(bArr, 545687)) // 7
	fmt.Println(LinearSearch(cArr, 1))      // 8
	fmt.Println(LinearSearch(cArr, 032))    // 9
	fmt.Println("end LinearSearch ............", time.Now().Nanosecond()-now)

	now = time.Now().Nanosecond()
	fmt.Println(LinearSearchRecusive(aArr, 0, len(aArr), 5))      // 0
	fmt.Println(LinearSearchRecusive(aArr, 0, len(aArr), 2))      // 6
	fmt.Println(LinearSearchRecusive(bArr, 0, len(bArr), 5))      // 9
	fmt.Println(LinearSearchRecusive(bArr, 0, len(bArr), 545687)) // 7
	fmt.Println(LinearSearchRecusive(cArr, 0, len(cArr), 1))      // 8
	fmt.Println(LinearSearchRecusive(cArr, 0, len(cArr), 032))    // 9
	fmt.Println("end LinearSearchRecusive ............", time.Now().Nanosecond()-now)

	now = time.Now().Nanosecond()
	fmt.Println(LinearSearch_2(aArr, 5))      // 0
	fmt.Println(LinearSearch_2(aArr, 2))      // 6
	fmt.Println(LinearSearch_2(bArr, 5))      // 9
	fmt.Println(LinearSearch_2(bArr, 545687)) // 7
	fmt.Println(LinearSearch_2(cArr, 1))      // 8
	fmt.Println(LinearSearch_2(cArr, 032))    // 9
	fmt.Println("end LinearSearch_2 ............", time.Now().Nanosecond()-now)

	now = time.Now().Nanosecond()
	fmt.Println(LinearSearchRecusive_2(aArr, 0, len(aArr), 5))      // 0
	fmt.Println(LinearSearchRecusive_2(aArr, 0, len(aArr), 2))      // 6
	fmt.Println(LinearSearchRecusive_2(bArr, 0, len(bArr), 5))      // 9
	fmt.Println(LinearSearchRecusive_2(bArr, 0, len(bArr), 545687)) // 7
	fmt.Println(LinearSearchRecusive_2(cArr, 0, len(cArr), 1))      // 8
	fmt.Println(LinearSearchRecusive_2(cArr, 0, len(cArr), 032))    // 9
	fmt.Println("End LinearSearchRecusive_2 ............", time.Now().Nanosecond()-now)
}
