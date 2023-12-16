package main

import (
	"github.com/pterm/pterm"
)

func main() {
	area, _ := pterm.DefaultArea.WithFullscreen().Start()
	contentBox := pterm.DefaultBox.WithTitle("default page").Sprint("content")
	addressBox := pterm.DefaultBox.Sprint("address")

	panels := [][]pterm.Panel{
		{{Data: contentBox}},
		{{Data: addressBox}},
	}

	pterm.DefaultPanel.WithPanels(panels).Render()
	area.Stop()
}
