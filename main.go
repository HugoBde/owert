package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

/*

Todo:

Do a bit of documenting and commenting maybe

work on global illumination

add new materials and reflection

*/

//Scene is global

var myObjs []Hittable
var myScene Scene

var vertical Vec3

const WIDTH = 1500
const HEIGHT = 750

func progress(c, q chan int) {
	for {
		select {
		case progress := <-c:
			fmt.Printf("\x1b[50D[%v/100%%]Rendering ...", progress)

		case stop := <-q:
			if stop == 1 {
				return
			}
		}

	}
}

func render(output *image.NRGBA, c, q chan int) {
	height := output.Rect.Dy()
	width := output.Rect.Dx()
	for i := 0; i < height; i++ {
		c <- ((i + 1) * 100) / height
		for j := 0; j < width; j++ {

			y := 2.0*(float64(i)/float64(height)-0.5)*math.Tan(myScene.camera.fov/2.0) + myScene.camera.direction.y
			x := 2.0*(float64(j)/float64(width)-0.5)*math.Tan(myScene.camera.fov/2.0)*float64(width)/float64(height) + myScene.camera.direction.x

			ray := NewRay(myScene.camera.position, Vec3{x, y, -1}, 0)

			/*



				The formula is y = 2 * (i/height - 0.5) * tan(FOV / 2)
							   x = 2 * (j/width - 0.5) * tan(FOV/2) need to adjust to aspect ratio


			*/
			c := ray.GetColor()
			output.Set(j, i, c)
		}
	}
	q <- 1
}

func liveRender(renderer *sdl.Renderer, output *image.NRGBA) {
	for {
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		for i := 0; i < HEIGHT; i++ {
			for j := 0; j < WIDTH; j++ {
				color := output.NRGBAAt(j, i)
				renderer.SetDrawColor(color.R, color.G, color.B, color.A)
				renderer.DrawPoint(int32(j), int32(i))
			}
		}
		renderer.Present()
	}
}

func save(output *image.NRGBA) {
	out, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Error creating output file.")
		return
	}
	err = png.Encode(out, output)
	if err != nil {
		fmt.Println("Output could not be encoded.")
	}
	out.Close()
}

func main() {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	defer sdl.Quit()

	window, err := sdl.CreateWindow("Owert", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	liveRenderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		panic(err)
	}
	vertical = Vec3{0, -1, 0}

	rand.Seed(time.Now().UnixNano())

	myObjs = make([]Hittable, 6)

	myObjs[0] = Sphere{Point3{-10, -10, -30}, 2, Material{ColorRGB{255, 0, 0}, false}}
	myObjs[1] = Sphere{Point3{-10, 10, -30}, 2, Material{ColorRGB{255, 0, 255}, false}}
	myObjs[2] = Sphere{Point3{10, 10, -30}, 2, Material{ColorRGB{125, 255, 212}, false}}
	myObjs[3] = Sphere{Point3{10, -10, -30}, 7, Material{ColorRGB{63, 63, 63}, true}}
	myObjs[4] = Sphere{Point3{0, 3000, 0}, 2988, Material{ColorRGB{63, 63, 63}, false}}
	myObjs[5] = Sphere{Point3{0, 0, 0}, 3000, Material{ColorRGB{135, 206, 235}, false}}

	myCamera := NewCamera(Point3{0, 0, 0}, Vec3{0, 0, -1}, math.Pi/3)

	myScene = Scene{myCamera, myObjs}
	output := image.NewNRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	c := make(chan int)
	q := make(chan int)

	go progress(c, q)
	go liveRender(liveRenderer, output)
	render(output, c, q)
	save(output)
	fmt.Print("\n")

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}
	}
}
