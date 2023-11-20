package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/ElioenaiFerrari/flappy-bird/components"
	"github.com/ElioenaiFerrari/flappy-bird/config"
	rl "github.com/gen2brain/raylib-go/raylib"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type History struct {
	Score     int        `json:"score" gorm:"column:score"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

func getDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite"))
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&History{})
	return db
}

func gameOver(
	score int,
	background rl.Texture2D,
	gameOverNoise rl.Sound,
) {
	history := &History{
		Score: score,
	}

	if err := db.Create(&history).Error; err != nil {
		panic(err)
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Game Over", config.ScreenWidth/2-50, config.ScreenHeight/2, 20, rl.Black)
	rl.DrawText(fmt.Sprintf("Score: %d", score), config.ScreenWidth/2-50, config.ScreenHeight/2+20, 20, rl.Black)
	rl.PlaySound(gameOverNoise)
	rl.EndDrawing()

	time.Sleep(2 * time.Second)
	rl.UnloadTexture(background)
}

func buildScenario() (background rl.Texture2D, ground rl.Texture2D, render func()) {
	background = rl.LoadTexture("assets/background.png")
	ground = rl.LoadTexture("assets/ground.png")
	ground.Width = config.ScreenWidth
	render = func() {
		rl.DrawTexture(background, 0, 0, rl.White)
		rl.DrawTexture(ground, 0, config.ScreenHeight-ground.Height, rl.White)
	}

	return
}

func play() {
	rl.InitAudioDevice()
	gameOverNoise := rl.LoadSound("assets/game-over.mp3")
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Flappy Bird")
	rl.SetTargetFPS(60)

	bird := components.NewBird()
	background, ground, render := buildScenario()

	apples := []*components.Apple{
		components.NewApple(ground.Height),
	}

	score := 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		render()
		rl.DrawTexture(bird.Texture, int32(bird.X), int32(bird.Y), rl.White)

		rl.DrawText(fmt.Sprintf("Score: %d", score), 10, 10, 20, rl.Black)
		if rl.IsKeyDown(rl.KeySpace) {
			bird.Up()
		} else {
			bird.Down()
		}

		for i, apple := range apples {
			rl.DrawTexture(apple.Texture, apple.X, apple.Y, rl.White)
			apples[i].Update(ground.Height)

			if rl.CheckCollisionRecs(
				bird.CollisionRec(),
				apple.CollisionRec(),
			) {
				apples[i].X = 0
				bird.Eat()
				apples[i].Update(ground.Height)
				score++
				if len(apples) < 5 {
					apples = append(apples, components.NewApple(ground.Height))
				}
			}
		}

		if rl.CheckCollisionRecs(
			bird.CollisionRec(),
			rl.Rectangle{
				X:      0,
				Y:      float32(config.ScreenHeight - ground.Height),
				Width:  float32(ground.Width),
				Height: float32(ground.Height),
			},
		) {
			gameOver(score, background, gameOverNoise)
			bird.Close()
			for _, apple := range apples {
				apple.Close()
			}
			break
		}

		rl.EndDrawing()
	}
	runtime.GC()
}

func history() {
	var history []History
	if err := db.
		Order("score desc").
		Limit(10).
		Find(&history).
		Error; err != nil {
		panic(err)
	}
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Flappy Bird")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("History", config.ScreenWidth/2-50, config.ScreenHeight/4, 20, rl.Black)

		text := ""

		for i, h := range history {
			// make text with borders
			text += fmt.Sprintf("%d. %d\n\n", i+1, h.Score)
		}

		rl.DrawText(text, config.ScreenWidth/2-50, config.ScreenHeight/3, 20, rl.Black)

		rl.EndDrawing()
	}
	runtime.GC()

}

func menu() {
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Flappy Bird")
	rl.SetTargetFPS(60)

	playButton := rl.Rectangle{
		X:      float32(config.ScreenWidth/2 - 50),
		Y:      float32(config.ScreenHeight/2 - 50),
		Width:  100,
		Height: 50,
	}

	historyButton := rl.Rectangle{
		X:      float32(config.ScreenWidth/2 - 50),
		Y:      float32(config.ScreenHeight / 2),
		Width:  100,
		Height: 50,
	}

	exitButton := rl.Rectangle{
		X:      float32(config.ScreenWidth/2 - 50),
		Y:      float32(config.ScreenHeight/2 + 50),
		Width:  100,
		Height: 50,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Flappy Bird", config.ScreenWidth/2-50, config.ScreenHeight/4, 20, rl.Black)
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), playButton) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			play()
		}

		if rl.CheckCollisionPointRec(rl.GetMousePosition(), historyButton) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			history()
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
		rl.DrawText("History", int32(historyButton.X+historyButton.Width/2-35), int32(historyButton.Y+historyButton.Height/2-10), 20, rl.Black)
		rl.DrawText("Exit", int32(exitButton.X+exitButton.Width/2-20), int32(exitButton.Y+exitButton.Height/2-10), 20, rl.Black)

		rl.EndDrawing()
	}
}

var (
	db = getDB()
)

func main() {
	rl.SetTraceLogLevel(rl.LogError)

	menu()
}
