package main
import (
	"fmt"
	//"strconv"
)

func main() {
	arr := []int{4,1,2,1,2, 6, 6, 8, 7, 7, 8}
	val := singleNumber(arr)
	fmt.Println(val)
}

func singleNumber(nums []int) int {

	// map, key, val, {}
	mymap := map[int]int{}
	//res := []int{}	
	for _, val :=  range nums {
		/*
		keyStr := strconv.Itoa(key)
		valStr := strconv.Itoa(val);
		fmt.Println(keyStr + ": " + valStr)
		*/

		// just check val
		if(mymap[val] == 1) {
			mymap[val] = mymap[val] + 1
		} else {
			mymap[val] = 1;
		}	
	}

	res := -1
	for key, val := range mymap {
		if(val == 1) {
			res = key 
			break	
		}
	}

	return  res 
}
