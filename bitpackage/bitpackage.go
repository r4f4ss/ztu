package bitpackage

import (
	"math"
)

type Pack struct {
	position int
	data     []byte
	dictSize int
}

func NewPack(n *int, dictSize int) *Pack {
	var data []byte
	if n != nil {
		data = make([]byte, *n)
	} else {
		data = make([]byte, 512)
	}

	data = append(data, 0)
	return &Pack{
		position: 0,
		data:     data,
		dictSize: dictSize,
	}
}

func (p *Pack) Packing(code *int, data *byte) {
	if code != nil && data == nil {
		p.packNBit(*code, 1)
		p.packNBit(1, 0)
	} else if code == nil && data != nil {
		p.packNBit(p.dictSize, 1)
		p.packNBit(1, 0)
		p.packByte(*data)
	}
}

func (p *Pack) GetData() []byte {
	return p.data
}

func (p *Pack) packNBit(n int, bit byte) {
	var totalSize float64 = float64((n + p.position)) / 8.0
	size := int(math.Ceil(totalSize))
	mask := make([]byte, size)
	if bit == 0 {
		for i := 0; i < size; i++ {
			mask[i] = 0
		}
	} else {
		for i := 0; i < size; i++ {
			mask[i] = 0xff
		}
	}
	mask[0] = mask[0] >> p.position
	p.data[len(p.data)-1] = mask[0]
	for i := 1; i < len(mask); i++ {
		p.data = append(p.data, mask[i])
	}
	p.position = p.position + n%8
	p.position = p.position % 8
}

func (p *Pack) packByte(by byte) {
	if p.position == 0 {
		p.data = append(p.data, by)
	} else {
		by0 := by >> p.position
		by1 := by << (8 - p.position)
		p.data[len(p.data)-1] = p.data[len(p.data)-1] | by0
		p.data = append(p.data, by1)
	}
}
