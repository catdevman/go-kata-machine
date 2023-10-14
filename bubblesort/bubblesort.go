package bubblesort

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func Bubblesort[T constraints.Ordered](items []T, order string) ([]T, error){
    if order != "asc" && order != "desc" {
        return nil, errors.New("order needs to be \"asc\" or \"desc\"")
    }
    if len(items) == 0 {
        return nil, errors.New("nothing to sort")
    }

    if len(items) == 1 {
        return items, nil
    }

    if order == "asc" {
        // asc
        for i := range items {
            for j := 0; j < len(items) - 1 - i; j++{
                if items[j] > items[j + 1] {
                    items[j], items[j + 1] = items[j + 1], items[j]
                }
            }
        }
    } else {
        // desc
        for i := range items {
            for j := 0; j < len(items) - 1 - i; j++{
                if items[j] < items[j + 1] {
                    items[j], items[j + 1] = items[j + 1], items[j]
                }
            }
        }
    }
    return items, nil
}
