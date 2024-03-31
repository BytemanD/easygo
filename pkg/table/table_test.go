package table

import (
	"fmt"
	"strings"
	"testing"
)

type Human struct {
	ID   int
	Name string
	Age  int
}

var humans = []Human{
	{ID: 1, Name: "Olivia Thompson"},
	{ID: 2, Name: "张三"},
	{ID: 3, Name: "●●！"},
	{ID: 4, Name: "ｎａｍｅ4"},
	{ID: 5, Name: "😊"},
	{ID: 6, Name: "Charlotte\nWilliams"},
	{ID: 7, Name: "Alexander\nGreen"},
	{ID: 8, Name: strings.Repeat("我！hello", 10)},
}

func TestItemsTable(t *testing.T) {
	itemsTable := ItemsTable{
		Headers: []H{
			{Field: "ID"},
			{Title: "Name", Color: true, MaxWidth: 20},
			{Title: "age", Field: "Age"},
		},
		Items: humans,
	}
	itemsTable.SetStyle(StyleDefault)
	fmt.Println(itemsTable.Render())

	itemsTable.InlineBorder = true
	fmt.Println(itemsTable.Render())
}
