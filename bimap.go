package bimap

//Uint8 is a bi-directional map with both key and value Uint8. ab being the normal map and ba the reverse.
//src: https://stackoverflow.com/questions/23190311/reverse-a-map-in-value-key-format-in-golang
type Uint8 struct {
	ab map[uint8]uint8
	ba map[uint8]uint8
}

//NewUint8 returns an empty Uint8 bimap.
func NewUint8() *Uint8 {
	return &Uint8{make(map[uint8]uint8), make(map[uint8]uint8)}
}

func (m *Uint8) Put(key, value uint8) *Uint8 {
	m.ab[key] = value
	m.ba[value] = key
	return m
}

func (m *Uint8) GetByKey(key uint8) (value uint8, exists bool) {
	value, exists = m.ab[key]
	return value, exists
}

func (m *Uint8) GetByValue(value uint8) (key uint8, exists bool) {
	key, exists = m.ba[value]
	return key, exists
}
