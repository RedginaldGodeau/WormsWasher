package fluit

import "github.com/hajimehoshi/ebiten/v2"

type Instance interface {
	Update(dt float32)
	Draw(screen *ebiten.Image)
	GetPosition() Vector2[float64]
	SetParent(v Instance)
	GetParent() Instance
}
