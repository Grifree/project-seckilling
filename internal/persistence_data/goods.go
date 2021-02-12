package pd

type IDGoods string
func (id IDGoods) String() string {
	return string(id)
}
