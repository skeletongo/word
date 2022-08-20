package app

import (
	"errors"
	"fmt"
	"image/color"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"

	"github.com/skeletongo/word/common"
	"github.com/skeletongo/word/model"
)

func NewWordTypeSelect(changed func(s string)) *widget.Select {
	items := make([]string, 0, len(common.WordTypeToStringMap))
	for k, v := range common.WordTypeToStringMap {
		items = append(items, fmt.Sprintf("%s %s", common.WordTypeToShortMap[k], v))
	}
	sort.Strings(items)
	ret := widget.NewSelect(items, changed)
	ret.PlaceHolder = "选择词性"
	return ret
}

func InputPage() fyne.CanvasObject {
	form := widget.NewForm()

	var enEntry, ukEntry, usaEntry *widget.Entry
	var find *widget.Button
	var meanVBox *fyne.Container
	addMeanFunc := func(typ int, selected, text string) func() {
		if typ > 0 {
			selected = fmt.Sprintf("%s %s", common.WordType(typ).Short(), common.WordType(typ).String())
		}
		return func() {
			var meanItem *fyne.Container
			mean := canvas.NewText(fmt.Sprintf("%s：%s", selected, text), color.White)
			meanButton := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
				meanVBox.Remove(meanItem)
			})
			meanItem = container.NewHBox(mean, meanButton)
			l := len(meanVBox.Objects)
			old := meanVBox.Objects
			meanVBox.Objects = make([]fyne.CanvasObject, len(meanVBox.Objects)+1)
			copy(meanVBox.Objects, old[:l-1])
			meanVBox.Objects[l-1] = meanItem
			meanVBox.Objects[l] = old[l-1]
		}
	}
	rangeMeans := func(f func(m *model.Mean) bool) {
		for _, v := range meanVBox.Objects[:len(meanVBox.Objects)-1] {
			t := v.(*fyne.Container).Objects[0].(*canvas.Text)
			arr := strings.Split(t.Text, "：")
			if !f(&model.Mean{
				Typ:  int(common.ShortToWordTypeMap[strings.Split(arr[0], " ")[0]]),
				Text: arr[1],
			}) {
				return
			}
		}
	}

	enEntry = widget.NewEntry()
	find = widget.NewButton("查询", func() {
		if enEntry.Text == "" {
			return
		}

		WaitShow()
		defer WaitHide()

		ukEntry.SetText("")
		usaEntry.SetText("")
		meanVBox.Objects = meanVBox.Objects[len(meanVBox.Objects)-1:]

		w := new(model.Word)
		r := model.Db.First(w, "text = ?", enEntry.Text)
		if r.RowsAffected == 0 {
			if r.Error != nil && !errors.Is(r.Error, gorm.ErrRecordNotFound) {
				dialog.ShowError(r.Error, MainWindow)
			}
			return
		}
		ukEntry.SetText(w.UK)
		usaEntry.SetText(w.USA)
		for _, v := range w.Means {
			addMeanFunc(v.Typ, "", v.Text)()
		}
	})
	form.Append("英文单词", container.NewGridWrap(fyne.NewSize(150, EntryHigh), enEntry, find))

	uk := canvas.NewText("英", color.White)
	usa := canvas.NewText("美", color.White)
	ukEntry = widget.NewEntry()
	ukEntry.Resize(fyne.NewSize(150, EntryHigh))
	usaEntry = widget.NewEntry()
	usaEntry.Resize(fyne.NewSize(150, EntryHigh))
	phoneticVBox := container.NewVBox(
		container.NewHBox(uk, container.NewWithoutLayout(ukEntry)),
		container.NewHBox(usa, container.NewWithoutLayout(usaEntry)))
	form.Append("音标", phoneticVBox)

	meanVBox = container.NewVBox()
	sel := NewWordTypeSelect(func(s string) {})
	meanEntry := widget.NewMultiLineEntry()
	meanEntry.Wrapping = fyne.TextWrapOff
	confirm := widget.NewButton("确认", func() {
		// 不能为空
		if sel.SelectedIndex() == -1 || meanEntry.Text == "" {
			return
		}
		// 不能重复
		var has bool
		rangeMeans(func(m *model.Mean) bool {
			if strings.Split(sel.Selected, " ")[0] == common.WordType(m.Typ).Short() {
				dialog.ShowInformation("提示", "已添加", MainWindow)
				has = true
				return false
			}
			return true
		})
		// 添加翻译
		if !has {
			addMeanFunc(0, sel.Selected, meanEntry.Text)()
		}
	})
	meanVBox.Add(container.NewBorder(nil, nil, sel, confirm, meanEntry))
	form.Append("中文翻译", meanVBox)

	form.SubmitText = "添加/修改"
	form.OnSubmit = func() {
		if enEntry.Text == "" {
			return
		}
		if ukEntry.Text == "" && usaEntry.Text == "" {
			return
		}
		if ukEntry.Text == "" {
			ukEntry.SetText(usaEntry.Text)
		} else if usaEntry.Text == "" {
			usaEntry.SetText(ukEntry.Text)
		}
		if len(meanVBox.Objects) <= 1 {
			return
		}

		WaitShow()
		defer WaitHide()

		w := model.Word{
			Text: enEntry.Text,
			UK:   ukEntry.Text,
			USA:  usaEntry.Text,
		}
		rangeMeans(func(m *model.Mean) bool {
			w.Means = append(w.Means, m)
			return true
		})
		o := &model.Word{Text: enEntry.Text}
		r := model.Db.First(o)
		if r.RowsAffected > 0 {
			r = model.Db.Model(o).Updates(w)
			if r.Error != nil {
				dialog.ShowError(r.Error, MainWindow)
			}
			return
		}
		if r.Error != nil && !errors.Is(r.Error, gorm.ErrRecordNotFound) {
			dialog.ShowError(r.Error, MainWindow)
			return
		}
		r = model.Db.Save(&w)
		if r.Error != nil {
			dialog.ShowError(r.Error, MainWindow)
		}
	}
	return form
}
