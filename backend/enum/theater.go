package enum

type Theater int

const (
	Small Theater = iota
	Medium
	Large
)

func (t Theater) GetSize(s string) int {
	m := map[string]int{"Small": 20, "Medium": 30, "Large": 60,}
	return m[s]
}

func (t Theater) String() string {
	return [...]string{"Small", "Medium", "Large"}[t]
}

