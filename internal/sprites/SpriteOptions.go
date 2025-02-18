package sprites

import (
	. "github.com/GianniBuoni/pirate-platformer/internal/interfaces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func WithImgPos(pos rl.Vector2) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		bs.imgRect.X = pos.X
		bs.imgRect.Y = pos.Y
	}
}

func WithImgWidth(w float32) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		bs.imgRect.Width = w
	}
}

func WithImgHeight(h float32) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		bs.imgRect.Height = h
	}
}

func WithAssetLib(al AssetLibrary) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		bs.assetLib = al
	}
}

func WithSpeed(f float32) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		bs.speed = f
	}
}

func WithDirection(v rl.Vector2) func(*BasicSprite) {
	return func(bs *BasicSprite) {
		bs.direction = v
	}
}
