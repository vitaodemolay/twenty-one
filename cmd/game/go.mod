module github.com/vitaodemolay/twenty-one/cmd/game

go 1.23.3

require github.com/pterm/pterm v0.12.80

require github.com/vitaodemolay/twenty-one/internal/model v1.0.0 //internal model

require (
	atomicgo.dev/cursor v0.2.0 // indirect
	atomicgo.dev/keyboard v0.2.9 // indirect
	atomicgo.dev/schedule v0.1.0 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/gookit/color v1.5.4 // indirect
	github.com/lithammer/fuzzysearch v1.1.8 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/vitaodemolay/twenty-one/internal/symbol v1.0.0 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/term v0.26.0 // indirect
	golang.org/x/text v0.20.0 // indirect
)

replace github.com/vitaodemolay/twenty-one/internal/model => ../../internal/model

replace github.com/vitaodemolay/twenty-one/internal/symbol => ../../internal/symbol
