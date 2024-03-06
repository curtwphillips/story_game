package game

// assets from https://kenney.nl/

import (
	"fmt"
	"image/color"
	"story_game/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 1400
	screenHeight = 1000
)

type Game struct {
	page *storyNode
	priorPages []*storyNode
}

func NewGame() *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	setGame()

	g := &Game{
		page: &House,
	}

	return g
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if len(g.priorPages) > 0 {
			g.page = g.priorPages[len(g.priorPages) - 1]
			g.priorPages = g.priorPages[0:len(g.priorPages) - 1]
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if len(g.page.choices) > 0 {
			g.priorPages = append(g.priorPages, g.page)
			g.page = g.page.choices[0].page
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		if len(g.page.choices) >= 2 {
			g.priorPages = append(g.priorPages, g.page)
			g.page = g.page.choices[1].page
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		if len(g.page.choices) >= 3 {
			g.priorPages = append(g.priorPages, g.page)
			g.page = g.page.choices[2].page
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if len(g.page.choices) >= 4 {
			g.priorPages = append(g.priorPages, g.page)
			g.page = g.page.choices[3].page
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// op.GeoM.Scale(.5, .5)
	// op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(g.page.image, op)
	screenText := g.page.text + "\n\n"
	for _, choice := range g.page.choices {
		screenText += choice.text + "\n"
	}
	text.Draw(screen, fmt.Sprint(screenText), assets.ScoreFont, 50, screenHeight - 300, color.White)
}

func (g *Game) Reset() {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
