package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type circle struct {
	positionX float64
	positionY float64
	speed     vector2D
	color     color.RGBA
	radius    float64
}

type objectSquareBounds struct {
	left   float64
	right  float64
	top    float64
	bottom float64
}

type vector2D struct {
	x float64
	y float64
}

func (circle *circle) getObjectSquareBounds() objectSquareBounds {
	return objectSquareBounds{
		circle.positionX - circle.radius,
		circle.positionX + circle.radius,
		circle.positionY + circle.radius,
		circle.positionY - circle.radius,
	}
}

func (vector vector2D) normalize() vector2D {
	magnitude := math.Sqrt(math.Pow(vector.x, 2) + math.Pow(vector.y, 2))
	xNormal := vector.x / magnitude
	yNormal := vector.y / magnitude
	return vector2D{
		xNormal, yNormal,
	}
}

func (u vector2D) dotProduct(v vector2D) float64 {
	return u.x*v.x + u.y*v.y
}

func (v vector2D) magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (c1 *circle) circleCollided(c2 *circle) bool {
	distanceBetweenTwoCentres := math.Sqrt(math.Pow(c2.positionX-c1.positionX, 2) + math.Pow(c2.positionY-c1.positionY, 2))
	return c1.radius <= distanceBetweenTwoCentres+c2.radius

}

func newBall(positionX, positionY, radius float64, color color.RGBA, speed vector2D) circle {
	return circle{
		positionX: positionX,
		positionY: positionY,
		speed:     speed,
		color:     color,
		radius:    radius,
	}
}

func (ball *circle) render(window *pixelgl.Window) {
	object := imdraw.New(nil)
	object.Color = ball.color
	object.Push(pixel.V(float64(ball.positionX), float64(ball.positionY)))
	object.Circle(ball.radius, 0)
	object.Draw(window)
}

func (ball *circle) updatePosition() {
	ball.positionX += ball.speed.x
	ball.positionY += ball.speed.y
}

func (ball *circle) checkColisionWithSquare(objectBounds objectSquareBounds) bool {
	circleBounds := ball.getObjectSquareBounds()

	if circleBounds.right >= objectBounds.right || circleBounds.left <= objectBounds.left || circleBounds.top >= objectBounds.top || circleBounds.bottom <= objectBounds.bottom {
		return true
	}
	return false
}

func (c1 *circle) handleCircleCollision(container *circle) {
	collided := container.circleCollided(c1)
	if collided {
		pointOfCollision := vector2D{
			x: c1.positionX + c1.radius,
			y: c1.positionY + c1.radius,
		}
		collisionNormal := vector2D{
			x: pointOfCollision.x - container.positionX,
			y: pointOfCollision.y - container.positionY,
		}.normalize()
		collisionTangent := vector2D{
			x: collisionNormal.y * -1,
			y: collisionNormal.x,
		}
		speedTangentDot := collisionTangent.x*c1.speed.x + collisionTangent.y*c1.speed.y
		speedTangent := vector2D{
			x: speedTangentDot * collisionTangent.x,
			y: speedTangentDot * collisionTangent.y,
		}
		speedNormalDot := c1.speed.x*collisionNormal.x + c1.speed.y*collisionNormal.y
		reflectedSpeedNormal := vector2D{
			x: speedNormalDot * collisionNormal.x * -1,
			y: speedNormalDot * collisionNormal.y * -1,
		}
		resultingSpeed := vector2D{
			x: speedTangent.x + reflectedSpeedNormal.x,
			y: speedTangent.y + reflectedSpeedNormal.y,
		}
		c1.speed = resultingSpeed
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Playball",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	windowBounds := objectSquareBounds{
		left:   0,
		right:  1024,
		top:    768,
		bottom: 0,
	}

	containerCircle := newBall(512, 384, 300, colornames.Burlywood, vector2D{0, 0})
	ball := newBall(500, 500, 15, colornames.Blueviolet, vector2D{
		5, 5,
	})
	// seed := rand.NewSource(1)
	// random := rand.New(seed)
	for !win.Closed() {
		win.Clear(colornames.Aliceblue)
		containerCircle.render(win)
		collided := ball.checkColisionWithSquare(windowBounds)
		ball.handleCircleCollision(&containerCircle)
		if collided {
			fmt.Println("Collision!")
		}
		ball.updatePosition()
		ball.render(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
