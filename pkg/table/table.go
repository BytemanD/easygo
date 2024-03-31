package table

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"

	"github.com/BytemanD/easygo/pkg/stringutils"
	"github.com/fatih/color"
)

type ItemsTable struct {
	Headers      []H
	Items        interface{}
	columnsWidth []int
	style        []string
}

func (t *ItemsTable) SetStyle(style []string) {
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
func (t ItemsTable) centerBorder() string {
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

func (t ItemsTable) Render() string {
	itemsValue := reflect.ValueOf(t.Items)
	rows := make([][]string, itemsValue.Len())
	rowHeader := make([]string, len(t.Headers))
	t.columnsWidth = make([]int, len(t.Headers))
	for i, h := range t.Headers {
		rowHeader[i] = h.title()
		t.columnsWidth[i] = max(t.columnsWidth[i], stringutils.TextWidth(h.title()))
	}
	for i := 0; i < itemsValue.Len(); i++ {
		item := itemsValue.Index(i)
		for j, h := range t.Headers {
			itemField := item.FieldByName(h.field())
			rows[i] = append(rows[i], fmt.Sprintf("%v", itemField))
			itemValue := fmt.Sprintf("%v", itemField)
			t.columnsWidth[j] = max(t.columnsWidth[j], stringutils.TextWidth(itemValue))
		}
	}

	lines := []string{
		t.topBorder(), t.headerRow(rowHeader), t.centerBorder(),
	}
	for _, row := range rows {
		lines = append(lines, t.bodyRow(row))
	}
	lines = append(lines, t.bottomBorder())
	fmt.Println(t.columnsWidth)
	return strings.Join(lines, "\n")
}
