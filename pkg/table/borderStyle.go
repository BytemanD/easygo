package table

// StyleDefault renders a table like below:
//
//	+----+-------+-----+
//	| Id | Name  | Age |
//	+----+-------+-----+
//	| 1  | name1 | 1   |
//	| 2  | name2 | 10  |
//	+----+-------+-----+
var StyleDefault = []string{
	"+", "-", "+",
	"+", "+", "+",
	"+", "-", "+",
	"-", "|",
}

// StyleLight renders a table like below:
//
//	┌────┬───────┬─────┐
//	│ Id │ Name  │ Age │
//	├────┼───────┼─────┤
//	│ 1  │ name1 │ 1   │
//	│ 2  │ name2 │ 10  │
//	└────┴───────┴─────┘
var StyleLight = []string{
	"┌", "┬", "┐",
	"├", "┼", "┤",
	"└", "┴", "┘",
	"─", "│",
}

// StyleRounded renders a table like below:
//
//	╭────┬───────┬─────╮
//	│ Id │ Name  │ Age │
//	├────┼───────┼─────┤
//	│ 1  │ name1 │ 1   │
//	│ 2  │ name2 │ 10  │
//	╰────┴───────┴─────╯
var StyleRounded = []string{
	"╭", "┬", "╮",
	"├", "┼", "┤",
	"╰", "┴", "╯",
	"─", "│",
}

// StyleBold renders a table like below:
//
//	┏━━━━┳━━━━━━━┳━━━━━┓
//	┃ Id ┃ Name  ┃ Age ┃
//	┃━━━━╋━━━━━━━╋━━━━━┫
//	┃ 1  ┃ name1 ┃ 1   ┃
//	┃ 2  ┃ name2 ┃ 10  ┃
//	┗━━━━┻━━━━━━━┻━━━━━┛
var StyleBold = []string{
	"┏", "┳", "┓",
	"┣", "╋", "┫",
	"┗", "┻", "┛",
	"━", "┃",
}

// StyleBold renders a table like below:
//
//	┏━━━━┳━━━━━━━┳━━━━━┓
//	┃ Id ┃ Name  ┃ Age ┃
//	┃━━━━╋━━━━━━━╋━━━━━┫
//	┃ 1  ┃ name1 ┃ 1   ┃
//	┃ 2  ┃ name2 ┃ 10  ┃
//	┗━━━━┻━━━━━━━┻━━━━━┛
var StyleDouble = []string{
	"╔", "╦", "╗",
	"╠", "╬", "╣",
	"╚", "╩", "╝",
	"═", "║",
}
