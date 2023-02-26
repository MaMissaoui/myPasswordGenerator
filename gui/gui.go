package gui

import (
	theme2 "my-password-generator/theme"

	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type MainScreen struct {
	mainWindow     fyne.Window
	application    fyne.App
	Password       *widget.Entry
	pwdEntropy     binding.String
	pwdOptionsBind binding.BoolList
	lengthBind     binding.Float
}

// Start .
func Start(s *MainScreen) {
	w := s.mainWindow

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Password Generator", theme.DocumentIcon(), container.NewPadded(mainWindow(s))),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), container.NewPadded(settingsWindow())),
	)

	tabs.OnSelected = func(t *container.TabItem) {
		t.Content.Refresh()
	}

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(w.Canvas().Size().Width, w.Canvas().Size().Height))
	w.CenterOnScreen()
	w.SetMaster()
	w.ShowAndRun()
}

// MakeMainScreen
func MakeMainScreen() *MainScreen {
	a := app.NewWithID("my-password-generator")

	t := a.Preferences().StringWithFallback("Theme", "Light")
	a.Settings().SetTheme(&theme2.MyTheme{Theme: t})
	a.SetIcon(theme2.Ico)
	w := a.NewWindow("Random Password Generator")

	return &MainScreen{
		mainWindow:  w,
		application: a,
	}
}
