// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parts

import (
	"gxui"
	"gxui/mixins/outer"
)

type PaintChildrenOuter interface {
	gxui.Container
	outer.Bounds
	outer.PaintChilder
}

type PaintChildren struct {
	outer PaintChildrenOuter
}

func (p *PaintChildren) Init(outer PaintChildrenOuter) {
	p.outer = outer
}

func (p *PaintChildren) Paint(c gxui.Canvas) {
	for i, v := range p.outer.Children() {
		if v.IsVisible() {
			c.Push()
			c.AddClip(v.Bounds())
			p.outer.PaintChild(c, v, i)
			c.Pop()
		}
	}
}

func (p *PaintChildren) PaintChild(c gxui.Canvas, child gxui.Control, idx int) {
	childCanvas := child.Draw()
	if childCanvas != nil {
		c.DrawCanvas(childCanvas, child.Bounds().Min)
	}
}