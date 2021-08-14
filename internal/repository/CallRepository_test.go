package repository

import (
	"fmt"
	"testing"
)

func TestCallRepository_GetItems(t *testing.T) {

	repo := NewCallRepository(nil)

	items, _ := repo.GetItems()

	fmt.Println(items)

}