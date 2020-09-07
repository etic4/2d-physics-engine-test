package main

import (
	ph "github.com/machinbrol/2d-physics-engine"
)

type bullet struct {
	*ph.Rectangle
}

func (b *bullet) String() string {
	return b.Name()
}

func addBullet(pos ph.Vec2) {
	bullet := &bullet{Rectangle: ph.NewRectangle(pos, 4, 2)}
	bullet.SetVelocity(ph.Vec2{X: 2, Y: 0})
	bullet.SetAccel(ph.Vec2{X: 0.016 * 10, Y: 0})
	bullet.SetMaxVel(ph.Vec2{X: 20, Y: 0})
	// bullet.SetMaxAccel(ph.Vec2{X: 20, Y: 0})
	bullet.SetGravity(ph.Vec2{X: 0, Y: 0})
	bullet.SetMass(10)
	bullet.SetTags([]string{"bullet"})

	bullets = append(bullets, bullet)
	space.AddShape(bullet.Rectangle)
}

func removeBullet(b ph.Shape) {
	i := 0
	for i < len(bullets) && bullets[i].Name() != b.Name() {
		i++
	}
	if i < len(bullets) {
		if i < len(bullets)-1 {
			copy(bullets[i:], bullets[i+1:])
		}
		bullets[len(bullets)-1] = nil
		bullets = bullets[:len(bullets)-1]
		space.RemoveShape(b)
	}
}
