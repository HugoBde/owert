package main

import (
)

type ColorRGB struct {
	r,g,b uint8
}

func backgroundColor() ColorRGB {
	return ColorRGB{255,255,255}
}

func (c ColorRGB) Scale(f float64) (ret ColorRGB) {
	ret.r = uint8(float64(c.r) * f)
	ret.g = uint8(float64(c.g) * f)
	ret.b = uint8(float64(c.b) * f)
	return
}

func (c ColorRGB) RGBA() (r, g, b, a uint32){
	r = uint32(c.r)
	r |= r << 8
	r *= uint32(255)
	r /= 0xff
	
	g = uint32(c.g)
	g |= g << 8
	g *= uint32(255)
	g /= 0xff

	b = uint32(c.b)
	b |= b << 8
	b *= uint32(255)
	b /= 0xff

	a = uint32(255)
	a |= a << 8

	return
}