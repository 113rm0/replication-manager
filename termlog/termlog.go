// replication-manager - Replication Manager Monitoring and CLI for MariaDB
// Authors: Guillaume Lefranc <guillaume.lefranc@mariadb.com>
//          Stephane Varoqui  <stephane.varoqui@mariadb.com>
// This source code is licensed under the GNU General Public License, version 3.
// Redistribution/Reuse of this code is permitted under the GNU v3 license, as
// an additional term, ALL code must carry the original Author(s) credit in comment form.
// See LICENSE in this directory for the integral text.

// Package termlog
// termbox logging package
package termlog

import (
	"time"

	"github.com/nsf/termbox-go"
)

type TermLog struct {
	Buffer []string
	Len    int
	Line   int
}

func NewTermLog(sz int) TermLog {
	tl := TermLog{}
	tl.Len = sz
	tl.Buffer = make([]string, tl.Len)
	return tl
}

func (tl *TermLog) Add(s string) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	s = " " + ts + " " + s
	tl.Shift(s)
}

func (tl *TermLog) Shift(e string) {
	ns := make([]string, 1)
	ns[0] = e
	tl.Buffer = append(ns, tl.Buffer[0:tl.Len]...)
}

func (tl *TermLog) Extend() {
	tl.Buffer = append(tl.Buffer, make([]string, tl.Len)...)
}

func (tl *TermLog) Shrink() {
	tl.Buffer = tl.Buffer[:tl.Len]
}

func (tl TermLog) Print() {
	for _, line := range tl.Buffer {
		x := 0
		for _, c := range line {
			termbox.SetCell(x, tl.Line, c, termbox.ColorWhite, termbox.ColorBlack)
			x++
		}
		tl.Line++
	}
}
