package proxylist

type ByFilter []Settings

func (s ByFilter) Len() int {
	return len(s)
}

func (s ByFilter) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByFilter) Less(i int, j int) bool {
	return s[i].Filter < s[j].Filter
}
