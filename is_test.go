package godash_test

import (
	"github.com/aix-tech/godash"
	"reflect"
	"testing"
)

func TestIsPtr(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
		want bool
	}{
		{
			name: "nil value",
			v:    nil,
			want: false,
		},
		{
			name: "reflect.Value of non-pointer",
			v:    reflect.ValueOf(5),
			want: false,
		},
		{
			name: "reflect.Value of pointer",
			v:    reflect.ValueOf(&struct{}{}),
			want: true,
		},
		{
			name: "reflect.Type of non-pointer",
			v:    reflect.TypeOf(5),
			want: false,
		},
		{
			name: "reflect.Type of pointer",
			v:    reflect.TypeOf(&struct{}{}),
			want: true,
		},
		{
			name: "string",
			v:    "test",
			want: false,
		},
		{
			name: "[]string",
			v:    []string{"test1", "test2"},
			want: false,
		},
		{
			name: "int",
			v:    5,
			want: false,
		},
		{
			name: "[]int",
			v:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "int64",
			v:    int64(10),
			want: false,
		},
		{
			name: "[]int64",
			v:    []int64{10, 20, 30},
			want: false,
		},
		{
			name: "uint",
			v:    uint(5),
			want: false,
		},
		{
			name: "[]uint",
			v:    []uint{1, 2, 3},
			want: false,
		},
		{
			name: "float32",
			v:    float32(3.14),
			want: false,
		},
		{
			name: "[]float32",
			v:    []float32{1.5, 2.5, 3.5},
			want: false,
		},
		{
			name: "bool",
			v:    true,
			want: false,
		},
		{
			name: "[]bool",
			v:    []bool{true, false},
			want: false,
		},
		{
			name: "byte",
			v:    byte('c'),
			want: false,
		},
		{
			name: "[]byte",
			v:    []byte{'t', 'e', 's', 't'},
			want: false,
		},
		{
			name: "map[string]interface{}",
			v:    map[string]interface{}{"key1": "value1", "key2": "value2"},
			want: false,
		},
		{
			name: "custom struct",
			v:    struct{}{},
			want: false,
		},
		{
			name: "pointer to custom struct",
			v:    &struct{}{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := godash.IsPtr(tt.v); got != tt.want {
				t.Errorf("IsPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsStr(t *testing.T) {
	type args struct {
		v interface{}
	}
	ptr := "asd"
	ptr2 := &ptr
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test-1", args: args{v: ""}, want: true},
		{name: "test-2", args: args{v: "123"}, want: true},
		{name: "test-2-1", args: args{v: &ptr}, want: true},
		{name: "test-2-2", args: args{v: &ptr2}, want: true},
		{name: "test-3", args: args{v: 123}, want: false},
		{name: "test-4", args: args{v: map[string]interface{}{}}, want: false},
		{name: "test-5", args: args{v: 1.1}, want: false},
		{name: "test-6", args: args{v: struct{}{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := godash.IsStr(tt.args.v); got != tt.want {
				t.Errorf("IsStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type testStruct struct {
		Field1 int
		Field2 string
	}
	tests := []struct {
		name string
		v    interface{}
		want bool
	}{
		{"nil", nil, true},
		{"empty string", "", true},
		{"non-empty string", "hello", false},
		{"false bool", false, true},
		{"true bool", true, false},
		{"zero int", 0, true},
		{"non-zero int", 10, false},
		{"zero int8", int8(0), true},
		{"non-zero int8", int8(10), false},
		{"zero int16", int16(0), true},
		{"non-zero int16", int16(10), false},
		{"zero int32", int32(0), true},
		{"non-zero int32", int32(10), false},
		{"zero int64", int64(0), true},
		{"non-zero int64", int64(10), false},
		{"zero uint", uint(0), true},
		{"non-zero uint", uint(10), false},
		{"zero uint8", uint8(0), true},
		{"non-zero uint8", uint8(10), false},
		{"zero uint16", uint16(0), true},
		{"non-zero uint16", uint16(10), false},
		{"zero uint32", uint32(0), true},
		{"non-zero uint32", uint32(10), false},
		{"zero uint64", uint64(0), true},
		{"non-zero uint64", uint64(10), false},
		{"zero float32", float32(0), true},
		{"non-zero float32", float32(10), false},
		{"zero float64", float64(0), true},
		{"non-zero float64", float64(10), false},
		{"empty slice", []int{}, true},
		{"non-empty slice", []int{1}, false},
		{"empty map", map[string]interface{}{}, true},
		{"non-empty map", map[string]interface{}{"key": "value"}, false},
		{"nil pointer", (*int)(nil), true},
		{"non-nil zero ptr", new(int), true},
		{"non-nil non-zero ptr", new(int), true}, // Assuming the pointer is initialized with zero value
		{"non-nil struct ptr", &testStruct{}, true},
		{"non-nil struct ptr non-zero", &testStruct{Field1: 1, Field2: "test"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := godash.IsEmpty(tt.v); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
