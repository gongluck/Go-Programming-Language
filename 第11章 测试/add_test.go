// add_test.go

package add

import (
	"fmt"
	"testing"
)

// go test -v -run="Add" -coverprofile="c.out"
// go tool cover -html="c.out"
func TestAdd2zero(t *testing.T) {
	if Add(0, 0) != 0 {
		t.Error("Add(0, 0) != 0")
	}
}

func TestAdd2positive(t *testing.T) {
	if Add(10, 10) != 20 {
		t.Error("Add(10, 10) != 20")
	}
}

func TestAddpositivenegative(t *testing.T) {
	if Add(10, -10) != 0 {
		t.Error("Add(10, -10) != 0")
	}
}

func TestAdd2negative(t *testing.T) {
	if Add(-10, -10) != -20 {
		t.Error("Add(-10, -10) != -20")
	}
}

// func TestAdd2char(t *testing.T) {
// 	if Add('a', 'b') != 'c' {
// 		t.Error("Add('a', 'b') != 'c'")
// 	}
// }

//go test -bench="." -benchmem
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, i)
	}
}

// go test -run=NONE -bench="." -cpuprofile="cpu.out"
// go tool pprof -text -nodecount=10 ".\cpu.out"

func ExampleAdd() {
	fmt.Println(Add(10, 10))
	// 输出:
	// 20
}
