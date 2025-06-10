package fluit

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button struct {
	Instance

	background  *Frame
	textLabel   *TextLabel
	parent      Instance
	ClickActive bool
}

func NewButton() Button {
	frame := NewFrame()
	txt := NewTextLabel()
	txt.SetOpacity(1)
	txt.SetAlign(text.AlignCenter, text.AlignCenter)
	frame.SetChildren(&txt)

	btn := Button{
		background: &frame,
		textLabel:  &txt,
	}

	return btn
}

func (e *Button) Update(dt float32) {
	e.background.Update(dt)
}

func (e Button) Draw(screen *ebiten.Image) {
	e.background.Draw(screen)
}

func (e *Button) SetText(v string) *Button {
	e.textLabel.SetText(v)
	return e
}

func (e *Button) SetClickActive(v bool) *Button {
	e.background.ClickActive = false
	return e
}

func (e *Button) SetTextColor(v color.Color) *Button {
	e.textLabel.color = v
	return e
}

func (e *Button) SetFontSize(v float64) *Button {
	e.textLabel.fontSize = v
	return e
}

func (e *Button) SetFontName(v string) *Button {
	e.textLabel.fontName = v
	return e
}

func (e *Button) SetSize(v Vector2[int]) *Button {
	e.background.size = v
	e.textLabel.size = v
	e.textLabel.position = NewVector2(float64(v.X)/2, float64(v.Y)/2)
	return e
}
func (e *Button) SetRounded(v float32) *Button {
	e.background.rounded = v
	return e
}
func (e *Button) SetPosition(v Vector2[float64]) *Button {
	e.background.position = v
	return e
}
func (e *Button) SetColor(v color.Color) *Button {
	e.background.color = v
	return e
}
func (e *Button) SetOpacity(v float64) *Button {
	e.background.opacity = v
	return e
}

func (e *Button) SetClickEvent(handler func(e *Frame)) *Button {
	e.background.Click = handler
	e.background.ClickActive = true
	return e
}
func (e *Button) SetMouseEnter(handler func(e *Frame)) *Button {
	e.background.MouseEnter = handler
	return e
}
func (e *Button) SetMouseLeave(handler func(e *Frame)) *Button {
	e.background.MouseLeave = handler
	return e
}

func (e Button) GetSize() Vector2[int] {
	return e.background.size
}
func (e Button) GetPosition() Vector2[float64] {
	return e.background.position
}
func (e Button) GetColor() color.Color {
	return e.background.color
}
func (e Button) GetOpacity() float64 {
	return e.background.opacity
}
func (e *Button) SetParent(v Instance) {
	e.parent = v
	e.background.SetParent(v)
}
func (e *Button) GetParent() Instance {
	return e.parent
}
