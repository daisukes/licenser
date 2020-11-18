// Copyright 2019 Liam White
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package license

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"text/template"
)

var _ Handler = &MIT{}

// NewMIT creates a new MIT license handler
func NewMIT(year int, owner string) *MIT {
	return &MIT{Year: year, Owner: owner}
}

// MIT is an MIT license handler
type MIT struct {
	Year  int
	Owner string

	licenseCache []byte
}

// Reader returns a reader populated with the MIT license file prefix
func (a *MIT) Reader() io.Reader {
	return bytes.NewReader(a.bytes())
}

// IsPresent verifies that an MIT license is present in the reader passed.
func (a *MIT) IsPresent(in io.Reader) bool {
	inScanner := bufio.NewScanner(in)
	// Check for presence of license in first 20 lines
	for i := 0; i < 20; i++ {
		if inScanner.Scan() {
			// We should definitely be more thorough here but this will do for now
			if strings.Contains(inScanner.Text(), "Permission is hereby granted, free of charge, to any person obtaining a copy") {
				return true
			}
		}
	}
	return false
}

func (a *MIT) SetOwner(in string) {
	a.Owner = in
}

func (a *MIT) bytes() []byte {
	if a.licenseCache != nil {
		return copyBytes(a.licenseCache)
	}
	tmpl, _ := template.New("mit").Parse(mitTemplate)
	b := bytes.NewBuffer([]byte{})
	_ = tmpl.Execute(b, a)
	a.licenseCache = b.Bytes()
	return copyBytes(a.licenseCache)
}

const mitTemplate = `Copyright (c) {{.Year}}  {{.Owner}}

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`
