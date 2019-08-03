package bimap

//bimap is a bi-directional map. Essentially just a struct with 2 maps, with one being a reverse of the other.
//Credit for the idea (and code) comes from this thread: https://stackoverflow.com/questions/23190311/reverse-a-map-in-value-key-format-in-golang

//Uint8 is a bimap with both key and value pairs being of type uint8.
type Uint8 struct {
	ab map[uint8]uint8
	ba map[uint8]uint8
}

//NewUint8 returns an empty Uint8 bimap.
func NewUint8() *Uint8 {
	return &Uint8{make(map[uint8]uint8), make(map[uint8]uint8)}
}

//Put will insert a key value pair to the Uint8 bimap.
func (m *Uint8) Put(key, value uint8) *Uint8 {
	m.ab[key] = value
	m.ba[value] = key
	return m
}

//GetByKey will return the value from the given key.
func (m *Uint8) GetByKey(key uint8) (value uint8, exists bool) {
	value, exists = m.ab[key]
	return value, exists
}

//GetByValue will return the key from the given value.
func (m *Uint8) GetByValue(value uint8) (key uint8, exists bool) {
	key, exists = m.ba[value]
	return key, exists
}
