package assets

import (
	"WormsWasher/src/core"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewRectangle(pos core.Vector2, size core.Vector2, color color.Color, opacity float64) (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	rect := ebiten.NewImage(size.XInt(), size.YInt())
	op.GeoM.Translate(pos.X(), pos.Y())
	rect.Fill(color)
	op.ColorM.Scale(1, 1, 1, opacity)

	return rect, op
}
