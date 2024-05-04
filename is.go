package godash

import (
	"math"
	"reflect"
)

// IsPtr 函数用于判断传入的 interface{} 类型的值是否为指针。
//
// 参数:
// v - 一个 interface{} 类型的值，代表需要检查的目标。
//
// 返回值:
// 返回一个 bool 类型的值，如果 v 是指针，则返回 true；否则返回 false。
func IsPtr(v interface{}) bool {
	// 检查值是否为 nil，为 nil 则直接返回 false
	if v == nil {
		return false
	}

	// 使用类型断言，根据 v 的实际类型进行不同的处理
	switch vl := v.(type) {
	case reflect.Value: // 如果 v 是 reflect.Value 类型，判断其是否为指针类型
		return vl.Kind() == reflect.Ptr
	case reflect.Type: // 如果 v 是 reflect.Type 类型，判断其是否为指针类型
		return vl.Kind() == reflect.Ptr
	// 列举一系列非指针类型，如果 v 是这些类型之一，则返回 false
	case string, []string,
		int, int8, int16, int32, int64,
		[]int, []int8, []int16, []int32, []int64,
		uint, uint16, uint32, uint64,
		float32, float64, []float32, []float64,
		bool, []bool,
		byte, []byte,
		map[string]interface{}:
		return false
	default: // 对于其他类型，使用 reflect.TypeOf(v) 判断其是否为指针类型
		return TypeOf(v).Kind() == reflect.Ptr
	}
}

// IsStr 检查提供的值是否是字符串或可转换为字符串。
// 函数接受一个接口类型，允许检查多种类型的字符串转换能力。
// 参数:
// v: 要检查的值，可以是任意类型。
//
// 返回值:
// bool: 如果值是字符串则返回 true，否则返回 false。
func IsStr(v interface{}) bool {
	// 检查值是否为 nil，如果是则返回 false
	if v == nil {
		return false
	}

	// 使用类型切换处理不同类型的值
	switch vl := v.(type) {
	case string, *string, **string: // 直接支持字符串、字符串指针和字符串双指针
		return true
	case reflect.Value: // 处理反射值
		// 检查 reflect.Value 的类型是否为字符串
		return vl.Kind() == reflect.String
	case reflect.Type: // 处理反射类型
		// 检查 reflect.Type 是否为字符串类型
		return vl.Kind() == reflect.String
	default:
		// 如果以上情况都不匹配，则返回 false
		return false
	}
}

// IsEmpty 检查给定的值是否可以认为是“空”的。
// 它支持多种类型，包括基本类型、切片、映射、指针、函数和结构体。
// 对于 nil 值、空字符串、false 布尔值等，该函数返回 true。
func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}

	switch vl := v.(type) {
	case string:
		return vl == ""
	case bool:
		return !vl
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return vl == 0
	case float32:
		return vl == 0
	case float64:
		return vl == 0 || math.IsNaN(vl)
	case []int:
		return len(vl) == 0
	case []int8:
		return len(vl) == 0
	case []int16:
		return len(vl) == 0
	case []int32:
		return len(vl) == 0
	case []int64:
		return len(vl) == 0
	case []uint:
		return len(vl) == 0
	case []uint8:
		return len(vl) == 0
	case []uint16:
		return len(vl) == 0
	case []uint32:
		return len(vl) == 0
	case []uint64:
		return len(vl) == 0
	case []float32:
		return len(vl) == 0
	case []float64:
		return len(vl) == 0
	case map[string]interface{}:
		return len(vl) == 0
	default:
		// Using reflection for handling unsupported types,
		// optimized with predefined handling functions.
		return checkEmptyByReflect(v)
	}
}

// checkEmptyByReflect 使用反射检查无法直接处理的类型的值是否为空。
// 这是一个优化过的实现，针对预定义的处理函数。
func checkEmptyByReflect(v interface{}) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0 || math.IsNaN(rv.Float())
	case reflect.String:
		return rv.Len() == 0
	case reflect.Struct:
		return isStructEmpty(rv)
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return rv.Len() == 0 || rv.IsNil()
	default:
		return isZero(rv)
	}
}

// isStructEmpty 检查一个结构体是否所有字段都为空。
// 如果所有字段都为空或等于其零值，返回 true。
func isStructEmpty(v reflect.Value) bool {
	if v.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < v.NumField(); i++ {
		if !isZero(v.Field(i)) {
			return false
		}
	}
	return true
}

// isZero 判断给定值是否等于其零值。
// 对于某些类型（如切片、映射、指针、函数等），这等同于检查它们是否为 nil
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Interface, reflect.Chan, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return isZero(v.Elem())
	case reflect.Struct:
		return isStructEmpty(v)
	default:
		zv := reflect.Zero(v.Type())
		return v.Interface() == zv.Interface()
	}
}
