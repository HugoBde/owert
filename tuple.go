package main

import (
	"math"
)

type Tuple struct {
	x, y, z float64
}

type Vec3 = Tuple

type Point3 = Tuple

func (v Tuple) Add(u Tuple) Tuple {
	return Tuple{v.x + u.x, v.y + u.y, v.z + u.z}
}

func (v Tuple) Sub(u Tuple) Tuple {
	return Tuple{v.x - u.x, v.y - u.y, v.z - u.z}
}

func (v Tuple) Scale(k float64) Tuple {
	return Tuple{v.x * k, v.y * k, v.z * k}
}

func (v Tuple) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Tuple) Normalize() Tuple {
	return v.Scale(1 / v.Magnitude())
}

func (v Tuple) Dot(u Tuple) float64 {
	return v.x*u.x + v.y*u.y + v.z*u.z
}

func (v Tuple) Cross(u Tuple) Tuple {
	return Tuple{v.y*u.z - v.z*u.y, v.z*u.x - v.x*u.z, v.x*u.y - v.y*u.x}
}

func (v Tuple) Mul(u Tuple) Tuple {
	return Tuple{v.x * u.x, v.y * u.y, v.z * u.z}
}
