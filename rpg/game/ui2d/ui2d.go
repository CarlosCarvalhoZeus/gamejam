package ui2d

import (
	"bufio"
	"fmt"
	"gamejam/rpg/game"
	"image/png"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight = 800, 600

var renderer *sdl.Renderer
var textureAtlas *sdl.Texture
var textureIndex map[game.Tile]sdl.Rect

func loadTextureIndex() {
	infile, err := os.Open("/home/zeus/Área de Trabalho/gamejam/rpg/game/maps/atlas-index.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(infile)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		tileRune := game.Tile(line[0])
		xy := line[1:]
		spliceXY := strings.Split(xy, ",")
		x, err := strconv.Atoi(strings.TrimSpace(spliceXY[0]))
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strings.TrimSpace(spliceXY[1]))
		if err != nil {
			panic(err)
		}
		fmt.Println(tileRune, x, y)
	}
}

func imgFileToTexture(filename string) *sdl.Texture {
	infile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer infile.Close()
	img, err := png.Decode(infile)
	if err != nil {
		panic(err)
	}
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y

	pixels := make([]byte, w*h*4)
	bIndex := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[bIndex] = byte(r / 256)
			bIndex++
			pixels[bIndex] = byte(g / 256)
			bIndex++
			pixels[bIndex] = byte(b / 256)
			bIndex++
			pixels[bIndex] = byte(a / 256)
			bIndex++
		}
	}
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STATIC, int32(w), int32(h))
	if err != nil {
		panic(err)
	}
	teste := unsafe.Pointer(&pixels[0]) //entender melhor o porque disso aqui
	tex.Update(nil, teste, w*4)
	err = tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		panic(err)
	}
	return tex

}
func init() {
	fmt.Println("isso primeiro")
	sdl.LogSetAllPriority(sdl.LOG_PRIORITY_VERBOSE)
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	window, err := sdl.CreateWindow("RPG", 200, 200, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	textureAtlas = imgFileToTexture("/home/zeus/Área de Trabalho/gamejam/rpg/game/maps/tiles1.png")
	loadTextureIndex()
}

type UI2d struct {
}

func (*UI2d) Draw(level *game.Level) {
	fmt.Println("entrou aqui")
	renderer.Copy(textureAtlas, nil, nil)
	renderer.Present()
	sdl.Delay(5000)
}
