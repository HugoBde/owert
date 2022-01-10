package main

import (
	"math"
)

type Sphere struct {
	center   Point3
	radius   float64
	material Material
}

func (s Sphere) Intersect(r Ray) (ret HitRecord) {
	oc := s.center.Sub(r.origin)
	projectedDistance := oc.Dot(r.direction)
	projectingVector := r.direction.Scale(projectedDistance).Sub(oc)
	dSquared := projectingVector.Dot(projectingVector)
	radiusSquared := s.radius * s.radius

	if dSquared > radiusSquared {
		ret.hit = false
	} else {
		ret.hit = true
		offset := math.Sqrt(radiusSquared - dSquared)
		if projectedDistance-offset > 0 {
			ret.hitDistance = projectedDistance - offset
		} else if projectedDistance+offset > 0 {
			ret.hitDistance = projectedDistance + offset
		} else {
			ret.hit = false
		}
	}

	if s.material.reflective {
		if r.depth > 8 {
			ret.color = ColorRGB{0, 0, 0}
		} else {
			hitPoint := r.At(ret.hitDistance)
			normal := hitPoint.Sub(s.center).Normalize()
			normal = normal.Scale(-normal.Dot(r.direction))

			reflectedDirection := r.direction.Sub(normal.Scale(-2))
			reflectedRay := NewRay(hitPoint.Add(normal.Scale(0.0000001)), reflectedDirection, r.depth+1)

			ret.color = reflectedRay.GetColor()
		}

	} else {
		normal := r.At(ret.hitDistance).Sub(s.center).Normalize()
		factor := (normal.Dot(vertical) + 1) / 2

		ret.color = s.material.color.Scale(factor)
	}

	return
}
