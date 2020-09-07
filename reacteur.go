package main

import (
	"fmt"
	"math/rand"

	ph "github.com/machinbrol/2d-physics-engine"
)

const (
	partwidth  float64 = 1
	partHeight float64 = 1
)

type particule struct {
	*ph.Rectangle
	count int
	dead  int //nombre d'animations après lesquelles particule est morte
}

func (p *particule) next() {
	p.count++
}

func (p *particule) isDead() bool {
	return p.count > p.dead
}

func (p *particule) String() string {
	return fmt.Sprintf("%v", p.Name())
}

func reacteur(dir ph.Vec2) {
	loc := ph.Vec2{X: 0, Y: 0}

	switch dir {
	case ph.Vec2{X: 1, Y: 0}:
		loc.X = circ1.Center().X + circ1.Width()/2
		loc.Y = circ1.Center().Y - 1
	case ph.Vec2{X: -1, Y: 0}:
		loc.X = circ1.Center().X - circ1.Width()/2
		loc.Y = circ1.Center().Y + 1
	case ph.Vec2{X: 0, Y: 1}:
		loc.X = circ1.Center().X - 1
		loc.Y = circ1.Center().Y + circ1.Height()/2
	case ph.Vec2{X: 0, Y: -1}:
		loc.X = circ1.Center().X + 1
		loc.Y = circ1.Center().Y - circ1.Height()/2
	}

	nbr := rand.Intn(7) // nombre de particules
	for i := 0; i < nbr; i++ {
		d := rand.Intn(2) //delta sur l'axe, pour répartition

		vDir := ph.Vec2{X: 0, Y: 0}

		// répartition se fait sur axe à 90° de celui de l'émission du jet
		vDir.Y = float64(d) * dir.X
		vDir.X = float64(d) * dir.Y

		pos := loc.Add(vDir)
		vel := dir.Mult(6) // insuffisant pour compenser la vitesse
		vel = vel.AddScalar(float64(rand.Intn(3) - 1))

		partic := &particule{Rectangle: ph.NewRectangle(pos, partwidth, partHeight)}
		partic.SetSolid(false) //Pas de check de collision
		partic.SetVelocity(vel)
		partic.dead = rand.Intn(2) + 1
		partic.SetTags([]string{"particule"})
		particules = append(particules, partic)
		space.AddShape(partic.Rectangle)
	}
}

func removeParticule(p *particule) {
	i := 0
	for i < len(particules) && particules[i].Name() != p.Name() {
		i++
	}
	if i < len(particules) {
		if i < len(particules)-1 {
			copy(particules[i:], particules[i+1:])
		}
		particules[len(particules)-1] = nil
		tmp := particules[:len(particules)-1]
		particules = tmp
		space.RemoveShape(p)
	}
	return
}
