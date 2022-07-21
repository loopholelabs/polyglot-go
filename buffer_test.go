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
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	p := *NewBuffer()

	b := make([]byte, 32)
	_, err := rand.Read(b)
	assert.NoError(t, err)

	p.Write(b)
	assert.EqualValues(t, b, p)

	p.Reset()
	assert.NotEqual(t, b, p)
	assert.Equal(t, 0, len(p))
	assert.Equal(t, 512, cap(p))

	b = make([]byte, 1024)
	_, err = rand.Read(b)
	assert.NoError(t, err)

	p.Write(b)

	assert.EqualValues(t, b, p)
	assert.Equal(t, 1024, len(p))
	assert.GreaterOrEqual(t, cap(p), 1024)

}
