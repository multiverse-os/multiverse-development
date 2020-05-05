package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"flag"
	"log"
	"os"
	"path/filepath"

	nt "./font"
	tt "./themetxt"

	gg "github.com/fogleman/gg"
)

var optThemeFile string
var optThemeDir string
var optDraw bool
var optDump bool
var optDrawOutline bool
var optOutput string

var optScreenWidth int
var optScreenHeight int

var globalThemeDir string

func init() {
	flag.StringVar(&optThemeFile, "theme", "", "theme file")
	flag.StringVar(&optThemeDir, "theme-dir", "", "theme dir")
	flag.BoolVar(&optDraw, "draw", false, "draw out.png")
	flag.StringVar(&optOutput, "out", "./out.png", "output image file")
	flag.BoolVar(&optDump, "dump", false, "dump theme")
	flag.BoolVar(&optDrawOutline, "outline", false, "draw outline")

	flag.IntVar(&optScreenWidth, "width", 1366, "screen width (px)")
	flag.IntVar(&optScreenHeight, "height", 768, "screen height (px)")
}

func testMain() {
	ec := newEvalContent()
	ec.setUnknown("screen-width", 500)
	ec.setUnknown("screen-height", 600)

	root := &Node{}
	dc := gg.NewContext(500, 600)

	c1 := &Node{
		left:   tt.AbsNum(0),
		top:    tt.AbsNum(0),
		width:  tt.RelNum(50),
		height: tt.RelNum(50),
	}

	c11 := &Node{
		left:   tt.RelNum(50),
		top:    tt.RelNum(50),
		width:  tt.RelNum(50),
		height: tt.CombinedNum{Rel: 50, Abs: 10, Op: tt.CombinedNumSub},
	}
	c1.addChild(c11)

	root.addChild(c1)

	root.DrawTo(dc, ec)
	dc.SavePNG("./test.png")
	os.Exit(0)
}

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()

	if optThemeDir == "" {
		globalThemeDir = filepath.Dir(optThemeFile)
	} else {
		globalThemeDir = optThemeDir
	}

	theme, err := tt.ParseThemeFile(optThemeFile)
	if err != nil {
		log.Fatal(err)
	}

	if optDump {
		theme.Dump()
	}

	if optDraw {
		loadAllFonts()
		// draw
		draw(theme)
	}
}

const globalFontFile = "/usr/share/fonts/truetype/noto/NotoSans-Regular.ttf"

func draw(theme *tt.Theme) {
	ec := newEvalContent()
	ec.setUnknown("screen-width", float64(optScreenWidth))
	ec.setUnknown("screen-height", float64(optScreenHeight))

	root := themeToNodeTree(theme, optScreenWidth, optScreenHeight)
	ctx := gg.NewContext(optScreenWidth, optScreenHeight)
	// 画背景
	root.draw = func(n *Node, ctx *gg.Context, ec *EvalContext) {
		n.drawImage(ctx, ec, "background.png")
	}

	root.DrawTo(ctx, ec)
	ctx.SavePNG(optOutput)
}

func getResourceFile(name string) string {
	dir := globalThemeDir
	return filepath.Join(dir, name)
}

func themeToNodeTree(theme *tt.Theme, w, h int) *Node {
	root := &Node{}
	for _, comp := range theme.Components {
		if comp.Type == "boot_menu" {
			log.Println("add child boot_menu")
			root.addChild(compBootMenuToNode(comp, root))
		} else if comp.Type == "label" {
			log.Println("add child label")
			root.addChild(compLabelToNode(comp, root))
		}
	}
	return root
}

type CompCommon struct {
	left   tt.Length
	top    tt.Length
	width  tt.Length
	height tt.Length
	id     string

	node *Node
}

func getFont(name string) *font.Face {
	face := getFontAux(name)
	if face == nil {
		panic("not found font face for " + name)
	}
	log.Printf("getFont %q -> %q\n", name, face.Name)
	return face
}

func getFontAux(name string) *font.Face {
	for _, face := range allFontFaces {
		if face.Name == name {
			return face
		}
	}

	fields := strings.Split(name, " ")
	var temp []string
	// match name
	for _, value := range fields {
		if value != "" {
			temp = append(temp, value)
		}
	}
	fields = temp

	if len(fields) >= 3 {
		// family style size
		family := strings.Join(fields[:len(fields)-2], " ")
		sizeStr := fields[len(fields)-1]
		size, _ := strconv.Atoi(sizeStr)

		// match family and size
		for _, face := range allFontFaces {
			if face.Family == family && size == face.PointSize {
				return face
			}
		}

		// match family
		for _, face := range allFontFaces {
			if face.Family == family {
				return face
			}
		}
	}

	if len(fields) >= 2 {

	}

	return nil
}

var allFontFaces []*font.Face

func loadAllFonts() {
	fileInfoList, err := ioutil.ReadDir(globalThemeDir)
	if err != nil {
		log.Println(err)
		return
	}

	for _, fileInfo := range fileInfoList {
		if fileInfo.IsDir() {
			continue
		}

		if filepath.Ext(fileInfo.Name()) == ".pf2" {
			fontFile := filepath.Join(globalThemeDir, fileInfo.Name())
			face, err := font.LoadFont(fontFile)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Printf("load font: %s %q %q\n", fileInfo.Name(), face.Name, face.Family)
			allFontFaces = append(allFontFaces, face)
		}
	}
}
