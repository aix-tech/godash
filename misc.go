package godash

import "reflect"

// ValueOf 返回给定参数的反射值。
// 如果参数已经是 reflect.Value 类型，则直接返回该值；
// 否则，将参数转换为 reflect.Value 并返回。
//
// 参数:
// v interface{} - 待处理的值，可以是任意类型。
//
// 返回值:
// reflect.Value - 给定参数的反射值。
func ValueOf(v interface{}) reflect.Value {
	// 尝试将 v 解释为 reflect.Value 类型，如果成功，则直接返回
	if rv, ok := v.(reflect.Value); ok {
		return rv
	}
	// 如果 v 不是 reflect.Value 类型，则通过 reflect.ValueOf 方法转换并返回
	return reflect.ValueOf(v)
}

// TypeOf 返回给定参数的反射类型。
//
// 参数:
// v interface{} - 任意类型的值。
//
// 返回值:
// reflect.Type - 给定参数的反射类型。如果参数v已经是reflect.Type类型，则直接返回该类型，
// 否则返回reflect.TypeOf(v)。
func TypeOf(v interface{}) reflect.Type {
	// 尝试将v断言为reflect.Type类型，如果成功，则直接返回该类型。
	if rt, ok := v.(reflect.Type); ok {
		return rt
	}
	// 如果断言失败，说明v不是reflect.Type类型，返回reflect.TypeOf(v)获取v的类型。
	return reflect.TypeOf(v)
}
