package main

import ()

type HitRecord struct {
	hit			bool
	hitDistance float64
	color		ColorRGB
}

type Hittable interface {
	Intersect(r Ray) HitRecord
}