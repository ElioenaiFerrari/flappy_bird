package components

import (
	"math/rand"
	"time"

	"github.com/ElioenaiFerrari/flappy-bird/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Apple struct {
	X, Y    int32
	Texture rl.Texture2D
	r       *rand.Rand
}

func NewApple() *Apple {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	a := &Apple{
		Texture: rl.LoadTexture("assets/apple.png"),
	}

	a.X = config.ScreenWidth
	a.Y = r.Int31n(config.ScreenHeight - 20)
	a.Texture.Width = 25
	a.Texture.Height = 25
	a.r = r

	return a
}

func (a *Apple) CollisionRec() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(a.X),
		Y:      float32(a.Y),
		Width:  float32(a.Texture.Width),
		Height: float32(a.Texture.Height),
	}
}

func (a *Apple) Update() {
	a.X -= 2
	if a.X <= 0 {
		a.X = config.ScreenWidth
		a.Y = a.r.Int31n(config.ScreenHeight - 20)
	}

}

func (a *Apple) Close() {
	rl.UnloadTexture(a.Texture)
}
