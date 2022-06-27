/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package polyglot

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestEncoderNil(t *testing.T) {
	t.Parallel()

	p := CNew()

	Encoder(p).Nil()

	assert.Equal(t, 1, len(*p))
	assert.Equal(t, NilKind, Kind(*p))

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Nil()
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderMap(t *testing.T) {
	t.Parallel()

	p := CNew()
	m := make(map[string]uint32)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3

	e := Encoder(p).Map(uint32(len(m)), StringKind, Uint32Kind)
	for k, v := range m {
		e.String(k).Uint32(v)
	}

	assert.Equal(t, 1+1+1+1+4+len(m)*(1+1+4+1+1+4), len(*p))

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		e = Encoder(p).Map(uint32(len(m)), StringKind, Uint32Kind)
		for k, v := range m {
			e.String(k).Uint32(v)
		}
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderSlice(t *testing.T) {
	t.Parallel()

	p := CNew()
	m := make(map[string]uint32)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3

	e := Encoder(p).Map(uint32(len(m)), StringKind, Uint32Kind)
	for k, v := range m {
		e.String(k).Uint32(v)
	}

	assert.Equal(t, 1+1+1+1+4+len(m)*(1+1+4+1+1+4), len(*p))

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		e = Encoder(p).Map(uint32(len(m)), StringKind, Uint32Kind)
		for k, v := range m {
			e.String(k).Uint32(v)
		}
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderBytes(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := []byte("Test String")

	Encoder(p).Bytes(v)

	assert.Equal(t, 1+1+4+len(v), len(*p))
	assert.Equal(t, (C)(v), (*p)[1+1+4:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Bytes(v)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderString(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := "Test String"
	e := []byte(v)

	Encoder(p).String(v)

	assert.Equal(t, 1+1+4+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1+1+4:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).String(v)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderError(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := errors.New("Test String")
	e := []byte(v.Error())

	Encoder(p).Error(v)

	assert.Equal(t, 1+1+1+4+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1+1+1+4:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Error(v)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderBool(t *testing.T) {
	t.Parallel()

	p := CNew()
	e := []byte{trueBool}

	Encoder(p).Bool(true)

	assert.Equal(t, 1+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Bool(true)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderUint8(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := uint8(32)
	e := []byte{v}

	Encoder(p).Uint8(v)

	assert.Equal(t, 1+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Uint8(v)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderUint16(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := uint16(1024)
	e := []byte{byte(v >> 8), byte(v)}

	Encoder(p).Uint16(v)

	assert.Equal(t, 1+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Uint16(v)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderUint32(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := uint32(4294967290)
	e := []byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}

	Encoder(p).Uint32(v)

	assert.Equal(t, 1+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Uint32(v)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderUint64(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := uint64(18446744073709551610)
	e := []byte{byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32), byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}

	Encoder(p).Uint64(v)

	assert.Equal(t, 1+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Uint64(v)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderInt32(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := int32(-2147483648)
	e := []byte{byte(uint32(v) >> 24), byte(uint32(v) >> 16), byte(uint32(v) >> 8), byte(uint32(v))}

	Encoder(p).Int32(v)

	assert.Equal(t, 1+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Int32(v)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderInt64(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := int64(-9223372036854775808)
	e := []byte{byte(uint64(v) >> 56), byte(uint64(v) >> 48), byte(uint64(v) >> 40), byte(uint64(v) >> 32), byte(uint64(v) >> 24), byte(uint64(v) >> 16), byte(uint64(v) >> 8), byte(uint64(v))}

	Encoder(p).Int64(v)

	assert.Equal(t, 1+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Int64(v)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderFloat32(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := float32(-214648.34432)
	e := []byte{byte(math.Float32bits(v) >> 24), byte(math.Float32bits(v) >> 16), byte(math.Float32bits(v) >> 8), byte(math.Float32bits(v))}

	Encoder(p).Float32(v)

	assert.Equal(t, 1+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Float32(v)
		p.Reset()
	})
	assert.Zero(t, n)
}

func TestEncoderFloat64(t *testing.T) {
	t.Parallel()

	p := CNew()
	v := -922337203685.2345
	e := []byte{byte(math.Float64bits(v) >> 56), byte(math.Float64bits(v) >> 48), byte(math.Float64bits(v) >> 40), byte(math.Float64bits(v) >> 32), byte(math.Float64bits(v) >> 24), byte(math.Float64bits(v) >> 16), byte(math.Float64bits(v) >> 8), byte(math.Float64bits(v))}

	Encoder(p).Float64(v)

	assert.Equal(t, 1+len(e), len(*p))
	assert.Equal(t, (C)(e), (*p)[1:])

	p.Reset()
	n := testing.AllocsPerRun(100, func() {
		Encoder(p).Float64(v)
		p.Reset()
	})
	assert.Zero(t, n)
}