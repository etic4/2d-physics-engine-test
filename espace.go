package main

import ph "github.com/machinbrol/2d-physics-engine"

//isVisible retourne true si la forme est dans les limites de la scène, false sinon
func isVisible(shape ph.Shape) bool {
	pos := shape.Pos()
	w := shape.Width()
	h := shape.Height()

	posLeft := pos.X + w
	posRight := pos.X
	posTop := pos.Y + h
	posBottom := pos.Y

	return posLeft < 0 || posRight > float64(screenWidth) || posTop < 0 || posBottom > float64(screenHeight)
}

func createBoundings() {
	top = ph.NewRectangle(ph.Vec2{X: 0, Y: -2}, screenWidth, 2)
	top.SetName("top")
	top.SetElasticity(1)
	top.SetTags([]string{"bord"})
	top.SetStatic(true)
	space.AddShape(top)

	right = ph.NewRectangle(ph.Vec2{X: screenWidth, Y: 0}, 2, screenHeight)
	right.SetName("right")
	right.SetElasticity(1)
	right.SetTags([]string{"bord"})
	right.SetStatic(true)
	space.AddShape(right)

	bottom = ph.NewRectangle(ph.Vec2{X: 0, Y: screenHeight}, screenWidth, 2)
	bottom.SetName("bottom")
	bottom.SetElasticity(1)
	bottom.SetTags([]string{"bord"})
	bottom.SetStatic(true)
	space.AddShape(bottom)

	left = ph.NewRectangle(ph.Vec2{X: -2, Y: 0}, 2, screenHeight)
	left.SetName("left")
	left.SetElasticity(1)
	left.SetTags([]string{"bord"})
	left.SetStatic(true)
	space.AddShape(left)
}
