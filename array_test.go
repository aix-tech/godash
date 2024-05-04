package godash_test

import (
	"fmt"
	"github.com/aix-tech/godash"
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args[T any] struct {
		arr  []T
		size int
	}
	type testCase[T any] struct {
		name       string
		args       args[T]
		wantResult [][]T
		wantErr    bool
	}
	tests := []testCase[int]{
		{name: "empty", args: args[int]{arr: nil, size: 0}, wantResult: nil, wantErr: true},
		{name: "empty-2", args: args[int]{arr: []int{}, size: 0}, wantResult: nil, wantErr: true},
		{name: "empty-3", args: args[int]{arr: []int{}, size: 0}, wantResult: nil, wantErr: true},
		{name: "one", args: args[int]{arr: []int{1, 2, 3}, size: 1}, wantResult: [][]int{{1}, {2}, {3}}, wantErr: false},
		{name: "two", args: args[int]{arr: []int{1, 2, 3, 4}, size: 2}, wantResult: [][]int{{1, 2}, {3, 4}}, wantErr: false},
		{name: "three", args: args[int]{arr: []int{1, 2, 3, 4}, size: 3}, wantResult: [][]int{{1, 2, 3}, {4}}, wantErr: false},
		{name: "equal", args: args[int]{arr: []int{1, 2, 3, 4}, size: 4}, wantResult: [][]int{{1, 2, 3, 4}}, wantErr: false},
		{name: "more than", args: args[int]{arr: []int{1, 2, 3, 4}, size: 8}, wantResult: [][]int{{1, 2, 3, 4}}, wantErr: false},
		{name: "-1", args: args[int]{arr: []int{1, 2, 3, 4}, size: -1}, wantResult: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := godash.Chunk(tt.args.arr, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chunk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Chunk() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func BenchmarkChunk(b *testing.B) {
	arr := make([]int, 100000000)
	for i := 0; i < b.N; i++ {
		_, _ = godash.Chunk[int](arr, 7)
	}
}

func TestCompact(t *testing.T) {
	// 测试用例1：原数组为空，返回空数组。
	arr1 := []int{}
	expected1 := []int{}
	if got1 := godash.Compact(arr1); !reflect.DeepEqual(got1, expected1) {
		t.Errorf("Compact(%v) = %v, want %v", arr1, got1, expected1)
	}

	// 测试用例2：原数组中所有元素为空，返回空数组。
	arr2 := []string{"", "", ""}
	expected2 := []string{}
	if got2 := godash.Compact(arr2); !reflect.DeepEqual(got2, expected2) {
		t.Errorf("Compact(%v) = %v, want %v", arr2, got2, expected2)
	}

	// 测试用例3：原数组中部分元素为空，返回非空数组。
	arr3 := []int{0, 1, 0, 2, 0, 3}
	expected3 := []int{1, 2, 3}
	if got3 := godash.Compact(arr3); !reflect.DeepEqual(got3, expected3) {
		t.Errorf("Compact(%v) = %v, want %v", arr3, got3, expected3)
	}

	// 测试用例4：原数组中所有元素非空，返回原数组。
	arr4 := []string{"a", "b", "c"}
	expected4 := []string{"a", "b", "c"}
	if got4 := godash.Compact(arr4); !reflect.DeepEqual(got4, expected4) {
		t.Errorf("Compact(%v) = %v, want %v", arr4, got4, expected4)
	}
}

func BenchmarkCompact(b *testing.B) {
	// 测试用例1：不同大小的原数组
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			arr := make([]int, size)
			for i := 0; i < size; i++ {
				if i%2 == 0 {
					arr[i] = i
				}
			}

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = godash.Compact(arr)
			}
		})
	}
}
