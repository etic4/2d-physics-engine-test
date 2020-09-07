package main

import (
	ph "github.com/machinbrol/2d-physics-engine"
)

func initShapes() {
	rect1 = ph.NewRectangle(ph.Vec2{X: 150, Y: 150}, 30, 30)
	rect1.SetName("rect1")
	rect1.SetVelocity(ph.Vec2{X: -2, Y: 8})
	rect1.SetMaxVel(ph.Vec2{X: 20, Y: 20})
	rect1.SetMaxAccel(ph.Vec2{X: 0.016 * 20, Y: 0.016 * 20})
	rect1.SetGravity(ph.Vec2{X: 0, Y: 0.016 * 2})
	rect1.SetMass(10)
	rect1.SetElasticity(1)
	rect1.SetFriction(0.01)
	rect1.SetTags([]string{"mobile"})
	space.AddShape(rect1)

	rect2 = ph.NewRectangle(ph.Vec2{X: 170, Y: 200}, 120, 10)
	rect2.SetName("rect immobile")
	rect2.SetElasticity(1)
	rect2.SetTags([]string{"immobile"})
	space.AddShape(rect2)

	circ1 = ph.NewCircle(ph.Vec2{X: 400, Y: 170}, 20)
	circ1.SetName("circ1")
	circ1.SetVelocity(ph.Vec2{X: 0, Y: 0})
	circ1.SetMaxVel(ph.Vec2{X: 0, Y: 10})
	circ1.SetMaxAccel(ph.Vec2{X: 0.016 * 10, Y: 0.016 * 10})
	circ1.SetMass(10)
	// circ1.SetGravity(ph.Vec2{X: 0, Y: 20})
	circ1.SetElasticity(0.2)
	circ1.SetFriction(0.2)
	circ1.SetTags([]string{"mobile"})
	space.AddShape(circ1)

	circ2 = ph.NewCircle(ph.Vec2{X: 500, Y: 200}, 40)
	circ2.SetName("circ immobile")
	circ2.SetMass(0)
	circ2.SetElasticity(0.5)
	circ2.SetTags([]string{"mobile"})
	space.AddShape(circ2)

	circ3 = ph.NewCircle(ph.Vec2{X: 500, Y: 300}, 40)
	circ3.SetName("circ statique 2")
	circ3.SetMass(0)
	circ3.SetElasticity(1)
	circ3.SetTags([]string{"immobile"})
	space.AddShape(circ3)

	bullets = []*bullet{}
	particules = []*particule{}
	rectangles = []*ph.Rectangle{top, bottom, left, right, rect1, rect2}
	cercles = []*ph.Circle{circ1, circ2, circ3}
}

func removeRectangle(r ph.Shape) {
	i := 0
	for i < len(rectangles) && rectangles[i].Name() != r.Name() {
		i++
	}
	if i < len(rectangles) {
		if i < len(rectangles)-1 {
			copy(rectangles[i:], rectangles[i+1:])
		}
		rectangles[len(rectangles)-1] = nil
		rectangles = rectangles[:len(rectangles)-1]
		space.RemoveShape(r)
	}
}
