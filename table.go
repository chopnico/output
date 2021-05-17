package output

type Entry struct {
	Label string
	Data string
}

type List struct {
	Name string
	Entries []Entry
}
