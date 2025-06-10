package entity

import (
	"WormsWasher/src/assets"
	"WormsWasher/src/core"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type MapCaseState string

const (
	MapCaseStateVoid = "void"
	MapCaseStateFood = "food"
	MapCaseStateWall = "walls"
)

type mapCase struct {
	position core.Vector2
	state    MapCaseState
	noise    float64
}

func (m mapCase) GetState() MapCaseState {
	return m.state
}
func (m *mapCase) SetState(state MapCaseState) {
	m.state = state
}

type MapEntity struct {
	size     core.Vector2
	cellSize core.Vector2

	cases []*mapCase

	spawnFoodTime  float32
	spawnFoodTimer float32
}

var groundColor = color.RGBA{181, 221, 112, 255}
var ground2Color = color.RGBA{145, 177, 90, 255}
var wallColor = color.RGBA{68, 73, 57, 255}
var foodColor = color.RGBA{177, 99, 47, 255}

func NewMapEntity(size, cellSize core.Vector2, spawnFoodTime float32, seed int64) *MapEntity {
	perlin := core.NewPerlin()
	perlin.Seed(seed)

	cases := make([]*mapCase, size.XInt()/cellSize.XInt()*size.YInt()/cellSize.YInt())
	for y := 0; y < size.YInt()/cellSize.YInt(); y++ {
		for x := 0; x < size.XInt()/cellSize.XInt(); x++ {
			noiseX := perlin.Noise(float64(x)*.2, float64(y)*.2, .1)
			gameCase := mapCase{position: core.NewVector(float64(x), float64(y)), state: MapCaseStateVoid, noise: noiseX}
			if x > 0 && float64(x) < size.X()/cellSize.X()-1 && y > 0 && float64(y) < size.Y()/cellSize.Y()-1.0 {
				if noiseX < 0 {
					gameCase.state = MapCaseStateWall
				}
			}
			cases[y*size.YInt()/cellSize.YInt()+x] = &gameCase
		}
	}

	return &MapEntity{
		size:     size,
		cellSize: cellSize,

		cases: cases,

		spawnFoodTime:  spawnFoodTime,
		spawnFoodTimer: spawnFoodTime,
	}
}

func (m *MapEntity) GetGroundNumbers() int {
	n := 0
	for _, c := range m.cases {
		if c.state == MapCaseStateWall {
			n++
		}
	}
	return n
}

func (m *MapEntity) GetMapCaseByVector2(vector core.Vector2) *mapCase {
	for _, c := range m.cases {
		if c.position.Equal(vector) {
			return c
		}
	}
	return nil
}

func (m *MapEntity) GetCaseNumber() (int, int) {
	return m.size.XInt() / m.cellSize.XInt(), m.size.YInt() / m.cellSize.YInt()
}

func (m *MapEntity) Update(dt float32) {
	m.spawnFoodTimer += dt
	if m.spawnFoodTimer < m.spawnFoodTime {
		return
	}

	foodX := rand.Intn(m.size.XInt() / m.cellSize.XInt())
	foodY := rand.Intn(m.size.YInt() / m.cellSize.YInt())
	c := m.GetMapCaseByVector2(core.NewVector(float64(foodX), float64(foodY)))

	if c != nil && c.state == MapCaseStateVoid {
		c.SetState(MapCaseStateFood)
		m.spawnFoodTimer = 0
	}
}

func (m *MapEntity) Draw(screen *ebiten.Image) {
	for _, c := range m.cases {
		switch c.state {
		case MapCaseStateVoid:
			{
				var a = 1.0 - c.noise
				color := ground2Color
				if (c.position.XInt()+c.position.YInt())%2 == 0 {
					color = groundColor
				}
				caseDraw, caseOpt := assets.NewRectangle(c.position.Mult(m.cellSize.X()), m.cellSize, color, a)
				screen.DrawImage(caseDraw, caseOpt)
			}
		case MapCaseStateFood:
			{
				caseDraw, caseOpt := assets.NewRectangle(c.position.Mult(m.cellSize.X()), m.cellSize, foodColor, 1)
				screen.DrawImage(caseDraw, caseOpt)
			}
		case MapCaseStateWall:
			{
				var a = 1.0 - c.noise
				caseDraw, caseOpt := assets.NewRectangle(c.position.Mult(m.cellSize.X()), m.cellSize, wallColor, a)
				screen.DrawImage(caseDraw, caseOpt)
			}
		}
	}
}
