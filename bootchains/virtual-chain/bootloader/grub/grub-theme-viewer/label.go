package main

import (
	"fmt"
	"image/color"
	"strings"

	tt "github.com/electricface/grub-theme-viewer/themetxt"
	"github.com/fogleman/gg"
)

type Label struct {
	CompCommon
	text    string
	font    string
	color   string
	align   string
	visible bool
}

func newLabel(comp *tt.Component) *Label {
	l := &Label{}
	l.node = &Node{}
	l.fillCommonOptions(comp)

	var ok bool
	l.visible, ok = comp.GetPropBool("visible")
	if !ok {
		l.visible = true
	}

	l.text, _ = comp.GetPropString("text")

	l.font, ok = comp.GetPropString("font")
	if !ok {
		l.font = "Unknown Regular 16"
	}

	l.color, ok = comp.GetPropString("color")
	if !ok {
		l.color = "black"
	}

	l.align, ok = comp.GetPropString("align")
	if !ok {
		l.align = "left"
	}

	return l
}

func (l *Label) getColor() color.Color {
	return parseColor(l.color)
}

func parseAlign(str string) gg.Align {
	switch str {
	case "left":
		return gg.AlignLeft
	case "right":
		return gg.AlignRight
	case "center":
		return gg.AlignCenter
	}
	return gg.AlignLeft
}

func (l *Label) getText() string {
	var text string
	text = l.text
	if l.id == "__timeout__" {
		if strings.Contains(l.text, "%d") {
			text = fmt.Sprintf(l.text, 5)
		}
	}
	return text
}

func (l *Label) getAlign() gg.Align {
	return parseAlign(l.align)
}

func compLabelToNode(comp *tt.Component, parent *Node) *Node {
	label := newLabel(comp)
	label.node.draw = func(n *Node, ctx *gg.Context, ec *EvalContext) {

		fontFace := getFont(label.font)
		width := n.getWidth().Eval(ec)
		n.drawText1(ctx, ec, label.getText(), label.getColor(), fontFace,
			width, label.getAlign())
	}
	return label.node
}
