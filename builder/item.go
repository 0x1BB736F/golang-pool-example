package pool

import "strings"

type Item struct {
	builder  strings.Builder
	idx      int
	poolChan chan int
}

func newItem(idx int, c chan int) *Item {
	c <- idx

	return &Item{
		builder:  strings.Builder{},
		idx:      idx,
		poolChan: c,
	}
}

func (i *Item) Write(b []byte) (int, error) {
	return i.builder.Write(b)
}

func (i *Item) WriteString(s string) (int, error) {
	return i.builder.WriteString(s)
}

func (i *Item) WriteByte(b byte) error {
	return i.builder.WriteByte(b)
}

func (i *Item) WriteRune(r rune) (int, error) {
	return i.builder.WriteRune(r)
}

func (i *Item) Reset() {
	i.builder.Reset()
}

func (i *Item) Close() {
	i.poolChan <- i.idx

	i.Reset()
}

func (i *Item) Index() int {
	return i.idx
}

func (i *Item) String() string {
	return i.builder.String()
}
