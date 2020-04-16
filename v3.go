package vector

import (
	"fmt"
	"math"
	"math/rand"
)

// V3 represents a 3 component vector (x, y, and z usually)
type V3 struct {
	X, Y, Z float64
}

func (v V3) LenSq() float64 {
	return v.Dot(v)
}

func (v V3) Len() float64 {
	return math.Sqrt(v.LenSq())
}

func (v V3) Dist(a V3) float64 {
	return v.Sub(a).Len()
}

func (v V3) Dot(a V3) float64 {
	return v.X*a.X + v.Y*a.Y + v.Z*a.Z
}

func (v V3) Cross(a V3) V3 {
	return V3{
		v.Y*a.Z - v.Z*a.Y,
		v.Z*a.X - v.X*a.Z,
		v.X*a.Y - v.Y*a.X}
}

// Reflect a direction vector with normal vector
func (v V3) Reflect(n V3) V3 {
	dist := 2.0 * v.Dot(n)
	return V3{v.X - dist*n.X,
		v.Y - dist*n.Y,
		v.Z - dist*n.Z}
}

func (v V3) Normalize() V3 {
	l := v.Len()
	if l == 0.0 {
		return V3{}
	}
	return v.Scale(1.0 / l)
}

func (v V3) Mult(a V3) V3 {
	return V3{v.X * a.X, v.Y * a.Y, v.Z * a.Z}
}

func (v V3) Scale(s float64) V3 {
	return V3{v.X * s, v.Y * s, v.Z * s}
}

func (v V3) Add(a V3) V3 {
	return V3{v.X + a.X, v.Y + a.Y, v.Z + a.Z}
}

func (v V3) AddS(s float64) V3 {
	return V3{v.X + s, v.Y + s, v.Z + s}
}

func (v V3) Sub(a V3) V3 {
	return V3{v.X - a.X, v.Y - a.Y, v.Z - a.Z}
}

func (v V3) SubS(s float64) V3 {
	return V3{v.X - s, v.Y - s, v.Z - s}
}

func (v V3) String() string {
	return fmt.Sprintf("\t{   %.4f,   \t%.4f,   \t%.4f}", v.X, v.Y, v.Z)
}

// Eq does floating point ==, so is only suitable for
// comparing to 0,0,0 or 1,1,1 etc
func (v V3) Eq(a V3) bool {
	if v.X == a.X && v.Y == a.Y && v.Z == a.Z {
		return true
	}
	return false
}

func (v V3) CartesianToHomogeneous() V4 {
	return V4{v.X, v.Y, v.Z, 1}
}

// Q will generate a non-normalized quaternion that will rotate by an angle vector of radians,
// good for being multiplied with a normalized quatnernion representing an orientation.
// UNTESTED + PROBABLY WRONG
func (v V3) Q() Q {
	return Q{0.0, v.X, v.Y, v.Z}
}

// https://karthikkaranth.me/blog/generating-random-points-in-a-sphere/
// My first attempt at this was pretty wrong. This blog post describes some much better alogorithms.
func RandV3(lr *rand.Rand) V3 {
	u := lr.Float64()
	v := lr.Float64()

	θ := τ * u
	φ := math.Acos(2*v - 1)
	r := math.Cbrt(lr.Float64())
	//r := math.Pow(lr.Float64(), 1.0/3.0)

	sθ := math.Sin(θ)
	cθ := math.Cos(θ)

	sφ := math.Sin(φ)
	cφ := math.Cos(φ)

	return V3{
		r * sφ * cθ,
		r * sφ * sθ,
		r * cφ,
	}
}
