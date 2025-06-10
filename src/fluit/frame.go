package fluit

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Frame struct {
	Instance

	size     Vector2[int]
	position Vector2[float64]
	color    color.Color
	opacity  float64
	rounded  float32
	Children []Instance

	parent Instance

	ClickActive bool

	Click      func(*Frame)
	MouseEnter func(*Frame)
	MouseLeave func(*Frame)
}

func NewFrame() Frame {
	return Frame{
		ClickActive: true,
	}
}

func (e *Frame) Update(dt float32) {
	if !e.ClickActive {
		return
	}
	mouseX, mouseY := ebiten.CursorPosition()
	parentX := 0.0
	parentY := 0.0

	currentParent := e.parent

	for currentParent != nil {
		pos := currentParent.GetPosition()
		parentX += pos.X
		parentY += pos.Y
		currentParent = currentParent.GetParent()
	}

	if mouseX >= int(e.position.X+parentX) && mouseX <= int(e.position.X+parentX)+e.size.X && mouseY >= int(e.position.Y+parentY) && mouseY <= int(e.position.Y+parentY)+e.size.Y {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			if e.Click != nil {
				e.Click(e)
			}
		}
		if e.MouseEnter != nil {
			e.MouseEnter(e)
		}
	} else {
		if e.MouseLeave != nil {
			e.MouseLeave(e)
		}
	}

	for _, child := range e.Children {
		child.Update(dt)
	}
}

func (e Frame) Draw(screen *ebiten.Image) {
	FrameOption := &ebiten.DrawImageOptions{}
	FrameImage := ebiten.NewImage(e.size.X, e.size.Y)
	FrameOption.GeoM.Translate(e.position.X, e.position.Y)

	if e.rounded > 0 {
		// Dessiner les coins arrondis
		vector.DrawFilledCircle(FrameImage, e.rounded, e.rounded, e.rounded, e.color, false)                                     // Coin supérieur gauche
		vector.DrawFilledCircle(FrameImage, float32(e.size.X)-e.rounded, e.rounded, e.rounded, e.color, false)                   // Coin supérieur droit
		vector.DrawFilledCircle(FrameImage, e.rounded, float32(e.size.Y)-e.rounded, e.rounded, e.color, false)                   // Coin inférieur gauche
		vector.DrawFilledCircle(FrameImage, float32(e.size.X)-e.rounded, float32(e.size.Y)-e.rounded, e.rounded, e.color, false) // Coin inférieur droit

		// Dessiner les côtés
		vector.DrawFilledRect(FrameImage, e.rounded, 0, float32(e.size.X)-2*e.rounded, e.rounded, e.color, false)                           // Côté supérieur
		vector.DrawFilledRect(FrameImage, e.rounded, float32(e.size.Y)-e.rounded, float32(e.size.X)-2*e.rounded, e.rounded, e.color, false) // Côté inférieur
		vector.DrawFilledRect(FrameImage, 0, e.rounded, e.rounded, float32(e.size.Y)-2*e.rounded, e.color, false)                           // Côté gauche
		vector.DrawFilledRect(FrameImage, float32(e.size.X)-e.rounded, e.rounded, e.rounded, float32(e.size.Y)-2*e.rounded, e.color, false) // Côté droit

		vector.DrawFilledRect(FrameImage, e.rounded, e.rounded, float32(e.size.X)-2*e.rounded, float32(e.size.Y)-2*e.rounded, e.color, false) // Centre

	} else {
		vector.DrawFilledRect(FrameImage, 0, 0, float32(e.size.X), float32(e.size.Y), e.color, false) // Côté supérieur
	}

	for _, instance := range e.Children {
		instance.Draw(FrameImage)
	}

	screen.DrawImage(FrameImage, FrameOption)
}

func (e *Frame) SetSize(v Vector2[int]) *Frame {
	e.size = v
	return e
}
func (e *Frame) SetRounded(v float32) *Frame {
	e.rounded = v
	return e
}
func (e *Frame) SetPosition(v Vector2[float64]) *Frame {
	e.position = v
	return e
}
func (e *Frame) SetColor(v color.Color) *Frame {
	e.color = v
	return e
}
func (e *Frame) SetOpacity(v float64) *Frame {
	e.opacity = v
	return e
}

func (e Frame) GetSize() Vector2[int] {
	return e.size
}
func (e Frame) GetPosition() Vector2[float64] {
	return e.position
}
func (e Frame) GetColor() color.Color {
	return e.color
}
func (e Frame) GetOpacity() float64 {
	return e.opacity
}

func (e *Frame) SetParent(v Instance) {
	e.parent = v
}
func (e *Frame) GetParent() Instance {
	return e.parent
}

func (e *Frame) AddChildren(v ...Instance) *Frame {
	for _, inst := range v {
		inst.SetParent(e)
		e.Children = append(e.Children, inst)
	}
	return e
}

func (e *Frame) SetChildren(v ...Instance) *Frame {
	for _, inst := range v {
		inst.SetParent(e)
	}
	e.Children = v
	return e
}
