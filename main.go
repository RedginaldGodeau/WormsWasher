package main

import (
	"WormsWasher/src/assets"
	"WormsWasher/src/core"
	"WormsWasher/src/scene"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene string

const (
	SceneGame = "game"
)

type Game struct {
	state  string
	scenes []core.SceneInterface
}

func (g Game) Update() error {
	var dt float32 = 1.0 / 60

	switch g.state {
	case SceneGame:
		g.scenes[0].Update(dt)
	}

	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case SceneGame:
		g.scenes[0].Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1080, 800
}

func main() {
	ebiten.SetWindowSize(1080, 800)
	ebiten.SetWindowTitle("Hello World")

	err := assets.AddFont("jersey", "./assets/fonts/Jersey.ttf")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		state: SceneGame,
		scenes: []core.SceneInterface{
			&scene.GameScene{},
		},
	}
	for _, scene := range game.scenes {
		scene.Init()
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
