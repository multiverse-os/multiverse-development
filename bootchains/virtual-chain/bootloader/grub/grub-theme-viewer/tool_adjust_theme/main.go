package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	font "./font"
	tt "./themetxt"

	locale "pkg.deepin.io/lib/locale"
	resize "github.com/nfnt/resize"

)

var optScreenHeight int
var optScreenWidth int
var optThemeDir string
var optLang string

func init() {
	flag.IntVar(&optScreenWidth, "width", 1366, "")
	flag.IntVar(&optScreenHeight, "height", 768, "")
	flag.StringVar(&optThemeDir, "theme-dir", "", "")
	flag.StringVar(&optLang, "lang", "en", "language")
}

func adjustBackground(theme *tt.Theme) {
	desktopImageFile, _ := theme.GetPropString("desktop-image")
	ext := filepath.Ext(desktopImageFile)
	originDesktopImageFile := strings.TrimSuffix(desktopImageFile, ext) + ".origin" + ext
	img, err := loadImage(filepath.Join(optThemeDir, originDesktopImageFile))
	if err != nil {
		log.Fatal(err)
	}
	img = resize.Resize(uint(optScreenWidth), uint(optScreenHeight), img, resize.Lanczos3)

	// save img
	err = savePng(img, filepath.Join(optThemeDir, desktopImageFile))
	if err != nil {
		log.Fatal(err)
	}
}

func adjustResourcesOsLogos(width, height int) {
	dir := filepath.Join(optThemeDir, "resources/os-logos")
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("WARN:", err)
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		file := filepath.Join(dir, fileInfo.Name())
		ext := filepath.Ext(fileInfo.Name())
		if ext != ".svg" {
			continue
		}
		outFileName := strings.TrimSuffix(fileInfo.Name(), ext) + ".png"
		outFile := filepath.Join(optThemeDir, "icons", outFileName)
		err = convertSvg(file, outFile, width, height)
		if err != nil {
			log.Println("WARN:", err)
		}
	}
}

func convertSvg(svgFile, outFile string, width, height int) error {
	cmd := exec.Command("rsvg-convert", "-o", outFile,
		"-w", strconv.Itoa(width),
		"-h", strconv.Itoa(height),
		"-f", "png",
		svgFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("$ rsvg-convert --output %s -width %d -height %d -f png %s\n",
		outFile, width, height, svgFile)
	return cmd.Run()
}

func loadImage(filename string) (image.Image, error) {
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	br := bufio.NewReader(fh)
	img, _, err := image.Decode(br)
	return img, err
}

func savePng(img image.Image, filename string) error {
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	bw := bufio.NewWriter(fh)
	err = png.Encode(bw, img)
	if err != nil {
		return err
	}
	err = bw.Flush()
	return err
}

// min 16px
func getFontSize(screenWidth int, screenHeight int) int {
	var x1 float64 = 768
	var y1 float64 = 16
	var x2 float64 = 2000
	var y2 float64 = 32
	y := (float64(screenHeight)-x1)/(x2-x1)*(y2-y1) + y1

	if y < 16 {
		y = 16
	}

	return round(y)
}

// copy from go source
func round(f float64) int {
	i := int(f)
	if f-float64(i) >= 0.5 {
		i += 1
	}
	return i
}

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	vars := map[string]float64{}

	themeFile := filepath.Join(optThemeDir, "theme.txt.tpl")
	theme, err := tt.ParseThemeFile(themeFile)
	if err != nil {
		log.Fatal(err)
	}

	stdFontSize := getFontSize(optScreenWidth, optScreenHeight)
	vars["std_font_size"] = float64(stdFontSize)
	vars["screen_width"] = float64(optScreenWidth)
	vars["screen_height"] = float64(optScreenHeight)

	adjustBackground(theme)

	for _, comp := range theme.Components {
		if comp.Type == tt.ComponentTypeBootMenu {
			adjustBootMenu(comp, vars)

			iconWidth0, _ := comp.GetProp("icon_width")
			iconWidth := iconWidth0.(int)
			iconHeight0, _ := comp.GetProp("icon_height")
			iconHeight := iconHeight0.(int)
			adjustResourcesOsLogos(iconWidth, iconHeight)

		} else if comp.Type == tt.ComponentTypeLabel {
			adjustLabel(comp, vars)
		}
	}

	themeOutput := filepath.Join(optThemeDir, "theme.txt")
	themeOutputFh, err := os.Create(themeOutput)
	if err != nil {
		log.Fatal(err)
	}
	defer themeOutputFh.Close()
	bw := bufio.NewWriter(themeOutputFh)
	theme.WriteTo(bw)
	bw.Flush()
}

func genFont(fontFile string, size int) (*font.Face, error) {
	// TODO cache support
	sizeStr := strconv.Itoa(size)

	fontBaseName := filepath.Base(fontFile)
	// trim ext
	fontBaseName = strings.TrimSuffix(fontBaseName, filepath.Ext(fontBaseName))
	fontBaseName = fmt.Sprintf("%s-%d.pf2", fontBaseName, size)
	output := filepath.Join(optThemeDir, fontBaseName)

	cmd := exec.Command("grub-mkfont", fontFile, "-s", sizeStr, "-o", output)
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	face, err := font.LoadFont(output)
	return face, err
}

func parseTplFont(str string) (fontFile string, sizeScale float64, err error) {
	fields := strings.SplitN(str, ";", 2)
	if len(fields) != 2 {
		return "", 0, errors.New("invalid font format")
	}
	fontFile = filepath.Join(optThemeDir, "fonts", fields[0])
	sizeScale, err = strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return "", 0, err
	}
	return fontFile, sizeScale, nil
}

func adjustFont(comp *tt.Component, propName string, vars map[string]float64) (*font.Face, error) {
	propFont, _ := comp.GetPropString(propName)
	fontFile, sizeScale, err := parseTplFont(propFont)
	if err != nil {
		return nil, err
	}

	fontSize := round(vars["std_font_size"] * sizeScale)
	face, err := genFont(fontFile, fontSize)

	comp.SetProp(propName, face.Name)
	return face, err
}

func adjustProp(comp *tt.Component, propName string, vars map[string]float64) {
	//propItemHeight, _ := comp.GetPropString(propName)
	propVal, ok := comp.GetProp(propName)
	if !ok {
		return
	}
	propValStr, ok := propVal.(string)
	if !ok {
		return
	}
	evalResult, err := eval(vars, propValStr)
	if err != nil {
		log.Fatal(err)
	}
	evalRet := round(evalResult)
	if evalRet < 0 {
		evalRet = 0
	}
	comp.SetProp(propName, evalRet)
}

func adjustBootMenu(comp *tt.Component, vars map[string]float64) {
	vars = copyVars(vars)
	face, err := adjustFont(comp, "item_font", vars)
	if err != nil {
		log.Fatal(err)
	}

	fontHeight := face.Height()
	vars["font_height"] = float64(fontHeight)

	for _, propName := range []string{
		"item_height", "item_spacing",
		"item_padding", "icon_width",
		"icon_height", "item_icon_space",
	} {

		adjustProp(comp, propName, vars)
	}
}

func copyVars(vars map[string]float64) map[string]float64 {
	varsCopy := make(map[string]float64, len(vars))
	for key, value := range vars {
		varsCopy[key] = value
	}
	return varsCopy
}

func adjustLabel(comp *tt.Component, vars map[string]float64) {
	vars = copyVars(vars)
	face, err := adjustFont(comp, "font", vars)
	if err != nil {
		log.Fatal(err)
	}

	fontHeight := face.Height()
	vars["font_height"] = float64(fontHeight)

	for _, propName := range []string{"left", "top", "width", "height"} {
		adjustProp(comp, propName, vars)
	}

	localeVars := locale.GetLocaleVariants(optLang)
	var text string
	var textFound bool
	for _, localeVar := range localeVars {
		var ok bool
		text, ok = comp.GetPropString("_text_" + localeVar)
		if ok {
			textFound = true
			break
		}
	}
	if !textFound {
		text, _ = comp.GetPropString("_text_en")
	}
	comp.SetProp("text", text)
}

func eval(vars map[string]float64, expr string) (float64, error) {
	bc := exec.Command("bc")
	var stdInBuf bytes.Buffer

	for key, value := range vars {
		fmt.Fprintf(&stdInBuf, "%s=%f\n", key, value)
	}

	stdInBuf.WriteString("scale=10\n")
	stdInBuf.WriteString(expr)
	stdInBuf.WriteByte('\n')
	log.Printf("stdin: %s", stdInBuf.Bytes())
	bc.Stdin = &stdInBuf
	out, err := bc.Output()
	if err != nil {
		return 0, err
	}
	out = bytes.TrimSuffix(out, []byte{'\n'})
	v, err := strconv.ParseFloat(string(out), 64)
	return v, err
}
