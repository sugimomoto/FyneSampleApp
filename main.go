package main

import (
	"database/sql"
	"fmt"
	"image/color"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "github.com/alexbrainman/odbc"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	w.SetContent(ZyankenIcon())

	w.Resize(fyne.NewSize(800, 800))
	w.ShowAndRun()

	fmt.Println("Hello")
}

func ZyankenIcon() *fyne.Container {
	goo := []int{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0}
	choki := []int{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0}
	per := []int{0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}

	mainContainer := container.New(layout.NewGridLayout(5),
		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),

		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),

		GetGridPixelIcon(goo, 10),
		widget.NewLabel(" "),
		GetGridPixelIcon(choki, 10),
		widget.NewLabel(" "),
		GetGridPixelIcon(per, 10),

		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),

		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),
		widget.NewLabel(" "),
	)

	return mainContainer
}

func GetGridPixelIcon(icon []int, size int) *fyne.Container {

	containerGoo := container.New(layout.NewGridLayout(size))
	for i := 0; i < len(icon); i++ {
		pixcel := icon[i]

		if pixcel == 0 {
			containerGoo.Add(canvas.NewRectangle(color.White))
		} else {
			containerGoo.Add(canvas.NewRectangle(color.Black))
		}
	}

	return containerGoo
}

func TabsSample() *container.AppTabs {

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", FormSample()),
	)
	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), RectangleSample()))
	tabs.SetTabLocation(container.TabLocationTop)

	return tabs
}

func FormSample() *widget.Form {
	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	entry3 := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "First Name", Widget: entry1},
			{Text: "Last Name", Widget: entry2},
			{Text: "Age", Widget: entry3},
			{Text: "Free Text", Widget: textArea},
		},
		OnSubmit: func() {
			log.Println("Form submitted:", entry1.Text)
			log.Println("multiline:", textArea.Text)
		},
	}

	return form
}

func EntryWidgetSample() *fyne.Container {

	labelText := binding.NewString()
	labelText.Set("Default Value")

	label := widget.NewLabelWithData(labelText)

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	button := widget.NewButton("Click me", func() {
		labelText.Set(input.Text)
	})

	check := widget.NewCheck("Please select", func(val bool) {
		labelText.Set(strconv.FormatBool(val))
	})

	return container.New(layout.NewGridLayout(1), label, input, button, check)
}

func RectangleSample() *fyne.Container {

	rect1 := canvas.NewRectangle(color.White)
	rect2 := canvas.NewRectangle(color.Black)
	rect3 := canvas.NewRectangle(color.White)
	rect4 := canvas.NewRectangle(color.Black)
	rect5 := canvas.NewRectangle(color.White)
	rect6 := canvas.NewRectangle(color.Black)
	rect7 := canvas.NewRectangle(color.White)
	rect8 := canvas.NewRectangle(color.Black)
	rect9 := canvas.NewRectangle(color.White)

	return container.New(layout.NewGridLayout(3),
		rect1,
		rect2,
		rect3,
		rect4,
		rect5,
		rect6,
		rect7,
		rect8,
		rect9)
}

// 拡張Widget
type tappableLabelWithData struct {
	widget.Label
}

func newTappableLabelWithData(data binding.String) *tappableLabelWithData {
	label := &tappableLabelWithData{}
	label.Bind(data)
	label.ExtendBaseWidget(label)

	return label
}

func (t *tappableLabelWithData) Tapped(_ *fyne.PointEvent) {
	if t.Text == "X" {
		t.Text = "O"
	} else {
		t.Text = "X"
	}
	t.Refresh()
}

func (t *tappableLabelWithData) TappedSecondary(_ *fyne.PointEvent) {
}

func BindingSample2() *fyne.Container {

	str := binding.NewString()
	str.Set("X")

	label1 := newTappableLabelWithData(str)

	button2 := widget.NewButton("O", func() {})
	button3 := widget.NewButton("O", func() {})
	button4 := widget.NewButton("O", func() {})
	button5 := widget.NewButton("O", func() {})
	button6 := widget.NewButton("O", func() {})
	button7 := widget.NewButton("O", func() {})
	button8 := widget.NewButton("O", func() {})
	button9 := widget.NewButton("O", func() {})

	return container.New(layout.NewGridLayout(3), label1, button2, button3, button4, button5, button6, button7, button8, button9)
}

func ChangeValue() {

}

func BindingSample() *fyne.Container {

	// 文字列をバインディング
	str := binding.NewString()
	str.Set("HH")

	return container.NewVBox(
		widget.NewLabelWithData(str),
		widget.NewButton("Button", func() {
			val, _ := str.Get()
			str.Set(val + val)
		}),
	)
}

type Key struct {
	Row, Col int
}

func GetTableFromODBC() *widget.Table {
	db, _ := sql.Open("odbc", "DSN=CData Salesforce Source")
	defer db.Close()

	rows, _ := db.Query("SELECT Id, Name, AccountNumber FROM Account limit 3")
	defer rows.Close()

	records := map[Key]string{}
	var count int = 0

	for rows.Next() {
		var (
			Id            string
			Name          string
			AccountNumber string
		)

		rows.Scan(&Id, &Name, &AccountNumber)
		records[Key{count, 0}] = Id
		records[Key{count, 1}] = Name
		records[Key{count, 2}] = AccountNumber

		count++
	}

	return widget.NewTable(
		func() (int, int) {
			return 3, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(records[Key{i.Row, i.Col}])
		})

}

func GetTable() *widget.Table {

	var data = [][]string{
		[]string{"top left", "top right"},
		[]string{"middle left", "middle right"},
		[]string{"under left", "under right"},
	}

	return widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})
}

func GetList() *widget.List {

	var data = []string{"Hello1", "Hello2", "Hello3"}

	return widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
}

func GetContainer() *fyne.Container {
	text1 := canvas.NewText("Hello", color.Black)
	text2 := canvas.NewText("There", color.Black)
	text3 := canvas.NewText("(right)", color.Black)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.Black)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())

	text5 := canvas.NewText("Bottom", color.Black)
	bottom := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), layout.NewSpacer(), text5)

	cont := container.New(layout.NewVBoxLayout(), content, centered, bottom)

	return cont
}
