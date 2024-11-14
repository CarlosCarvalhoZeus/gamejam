package main

import (
	"gamejam/rpg/game"
	"gamejam/rpg/game/ui2d"
)

func main() {
	ui := &ui2d.UI2d{}
	game.Run(ui)
	// if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
	// 	panic(err)
	// }
	// defer sdl.Quit()

	// window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	// if err != nil {
	// 	panic(err)
	// }
	// defer window.Destroy()

	// surface, err := window.GetSurface()
	// if err != nil {
	// 	panic(err)
	// }
	// surface.FillRect(nil, 0)

	// rect := sdl.Rect{X: 150, Y: 50, W: 200, H: 200}
	// colour := sdl.Color{R: 125, G: 255, B: 255, A: 255} // purple
	// pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	// surface.FillRect(&rect, pixel)
	// window.UpdateSurface()

	// running := true
	// for running {
	// 	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	// 		switch event.(type) {
	// 		case *sdl.QuitEvent: // NOTE: Please use *sdl.QuitEvent for v0.4.x (current version).
	// 			println("Quit")
	// 			running = false
	// 			break
	// 		}
	// 	}

	// 	sdl.Delay(33)
	// }
}
