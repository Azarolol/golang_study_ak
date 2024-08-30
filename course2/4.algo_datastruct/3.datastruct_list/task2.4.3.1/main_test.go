package main

import (
	"reflect"
	"testing"
)

var d DoubleLinkedList
var _ = d.LoadData("data.json")

func TestDoubleLinkedList_Len(t *testing.T) {
	tests := []struct {
		name string
		d    *DoubleLinkedList
		want int
	}{
		{"Test len", &d, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Len(); got != tt.want {
				t.Errorf("DoubleLinkedList.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Current(t *testing.T) {
	tests := []struct {
		name string
		d    *DoubleLinkedList
		want *Node
	}{
		{"Test current", &d, d.curr},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Current(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleLinkedList.Current() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Next(t *testing.T) {
	tests := []struct {
		name string
		d    *DoubleLinkedList
		want *Node
	}{
		{"Test next", &d, d.curr.next},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleLinkedList.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Prev(t *testing.T) {
	tests := []struct {
		name string
		d    *DoubleLinkedList
		want *Node
	}{
		{"Test prev", &d, d.curr.prev},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleLinkedList.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Insert(t *testing.T) {
	type args struct {
		n int
		c Commit
	}
	tests := []struct {
		name    string
		d       *DoubleLinkedList
		args    args
		wantErr bool
	}{
		{"Test insert", &d, args{10, Commit{}}, false},
		{"Test insert err 1", &d, args{-1, Commit{}}, true},
		{"Test insert err 2", &d, args{100, Commit{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Insert(tt.args.n, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DoubleLinkedList.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_GetByIndex(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		d       *DoubleLinkedList
		args    args
		want    *Node
		wantErr bool
	}{
		{"Test get by index", &d, args{0}, d.head, false},
		{"Test get by index err 1", &d, args{-1}, nil, true},
		{"Test get by index err 2", &d, args{100}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.GetByIndex(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoubleLinkedList.GetByIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleLinkedList.GetByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_SetCurrent(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		d       *DoubleLinkedList
		args    args
		wantErr bool
	}{
		{"Test set current", &d, args{0}, false},
		{"Test set current err 1", &d, args{-1}, true},
		{"Test set current err 2", &d, args{100}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.SetCurrent(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("DoubleLinkedList.SetCurrent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Push(t *testing.T) {
	type args struct {
		c Commit
	}
	tests := []struct {
		name    string
		d       *DoubleLinkedList
		args    args
		wantErr bool
	}{
		{"Test push", &d, args{Commit{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Push(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DoubleLinkedList.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Delete(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		d       *DoubleLinkedList
		args    args
		wantErr bool
	}{
		{"Test delete", &d, args{0}, false},
		{"Test delete err 1", &d, args{-1}, true},
		{"Test delete err 2", &d, args{100}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Delete(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("DoubleLinkedList.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_DeleteCurrent(t *testing.T) {
	tests := []struct {
		name    string
		d       *DoubleLinkedList
		wantErr bool
	}{
		{"Test delete current", &d, false},
		{"Test delete current err", &DoubleLinkedList{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.DeleteCurrent(); (err != nil) != tt.wantErr {
				t.Errorf("DoubleLinkedList.DeleteCurrent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Index(t *testing.T) {
	tests := []struct {
		name    string
		d       *DoubleLinkedList
		want    int
		wantErr bool
	}{
		{"Test index", &d, 0, false},
		{"Test index err", &DoubleLinkedList{}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Index()
			if (err != nil) != tt.wantErr {
				t.Errorf("DoubleLinkedList.Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoubleLinkedList.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Pop(t *testing.T) {
	tests := []struct {
		name string
		d    *DoubleLinkedList
		want *Node
	}{
		{"Test pop", &d, d.tail},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleLinkedList.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Shift(t *testing.T) {
	tests := []struct {
		name string
		d    *DoubleLinkedList
		want *Node
	}{
		{"Test shift", &d, d.head},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Shift(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleLinkedList.Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_SearchUUID(t *testing.T) {
	type args struct {
		uuID string
	}
	tests := []struct {
		name string
		d    *DoubleLinkedList
		args args
		want *Node
	}{
		{"Test search UUID", &d, args{d.head.data.UUID}, d.head},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.SearchUUID(tt.args.uuID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleLinkedList.SearchUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Search(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		d    *DoubleLinkedList
		args args
		want *Node
	}{
		{"Test search", &d, args{d.head.data.Message}, d.head},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Search(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleLinkedList.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
