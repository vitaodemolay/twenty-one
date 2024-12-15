package main

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/vitaodemolay/twenty-one/internal/model"
)

const (
	game_title        = "Twent One"
	continue_response = "sim"
)

func main() {
	showStartingScreen := true
	clear()
	showStartWindow(&showStartingScreen)
	fmt.Scanln()
	showStartingScreen = false
	time.Sleep(time.Second / 2)
	questionResult := showContinueGaming()
	for questionResult == continue_response {
		game := model.NewGame()
		showcase(game_title, func() {
			pterm.Info.Println("Iniciando um Jogo:")
			pterm.Println()
			gamerName, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Nome do Jogador").Show()
			game.CreatePlayer(gamerName)
		})

		gamerName, _ := game.GetPlayerName(1)
		showcase(game_title+" - Dealer X "+gamerName, func() {

		})
		questionResult = showContinueGaming()
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

func showContinueGaming() string {
	questionResult := ""
	showcase(game_title, func() {
		pterm.Info.Println("Iniciando um Jogo:")
		pterm.Println()
		questionResult, _ = pterm.DefaultInteractiveContinue.WithDefaultText("Você quer continuar jogando contra o Dealer?").WithOptions([]string{"sim", "não"}).Show()
	})

	return questionResult
}

func showFinalScreen() {
	showcase(game_title, func() {
		pterm.Println()
		pterm.Println()
		pterm.Println()
		text := "\n Obrigado! Até uma próxima... \n"
		area := pterm.DefaultArea.WithCenter()
		area.Update(pterm.DefaultBox.Sprintln(text))
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
		for !*showStartingScreen {
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
