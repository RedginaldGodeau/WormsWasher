package entity

import (
	"WormsWasher/src/assets"
	"WormsWasher/src/core"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerEntity struct {
	size  core.Vector2
	speed float32

	direction     core.Vector2
	position      core.Vector2
	movementTimer float32

	food      int
	tail      []core.Vector2
	touchTail bool
}

var playerColor = color.RGBA{0, 234, 255, 255}
var tailColor = color.RGBA{0, 234, 255, 255}

func NewPlayerEntity(speed float32, size, position core.Vector2) *PlayerEntity {
	return &PlayerEntity{
		size:  size,
		speed: speed,

		direction:     core.NewVector(0, 0),
		position:      position,
		movementTimer: 0,

		food:      0,
		tail:      []core.Vector2{},
		touchTail: false,
	}
}

func (p *PlayerEntity) GetDirection() core.Vector2 {
	return p.direction
}

func (p *PlayerEntity) SetDirection(direction core.Vector2) {
	p.direction = direction
}

func (p *PlayerEntity) GetPosition() core.Vector2 {
	return p.position
}

func (p *PlayerEntity) SetPosition(position core.Vector2) {
	p.position = position
}

func (p *PlayerEntity) GetFood() int {
	return p.food
}

func (p *PlayerEntity) AddFood(food int) {
	p.food += food
}

func (p *PlayerEntity) SetSpeed(speed float32) {
	p.speed = speed
}

func (p *PlayerEntity) Update(dt float32) {
	p.movementTimer += dt

	if p.movementTimer < p.speed {
		return
	}

	p.tail = append(p.tail, p.position)
	for len(p.tail) > p.food {
		p.tail = p.tail[1:]
	}

	p.position.Add(p.direction)

	for _, tail := range p.tail {
		if tail.Equal(p.position) {
			p.touchTail = true
			break
		}
	}

	p.movementTimer = 0
}

func (p *PlayerEntity) Draw(screen *ebiten.Image) {
	for _, tail := range p.tail {
		rect, op := assets.NewRectangle(tail.Mult(p.size.X()), p.size, tailColor, .8)
		screen.DrawImage(rect, op)
	}

	player, player_op := assets.NewRectangle(p.position.Mult(p.size.X()), p.size, playerColor, 1)

	if p.direction.X() != 0 {
		eyes1, eyes1_op := assets.NewRectangle(
			core.NewVector(
				p.direction.X()*p.size.X()*2,
				p.size.Y()+10,
			),
			core.NewVector(7, 5),
			color.RGBA{0, 0, 0, 255},
			1,
		)
		player.DrawImage(eyes1, eyes1_op)

		eyes2, eyes2_op := assets.NewRectangle(
			core.NewVector(
				p.direction.X()*p.size.X()*2,
				p.size.Y()-7-10,
			),
			core.NewVector(7, 5),
			color.RGBA{0, 0, 0, 255},
			1,
		)
		player.DrawImage(eyes2, eyes2_op)
	}

	if p.direction.Y() != 0 {
		eyes1, eyes1_op := assets.NewRectangle(
			core.NewVector(
				p.size.X()+10,
				p.direction.Y()*p.size.Y()*2,
			),
			core.NewVector(5, 7),
			color.RGBA{0, 0, 0, 255},
			1,
		)
		player.DrawImage(eyes1, eyes1_op)

		eyes2, eyes2_op := assets.NewRectangle(
			core.NewVector(
				p.size.X()-7-10,
				p.direction.Y()*p.size.Y()*2,
			),
			core.NewVector(5, 7),
			color.RGBA{0, 0, 0, 255},
			1,
		)
		player.DrawImage(eyes2, eyes2_op)
	}
	screen.DrawImage(player, player_op)
}
