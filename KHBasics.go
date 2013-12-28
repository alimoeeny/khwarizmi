package khwarizmi

type IntArray []int
type Int64Array []int64

func (a IntArray) Sum() int64 {
	sum := int64(0)
	for _, v := range a {
		sum = sum + int64(v)
	}
	return sum
}

func (a Int64Array) Sum() int64 {
	sum := int64(0)
	for _, v := range a {
		sum = sum + v
	}
	return sum
}
