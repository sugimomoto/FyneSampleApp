package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"database/sql"

	_ "github.com/alexbrainman/odbc"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	w.SetContent(GetTableFromODBC())

	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()

	fmt.Println("Hello")
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
