package valueobject

type SKU struct {
	value string
}

func NewSKU(val string) SKU {
	return SKU{value: val}
}

func (s SKU) String() string {
	return s.value
}
