package table

import (
	"fmt"
	"testing"
)

type Human struct {
	ID   int
	Name string
	Age  int
}

var humans = []Human{
	{ID: 1, Name: "Jack"},
	{ID: 2, Name: "张三"},
	{ID: 3, Name: "●●"},
	{ID: 4, Name: "ｎａｍｅ4"},
	{ID: 5, Name: "😊"},
	{ID: 6, Name: "name\n6"},
	{ID: 7, Name: "name\n7"},
	// {ID: 2, Name: strings.Repeat("name3", 100)},
}

func TestItemsTable(t *testing.T) {
	itemsTable := ItemsTable{
		Headers: []H{
			{Field: "ID"},
			{Title: "Name", Color: true},
			{Title: "age", Field: "Age"},
		},
		Items: humans,
	}
	itemsTable.SetStyle(StyleDefault)
	fmt.Println(itemsTable.Render())

	itemsTable.InlineBorder = true
	fmt.Println(itemsTable.Render())
}
