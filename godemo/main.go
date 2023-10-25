package main

import (
	"os"

	"github.com/getlantern/systray"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

type Game struct {
	carImage *ebiten.Image
	x        float64
	y        float64
	carSpeed float64
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.carSpeed -= 0.5
		if g.carSpeed < 0 {
			g.carSpeed = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.carSpeed += 0.5
	}
	g.x += g.carSpeed
	g.y += g.carSpeed
	if g.x > float64(screenWidth) {
		g.x = 0
	}
	if g.y > float64(screenHeight) {
		g.y = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)
	screen.DrawImage(g.carImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func onReady() {
	iconData, _ := os.ReadFile("car.ico")
	systray.SetIcon(iconData)
	systray.SetTitle("LTC's Car")
	systray.SetTooltip("LTC's Car")
	Restore := systray.AddMenuItem("恢复", "恢复程序")
	Quit := systray.AddMenuItem("退出", "退出程序")
	go func() {
		for {
			select {
			case <-Restore.ClickedCh:
				ebiten.RestoreWindow()
			case <-Quit.ClickedCh:
				systray.Quit()
				os.Exit(0)
			}
		}
	}()
}

func onExit() {}
func main() {
	img, _, _ := ebitenutil.NewImageFromFile("sf.png", ebiten.FilterDefault)
	game := &Game{
		carImage: img,
		x:        0,
		y:        0,
		carSpeed: 1,
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("LTC's Car")
	ebiten.SetRunnableOnUnfocused(true)
	go func() {
		systray.Run(onReady, onExit)
	}()
	ebiten.RunGame(game)
}
