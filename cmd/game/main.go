package main

import (
	"fmt"
	"time"

	"github.com/vitaodemolay/twenty-one/internal/model"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const (
	game_title = "Twent One"
)

func main() {
	showStartingScreen := true
	clear()
	showStartWindow(&showStartingScreen)
	fmt.Scanln()
	showStartingScreen = false
	time.Sleep(time.Second / 2)
	questionResult := ""
	showcase(game_title, func() {
		pterm.Info.Println("Iniciando um Jogo:")
		pterm.Println()
		prompt := pterm.DefaultInteractiveContinue
		questionResult, _ = prompt.WithDefaultText("Você quer jogar contra o Dealer?").WithOptions([]string{"sim", "não"}).Show()
		pterm.Println()
	})
	for questionResult == "sim" {
		game := model.NewGame()

		showcase(game_title, func() {
			pterm.Info.Println("Iniciando um Jogo:")
			pterm.Println()
			game.CreatePlayer("Player one")

			questionResult, _ = pterm.DefaultInteractiveContinue.WithDefaultText("Você quer jogar contra o Dealer novamente?").WithOptions([]string{"sim", "não"}).Show()
		})
	}
	showFinalScreen()
}

func clear() {
	print("\033[H\033[2J")
}

func showcase(title string, content func()) {
	clear()
	pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgBlue)).Println(title)
	pterm.Println()
	content()
}

func showFinalScreen() {
	showcase(game_title, func() {
		pterm.Println()
		pterm.Println()
		pterm.Println()
		area := pterm.DefaultArea.WithCenter()
		area.Update(pterm.DefaultBox.Sprintln("\n Obrigado! Até uma proxima... \n"))
		area.Stop()
	})
}

func showStartWindow(showStartingScreen *bool) {
	area, _ := pterm.DefaultArea.
		WithFullscreen().
		WithCenter().
		Start()
	title, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle(game_title, pterm.FgCyan.ToStyle())).Srender()
	subtleStyle := pterm.NewStyle(pterm.FgWhite, pterm.BgMagenta, pterm.Bold)

	go func() {
		show := true
		for *showStartingScreen {
			panels := pterm.Panels{
				{
					{Data: pterm.DefaultCenter.Sprint(title)},
				},
				{
					{Data: pterm.Sprintln()},
				},
				{
					{Data: pterm.Sprintln()},
				},
				{
					{Data: pterm.DefaultCenter.WithCenterEachLineSeparately().Sprint(subtleStyle.Sprint(func() string {
						time.Sleep(time.Second / 2)
						if invert(&show) {
							return "Welcome!\n\nPress Enter to continue"
						}
						return "Welcome!\n\n"
					}()))},
				},
			}
			text_area, _ := pterm.DefaultPanel.WithPanels(panels).WithPadding(60).Srender()
			area.Update(text_area)
		}
	}()
	area.Stop()
}

func invert(test *bool) bool {
	*test = !*test
	return *test
}
