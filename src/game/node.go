package gobject

import (
  "fmt"
  "math"
  "gmath"
  "github.com/gopherjs/gopherjs/js"
)

// Node structure
type Node struct {
  Attract_distance, Attract_strength float64
  Bounce_friction, Dampen float64
  Position gmath.Vector2
  Velicity gmath.Vector2
  Radius, Mass float64
}

// JS Constructor
func NewJS() *js.Object {
  return js.MakeWrapper(&New())
}

// Constructor for Nodes
func New() (Node) {
  return Node{90, 0.0003, -0.9, 1, gmath.Vector2{10, 10}, gmath.Vector2{10, 10}, 10, 10}
}

//Add motion to Nodes
func (t *Node) Move() {
  t.Velicity.SelfMultiplyScalar(t.dampen)
  t.Position.SelfAdd(t.velocity)
}

func (t *Node, farX, farY int) CheckBoundary() {
  if (t.Position.X - t.Radius < 0) {
    t.Position.X = t.Radius
    t.Velocity.X *= t.Bounce_friction
  } else if (t.Position.X + t.Radius > farX) {
    t.Position.X = farX - t.Radius
    t.Velocity.X *= t.Bounce_friction
  }
  if (t.Position.Y - t.Radius < 0) {
    t.Position.Y = t.Radius
    t.Velocity.Y *= t.Bounce_friction
  } else if (t.Position.Y + t.Radius >= farY) {
    t.Position.Y = farY - t.Radius
    t.Velocity.Y *= t.Bounce_friction
  }
}

func (t *Node, radiusIncrease int) MakeBigger() {
  t.Radius += radiusIncrease
  t.Mass = t.Radius / 10
}

func (t *Node) HasCollided(other *Node) {
  dist = t.Distance(other)
  return (dist < (t.Radius + other.Radius))
}

func (t *Node) CheckAttraction(other *Node) (bool) {
  dist = t.Distance(other)
  if (dist > t.Attract_distance) {
    return false
  }

  attractVector = dist.MultiplyScalar(t.Attract_strength)

  other.Velocity.SelfSubtract(attractVector.DivideScalar(other.Mass))
  t.Velocity.SelfAdd(attractVector.DivideScalar(t.Mass))

  // TODO what is this doing?
  // alpha = 1 - dist / t.attract_strength
  return true
}

func (t* Node) CheckCollision(other *Node) (bool) {
  dist = t.Position.Distance(other.Position)
  if (dist > (t.Radius + other.Radius)) {
    return false
  }

  angle = math.Atan2(dist.Y, dist.X)
  sin = math.Sin(angle)
  cos = math.Cos(angle)

  //Rotate positions
  p0 = gmath.Vector2{0, 0}
  p1 = pos0.Rotate(cos, sin, true)

  //Rotate velocities
  v0 = t.Rotate(cos, sin, true)
  v1 = other.Rotate(cos, sin, true)

  //Conserve Momentum
  vxReaction = v0.X - v1.X
  v0.X = ((t.Mass - other.Mass) * v0.X + 2 * other.Mass * v1.X) / (t.Mass + other.Mass)
  v1.X = vxReaction + v0.X

  //Prevent Emoji from sticking
  absV = math.Abs(v0.X) + math.Abs(v1.X)
  overlap = (t.Radius + other.Radius) - math.Abs(p0.X - p1.X)
  p0.X += v0.X / absV * overlap
  p1.X += v1.X / absV * overlap

  //Rotate positions back
  t.Position.SelfAdd(p0.Rotate(cos, sin, false))
  other.Position.SelfAdd(p1.Rotate(cos, sin, false))

  //Rotate velocities back
  t.Velocity = v0.Rotate(cos, sin, false)
  other.Velocity = v1.Rotate(cos, sin, false)

  return true
}


func main() {
  t := new(Node)
  t.Init()
  fmt.Printf("%d\n", t.dampen)
}

