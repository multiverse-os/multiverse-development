package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/color"
	"log"
	"strings"

)

var svgColorMap map[string]color.Color

func loadSvgColorMap() error {
	if svgColorMap != nil {
		return nil
	}

	file, err := assets.FS.Open("/svgcolors.txt")
	if err != nil {
		return err
	}

	svgColorMap = make(map[string]color.Color)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		fields := bytes.SplitN(line, []byte("\t"), 3)
		name := string(fields[0])
		hexColor := string(fields[1])

		r, g, b := parseHexColor(hexColor)
		c := color.NRGBA{
			R: byte(r),
			G: byte(g),
			B: byte(b),
			A: 255,
		}
		svgColorMap[name] = c
	}
	return nil
}

func parseColor(str string) color.Color {
	if strings.HasPrefix(str, "#") {
		r, g, b := parseHexColor(str)
		return color.NRGBA{
			R: byte(r),
			G: byte(g),
			B: byte(b),
			A: 255,
		}
	} else if strings.Contains(str, ",") {
		// r,g,b
		// remove space
		str = strings.Replace(str, " ", "", -1)
		var r int
		var g int
		var b int
		_, err := fmt.Sscanf(str, "%d,%d,%d", &r, &g, &b)
		if err != nil {
			log.Printf("fmt.SScanf failed: str %q, err: %v\n", str, err)
			return color.Black
		}

		return color.NRGBA{
			R: byte(r),
			G: byte(g),
			B: byte(b),
			A: 255,
		}

	} else {
		loadSvgColorMap()
		return svgColorMap[str]
	}
	return nil
}

func parseHexColor(x string) (r, g, b int) {
	x = strings.TrimPrefix(x, "#")
	if len(x) == 3 {
		format := "%1x%1x%1x"
		fmt.Sscanf(x, format, &r, &g, &b)
		r |= r << 4
		g |= g << 4
		b |= b << 4
	}
	if len(x) == 6 {
		format := "%02x%02x%02x"
		fmt.Sscanf(x, format, &r, &g, &b)
	}
	return
}
