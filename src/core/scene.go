package core

import "github.com/hajimehoshi/ebiten/v2"

type SceneInterface interface {
	Init()
	Update(dt float32)
	Draw(screen *ebiten.Image)
}
