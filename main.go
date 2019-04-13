package main

import (
	"strings"

	"github.com/fogleman/gg"
)

var (
	fontPath      = "font/DejaVuSansMono.ttf"
	fontSize      = 16.0
	lineHeight    = 1.5
	imageFileName = "image.png"
)

func main() {
	lorem := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	width := 400.0
	margin := 10.0
	innerWidth := float64(width - margin*2)

	// measure image size
	measure := gg.NewContext(int(width), 100)
	if err := measure.LoadFontFace(fontPath, fontSize); err != nil {
		panic(err)
	}
	formattedOutput := strings.Join(measure.WordWrap(lorem, innerWidth), "\n")
	_, innerHeight := measure.MeasureMultilineString(formattedOutput, 2)
	height := innerHeight + margin*2

	// create image
	dc := gg.NewContext(int(width), int(height))

	// background
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// text
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped(lorem, margin, margin, 0, 0, innerWidth, 2, gg.AlignLeft)
	dc.SavePNG(imageFileName)
}
