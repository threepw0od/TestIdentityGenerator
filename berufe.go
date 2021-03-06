// Test Identity Generator
// Copyright (C) 2019  Florian Probst
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

func holeBerufe(filename string) (berufe []string, err error) {
	berufe = make([]string, 0, 0)

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return berufe, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		line := scanner.Text()
		if !utf8.ValidString(line) {
			fmt.Println("Berufe: no valid UTF8!")
		}

		if strings.HasPrefix(line, "    ") {
			beruf := line[4:]
			berufe = append(berufe, beruf)
		}
	}

	if err := scanner.Err(); err != nil {
		return berufe, err
	}

	return berufe, nil
}
