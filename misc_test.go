package godash_test

import (
	"github.com/aix-tech/godash"
	"reflect"
	"testing"
)

func TestValueOf(t *testing.T) {
	// Test case 1: v is a reflect.Value type
	rv := reflect.ValueOf(42)
	res := godash.ValueOf(rv)
	if !reflect.DeepEqual(res, rv) {
		t.Errorf("Expected %v, but got %v", rv, res)
	}

	// Test case 2: v is not a reflect.Value type
	v := 42
	res = godash.ValueOf(v)
	expected := reflect.ValueOf(v)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, but got %v", expected, res)
	}
}

func TestTypeOf(t *testing.T) {
	// 测试传入reflect.Type类型时的返回值
	expected1 := reflect.TypeOf((*reflect.Type)(nil)).Elem()
	result1 := godash.TypeOf(expected1)
	if result1 != expected1 {
		t.Errorf("Expected: %v, but got: %v", expected1, result1)
	}

	// 测试传入非reflect.Type类型时的返回值
	expected2 := reflect.TypeOf("")
	result2 := godash.TypeOf("test")
	if result2 != expected2 {
		t.Errorf("Expected: %v, but got: %v", expected2, result2)
	}
}
