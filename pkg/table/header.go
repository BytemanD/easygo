package table

type H struct {
	Title    string
	Field    string
	MaxWidth int
	isIndex  bool
}

func (h H) title() string {
	var t string
	if h.Title != "" {
		t = h.Title
	} else {
		t = h.Field
	}
	return t
}

func (h H) field() string {
	if h.Field != "" {
		return h.Field
	}
	return h.Title
}
