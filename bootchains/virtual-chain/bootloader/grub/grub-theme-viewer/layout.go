package main

import (
	"image"
	"image/color"
	"log"

	font "./font"

	tt "./themetxt"

	gg "github.com/fogleman/gg"
	resize "github.com/nfnt/resize"
)

type Node struct {
	parent   *Node
	Children []*Node

	left   tt.Length
	top    tt.Length
	width  tt.Length
	height tt.Length

	leftExpr   Expr
	topExpr    Expr
	widthExpr  Expr
	heightExpr Expr

	draw func(n *Node, ctx *gg.Context, ec *EvalContext)
}

func getLengthExpr(l tt.Length, val Expr) Expr {
	switch ll := l.(type) {
	case tt.AbsNum:
		return AbsNum(int(ll))
	case tt.RelNum:
		// (val * (ll  / 100))
		return mul(val, div(AbsNum(int(ll)), AbsNum(100)))

	case tt.CombinedNum:
		a := mul(val, div(AbsNum(int(ll.Rel)), AbsNum(100)))
		switch ll.Op {
		case tt.CombinedNumSub:
			return sub(a, AbsNum(ll.Abs))
		case tt.CombinedNumAdd:
			return add(a, AbsNum(ll.Abs))
		}
	}
	panic("not expect")
	return nil
}

func (n *Node) getLeft() Expr {
	if n.parent == nil {
		// root
		return AbsNum(0)
	}

	pl := n.parent.getLeft()
	if n.leftExpr != nil {
		return add(pl, n.leftExpr)
	}

	pw := n.parent.getWidth()
	return add(pl, getLengthExpr(n.left, pw))
}

func (n *Node) getTop() Expr {
	if n.parent == nil {
		// root
		return AbsNum(0)
	}

	pt := n.parent.getTop()
	if n.topExpr != nil {
		return add(pt, n.topExpr)
	}

	ph := n.parent.getHeight()
	return add(pt, getLengthExpr(n.top, ph))
}

func (n *Node) getWidth() Expr {
	if n.widthExpr != nil {
		return n.widthExpr
	}

	if n.parent == nil {
		// root
		return &Unknown{name: "screen-width"}
	}
	pw := n.parent.getWidth()
	return getLengthExpr(n.width, pw)
}

func (n *Node) getHeight() Expr {
	if n.heightExpr != nil {
		return n.heightExpr
	}

	if n.parent == nil {
		// root
		return &Unknown{name: "screen-height"}
	}
	ph := n.parent.getHeight()
	return getLengthExpr(n.height, ph)
}

func (n *Node) addChild(child *Node) {
	child.parent = n
	n.Children = append(n.Children, child)
}

func (n *Node) drawImage(ctx *gg.Context, ec *EvalContext, name string) error {
	img, err := gg.LoadImage(getResourceFile(name))
	if err != nil {
		return err
	}

	x := n.getLeft().Eval(ec)
	y := n.getTop().Eval(ec)
	width := n.getWidth().Eval(ec)
	height := n.getHeight().Eval(ec)

	img = resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	ctx.DrawImage(img, int(x), int(y))
	return nil
}

func loadStyleBoxSlice(name string, part int) (image.Image, error) {
	return gg.LoadImage(getResourceFile(getPixmapName(name, part)))
}

func getPads(name string) (padLeft, padRight, padTop, padBottom int) {
	// nw
	imgNW, _ := loadStyleBoxSlice(name, styleBoxNW)
	if imgNW != nil {
		padLeft = imgNW.Bounds().Dx() // width
		padTop = imgNW.Bounds().Dy()  // height
	}

	// n
	imgN, _ := loadStyleBoxSlice(name, styleBoxN)
	if imgN != nil {
		if padTop == 0 {
			padTop = imgN.Bounds().Dy() // height
		}
	}

	// ne
	imgNE, _ := loadStyleBoxSlice(name, styleBoxNE)
	if imgNE != nil {
		if padTop == 0 {
			padTop = imgNE.Bounds().Dy() // height
		}

		padRight = imgNE.Bounds().Dx() // width
	}

	// w
	imgW, _ := loadStyleBoxSlice(name, styleBoxW)
	if imgW != nil {
		if padLeft == 0 {
			padLeft = imgW.Bounds().Dx() // width
		}
	}

	// e
	imgE, _ := loadStyleBoxSlice(name, styleBoxE)
	if imgE != nil {
		if padRight == 0 {
			padRight = imgE.Bounds().Dx() // width
		}
	}

	// sw
	imgSW, _ := loadStyleBoxSlice(name, styleBoxSW)
	if imgSW != nil {
		if padLeft == 0 {
			padLeft = imgSW.Bounds().Dx() // width
		}

		padBottom = imgSW.Bounds().Dy() // height
	}

	// s
	imgS, _ := loadStyleBoxSlice(name, styleBoxS)
	if imgS != nil {
		if padBottom == 0 {
			padBottom = imgS.Bounds().Dy() // height
		}
	}

	// se
	imgSE, _ := loadStyleBoxSlice(name, styleBoxSE)
	if imgSE != nil {
		if padBottom == 0 {
			padBottom = imgSE.Bounds().Dy() // height
		}
		if padRight == 0 {
			padRight = imgSE.Bounds().Dx() // width
		}
	}
	return
}

func (n *Node) drawStyleBox(ctx *gg.Context, ec *EvalContext, name string) {
	if name == "" {
		return
	}

	x := int(n.getLeft().Eval(ec))
	y := int(n.getTop().Eval(ec))
	width := int(n.getWidth().Eval(ec))
	height := int(n.getHeight().Eval(ec))

	color1 := "#f9f806"
	color2 := "#f97306"

	var padLeft int
	var padRight int
	var padTop int
	var padBottom int

	// nw
	imgNW, _ := loadStyleBoxSlice(name, styleBoxNW)
	if imgNW != nil {
		padLeft = imgNW.Bounds().Dx() // width
		padTop = imgNW.Bounds().Dy()  // height
	}

	// n
	imgN, _ := loadStyleBoxSlice(name, styleBoxN)
	if imgN != nil {
		if padTop == 0 {
			padTop = imgN.Bounds().Dy() // height
		}
	}

	// ne
	imgNE, _ := loadStyleBoxSlice(name, styleBoxNE)
	if imgNE != nil {
		if padTop == 0 {
			padTop = imgNE.Bounds().Dy() // height
		}

		padRight = imgNE.Bounds().Dx() // width
	}

	// w
	imgW, _ := loadStyleBoxSlice(name, styleBoxW)
	if imgW != nil {
		if padLeft == 0 {
			padLeft = imgW.Bounds().Dx() // width
		}
	}

	// c
	imgC, _ := loadStyleBoxSlice(name, styleBoxC)

	// e
	imgE, _ := loadStyleBoxSlice(name, styleBoxE)
	if imgE != nil {
		if padRight == 0 {
			padRight = imgE.Bounds().Dx() // width
		}
	}

	// sw
	imgSW, _ := loadStyleBoxSlice(name, styleBoxSW)
	if imgSW != nil {
		if padLeft == 0 {
			padLeft = imgSW.Bounds().Dx() // width
		}

		padBottom = imgSW.Bounds().Dy() // height
	}

	// s
	imgS, _ := loadStyleBoxSlice(name, styleBoxS)
	if imgS != nil {
		if padBottom == 0 {
			padBottom = imgS.Bounds().Dy() // height
		}
	}

	// se
	imgSE, _ := loadStyleBoxSlice(name, styleBoxSE)
	if imgSE != nil {
		if padBottom == 0 {
			padBottom = imgSE.Bounds().Dy() // height
		}
		if padRight == 0 {
			padRight = imgSE.Bounds().Dx() // width
		}
	}

	// ---
	// draw
	// nw
	if imgNW != nil {
		ctx.DrawImage(imgNW, x, y)

		if optDrawOutline {
			ctx.SetHexColor(color1)
			ctx.DrawRectangle(float64(x), float64(y),
				float64(imgNW.Bounds().Dx()), float64(imgNW.Bounds().Dy()))
			ctx.Stroke()
		}
	}

	// n
	if imgN != nil {
		imgN = resize.Resize(uint(width-padLeft-padRight), uint(padTop),
			imgN, resize.Lanczos3)
		ctx.DrawImage(imgN, x+padLeft, y)

		if optDrawOutline {
			ctx.SetHexColor(color2)
			ctx.DrawRectangle(float64(x+padLeft), float64(y),
				float64(width-padLeft-padRight), float64(padTop))
			ctx.Stroke()
		}
	}

	// ne
	if imgNE != nil {
		ctx.DrawImage(imgNE, x+width-padRight, y)

		if optDrawOutline {
			ctx.SetHexColor(color1)
			ctx.DrawRectangle(float64(x+width-padRight), float64(y),
				float64(imgNE.Bounds().Dx()), float64(imgNE.Bounds().Dy()))
			ctx.Stroke()
		}
	}

	// w
	if imgW != nil {
		imgW = resize.Resize(uint(padLeft), uint(height-padTop-padBottom),
			imgW, resize.Lanczos3)
		ctx.DrawImage(imgW, x, y+padTop)

		if optDrawOutline {
			ctx.SetHexColor(color2)
			ctx.DrawRectangle(float64(x), float64(y+padTop),
				float64(padLeft), float64(height-padTop-padBottom))
			ctx.Stroke()
		}
	}

	// c
	if imgC != nil {
		imgC = resize.Resize(uint(width-padLeft-padRight),
			uint(height-padTop-padBottom),
			imgC, resize.Lanczos3)
		ctx.DrawImage(imgC, x+padLeft, y+padTop)

		if optDrawOutline {
			ctx.SetHexColor(color1)
			ctx.DrawRectangle(float64(x+padLeft), float64(y+padTop),
				float64(width-padLeft-padRight),
				float64(height-padTop-padBottom))
			ctx.Stroke()
		}
	}

	// e
	if imgE != nil {
		imgE = resize.Resize(uint(padRight), uint(height-padTop-padBottom),
			imgE, resize.Lanczos3)
		ctx.DrawImage(imgE, x+width-padRight, y+padTop)

		if optDrawOutline {
			ctx.SetHexColor(color2)
			ctx.DrawRectangle(float64(x+width-padRight), float64(y+padTop),
				float64(padRight), float64(height-padTop-padBottom))
			ctx.Stroke()
		}
	}

	// sw
	if imgSW != nil {
		ctx.DrawImage(imgSW, x, y+height-padBottom)

		if optDrawOutline {
			ctx.SetHexColor(color1)
			ctx.DrawRectangle(float64(x), float64(y+height-padBottom),
				float64(imgSW.Bounds().Dx()), float64(imgSW.Bounds().Dy()))
			ctx.Stroke()
		}
	}

	// s
	if imgS != nil {
		imgS = resize.Resize(uint(width-padLeft-padRight),
			uint(padBottom),
			imgS, resize.Lanczos3)
		ctx.DrawImage(imgS, x+padLeft, y+height-padBottom)

		if optDrawOutline {
			ctx.SetHexColor(color2)
			ctx.DrawRectangle(float64(x+padLeft), float64(y+height-padBottom),
				float64(width-padLeft-padRight),
				float64(padBottom))
			ctx.Stroke()
		}
	}

	// se
	if imgSE != nil {
		ctx.DrawImage(imgSE, x+width-padRight, y+height-padBottom)

		if optDrawOutline {
			ctx.SetHexColor(color1)
			ctx.DrawRectangle(float64(x+width-padRight), float64(y+height-padBottom),
				float64(imgSE.Bounds().Dx()), float64(imgSE.Bounds().Dy()))
			ctx.Stroke()
		}
	}
}

func (n *Node) drawText(ctx *gg.Context, ec *EvalContext, str string, color color.Color, fontFace *font.Face) {
	x := n.getLeft().Eval(ec)
	y := n.getTop().Eval(ec)
	ctx.SetColor(color)

	ctx.SetFontFace(fontFace)
	log.Printf("drawText str: %q, x: %g, y: %g\n", str, x, y)
	ctx.DrawStringAnchored(str, x, y, 0, 1)
}

func (n *Node) drawText1(ctx *gg.Context, ec *EvalContext, str string, color color.Color, fontFace *font.Face, width float64, align gg.Align) {
	x := n.getLeft().Eval(ec)
	y := n.getTop().Eval(ec)
	ctx.SetColor(color)

	ctx.SetFontFace(fontFace)
	log.Printf("drawText1 str: %q, x: %g, y: %g, width: %g\n", str, x, y, width)
	ctx.DrawStringWrapped(str, x, y, 0, 0, width, 1, align)
}

func (n *Node) DrawTo(ctx *gg.Context, ec *EvalContext) {
	if optDrawOutline {
		x := n.getLeft().Eval(ec)
		y := n.getTop().Eval(ec)
		w := n.getWidth().Eval(ec)
		h := n.getHeight().Eval(ec)
		log.Printf("drawOutline x: %g, y: %g, w: %g, h: %g\n", x, y, w, h)
		ctx.DrawRectangle(x, y, w, h)
		ctx.SetRGB(1, 0, 0)
		ctx.Stroke()
	}

	if n.draw != nil {
		n.draw(n, ctx, ec)
	}

	for _, c := range n.Children {
		c.DrawTo(ctx, ec)
	}
}
