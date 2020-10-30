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
