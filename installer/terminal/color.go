package terminal

import (
	"github.com/multiverse-os/color"
)

func Header(text string) string  { return color.Fuchsia(text) }
func Accent(text string) string  { return color.Silver(text) }
func Strong(text string) string  { return color.Aqua(text) }
func Text(text string) string    { return color.SkyBlue(text) }
func Success(text string) string { return color.Lime(text) }
func Warning(text string) string { return color.Yellow(text) }
func Fail(text string) string    { return color.Red(text) }
