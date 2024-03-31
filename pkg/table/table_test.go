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
	{ID: 111, Name: "Olivia Thompson"},
	{ID: 222, Name: "张三"},
	{ID: 333, Name: "●●！"},
	{ID: 444, Name: "ｎａｍｅ4", Age: 10},
	{ID: 555, Name: "😊"},
	{ID: 666, Name: "Charlotte\nWilliams"},
	{ID: 777, Name: "Alexander\nGreen"},
	{ID: 888, Name: strings.Repeat("我！hello", 10)},
	{ID: 999, Name: "李四"},
	{ID: 101010, Name: "王五"},
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
	expect := `+--------+----------------------------+-----+
| ID     | Name                       | age |
+--------+----------------------------+-----+
| 111    | Olivia Thompson            | 0   |
| 222    | 张三                       | 0   |
| 333    | ●●！                       | 0   |
| 444    | ｎａｍｅ4                  | 10  |
| 555    | 😊                         | 0   |
| 666    | Charlotte                  | 0   |
|        | Williams                   |     |
| 777    | Alexander                  | 0   |
|        | Green                      |     |
| 888    | 我！hello我！hello我！hell | 0   |
|        | o我！hello我！hello我！hel |     |
|        | lo我！hello我！hello我！he |     |
|        | llo我！hello               |     |
| 999    | 李四                       | 0   |
| 101010 | 王五                       | 0   |
+--------+----------------------------+-----+`
	result := itemsTable.Render()
	t.Logf("result:\n%v", result)
	if result != expect {
		t.Errorf("itemsTable.Render() = \n%v, not \n%v", result, expect)
		return
	}
}
func TestItemsTableInlineBorder(t *testing.T) {
	itemsTable := ItemsTable{
		Headers: []H{
			{Field: "ID"},
			{Title: "Name", MaxWidth: 20},
			{Title: "age", Field: "Age"},
		},
		Items:        humans,
		InlineBorder: true,
	}
	expect := `+--------+----------------------------+-----+
| ID     | Name                       | age |
+--------+----------------------------+-----+
| 111    | Olivia Thompson            | 0   |
+--------+----------------------------+-----+
| 222    | 张三                       | 0   |
+--------+----------------------------+-----+
| 333    | ●●！                       | 0   |
+--------+----------------------------+-----+
| 444    | ｎａｍｅ4                  | 10  |
+--------+----------------------------+-----+
| 555    | 😊                         | 0   |
+--------+----------------------------+-----+
| 666    | Charlotte                  | 0   |
|        | Williams                   |     |
+--------+----------------------------+-----+
| 777    | Alexander                  | 0   |
|        | Green                      |     |
+--------+----------------------------+-----+
| 888    | 我！hello我！hello我！hell | 0   |
|        | o我！hello我！hello我！hel |     |
|        | lo我！hello我！hello我！he |     |
|        | llo我！hello               |     |
+--------+----------------------------+-----+
| 999    | 李四                       | 0   |
+--------+----------------------------+-----+
| 101010 | 王五                       | 0   |
+--------+----------------------------+-----+`
	result := itemsTable.Render()
	t.Logf("result:\n%v", result)
	if result != expect {
		t.Errorf("itemsTable.Render() = \n%v, not \n%v", result, expect)
		return
	}
}
func TestItemsTableAutoIndex(t *testing.T) {
	itemsTable := ItemsTable{
		Headers: []H{
			{Field: "ID"},
			{Title: "Name", MaxWidth: 20},
			{Title: "age", Field: "Age"},
		},
		Items:     humans,
		AutoIndex: true,
	}
	expect := `+----+--------+----------------------------+-----+
| #  | ID     | Name                       | age |
+----+--------+----------------------------+-----+
| 1  | 111    | Olivia Thompson            | 0   |
| 2  | 222    | 张三                       | 0   |
| 3  | 333    | ●●！                       | 0   |
| 4  | 444    | ｎａｍｅ4                  | 10  |
| 5  | 555    | 😊                         | 0   |
| 6  | 666    | Charlotte                  | 0   |
|    |        | Williams                   |     |
| 7  | 777    | Alexander                  | 0   |
|    |        | Green                      |     |
| 8  | 888    | 我！hello我！hello我！hell | 0   |
|    |        | o我！hello我！hello我！hel |     |
|    |        | lo我！hello我！hello我！he |     |
|    |        | llo我！hello               |     |
| 9  | 999    | 李四                       | 0   |
| 10 | 101010 | 王五                       | 0   |
+----+--------+----------------------------+-----+`
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
