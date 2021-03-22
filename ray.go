package main

import (
	"math"
)


type Ray struct {
	origin    	Point3
	direction 	Vec3
	depth		uint
}

//Always create rays using this function to ensure that the direction vector is always a unit vector in order to avoid having to normalize every ray.direction that will be used later on in the program
func NewRay(o Point3, d Vec3, depth uint) Ray {
	if magn := d.Magnitude(); magn != 1.0 {
		return Ray{o, d.Scale(1 / magn), depth}
	} else {
		return Ray{o, d, depth}
	}
}

func (r Ray) At(t float64) Point3 {
	return r.origin.Add(r.direction.Scale(t))
}

func (r Ray) GetColor() (ret ColorRGB) {
	closestHit := math.MaxFloat64
	ret = backgroundColor()
	for _, obj := range myScene.objects {
		hit := obj.Intersect(r)

		if hit.hit && hit.hitDistance < closestHit {
			closestHit = hit.hitDistance
			ret = hit.color
		}
	}
	
	

	


	return
}

