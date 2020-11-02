package mytree

import (
	"reflect"
	"testing"
)

var tree = New()

func Test_preprocedural(t *testing.T) {
	type args struct {
		t    *Tree
		path []int
	}
	TreePath = TreePath[0:0]
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			// 1->2->4->3->5->6
			args: args{
				t: tree,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preprocedural(tt.args.t)
			want := []int{1, 2, 4, 3, 5, 6}
			if !reflect.DeepEqual(TreePath, want) {
				t.Errorf("want: %v, compute: %v", want, TreePath)
			}
		})
	}
}

func Test_midsequent(t *testing.T) {
	type args struct {
		t *Tree
	}
	TreePath = TreePath[0:0]
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			// 1->2->4->3->5->6
			args: args{
				t: tree,
			},
		},
	}
	for _, tt := range tests {
		want := []int{4, 2, 1, 5, 3, 6}
		t.Run(tt.name, func(t *testing.T) {
			midsequent(tt.args.t)
			if !reflect.DeepEqual(TreePath, want) {
				t.Errorf("want: %v, compute: %v", want, TreePath)
			}
		})
	}
}

func Test_subsequent(t *testing.T) {
	type args struct {
		t *Tree
	}
	TreePath = TreePath[0:0]
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{t: tree},
		},
	}
	for _, tt := range tests {
		want := []int{4, 2, 5, 6, 3, 1}
		t.Run(tt.name, func(t *testing.T) {
			subsequent(tt.args.t)
			if !reflect.DeepEqual(TreePath, want) {
				t.Errorf("want: %v, compute: %v", want, TreePath)
			}
		})
	}
}

func Test_preproceduralv2(t *testing.T) {
	type args struct {
		t *Tree
	}
	TreePath = TreePath[0:0]
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				t: tree,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v2 := preproceduralv2(tt.args.t)
			preprocedural(tt.args.t)
			if !reflect.DeepEqual(v2, TreePath) {
				t.Errorf("want: %v, but compute %v", v2, TreePath)
			}

		})
	}
}

func Test_midsequentv2(t *testing.T) {
	type args struct {
		t *Tree
	}
	TreePath = TreePath[0:0]
	midsequent(tree)
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				t: tree,
			},
			want: TreePath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := midsequentv2(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("midsequentv2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subsequentv2(t *testing.T) {
	type args struct {
		t *Tree
	}
	TreePath = TreePath[0:0]
	subsequent(tree)
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				t: tree,
			},
			want: TreePath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsequentv2(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsequentv2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDFS(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				t: tree,
			},
			want: []int{1, 2, 4, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := DFS(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDFSv2(t *testing.T) {
	type args struct {
		t      *Tree
		result *[]int
	}
	result := make([]int, 0)
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				t:      tree,
				result: &result,
			},
		},
	}
	for _, tt := range tests {
		r := DFS(tree)
		if DFSv2(tt.args.t, tt.args.result); !reflect.DeepEqual(*tt.args.result, r) {
			t.Errorf("DFS() = %v, want %v", tt.args.result, r)
		}
	}
}

func TestBFS(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				t: tree,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFS(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		nums []int
		i    int
		j    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums: []int{1, 2},
				i:    0,
				j:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.nums, tt.args.i, tt.args.j)
			t.Log(tt.args.nums)
		})
	}
}

var nums = []int{3, 4, 5, 2, 1, 8}
var sortNums = []int{1, 2, 3, 4, 5, 8}

func Test_helper(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums:  nums,
				left:  0,
				right: len(nums),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := helper(tt.args.nums, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("helper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums: nums,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, sortNums) {
				t.Errorf("tt.args.num %p, sortNums %p \n", tt.args.nums, sortNums)
				t.Errorf("want %v, compute %v", tt.args.nums, sortNums)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		n1 []int
		n2 []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				n1: []int{2, 7},
				n2: []int{3, 8},
			},
			want:    []int{2, 3, 7, 8},
			wantErr: false,
		},
		{
			args: args{
				n1: []int{1, 2},
				n2: []int{8},
			},
			want:    []int{1, 2, 8},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := merge(tt.args.n1, tt.args.n2)
			if (err != nil) != tt.wantErr {
				t.Errorf("merge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums: nums,
			},
			want: sortNums,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
