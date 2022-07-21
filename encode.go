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
	"math"
	"reflect"
	"unsafe"
)

func encodeNil(b *Buffer) {
	b.Write(NilKind)
}

func encodeMap(b *Buffer, size uint32, keyKind, valueKind Kind) {
	b.Write(MapKind)
	b.Write(keyKind)
	b.Write(valueKind)
	encodeUint32(b, size)
}

func encodeSlice(b *Buffer, size uint32, kind Kind) {
	b.Write(SliceKind)
	b.Write(kind)
	encodeUint32(b, size)
}

func encodeBytes(b *Buffer, value []byte) {
	b.Write(BytesKind)
	encodeUint32(b, uint32(len(value)))
	b.Write(value)
}

func encodeString(b *Buffer, value string) {
	var nb []byte
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&nb))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&value))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	b.Write(StringKind)
	encodeUint32(b, uint32(len(nb)))
	b.Write(nb)
}

func encodeError(b *Buffer, err error) {
	b.Write(ErrorKind)
	encodeString(b, err.Error())
}

func encodeBool(b *Buffer, value bool) {
	b.Write(BoolKind)
	if value {
		*b = append(*b, trueBool)
	} else {
		*b = append(*b, falseBool)
	}
}

func encodeUint8(b *Buffer, value uint8) {
	b.Write(Uint8Kind)
	*b = append(*b, value)
}

func encodeUint16(b *Buffer, value uint16) {
	b.Write(Uint16Kind)
	*b = append(*b, byte(value>>8), byte(value))
}

func encodeUint32(b *Buffer, value uint32) {
	b.Write(Uint32Kind)
	*b = append(*b, byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
}

func encodeUint64(b *Buffer, value uint64) {
	b.Write(Uint64Kind)
	*b = append(*b, byte(value>>56), byte(value>>48), byte(value>>40), byte(value>>32), byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
}

func encodeInt32(b *Buffer, value int32) {
	b.Write(Int32Kind)
	castValue := uint32(value)
	*b = append(*b, byte(castValue>>24), byte(castValue>>16), byte(castValue>>8), byte(castValue))
}

func encodeInt64(b *Buffer, value int64) {
	b.Write(Int64Kind)
	castValue := uint64(value)
	*b = append(*b, byte(castValue>>56), byte(castValue>>48), byte(castValue>>40), byte(castValue>>32), byte(castValue>>24), byte(castValue>>16), byte(castValue>>8), byte(castValue))
}

func encodeFloat32(b *Buffer, value float32) {
	b.Write(Float32Kind)
	castValue := math.Float32bits(value)
	*b = append(*b, byte(castValue>>24), byte(castValue>>16), byte(castValue>>8), byte(castValue))
}

func encodeFloat64(b *Buffer, value float64) {
	b.Write(Float64Kind)
	castValue := math.Float64bits(value)
	*b = append(*b, byte(castValue>>56), byte(castValue>>48), byte(castValue>>40), byte(castValue>>32), byte(castValue>>24), byte(castValue>>16), byte(castValue>>8), byte(castValue))
}
