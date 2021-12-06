package main

import ph "github.com/etic4/2d-physics-engine"

//ajusteVitesse Freine progressivement
func ajusteVitesse(obj ph.Shape) {
	v0 := ph.Vec2{X: 0, Y: 0}
	if circ1.Accel() == v0 {
		vel := circ1.Velocity()

		//Intensité du freinage
		corrX := vel.X * 0.05
		corrY := vel.Y * 0.05
		obj.SetVelocity(circ1.Velocity().Sub(ph.Vec2{X: corrX, Y: corrY}))
	}
}
