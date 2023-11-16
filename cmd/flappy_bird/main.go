package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Apple struct {
	x, y          int32
	width, height int32
	Color         rl.Color
}

func unload(
	texture rl.Texture2D,
	bird *rl.Image,
	eatNoise rl.Sound,
	jumpNoise rl.Sound,
	gameOverNoise rl.Sound,
) {
	rl.UnloadTexture(texture)
	rl.UnloadImage(bird)
	rl.StopSound(jumpNoise)
	rl.StopSound(eatNoise)
	rl.StopSound(gameOverNoise)
	rl.UnloadSound(eatNoise)
	rl.UnloadSound(jumpNoise)
	rl.UnloadSound(gameOverNoise)
	rl.CloseAudioDevice()
}

func gameOver(
	screenWidth int32,
	screenHeight int32,
	texture rl.Texture2D,
	bird *rl.Image,
	score int,
	eatNoise rl.Sound,
	jumpNoise rl.Sound,
	gameOverNoise rl.Sound,
) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Game Over", screenWidth/2-50, screenHeight/2, 20, rl.Black)
	rl.DrawText(fmt.Sprintf("Score: %d", score), screenWidth/2-50, screenHeight/2+20, 20, rl.Black)
	rl.PlaySound(gameOverNoise)
	rl.EndDrawing()

	time.Sleep(2 * time.Second)
	unload(texture, bird, eatNoise, jumpNoise, gameOverNoise)
}

func play(
	screenWidth int32,
	screenHeight int32,
) {
	rl.InitAudioDevice()
	eatNoise := rl.LoadSound("assets/eat.mp3")
	jumpNoise := rl.LoadSound("assets/jump.mp3")
	gameOverNoise := rl.LoadSound("assets/game-over.mp3")
	rl.InitWindow(screenWidth, screenHeight, "Flappy Bird")
	rl.SetTargetFPS(60)

	bird := rl.LoadImage("assets/bird.png")
	texture := rl.LoadTextureFromImage(bird)
	xCoords := int32(screenWidth/2 - texture.Width/2)
	yCoords := int32(screenHeight/2 - texture.Height/2)

	texture.Width = 50
	texture.Height = 30

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	appleLoc := r.Int31n(screenHeight - 20)
	apples := []Apple{
		{
			x:      screenWidth,
			y:      appleLoc,
			width:  20,
			height: 20,
			Color:  rl.Red,
		},
	}

	score := 0

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawTexture(texture, xCoords, yCoords, rl.White)
		rl.DrawText(fmt.Sprintf("Score: %d", score), 10, 10, 20, rl.Black)
		rl.ClearBackground(rl.RayWhite)
		if rl.IsKeyDown(rl.KeySpace) {
			rl.PlaySound(jumpNoise)
			yCoords -= 5
		} else {
			yCoords++
		}

		for i, apple := range apples {
			rl.DrawRectangle(apple.x, apple.y, apple.width, apple.height, apple.Color)
			apples[i].x -= 2

			if apple.x < 0 {
				apples[i].x = screenWidth
				apples[i].y = r.Int31n(screenHeight - 20)
			}

			if rl.CheckCollisionRecs(
				rl.Rectangle{
					X:      float32(xCoords),
					Y:      float32(yCoords),
					Width:  float32(texture.Width),
					Height: float32(texture.Height)},
				rl.Rectangle{
					X:      float32(apple.x),
					Y:      float32(apple.y),
					Width:  float32(apple.width),
					Height: float32(apple.height),
				},
			) {
				apples[i].x = screenWidth
				apples[i].y = r.Int31n(screenHeight - 20)
				rl.PlaySound(eatNoise)
				score++
			}
		}

		if yCoords > screenHeight {
			gameOver(screenWidth, screenHeight, texture, bird, score, eatNoise, jumpNoise, gameOverNoise)
			break
		}

		rl.EndDrawing()
	}

}

func history(
	screenWidth int32,
	screenHeight int32,
) {
	history := []int{
		10,
		20,
		30,
		100,
	}

	rl.InitWindow(screenWidth, screenHeight, "Flappy Bird")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("History", screenWidth/2-50, screenHeight/4, 20, rl.Black)

		text := ""

		sort.Slice(history, func(i, j int) bool {
			return history[i] > history[j]
		})
		for i, score := range history {
			// make text with borders

			text += fmt.Sprintf("%d. %d\n\n", i+1, score)
		}

		rl.DrawText(text, screenWidth/2-50, screenHeight/3, 20, rl.Black)

		rl.EndDrawing()
	}
}

func menu(
	screenWidth int32,
	screenHeight int32,
) {

	rl.InitWindow(screenWidth, screenHeight, "Flappy Bird")
	rl.SetTargetFPS(60)

	playButton := rl.Rectangle{
		X:      float32(screenWidth/2 - 50),
		Y:      float32(screenHeight/2 - 50),
		Width:  100,
		Height: 50,
	}

	historyButton := rl.Rectangle{
		X:      float32(screenWidth/2 - 50),
		Y:      float32(screenHeight / 2),
		Width:  100,
		Height: 50,
	}

	exitButton := rl.Rectangle{
		X:      float32(screenWidth/2 - 50),
		Y:      float32(screenHeight/2 + 50),
		Width:  100,
		Height: 50,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Flappy Bird", screenWidth/2-50, screenHeight/4, 20, rl.Black)

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), playButton) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			play(screenWidth, screenHeight)
		}

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), historyButton) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			history(screenWidth, screenHeight)
		}

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), exitButton) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			rl.CloseAudioDevice()
			rl.CloseWindow()
			break
		}

		rl.DrawRectangleRec(playButton, rl.Green)
		rl.DrawRectangleRec(historyButton, rl.Blue)
		rl.DrawRectangleRec(exitButton, rl.Red)

		rl.DrawText("Play", int32(playButton.X+playButton.Width/2-20), int32(playButton.Y+playButton.Height/2-10), 20, rl.Black)
		rl.DrawText("History", int32(historyButton.X+historyButton.Width/2-25), int32(historyButton.Y+historyButton.Height/2-10), 20, rl.Black)
		rl.DrawText("Exit", int32(exitButton.X+exitButton.Width/2-20), int32(exitButton.Y+exitButton.Height/2-10), 20, rl.Black)

		rl.EndDrawing()
	}
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(600)

	menu(screenWidth, screenHeight)
}
