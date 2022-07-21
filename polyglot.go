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

type Kind []byte

var (
	NilKind     = Kind([]byte{0})
	SliceKind   = Kind([]byte{1})
	MapKind     = Kind([]byte{2})
	AnyKind     = Kind([]byte{3})
	BytesKind   = Kind([]byte{4})
	StringKind  = Kind([]byte{5})
	ErrorKind   = Kind([]byte{6})
	BoolKind    = Kind([]byte{7})
	Uint8Kind   = Kind([]byte{8})
	Uint16Kind  = Kind([]byte{9})
	Uint32Kind  = Kind([]byte{10})
	Uint64Kind  = Kind([]byte{11})
	Int32Kind   = Kind([]byte{12})
	Int64Kind   = Kind([]byte{13})
	Float32Kind = Kind([]byte{14})
	Float64Kind = Kind([]byte{15})
)

type Error string

func (e Error) Error() string {
	return string(e)
}

func (e Error) Is(err error) bool {
	return e.Error() == err.Error()
}

var (
	falseBool = byte(0)
	trueBool  = byte(1)
)
