package app

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/skeletongo/word/model"
	myTheme "github.com/skeletongo/word/theme"
)

var App fyne.App
var Store fyne.Preferences
var Storage fyne.Storage
var Tabs *container.AppTabs
var Infinite *widget.ProgressBarInfinite
var Wait *fyne.Container
var Content fyne.CanvasObject
var MainWindow fyne.Window

func Run() {
	App = app.NewWithID("word")
	App.Settings().SetTheme(&myTheme.MyTheme{})
	Store = App.Preferences()
	Storage = App.Storage()

	MainWindow = App.NewWindow(MainWindowTitle)
	MainWindow.Resize(fyne.NewSize(MainWindowWidth, MainWindowHigh))

	Tabs = container.NewAppTabs()
	Tabs.SetTabLocation(container.TabLocationLeading)

	Infinite = widget.NewProgressBarInfinite()
	Wait = container.NewVBox(layout.NewSpacer(), Infinite, layout.NewSpacer())

	Content = container.NewPadded(Tabs)
	WaitHide()

	// sqlite初始化
	if err := model.InitSqlite(GetFilePath("word.db")); err != nil {
		dialog.ShowError(err, MainWindow)
	}

	// 添加功能列表
	AddTabs()

	MainWindow.ShowAndRun()
}

func WaitHide() {
	Wait.Hide()
	Infinite.Stop()
	MainWindow.SetContent(Content)
}

func WaitShow() {
	MainWindow.SetContent(Wait)
	Wait.Resize(Content.Size())
	Infinite.Start()
	Wait.Show()
}

func GetFilePath(name string) string {
	_, err := Storage.Create(name)
	if err != nil && os.IsExist(err) {
		dialog.ShowError(err, MainWindow)
		return ""
	}
	uri, err := Storage.Open(name)
	if err != nil {
		dialog.ShowError(err, MainWindow)
		return ""
	}
	return uri.URI().Path()
}

func AddTabs() {
	Tabs.Append(container.NewTabItem(TabList, ListPage()))
	Tabs.Append(container.NewTabItem(TabInput, InputPage()))
	Tabs.Append(container.NewTabItem(TabFile, FilePage()))
	Tabs.Append(container.NewTabItem(TabUrl, UrlPage()))
	Tabs.Append(container.NewTabItem(TabNotice, NoticePage()))
	Tabs.Append(container.NewTabItem(TabSetting, SettingPage()))
	Tabs.Append(container.NewTabItem("测试", TestPage()))
}
