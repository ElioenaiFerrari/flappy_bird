package components

import (
	"github.com/ElioenaiFerrari/flappy-bird/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bird struct {
	Texture             rl.Texture2D
	X, Y                float32
	jumpNoise, eatNoise rl.Sound
	acceleration        float32
}

func NewBird() *Bird {
	eatNoise := rl.LoadSound("assets/eat.mp3")
	jumpNoise := rl.LoadSound("assets/jump.mp3")

	b := &Bird{
		Texture:      rl.LoadTexture("assets/bird.png"),
		jumpNoise:    jumpNoise,
		eatNoise:     eatNoise,
		acceleration: 0.25,
	}

	b.X = float32(config.ScreenWidth/2 - b.Texture.Width/2)
	b.Y = float32(config.ScreenHeight/2 - b.Texture.Height/2)
	b.Texture.Width = 50
	b.Texture.Height = 30

	return b
}

func (b *Bird) CollisionRec() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(b.X),
		Y:      float32(b.Y),
		Width:  float32(b.Texture.Width),
		Height: float32(b.Texture.Height),
	}
}

func (b *Bird) Close() {
	rl.StopSound(b.jumpNoise)
	rl.StopSound(b.eatNoise)
	rl.UnloadTexture(b.Texture)
	rl.UnloadSound(b.jumpNoise)
	rl.UnloadSound(b.eatNoise)
}

func (b *Bird) Up() {
	image := rl.LoadImage("assets/bird.png")
	rl.ImageRotate(image, -10)
	texture := rl.LoadTextureFromImage(image)
	texture.Width = 50
	texture.Height = 40

	b.Texture = texture
	rl.PlaySound(b.jumpNoise)
	b.Y -= 13
	b.acceleration = 0.25
}

func (b *Bird) Eat() {
	rl.PlaySound(b.eatNoise)
}

func (b *Bird) Down() {
	image := rl.LoadImage("assets/bird.png")
	rl.ImageRotate(image, 10)
	texture := rl.LoadTextureFromImage(image)
	texture.Width = 50
	texture.Height = 40

	b.Texture = texture
	b.Y += b.acceleration
	b.acceleration += 0.25
}
