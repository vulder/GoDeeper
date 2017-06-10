package gui

import (
	"os"
	"log"
	"image"
	"image/draw"
	_ "image/png"
	_ "image/jpeg"
	"github.com/go-gl/gl/v2.1/gl"
)

const (
	Test   int8 = iota
	Gopher int8 = iota
	Earth1 int8 = iota
	Tunnel int8 = iota
	Water  int8 = iota
	Badger int8 = iota
)

var TextureMap = make(map[int8]uint32)

func LoadTextures() {
	TextureMap[Test] = NewTexture("img/test.png")
	TextureMap[Gopher] = NewTexture("img/gopherb.png")
	TextureMap[Earth1] = NewTexture("img/earth1.jpg")
	TextureMap[Tunnel] = NewTexture("img/tunnel.jpg")
	TextureMap[Water] = NewTexture("img/water.png")
	TextureMap[Badger] = NewTexture("img/badger.png")
}

func FreeTextures() {
	for _, v := range TextureMap {
		gl.DeleteTextures(1, &v)
	}
}

func DrawTextureTS(top, left int, textureName int8) {
	DrawTexture(top, left, tile_size, tile_size, textureName)
}

func DrawTexture(top, left, w, h int, textureName int8) {
	gl.ClearColor(0,0,0,0)

	gl.BindTexture(gl.TEXTURE_2D, TextureMap[textureName])

	gl.Enable(gl.TEXTURE_2D)
	defer gl.Disable(gl.TEXTURE_2D)

	gl.Enable(gl.LIGHTING)
	defer gl.Disable(gl.LIGHTING)

	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0,0)
	gl.Vertex2i(int32(left), int32(top))

	gl.TexCoord2f(1, 0)
	gl.Vertex2i(int32(left + w), int32(top))

	gl.TexCoord2f(1, 1)
	gl.Vertex2i(int32(left + w),int32(top + h))

	gl.TexCoord2f(0, 1)
	gl.Vertex2i(int32(left),int32(top + h))
	gl.End()
}

func NewTexture(file string) uint32 {
	imgFile, err := os.Open(file)
	if err != nil {
		log.Fatalf("texture %q not found on disk: %v\n", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))
	gl.Disable(gl.TEXTURE_2D)

	return texture
}