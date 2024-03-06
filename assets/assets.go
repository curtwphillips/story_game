package assets

// assets from pixabay
// assets from https://kenney.nl/

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var ScoreFont = mustLoadFont("font.ttf")
var HauntedHouseImage = mustLoadImage("haunted_house.png")
var FrontDoorEntranceImage = mustLoadImage("house_entrance.png")
var LivingRoomSeatsImage = mustLoadImage("living_room_seats.png")
var SpiderFlowersImage = mustLoadImage("spider_flowers.png")
var SpiderJuiceImage = mustLoadImage("spider_juice.png")
var BoxSkeletonImage = mustLoadImage("box_skeleton.png")
var OldStairsImage = mustLoadImage("old_stairs.png")
var TrinketsRoomImage = mustLoadImage("trinkets_room.png")
var SkullMirrorImage = mustLoadImage("skull_mirror.png")
var PumpkinImage = mustLoadImage("pumpkin.png")
var ForestDarkImage = mustLoadImage("forest_dark.png")
var CryptHallwayImage = mustLoadImage("crypt_hallway.png")
var TombEntranceImage = mustLoadImage("tomb_entrance.png")
var DarkGraveyardImage = mustLoadImage("dark_graveyard.png")
var DeskRoomImage = mustLoadImage("desk_room.png")
var BonePileImage = mustLoadImage("bone_pile.png")
var SnakeImage = mustLoadImage("snake.png")
var SnakeDedImage = mustLoadImage("snake_ded.png")
var MarioLuigiYoshiImage = mustLoadImage("mario_luigi_yoshi.png")

//go:embed *
var assets embed.FS

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}

	return images
}

func mustLoadFont(name string) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}
