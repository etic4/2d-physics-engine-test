package main

import (
	ph "github.com/etic4/2d-physics-engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawAll() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for _, s := range particules { // D'abord les particules, les reste est dessiné par dessus
		drawRect(s.Rectangle, rl.White)
	}

	for _, s := range rectangles {
		drawRect(s, rl.White)
	}

	for _, s := range cercles {
		drawCircle(s, rl.White)
	}

	for _, s := range bullets {
		drawRect(s.Rectangle, rl.White)
	}

	rl.EndDrawing()
}

func drawRect(rect *ph.Rectangle, clr rl.Color) {
	x := int32(rect.Pos().X)
	y := int32(rect.Pos().Y)
	w := int32(rect.Width())
	h := int32(rect.Height())

	rl.DrawRectangle(x, y, w, h, rl.Black)
	rl.DrawRectangleLines(x, y, w, h, clr)
}

func drawCircle(circ *ph.Circle, clr rl.Color) {
	center := circ.Center()
	rl.DrawCircle(int32(center.X), int32(center.Y), float32(circ.Radius()), rl.Black)
	rl.DrawCircleLines(int32(center.X), int32(center.Y), float32(circ.Radius()), clr)
}
