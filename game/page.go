package game

import (
	"story_game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type storyNode struct{
	text string
	image *ebiten.Image
	choices []*Choice
}

type Choice struct{
	text string
	page *storyNode
}

func (node *storyNode) addChoice(text string, page *storyNode) {
	choice := &Choice{text, page}
	if node.choices != nil {
		node.choices = append(node.choices, choice)
	} else {
		node.choices = []*Choice{choice}
	}
}

var House = storyNode{
	image: assets.HauntedHouseImage,
	text: "You have been lost for several days.\nYou find a house in a field.",
}

var FrontEntrance = storyNode{
	image: assets.FrontDoorEntranceImage,
	text: "You try the door handle and the door creaks open.\nDo you enter?",
}

var DarkForest = storyNode{
	image: assets.ForestDarkImage,
	text: "You walk away.\nIt's getting dark.",
}

var DarkGraveyard = storyNode{
	image: assets.DarkGraveyardImage,
	text: "You walk through the yard.\nNothing seems out of the ordinary.",
}

var TombEntrance = storyNode{
	image: assets.TombEntranceImage,
	text: "It looks like this place was forgotten long ago.\nIt must be an entrance to a crypt.",
}

var CryptHallway = storyNode{
	image: assets.CryptHallwayImage,
	text: "The entrance opens to a long hallway lined with tombs.",
}

var SleepSkeleton = storyNode{
	image: assets.BoxSkeletonImage,
	text: "You finally rest after days of searching.\nYou feel like you could sleep forever.",
}

var PoisonSkeleton = storyNode{
	image: assets.BoxSkeletonImage,
	text: "The poison floods your system.\nYou feel your soul ripped from your body.",
}

var ThirstySkeleton = storyNode{
	image: assets.BoxSkeletonImage,
	text: "You die of thirst thinking about what was in the bottle.",
}

var EntryWay = storyNode{
	image: assets.LivingRoomSeatsImage,
	text: "You try the door handle and the door creaks open.\nThere is a seating area inside.",
}

var SpiderFlowers = storyNode{
	image: assets.SpiderFlowersImage,
	text: "There are flowers on the table.\nDo you smell them?",
}

var SpiderJuice = storyNode{
	image: assets.SpiderJuiceImage,
	text: "A spider scatters off to a peculiar bottle.",
}

var Stairs = storyNode{
	image: assets.OldStairsImage,
	text: "You see some stairs.\nThey look inviting.",
}

var Upstairs = storyNode{
	image: assets.TrinketsRoomImage,
	text: "The room is filled with trinkets. You see some cabinets.",
}

var UpstairsCabinets = storyNode{
	image: assets.SkullMirrorImage,
	text: "You find an old mirror.\nThere's something strange about it.",
}

var SkullMirror = storyNode{
	image: assets.PumpkinImage,
	text: "There is something off in your reflection.",
}

var MirrorSkeleton = storyNode{
	image: assets.BoxSkeletonImage,
	text: "You feel strange.\nYou go to rest with the mirror.",
}

var SecretHouseEntrance = storyNode{
  image: assets.DeskRoomImage,
	text: "You've entered a strange room with a desk and a furnace.\nThere is only one door.",
}

var DeadlyTomb = storyNode{
	image: assets.BonePileImage,
	text: "A pile of bones blocks the way. And there is a broken window.",
}

var ParalizedSkeleton = storyNode{
	image: assets.BoxSkeletonImage,
	text: "You feel paralized.\nYour flesh melts away.\nYou join the pile of bones.",
}
var Snake = storyNode{
	image: assets.SnakeImage,
	text: "There is a door.",
}
var SnakeDed = storyNode{
	image: assets.SnakeDedImage,
	text: "You had a heart attack in your sleep.",
}

var MarioLuigiYoshi = storyNode{
	image: assets.MarioLuigiYoshiImage,
	text: "It's a me! Mario luigi yoshi with\nspaghetti and scrotum toad and isaac.",
}

func setGame(){
	House.addChoice("A) Look for a way in.", &FrontEntrance)
	House.addChoice("B) Leave this place.", &DarkForest)
	House.addChoice("C) Look around.", &DarkGraveyard)

	DarkGraveyard.addChoice("A) Look closer.", &TombEntrance)
	DarkGraveyard.addChoice("B) Leave this place", &DarkForest)
	DarkGraveyard.addChoice("C) Enter the Tomb.", &CryptHallway)

	TombEntrance.addChoice("A) Enter.", &CryptHallway)
	TombEntrance.addChoice("B) Leave this place.", &DarkForest)
	TombEntrance.addChoice("C) Check for clues.", &MarioLuigiYoshi)

	CryptHallway.addChoice("A) Walk down the hall.", &SecretHouseEntrance)
	CryptHallway.addChoice("B) Inspect the tombs.", &DeadlyTomb)

	DeadlyTomb.addChoice("A) Take a bone.", &ParalizedSkeleton)
	DeadlyTomb.addChoice("B) Find another way.", &CryptHallway)
	DeadlyTomb.addChoice("C) Look at the window.", &Snake)

	Snake.addChoice("A) Go to sleep.",  &SnakeDed)
	Snake.addChoice("B) Go through the passageway.", &House)

	SecretHouseEntrance.addChoice("A) Go to the next room.", &Stairs)

	DarkForest.addChoice("A) Rest here until morning.", &SleepSkeleton)
	DarkForest.addChoice("B) Keep going.", &House)

	FrontEntrance.addChoice("A) Check the door.", &EntryWay)
	FrontEntrance.addChoice("B) Look elsewhere.", &DarkGraveyard)

	EntryWay.addChoice("A) Have a seat.", &SpiderFlowers)
	EntryWay.addChoice("B) Continue searching.", &Stairs)

	SpiderFlowers.addChoice("A) Smell a flower.", &SpiderJuice)

	SpiderJuice.addChoice("A) Drink from the peculiar bottle.", &PoisonSkeleton)

	Stairs.addChoice("A) Climb the stairs.", &Upstairs)
	Stairs.addChoice("B) Look around.", &SpiderJuice)

	Upstairs.addChoice("A) Check the cabinets.", &UpstairsCabinets)

	UpstairsCabinets.addChoice("A) Look in the mirror.", &SkullMirror)
	UpstairsCabinets.addChoice("B) Do not touch the mirror", &Upstairs)

	SkullMirror.addChoice("A) Take the mirror.", &MirrorSkeleton)
	SkullMirror.addChoice("B) Leave the mirror.", &Upstairs)

}
