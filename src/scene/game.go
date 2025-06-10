package scene

import (
	"WormsWasher/src/assets"
	"WormsWasher/src/core"
	"WormsWasher/src/entity"
	"WormsWasher/src/fluit"
	"WormsWasher/src/gui"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	core.SceneInterface

	Player  entity.PlayerEntity
	Map     entity.MapEntity
	Started bool

	BasePosition core.Vector2
	BaseSize     core.Vector2

	Loose bool

	Scoreboard gui.Scoreboard
	GameOver   gui.GameOver
}

func (s *GameScene) Reset() {
	s.Player.SetPosition(core.NewVector(0, 0))
	s.Player.SetDirection(core.NewVector(0, 0))

	s.Map = *entity.NewMapEntity(
		core.NewVector(800, 800),
		core.NewVector(40, 40), 3,
		time.Now().UnixNano(),
	)

	s.Scoreboard = *gui.NewScoreboard(s.Map.GetGroundNumbers())
	s.GameOver = *gui.NewGameOver(func(e *fluit.Frame) {
		s.Init()
	})

	s.Started = false
	s.Loose = false
}

func (s *GameScene) Init() {
	s.BasePosition = core.NewVector(1080-800, 800-800)
	s.BaseSize = core.NewVector(800, 800)

	s.Player = *entity.NewPlayerEntity(.15, core.NewVector(40, 40), core.NewVector(0, 0))
	s.Map = *entity.NewMapEntity(
		core.NewVector(800, 800),
		core.NewVector(40, 40), 3,
		time.Now().UnixNano(),
	)
	s.Scoreboard = *gui.NewScoreboard(s.Map.GetGroundNumbers())
	s.GameOver = *gui.NewGameOver(func(e *fluit.Frame) {
		s.Init()
	})
}

func (s *GameScene) Update(dt float32) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		s.Player.SetDirection(core.NewVector(0, -1))
		s.Started = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		s.Player.SetDirection(core.NewVector(0, 1))
		s.Started = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		s.Player.SetDirection(core.NewVector(-1, 0))
		s.Started = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		s.Player.SetDirection(core.NewVector(1, 0))
		s.Started = true
	}

	playerPosition := s.Player.GetPosition()
	maxCaseX, maxCaseY := s.Map.GetCaseNumber()
	if playerPosition.XInt() < 0 {
		s.Player.SetPosition(core.NewVector(float64(maxCaseX-1), playerPosition.Y()))
	}
	if playerPosition.XInt() >= maxCaseX {
		s.Player.SetPosition(core.NewVector(0, playerPosition.Y()))
	}

	if playerPosition.YInt() < 0 {
		s.Player.SetPosition(core.NewVector(playerPosition.X(), float64(maxCaseY-1)))
	}
	if playerPosition.YInt() >= maxCaseY {
		s.Player.SetPosition(core.NewVector(playerPosition.X(), 0))
	}

	playerPosition = s.Player.GetPosition()
	playerDirection := s.Player.GetDirection()
	playerCase := s.Map.GetMapCaseByVector2(playerPosition)
	futurePlayerCase := s.Map.GetMapCaseByVector2(core.NewVector(playerPosition.X()+playerDirection.X(), playerPosition.Y()+playerDirection.Y()))

	if futurePlayerCase != nil && futurePlayerCase.GetState() == entity.MapCaseStateWall && s.Player.GetFood() == 0 {
		s.Player.SetSpeed(.35)
	} else {
		s.Player.SetSpeed(.15)
	}

	switch playerCase.GetState() {
	case entity.MapCaseStateFood:
		{
			playerCase.SetState(entity.MapCaseStateVoid)
			s.Player.AddFood(1)
		}
	case entity.MapCaseStateWall:
		{
			if s.Player.GetFood() > 0 {
				playerCase.SetState(entity.MapCaseStateVoid)
				s.Player.AddFood(-1)
				s.Scoreboard.AddBreakedBlock(1)
				break
			}
			s.Loose = true
		}
	}

	if s.Loose {
		s.GameOver.SetVisible(true)
	}

	if s.Started && !s.Loose {
		s.Map.Update(dt)
		s.Player.Update(dt)
		s.Scoreboard.Update(dt)
	}
	s.GameOver.Update(dt, s.Scoreboard.GetBreakedBlocks(), s.Scoreboard.GetMaxBlocks(), int(s.Scoreboard.GetLifeTime()))
}

func (s GameScene) Draw(screen *ebiten.Image) {
	base, baseOpt := assets.NewRectangle(s.BasePosition, s.BaseSize, color.RGBA{0, 0, 0, 255}, 1)
	s.Map.Draw(base)
	s.Player.Draw(base)
	screen.DrawImage(base, baseOpt)
	s.Scoreboard.Draw(screen)
	s.GameOver.Draw(screen)
}
