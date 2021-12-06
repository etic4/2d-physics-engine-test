package main

import (
	"fmt"

	ph "github.com/etic4/2d-physics-engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth          = 640
	screenHeight         = 480
	accel        float64 = 0.016
)

var (
	err        error
	space      *ph.Space
	top        *ph.Rectangle
	right      *ph.Rectangle
	bottom     *ph.Rectangle
	left       *ph.Rectangle
	bords      []*ph.Rectangle
	rect1      *ph.Rectangle
	rect2      *ph.Rectangle
	circ1      *ph.Circle
	circ2      *ph.Circle
	circ3      *ph.Circle
	bullets    []*bullet
	particules []*particule
	rectangles []*ph.Rectangle
	cercles    []*ph.Circle
)

func init() {
	space = &ph.Space{}
	space.SetGravity(ph.Vec2{X: 0, Y: 20})

	createBoundings()
	initShapes()
}

func handleKeyStroke() {
	if rl.IsKeyDown(rl.KeyRight) {
		circ1.SetAccel(ph.Vec2{X: circ1.Accel().X + accel, Y: circ1.Accel().Y})
		reacteur(ph.Vec2{X: -1, Y: 0})
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		circ1.SetAccel(ph.Vec2{X: circ1.Accel().X - accel, Y: circ1.Accel().Y})
		reacteur(ph.Vec2{X: 1, Y: 0})
	}
	if rl.IsKeyDown(rl.KeyUp) {
		circ1.SetAccel(ph.Vec2{X: circ1.Accel().X, Y: circ1.Accel().Y - (accel * 5)})
		reacteur(ph.Vec2{X: 0, Y: 1})
	}
	if rl.IsKeyDown(rl.KeyDown) {
		circ1.SetAccel(ph.Vec2{X: circ1.Accel().X, Y: circ1.Accel().Y + accel})
		reacteur(ph.Vec2{X: 0, Y: -1})
	}
	// Saut
	// if rl.IsKeyPressed(rl.KeySpace) {
	// 	dirX := ph.Sign(circ1.Velocity().X)
	// 	circ1.SetVelocity(ph.Vec2{X: dirX * 2, Y: -10})
	// }

	if rl.IsKeyPressed(rl.KeySpace) {
		addBullet(ph.Vec2{X: circ1.Pos().X + circ1.Width(), Y: circ1.Center().Y})
	}

	if rl.IsKeyReleased(rl.KeyRight) {
		circ1.SetAccel(ph.Vec2{X: 0, Y: 0})
	}
	if rl.IsKeyReleased(rl.KeyLeft) {
		circ1.SetAccel(ph.Vec2{X: 0, Y: 0})
	}
	if rl.IsKeyReleased(rl.KeyUp) {
		circ1.SetAccel(ph.Vec2{X: 0, Y: 0})
	}
	if rl.IsKeyReleased(rl.KeyDown) {
		circ1.SetAccel(ph.Vec2{X: 0, Y: 0})
	}
}

func update() {
	handleKeyStroke()
	ajusteVitesse(circ1)

	for i := 0; i < len(particules); i++ { // faut faire comme ça pour supprimer pendant l'itération
		p := particules[i]
		if p.isDead() {
			removeParticule(p)
		} else {
			p.next()
		}
	}

	space.Update()
	coll := space.Collisions()

	for _, info := range coll.GetAll("bullet") {
		bullet, err := info.GetShapeForTag("bullet")
		if err == nil {
			b := bullet[0]
			removeBullet(b)

		}
		rect, err := info.GetShapeForName("rect1")
		if err == nil {
			removeRectangle(rect)
		}
		info.SetResolved(true)

	}

	// résoud toutes les collision impliquant des mobiles
	for i, info := range coll.GetAll("mobile") {
		//Pour éviter que les rectangles de bords soient pris en compte
		if !(info.First().IsStatic() && info.Second().IsStatic()) {
			if info.Resolved() {
				fmt.Println(i, info)
				panic("Devrait pas être résolue!")
			}
			info.Resolv()
		}
	}
	drawAll()

}

func main() {
	Fps := int32(60)

	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "raylib")

	rl.SetTargetFPS(Fps)

	for !rl.WindowShouldClose() {
		update()
	}

	rl.CloseWindow()
}
