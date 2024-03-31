package table

type Style []string

// StyleDefault renders a table like below:
//
//	+----+-------+-----+
//	| Id | Name  | Age |
//	+----+-------+-----+
//	| 1  | name1 | 1   |
//	| 2  | name2 | 10  |
//	+----+-------+-----+
var StyleDefault = Style{
	"+", "-", "+",
	"+", "+", "+",
	"+", "+", "+",
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
var StyleLight = Style{
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
var StyleRounded = Style{
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
var StyleBold = Style{
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
var StyleDouble = Style{
	"╔", "╦", "╗",
	"╠", "╬", "╣",
	"╚", "╩", "╝",
	"═", "║",
}
