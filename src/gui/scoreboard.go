package gui

import (
	"WormsWasher/src/fluit"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scoreboard struct {
	breakedBlocks int
	maxBlocks     int
	lifeTime      float32

	scoreLabel    *fluit.TextLabel
	lifeTimeLabel *fluit.TextLabel
	gui           fluit.Instance
}

func NewScoreboard(maxBlocks int) *Scoreboard {
	var scoreLabel = fluit.NewTextLabel()
	var lifeTimeLabel = fluit.NewTextLabel()

	scoreLabel.
		SetFontName("jersey").
		SetText(fmt.Sprintf("Blocks: %d/%d", 0, maxBlocks)).
		SetFontSize(25).
		SetSize(fluit.NewVector2(100, 50)).
		SetPosition(fluit.NewVector2(5.0, 5.0)).
		SetColor(color.White).SetOpacity(1)
	lifeTimeLabel.
		SetFontName("jersey").
		SetText("LifeTime: 0").
		SetFontSize(25).
		SetSize(fluit.NewVector2(100, 50)).
		SetPosition(fluit.NewVector2(5.0, 30)).
		SetColor(color.White).SetOpacity(1)

	frame := fluit.NewFrame()
	frame.
		SetColor(color.RGBA{145, 177, 90, 255}).
		SetPosition(fluit.NewVector2(0, 0.0)).
		SetSize(fluit.NewVector2(280, 800)).AddChildren(
		&scoreLabel, &lifeTimeLabel,
	)

	return &Scoreboard{
		maxBlocks:     maxBlocks,
		scoreLabel:    &scoreLabel,
		lifeTimeLabel: &lifeTimeLabel,
		gui:           &frame,
	}
}

func (gui *Scoreboard) Update(dt float32) {
	gui.lifeTime += dt
	gui.scoreLabel.SetText(fmt.Sprintf("Blocks: %d/%d", gui.breakedBlocks, gui.maxBlocks))
	gui.lifeTimeLabel.SetText(fmt.Sprintf("Lifetime: %d", int(gui.lifeTime)))
}

func (gui *Scoreboard) Draw(screen *ebiten.Image) {
	gui.gui.Draw(screen)
}

func (gui *Scoreboard) AddBreakedBlock(v int) {
	gui.breakedBlocks += v
}

func (gui Scoreboard) GetBreakedBlocks() int {
	return gui.breakedBlocks
}

func (gui Scoreboard) GetMaxBlocks() int {
	return gui.maxBlocks
}

func (gui Scoreboard) GetLifeTime() float32 {
	return gui.lifeTime
}
