// ItemsTable 用于展示 Array 和 Slice 类型数据
//
// 支持以下特性：
//
// 样式
//  1. 设置边框格式
//  1. 自定义颜色
//  1. 开启/关闭表格内边框
//
// 列
//  1. 开启/关闭自动编号
//  1. 设置最大宽度
package table

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/BytemanD/easygo/pkg/stringutils"
	"github.com/fatih/color"
)

type ItemsTable struct {
	Name           string
	Headers        []H
	Items          interface{}
	columnsWidth   []int
	InlineBorder   bool
	AutoIndex      bool
	style          Style
	colorFormatter *color.Color
}

func (t *ItemsTable) SetStyle(style Style, colors ...color.Attribute) *ItemsTable {
	t.style = style
	if len(colors) > 0 {
		t.SetColor(colors...)
	}
	return t
}
func (t *ItemsTable) SetColor(attributes ...color.Attribute) *ItemsTable {
	t.colorFormatter = color.New(attributes...)
	return t
}
func (t *ItemsTable) EnableAutoIndex() *ItemsTable {
	t.AutoIndex = true
	return t
}
func (t ItemsTable) borderStr(i int) string {
	return t.colorStr(t.style[i])
}
func (t ItemsTable) colorStr(s string) string {
	if t.colorFormatter == nil {
		return s
	}
	return t.colorFormatter.Sprintf(s)
}

func (t ItemsTable) rowString(row []string, enableColor bool) string {
	columes := []string{}
	for i, column := range row {
		renderCol := fmt.Sprintf("%-*s",
			t.columnsWidth[i]-stringutils.TextWidth(column)+utf8.RuneCountInString(column),
			column)
		if enableColor {
			renderCol = t.colorStr(renderCol)
		}
		columes = append(columes, renderCol)
	}
	return fmt.Sprintf("%s %s %s",
		t.borderStr(10),
		strings.Join(columes, " "+t.borderStr(10)+" "),
		t.borderStr(10),
	)
}
func (t ItemsTable) titleRow() string {
	titleWidth := 0
	for _, w := range t.columnsWidth {
		titleWidth += w + 2
	}
	return fmt.Sprintf("%s %-*s %s",
		t.borderStr(10),
		titleWidth-1, t.Name,
		t.borderStr(10),
	)
}
func (t ItemsTable) headerRow(row []string) string {
	return t.rowString(row, true)
}
func (t ItemsTable) bodyRow(row []string) string {
	return t.rowString(row, false)
}
func (t ItemsTable) titleTopBorder() string {
	columes := []string{}
	for i := range t.Headers {
		columes = append(columes, strings.Repeat(t.borderStr(9), t.columnsWidth[i]+2))
	}
	return fmt.Sprintf("%s%s%s",
		t.borderStr(0), strings.Join(columes, t.borderStr(9)),
		t.borderStr(2),
	)
}
func (t ItemsTable) topBorder() string {
	columes := []string{}
	for i := range t.Headers {
		columes = append(columes, strings.Repeat(t.borderStr(9), t.columnsWidth[i]+2))
	}
	return fmt.Sprintf("%s%s%s",
		t.borderStr(0), strings.Join(columes, t.borderStr(1)),
		t.borderStr(2),
	)
}
func (t ItemsTable) inlineBorder() string {
	columes := []string{}
	for i := range t.Headers {
		columes = append(columes, strings.Repeat(
			t.borderStr(9), t.columnsWidth[i]+2))
	}
	return fmt.Sprintf("%s%s%s",
		t.borderStr(3), strings.Join(columes, t.borderStr(4)), t.borderStr(5),
	)
}
func (t ItemsTable) bottomBorder() string {
	columes := []string{}
	for i := range t.Headers {
		columes = append(columes, strings.Repeat(t.borderStr(9), t.columnsWidth[i]+2))
	}
	return fmt.Sprintf("%s%s%s",
		t.borderStr(6), strings.Join(columes, t.borderStr(7)), t.borderStr(8),
	)
}
func (t *ItemsTable) parseToRows(row []string) [][]string {
	rows := [][]string{}

	for x, column := range row {
		lines := strings.Split(column, "\n")
		for y, line := range lines {
			var subColumes []string
			// 如果设置了MaxWidth, 分割字符串
			if x < len(t.Headers) && t.Headers[x].MaxWidth != 0 {
				subColumes = stringutils.SubStrings(line, t.Headers[x].MaxWidth)
			} else {
				subColumes = []string{column}
			}
			for _, subColume := range subColumes {
				if len(rows) <= y {
					rows = append(rows, make([]string, len(row)))
				}
				rows[y][x] = subColume
				y++
			}
		}
	}
	return rows
}
func (t ItemsTable) header() []string {
	header := make([]string, len(t.Headers))
	for i, h := range t.Headers {
		header[i] = h.title()
	}
	return header
}
func (t *ItemsTable) Render() (string, error) {
	itemsValue := reflect.ValueOf(t.Items)
	if itemsValue.Kind() != reflect.Slice && itemsValue.Kind() != reflect.Array {
		return "", fmt.Errorf("items must be Slice or Array type")
	}
	if t.style == nil {
		t.SetStyle(StyleDefault)
	}

	if t.AutoIndex {
		t.Headers = append([]H{{Title: "#", isIndex: true}}, t.Headers...)
		t.columnsWidth = make([]int, len(t.Headers)+1)
	} else {
		t.columnsWidth = make([]int, len(t.Headers))
	}

	rows := [][]string{t.header()}
	borderPosition := map[int]bool{}
	for i := 0; i < itemsValue.Len(); i++ {
		item := itemsValue.Index(i)
		if item.Kind() != reflect.Struct {
			return "", fmt.Errorf("render failed, item is not Struct type")
		}

		tmpRows := []string{}
		for _, h := range t.Headers {
			if h.isIndex {
				tmpRows = append(tmpRows, strconv.Itoa(i+1))
				continue
			}
			itemField := item.FieldByName(h.field())
			if !itemField.IsValid() {
				tmpRows = append(tmpRows, "")
				continue
			}
			tmpRows = append(tmpRows, fmt.Sprintf("%v", itemField))
		}
		rows = append(rows, t.parseToRows(tmpRows)...)
		borderPosition[len(rows)-1] = true
	}
	// 计算每列合适的宽度
	for _, row := range rows {
		for x, column := range row {
			t.columnsWidth[x] = max(t.columnsWidth[x], stringutils.TextWidth(column))
		}
	}
	// t.fixMaxWidth()
	// 渲染
	lines := []string{}
	if t.Name != "" {
		lines = append(lines, t.titleTopBorder(), t.titleRow(), t.inlineBorder())
	} else {
		lines = append(lines, t.topBorder())
	}
	lines = append(lines, t.headerRow(rows[0]), t.inlineBorder())
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		lines = append(lines, t.bodyRow(row))
		if t.InlineBorder && i < len(rows)-1 && borderPosition[i] {
			lines = append(lines, t.inlineBorder())
		}
	}

	lines = append(lines, t.bottomBorder())
	return strings.Join(lines, "\n"), nil
}

func NewItemsTable(titles []string, items interface{}) *ItemsTable {
	header := make([]H, len(titles))
	for i, title := range titles {
		header[i] = H{Title: title}
	}
	table := ItemsTable{
		Headers: header,
		Items:   items,
	}
	return &table
}
