package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	contentBox := tview.NewBox().SetBorder(true).SetTitle("TEXTNET")

	cmdInput := tview.NewInputField().
		SetPlaceholder("Enter address")

	cmdInput.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			// loadSite(contentBox)
			go func() {
				address := cmdInput.GetText()
				contentBox.SetTitle("Loading " + address)
				cmdInput.SetDisabled(true)
				app.ForceDraw()

				// should make request here
				time.Sleep(5 * time.Second)

				contentBox.SetTitle("TEXTNET")
				cmdInput.SetDisabled(false)
				app.ForceDraw()
			}()
		}
	})

	inputFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(cmdInput, 0, 1, false)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(contentBox, 0, 1, false).
		AddItem(inputFlex, 1, 0, true)

	if err := app.SetRoot(flex, true).SetFocus(cmdInput).Run(); err != nil {
		panic(err)
	}
}
