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

func TestItemsTable(t *testing.T) {
	items := []Human{
		{ID: 1, Name: "Jack"},
		{ID: 2, Name: "张三"},
		{ID: 2, Name: "●●"},
		{ID: 2, Name: "ｎａｍｅ2"},
		{ID: 2, Name: "😊"},
	}
	itemsTable := ItemsTable{
		Headers: []H{
			{Title: "Id", Field: "ID"},
			{Title: "Name", Color: true},
			{Title: "Age"},
		},
		Items: items,
	}
	itemsTable.SetStyle(StyleDefault)
	fmt.Println(itemsTable.Render())
	itemsTable.SetStyle(StyleLight)
	fmt.Println(itemsTable.Render())

	itemsTable.SetStyle(StyleRounded)
	fmt.Println(itemsTable.Render())

	itemsTable.SetStyle(StyleBold)
	fmt.Println(itemsTable.Render())
	itemsTable.SetStyle(StyleDouble)
	fmt.Println(itemsTable.Render())
}
