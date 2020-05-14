/*
ECS 140a Hw2
Faustine Yiu 913062973
*/
package min

// Min returns the minimum value in the arr,
// and 0 if arr is nil.
func Min(arr []int) int {
	//if arr is empty, return 0
	if len(arr) == 0 {
		return 0
	}

	min := arr[0] //set first value as the smallest for now

	//loop through arr and check the next min
	for i := 0; i < len(arr); i++ {
	  if arr[i] < min {
	    min = arr[i] //swap min to lowest value
    }
	}
	return min //min will be the smallest after the loop, so return it
}
