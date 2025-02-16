package main

import (
	"github.com/rivo/tview"
)

func main() {
	contentBox := tview.NewBox().SetBorder(true).SetTitle("TEXTNET")

	addressInput := tview.NewInputField().SetPlaceholder("Enter address")

	inputFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(addressInput, 0, 1, false)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(contentBox, 0, 1, false).
		AddItem(inputFlex, 1, 0, true)

	if err := tview.NewApplication().SetRoot(flex, true).SetFocus(addressInput).Run(); err != nil {
		panic(err)
	}
}
