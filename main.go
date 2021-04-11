package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	rect := canvas.NewRectangle(color.Black)
	w.SetContent(rect)
	w.Resize(fyne.NewSize(300, 400))

	w.ShowAndRun()
}
