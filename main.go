package main 

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"math"
	"math/rand"
	"time"
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


const HEIGHT = 3840
const WIDTH = 2160


func progress(c, q chan int) {
	for {
		select {
		case progress := <- c :
			fmt.Printf("\x1b[50D[%v/100%%]Rendering ...",progress)

		case stop := <- q:
			if stop == 1 {
				return
			}
		}
		
	}
}

func render(output *image.NRGBA, c , q chan int) {
	height := output.Rect.Dy()
	width := output.Rect.Dx()
	for i := 0; i < height; i++ {
		c <- ((i+1) * 100) / height
		for j := 0; j < width; j++ {
			
			y := 2.0 * (float64(i) / float64(height) - 0.5) * math.Tan(myScene.camera.fov / 2.0) + myScene.camera.direction.y
			x := 2.0 * (float64(j) / float64(width) - 0.5) * math.Tan(myScene.camera.fov / 2.0) * float64(width) / float64(height) + myScene.camera.direction.x

			ray := NewRay(myScene.camera.position, Vec3{x,y,-1},0)


			/*



			The formula is y = 2 * (i/height - 0.5) * tan(FOV / 2)
						   x = 2 * (j/width - 0.5) * tan(FOV/2) need to adjust to aspect ratio


			*/
			c := ray.GetColor()
			output.Set(j,i,c)
		}
	}
	q <- 1
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

	vertical = Vec3{0,-1,0}

	rand.Seed(time.Now().UnixNano())


	myObjs = make([]Hittable,6)

	myObjs[0] = Sphere {Point3{-10,-10,-30}, 2, Material{ColorRGB{255,0,0},false}}
	myObjs[1] = Sphere {Point3{-10,10,-30}, 2, Material{ColorRGB{255,0,255},false}}
	myObjs[2] = Sphere {Point3{10,10,-30}, 2, Material{ColorRGB{125,255,212},false}}
	myObjs[3] = Sphere {Point3{10,-10,-30}, 7, Material{ColorRGB{63,63,63},true}}
	myObjs[4] = Sphere {Point3{0,3000,0}, 2988, Material{ColorRGB{63,63,63},false}}
	myObjs[5] = Sphere {Point3{0,0,0}, 3000, Material{ColorRGB{135,206,235},false}}


	myCamera := NewCamera(Point3{0,0,0} , Vec3{0,0,-1} , math.Pi / 3)

	myScene = Scene{myCamera, myObjs}
	output := image.NewNRGBA(image.Rect(0,0,HEIGHT,WIDTH))

	c := make(chan int)
	q := make(chan int)

	go progress(c,q)
	render(output, c, q)
	save(output)
	fmt.Print("\n")
}