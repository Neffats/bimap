package bimap

import "testing"

func TestPut(t *testing.T) {
	bi := NewUint8()

	var PutTests = []struct {
		name  string
		key   uint8
		value uint8
	}{
		{name: "key/value: 3/3", key: 3, value: 3},
		{name: "key/value: 0xF/3", key: 0xF, value: 3},
		{name: "key/value: 7/0xE", key: 7, value: 0xE},
	}

	for _, tc := range PutTests {
		t.Run(tc.name, func(t *testing.T) {
			bi.Put(tc.key, tc.value)
			if bi.ab[tc.key] != tc.value {
				t.Errorf("expected key/value: %d/%d, but got value: %d", tc.key, tc.value, bi.ab[tc.key])
			}
			if bi.ba[tc.value] != tc.key {
				t.Errorf("expected key/value: %d/%d, but got key: %d", tc.key, tc.value, bi.ab[tc.key])
			}
		})
	}
}

func TestGetByKey(t *testing.T) {
	bi := NewUint8()

	var GetByKeyTests = []struct {
		name   string
		key    uint8
		value  uint8
		arg    uint8
		exsits bool
	}{
		{name: " key/value/arg: 3/3/3", key: 3, value: 3, arg: 3, exsits: true},
		{name: " key/value/arg: 0xF/3/5", key: 0xF, value: 3, arg: 5, exsits: false},
		{name: " key/value/arg: 7/0xE/7", key: 7, value: 0xE, arg: 7, exsits: true},
	}

	for _, tc := range GetByKeyTests {
		t.Run(tc.name, func(t *testing.T) {
			bi.Put(tc.key, tc.value)
			r, exists := bi.GetByKey(tc.arg)
			if exists != tc.exsits {
				t.Errorf("exists: expected: %t got: %t", tc.exsits, exists)
			}
			if exists == false {
				return
			}
			if r != tc.value {
				t.Errorf("expected key/value: %d/%d, but got value: %d", tc.key, tc.value, r)
			}
		})
	}
}

func TestGetByValue(t *testing.T) {
	bi := NewUint8()

	var GetByValueTests = []struct {
		name   string
		key    uint8
		value  uint8
		arg    uint8
		exsits bool
	}{
		{name: " key/value/arg: 3/3/3", key: 3, value: 3, arg: 3, exsits: true},
		{name: " key/value/arg: 0xF/3/5", key: 0xF, value: 3, arg: 5, exsits: false},
		{name: " key/value/arg: 7/0xE/7", key: 7, value: 0xE, arg: 0xE, exsits: true},
	}

	for _, tc := range GetByValueTests {
		t.Run(tc.name, func(t *testing.T) {
			bi.Put(tc.key, tc.value)
			r, exists := bi.GetByValue(tc.arg)
			if exists != tc.exsits {
				t.Errorf("exists: expected: %t got: %t", tc.exsits, exists)
			}
			if exists == false {
				return
			}
			if r != tc.key {
				t.Errorf("expected key/value: %d/%d, but got key: %d", tc.key, tc.value, r)
			}
		})
	}
}
