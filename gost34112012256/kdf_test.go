// GoGOST -- Pure Go GOST cryptographic functions library
// Copyright (C) 2015-2023 Sergey Matveev <stargrave@stargrave.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 of the License.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package gost34112012256

import (
	"bytes"
	"testing"
)

func TestKDFGOSTR34112012256(t *testing.T) {
	kdf := NewKDF([]byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
	})
	derived := kdf.Derive(
		nil,
		[]byte{0x26, 0xbd, 0xb8, 0x78},
		[]byte{0xaf, 0x21, 0x43, 0x41, 0x45, 0x65, 0x63, 0x78},
	)
	if !bytes.Equal(derived, []byte{
		0xa1, 0xaa, 0x5f, 0x7d, 0xe4, 0x02, 0xd7, 0xb3,
		0xd3, 0x23, 0xf2, 0x99, 0x1c, 0x8d, 0x45, 0x34,
		0x01, 0x31, 0x37, 0x01, 0x0a, 0x83, 0x75, 0x4f,
		0xd0, 0xaf, 0x6d, 0x7c, 0xd4, 0x92, 0x2e, 0xd9,
	}) {
		t.FailNow()
	}
}
