package bubblesort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T){
    t.Run("test for empty items", func(t *testing.T){
        want := []int(nil)
        got, _ := Bubblesort([]int{}, "asc")

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })
    t.Run("test ascending order sort", func(t *testing.T){
        want := []int{1, 2, 3, 4, 5}
        got, _ := Bubblesort([]int{5, 4, 3, 2, 1}, "asc")
        if !reflect.DeepEqual(got, want){
            t.Errorf("got: %v, want: %v", got, want)
        }
    })
    t.Run("test descending order sort", func(t *testing.T){
        want := []int{5, 4, 3, 2, 1}
        got, _ := Bubblesort([]int{1, 2, 3, 4, 5}, "desc")
        if !reflect.DeepEqual(got, want){
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

    t.Run("test invalid ordering flag", func(t *testing.T){
        want := []int(nil)
        got, err := Bubblesort([]int{1, 2, 3, 4, 5}, "error")
        if err != nil {
            fmt.Printf("error: %v\n", err)
        }

        if !reflect.DeepEqual(got, want){
            t.Errorf("got: %v, want: %v", got, want)
        }
    })
}
