package main

import "github.com/veandco/go-sdl2/sdl"
import "time"
import "math"

const wWidth int32 = 100
const wHeight int32 = 100
const scale int32 = 8

var surface *sdl.Surface

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Raycast", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, wWidth*scale, wHeight*scale, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		renderFrame()
		window.UpdateSurface()
		time.Sleep(100 * 1E6)
	}
}

var zPos float64 = 0
var xPos float64 = 0
func renderFrame() {
	xFov := float64(45)
	yFov := float64(90) / 180 * math.Pi
	h := int32(1)
	for y := int32(0); y <= wHeight; y++ {
		yV := float64(y) / float64(wHeight) * yFov - yFov / 2
		z := float64(h) / math.Tan(math.Abs(yV)) + zPos

		for x := int32(0); x <= wWidth; x++ {
			xV := float64(x) / float64(wWidth) * xFov - xFov / 2
			xOffs := math.Tan(math.Abs(xV)) * float64(z) + xPos

			col := (uint32(z * 0xFF) & 0xFF) << 16
			col = 0
			col |= (uint32(xOffs * 0xFF) & 0xFF) << 8
			drawPixel(x, y, col)
		}
	}

	xPos += 0.1
}

func drawPixel(x int32, y int32, color uint32) {
	rect := sdl.Rect{scale * x, scale * y, scale, scale}
	surface.FillRect(&rect, color)
}

