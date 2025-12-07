package main

import (
	"fmt"
	"slices"
	"testing"
)

func testData() []int {	
	return []int{1,2,3,4,5,6,7,8,9,10}
}

func Test_sliceExample(t *testing.T) {
	data := testData()

	got := sliceExample(data)

	want := []int{2,4,6,8,10}

	if slices.Equal(got, want) != true {
	  t.Fatalf("Must be %d, want %d", got, want)
	}

	t.Log("Test_sliceExample passed")
}

func Test_addElements(t *testing.T) {
	data := testData()
	el := 99

	got := addElements(data, el)
	want := []int{1,2,3,4,5,6,7,8,9,10,el}

  if slices.Equal(got, want) != true  {
	  t.Fatalf("Got slice %d, want %d", el, want[len(want) - 1])
	}

	if want[len(want) - 1] != el {
	  t.Fatalf("Added element must be %d, want %d", el, want[len(want) - 1])
	}

	t.Log("Test_addElements passed")
}

func Test_copySlice(t *testing.T) {
	data := testData()

	got := copySlice(data)

	if fmt.Sprintf("%p\n", data) == fmt.Sprintf("%p\n", got) {
		t.Fatalf("shouldn't point to one memory cell")
	}

	data = append(data, 100)

  if &got[0] == &data[0] {
		t.Fatalf("Changes original slice should not influence copy slice")
	}

	t.Log("Test passed")

}
