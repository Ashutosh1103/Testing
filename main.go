//Package ashu gives the sum of any number of integers passed to it.
package ashu

//Sum function operation
func Sum(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}
