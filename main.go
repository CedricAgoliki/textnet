package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	header := pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgRed))

	pterm.DefaultCenter.Println(header.Sprint("textnet"))

	_ = pterm.DefaultBigText.WithLetters(putils.LettersFromString("textnet")).Render()
}
