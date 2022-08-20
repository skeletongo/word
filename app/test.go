package app

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func TestPage() fyne.CanvasObject {
	box := container.NewHBox(widget.NewButton("提交", func() {
		WaitShow()
		time.AfterFunc(time.Second*2, func() {
			WaitHide()
		})
	}))
	return box
}
