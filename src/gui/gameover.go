package gui

import (
	"WormsWasher/src/fluit"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameOver struct {
	scoreLabel    *fluit.TextLabel
	lifeTimeLabel *fluit.TextLabel
	restartButton *fluit.Button
	gui           fluit.Instance
	visible       bool
}

func NewGameOver(resetHandler func(e *fluit.Frame)) *GameOver {
	var scoreLabel = fluit.NewTextLabel()
	var lifeTimeLabel = fluit.NewTextLabel()
	var restartBtn = fluit.NewButton()
	var looseTitle = fluit.NewTextLabel()
	looseTitle.
		SetFontName("jersey").
		SetText("You Loose...").
		SetFontSize(42).
		SetSize(fluit.NewVector2(200, 10)).
		SetPosition(fluit.NewVector2(200, 25.0)).
		SetColor(color.White).
		SetAlign(text.AlignCenter, text.AlignStart)

	scoreLabel.
		SetFontName("jersey").
		SetText(fmt.Sprintf("Blocks: %d/%d", 0, 0)).
		SetFontSize(25).
		SetSize(fluit.NewVector2(100, 50)).
		SetPosition(fluit.NewVector2(200, 90.0)).
		SetColor(color.White).
		SetAlign(text.AlignCenter, text.AlignCenter)

	lifeTimeLabel.
		SetFontName("jersey").
		SetText("LifeTime: 0").
		SetFontSize(25).
		SetSize(fluit.NewVector2(100, 50)).
		SetPosition(fluit.NewVector2(200, 115.0)).
		SetColor(color.White).
		SetAlign(text.AlignCenter, text.AlignCenter)

	restartBtn.
		SetColor(color.RGBA{251, 197, 114, 255}).
		SetPosition(fluit.NewVector2(100.0, 200.0)).
		SetSize(fluit.NewVector2(200, 75)).
		SetRounded(10).
		SetText("Restart").
		SetFontName("jersey").
		SetFontSize(24).
		SetTextColor(color.Black).
		SetClickEvent(resetHandler).
		SetMouseEnter(func(e *fluit.Frame) {
			e.SetColor(color.RGBA{246, 213, 96, 255})
		}).
		SetMouseLeave(func(e *fluit.Frame) {
			e.SetColor(color.RGBA{251, 197, 114, 255})
		})

	frame := fluit.NewFrame()
	frame.
		SetColor(color.RGBA{68, 73, 57, 255}).
		SetPosition(fluit.NewVector2(400.0, 200.0)).
		SetRounded(10).
		SetSize(fluit.NewVector2(400, 400)).SetChildren(
		&scoreLabel, &lifeTimeLabel, &restartBtn, &looseTitle,
	)

	return &GameOver{
		scoreLabel:    &scoreLabel,
		lifeTimeLabel: &lifeTimeLabel,
		restartButton: &restartBtn,
		gui:           &frame,
		visible:       false,
	}
}

func (gui *GameOver) Update(dt float32, breakedBlocks, maxBlocks, lifeTime int) {
	gui.scoreLabel.SetText(fmt.Sprintf("Blocks: %d/%d", breakedBlocks, maxBlocks))
	gui.lifeTimeLabel.SetText(fmt.Sprintf("Lifetime: %d", lifeTime))
	gui.restartButton.Update(dt)
}

func (gui *GameOver) Draw(screen *ebiten.Image) {
	if !gui.visible {
		return
	}
	gui.gui.Draw(screen)
}

func (gui *GameOver) SetVisible(v bool) {
	gui.visible = v
}
