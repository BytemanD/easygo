package table

import (
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

func TestItemsTableDefault(t *testing.T) {
	itemsTable := ItemsTable{
		Headers: []H{
			{Field: "ID"},
			{Title: "Name", MaxWidth: 20},
			{Title: "age", Field: "Age"},
		},
		Items: humans,
	}
	expect := `+---------------------------------------+
| ID | Name                       | age |
+----+----------------------------+-----+
| 1  | Olivia Thompson            | 0   |
| 2  | 张三                       | 0   |
| 3  | ●●！                       | 0   |
| 4  | ｎａｍｅ4                  | 0   |
| 5  | 😊                         | 0   |
| 6  | Charlotte                  | 0   |
|    | Williams                   |     |
| 7  | Alexander                  | 0   |
|    | Green                      |     |
| 8  | 我！hello我！hello我！hell | 0   |
|    | o我！hello我！hello我！hel |     |
|    | lo我！hello我！hello我！he |     |
|    | llo我！hello               |     |
+----+----------------------------+-----+`

	result := itemsTable.Render()
	t.Logf("result:\n%v", result)
	if result != expect {
		t.Errorf("itemsTable.Render() = \n%v, not \n%v", result, expect)
		return
	}
}

func BenchmarkItemsTable(b *testing.B) {
	items := []Human{}
	for i := 0; i <= b.N; i++ {
		items = append(items, Human{
			ID: i, Name: strings.Repeat("Olivia Thompson", 10),
			Age: i,
		})
	}
	b.ResetTimer()
	itemsTable := ItemsTable{
		Headers: []H{
			{Field: "ID"},
			{Title: "Name", Color: true, MaxWidth: 20},
			{Title: "age", Field: "Age"},
		},
		Items: items,
	}
	itemsTable.Render()
}
