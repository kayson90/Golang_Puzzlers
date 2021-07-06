package main

//实际上就是hash表，需要判断是否等于来操作
func main() {
	// 示例1。
	// key是切片
	//var badMap1 = map[[]int]int{} // 这里会引发编译错误。
	//_ = badMap1

	// 示例2。
	// key是切片
	//var badMap2 = map[interface{}]int{
	//	"1":      1,
	//	[]int{2}: 2, // 这里会引发panic。
	//	3:        3,
	//}
	//_ = badMap2

	// 示例3。
	// 数组的元素不能为切片
	//var badMap3 map[[1][]string]int // 这里会引发编译错误。
	//_ = badMap3

	// 示例4。
	// 结构体的元素不能为切片
	//type BadKey1 struct {
	//	slice []string
	//}
	//var badMap4 map[BadKey1]int // 这里会引发编译错误。
	//_ = badMap4

	// 示例5。
	// 多维数组的元素为切片，不管隐藏多深
	//var badMap5 map[[1][2][3][]string]int // 这里会引发编译错误。
	//_ = badMap5

	// 示例6。
	//type BadKey2Field1 struct {
	//	slice []string
	//}
	//type BadKey2 struct {
	//	field BadKey2Field1
	//}
	//var badMap6 map[BadKey2]int // 这里会引发编译错误。
	//_ = badMap6

}
