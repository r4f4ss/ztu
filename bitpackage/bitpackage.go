package bitpackage

import (
	"math"
)

type Pack struct {
	positionIn  int
	positionOut int
	indexOut    int
	data        []byte
	dictSize    int
}

func NewPack(dictSize int, data []byte) *Pack {
	var dataP []byte
	if data == nil {
		dataP = append(dataP, 0)
	} else {
		dataP = data
	}
	return &Pack{
		positionIn:  0,
		positionOut: 0,
		indexOut:    0,
		data:        dataP,
		dictSize:    dictSize,
	}
}

func (p *Pack) Packing(code *int, data *byte) {
	if code != nil && data == nil {
		p.packNBit(*code+1, 1)
		p.packNBit(1, 0)
	} else if code == nil && data != nil {
		p.packNBit(p.dictSize+1, 1)
		p.packNBit(1, 0)
		p.packByte(*data)
	}
}

func (p *Pack) UnpackingNext() (*int, *byte) {
	code := p.getNextCode()
	var data *byte = nil
	if code == nil {
		return nil, nil
	} else if *code == p.dictSize {
		data = p.getNextByte()
	}
	return code, data
}

func (p *Pack) GetData() []byte {
	return p.data
}

func (p *Pack) packNBit(n int, bit byte) {
	if p.positionIn == 8 {
		p.data = append(p.data, 0)
		p.positionIn = 0
	}
	var totalSize float64 = float64((n + p.positionIn)) / 8.0
	size := int(math.Ceil(totalSize))
	mask := make([]byte, size)
	mask[0] = 0xff
	if bit == 0 {
		for i := 1; i < size; i++ {
			mask[i] = 0
		}
		mask[0] = mask[0] << (8 - p.positionIn)
		p.data[len(p.data)-1] = p.data[len(p.data)-1] & mask[0]
	} else {
		for i := 1; i < size; i++ {
			mask[i] = 0xff
		}
		mask[0] = mask[0] >> p.positionIn
		p.data[len(p.data)-1] = p.data[len(p.data)-1] | mask[0]
	}
	for i := 1; i < len(mask); i++ {
		p.data = append(p.data, mask[i])
	}
	p.positionIn = (p.positionIn + n) % 8
	if p.positionIn == 0 {
		p.data = append(p.data, 0)
	}
}

func (p *Pack) packByte(by byte) {
	if p.positionIn == 0 {
		p.data[len(p.data)-1] = by
		p.positionIn = 8
	} else {
		by0 := by >> p.positionIn
		by1 := by << (8 - p.positionIn)
		p.data[len(p.data)-1] = p.data[len(p.data)-1] | by0
		p.data = append(p.data, by1)
	}
}

func (p *Pack) getNextCode() *int {
	if p.indexOut >= len(p.data) || (p.indexOut == len(p.data) && p.positionOut == 7) {
		return nil
	}

	var code int = 0
	var mask byte = 0x01
	for b := (p.data[p.indexOut] >> (8 - p.positionOut - 1)) & mask; b != 0; b = (p.data[p.indexOut] >> (8 - p.positionOut - 1)) & mask {
		code++
		p.positionOut++
		if p.positionOut == 8 {
			p.indexOut++
			p.positionOut = 0
		}
	}
	p.positionOut++
	if p.positionOut == 8 {
		p.indexOut++
		p.positionOut = 0
	}
	code--
	return &code
}

func (p *Pack) getNextByte() *byte {
	var dataB byte = p.data[p.indexOut]
	if p.positionOut != 0 {
		var dataB2 byte = p.data[p.indexOut+1]
		dataB = dataB << p.positionOut
		dataB2 = dataB2 >> (8 - p.positionOut)
		dataB = dataB | dataB2
	}
	p.indexOut++
	return &dataB
}
