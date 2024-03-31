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
	Headers      []H
	Items        interface{}
	columnsWidth []int
	style        Style
	InlineBorder bool
	AutoIndex    bool
}

func (t *ItemsTable) SetStyle(style Style) {
	t.style = style
}

func (t ItemsTable) rowString(row []string, enableColor bool) string {
	columes := []string{}
	for i, column := range row {
		renderCol := fmt.Sprintf("%-*s",
			t.columnsWidth[i]-stringutils.TextWidth(column)+utf8.RuneCountInString(column),
			column)
		if enableColor && t.Headers[i].Color {
			renderCol = color.CyanString(renderCol)
		}
		columes = append(columes, renderCol)
	}
	return fmt.Sprintf("%s %s %s",
		t.style[10],
		strings.Join(columes, " "+t.style[10]+" "),
		t.style[10],
	)
}
func (t ItemsTable) headerRow(row []string) string {
	return t.rowString(row, true)
}
func (t ItemsTable) bodyRow(row []string) string {
	return t.rowString(row, false)
}
func (t ItemsTable) topBorder() string {
	columes := []string{}
	for i := range t.Headers {
		columes = append(columes, strings.Repeat(t.style[9], t.columnsWidth[i]+2))
	}
	return fmt.Sprintf("%s%s%s",
		t.style[0], strings.Join(columes, t.style[1]), t.style[2],
	)
}
func (t ItemsTable) inlineBorder() string {
	columes := []string{}
	for i := range t.Headers {
		columes = append(columes, strings.Repeat(t.style[9], t.columnsWidth[i]+2))
	}
	return fmt.Sprintf("%s%s%s",
		t.style[3], strings.Join(columes, t.style[4]), t.style[5],
	)
}
func (t ItemsTable) bottomBorder() string {
	columes := []string{}
	for i := range t.Headers {
		columes = append(columes, strings.Repeat(t.style[9], t.columnsWidth[i]+2))
	}
	return fmt.Sprintf("%s%s%s",
		t.style[6], strings.Join(columes, t.style[7]), t.style[8],
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
func (t ItemsTable) Render() (string, error) {
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
	lines := []string{t.topBorder(), t.headerRow(rows[0]), t.inlineBorder()}
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
