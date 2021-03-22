package main

import ()


type Camera struct {
	position 	Point3
	direction 	Vec3
	fov			float64			// Vertical FOV
}

func NewCamera(p Point3, d Vec3, fov float64) (ret Camera) {
	ret.position = p
	ret.fov = fov
	if magn := d.Magnitude(); magn == 1 {
		ret.direction = d
	} else {
		ret.direction = d.Scale(1/magn)
	}
	return
}


type Scene struct {
	camera		Camera
	objects 	[]Hittable		//Right now we only have a single sphere, later on this will be a slice of Hittables, or maybe a BVH
}