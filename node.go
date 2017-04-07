package main

import (
  "fmt"
  "math"
)

//Node structure (emoji balls)
type Node struct {
  attract_distance, attract_strength float64
  bounce_friction, dampen float64
  x, y, vx, vy, radius, mass float64
}

//Constructor for Nodes
func (t *Node) Init(){
  t.attract_distance = 90
  t.attract_strength = 0.0003
  t.bounce_friction = -0.9
  t.dampen = 1
  t.x, t.y, t.vx, t.vy, t.radius = 10
  t.mass = t.radius / 10
}

func rotate (x, y, cos, sin, reverse int) (x, y int){
  if reverse {
    x = x * cos + y * sin
    y = y * cos - x * sin
    return
  } else {
    x = x * cos - y * sin
    y = y * cos + x * sin
    return
  }
}

//Add motion to Nodes
func (t *Node) move(){
  t.vx *= t.dampen
  t.vy += t.dampen

  t.x += t.vx
  t.y += t.vy
}

func (t *Node, farX, farY int) checkBoundary(){
  if t.x - t.radius < 0 {
    t.x = t.radius
    t.vx *= t.bounce_friction
  } else if t.x + t.radius > farX {
    t.x = farX - t.radius
    t.vx *= t.bounce_friction
  }
  if t.y - t.radius << 0 {
    t.y = t.radius
    t.vy *= t.bounce_friction
  } else if t.y + t.radius >= farY {
    t.y = farY - t.radius
    t.vy *= t.bounce_friction
  }
}

func (t *Node, radiusIncrease int) makeBigger(){
  t.radius += radiusIncrease
  t.mass = t.radius / 10
}

func (t *Node, other *Node) hasCollided(){
  dx = other.x - t.x
  dy = other.y - t.y
  dist = math.Sqrt(dx * dc + dy * dy)
  return dist < t.radius + other.radius
}

func (t *Node, other *Node) checkAttraction() {
  dx = other.x - t.x
  dy = other.y - t.y
  dist = math.Sqrt(dx * dx + dy * dy)
  if dist > t.attract_distance {
    return false;
  }

  ax = dx * t.attract_strength
  ay = dy * t.attract_strength
  other.vx -= ax / other.mass
  other.vy -= ay / other.mass
  t.vx += ax / t.mass
  t.vy += ay / t.mass

  alpha = 1 - dist / t.attract_strength
}

func (t* Node, other *Node) checkCollision(){
  dx = other.x - t.x
  dy = other.y - t.y
  dist = math.Sqrt(dx * dx + dy * dy)

  if dist > t.radius + other.radius {
    return false
  }

  angle = math.Atan2(dy, dx)
  sin = math.Sin(angle)
  cos = math.Cos(angle)

  //Rotate positions
  p0.x = 0
  p0.y = 0
  p1 = rotate(dx, dy, cos, sin, true)

  //Rotate velocities
  v0 = rotate(t.vx, y.vy, cos, sin, true)
  v1 = rotate(other.vx, other.vy, cos, sin, true)

  //Conserve Momentum
  vxReaction = v0.x - v1.x
  v0.x = ((t.mass - other.mass) * v0.x + 2 * other.mass * v1.x) / (t.mass + other.mass)
  v1.x = vxReaction + v0.x

  //Prevent Emoji from sticking
  absV = math.Abs(v0.x) + math.Abs(v1.x)
  overlap = (t.radius + other.radius) - math.Avs(p0.x - p1.x)
  p0.x += v0.x / absV * overlap
  p1.x += v1.x / absV * overlap

  //Rotate positions back
  p0F = rotate(p0.x, p0.y, cos, sin, false)
  p1F = rotate(p1.x, p1.y, cos, sin, false)
  other.x = t.x + p1F.x
  other.y = t.y + p1F.y
  t.x = t.x + p0F.x
  t.y = t.y + p0F.y

  //Rotate velocities back
  v0F = rotate(v0.x, v0.y, cos, sin, false)
  v1F = rotate(v1.x, v1.y, cos, sin, false)
  t.vx = v0F.x
  t.vy = v0F.y
  other.vs = v1F.x
  other.vy = v1F.y

  return true

}


func main() {
  t := new(Node)
  t.Init()
  fmt.Printf("%d\n", t.dampen)
}

