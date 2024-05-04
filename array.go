package godash

import "fmt"

// Chunk 将给定的切片 arr 分割成大小为 size 的小切片。
// 如果 size 小于 1 或 arr 为 nil，将返回错误。
// T 可以是任何类型，但函数调用者需要确保 T 的类型支持基本的赋值操作。
func Chunk[T any](arr []T, size int) (result [][]T, err error) {
	// 检查 arr 是否为 nil
	if arr == nil {
		err = fmt.Errorf("the input array cannot be nil")
		return
	}

	// 检查 size 是否小于 1
	if size < 1 {
		err = fmt.Errorf("the chunk size of an array cannot be smaller than %d", 1)
		return
	}

	// 获取原切片的长度
	l := len(arr)

	// 计算需要分割成多少个小切片
	s := (l + size - 1) / size
	result = make([][]T, s)

	for i := range result {
		start := i * size
		end := start + size
		if end > l {
			end = l
		}
		// 为了更好地处理潜在的边界条件和类型安全性，
		// 明确地使用切片操作来创建结果切片的部分。
		result[i] = arr[start:end:end]
	}

	return
}

// Compact 创建一个新数组，包含原数组中所有的非假值元素。例如false, null,0, "", undefined, 和 NaN 都是被认为是“假值”。
// 该函数支持泛型，可以适用于任何类型[T any]的数组。
//
// 参数：
// arr []T：一个泛型数组，其中的元素类型为T。
//
// 返回值：
// []T：一个新的数组，其中不包含任何空值。
func Compact[T any](arr []T) []T {
	// 创建一个空的新数组n，用于存储非空元素。
	n := make([]T, 0)
	// 遍历原数组arr，检查每个元素是否为空。
	for _, v := range arr {
		// 如果元素不为空，则将其添加到新数组n中。
		if !IsEmpty(v) {
			n = append(n, v)
		}
	}
	// 返回新数组n。
	return n
}
