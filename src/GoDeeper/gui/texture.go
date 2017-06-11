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
	Test    	int8 = iota
	Gopher  	int8 = iota
	Earth1  	int8 = iota
	Tunnel  	int8 = iota
	Water   	int8 = iota
	Badger  	int8 = iota
	TileSet  	int8 = iota
	PowerG 	 	int8 = iota
	PowerP 	 	int8 = iota
	Wall1  	 	int8 = iota
	WormHole 	int8 = iota
	Strawberry	int8 = iota
	Dead		int8 = iota
	WormSpawn	int8 = iota

	// number textures
	NUM_0		int8 = iota
	NUM_1		int8 = iota
	NUM_2		int8 = iota
	NUM_3		int8 = iota
	NUM_4		int8 = iota
	NUM_5		int8 = iota
	NUM_6		int8 = iota
	NUM_7		int8 = iota
	NUM_8		int8 = iota
	NUM_9		int8 = iota
)

var TextureMap = make(map[int8]uint32)

func LoadTextures() {
	TextureMap[Dead] = NewTexture("img/dead.jpg")
	TextureMap[Test] = NewTexture("img/test.png")
	TextureMap[Gopher] = NewTexture("img/gopherb.png")
	TextureMap[Earth1] = NewTexture("img/earth1.jpg")
	TextureMap[Tunnel] = NewTexture("img/tunnel.jpg")
	TextureMap[Water]  = NewTexturePiece("img/Tileset.png",
		864,608,30,30)
	//TextureMap[Badger] = NewTexture("img/badger.png")
	TextureMap[Badger] = NewTexturePiece("img/Tileset.png",
		641, 98, 30, 30)
	TextureMap[PowerG] = NewTexturePiece("img/Tileset.png",
		160,320,30,30)
	TextureMap[PowerP] = NewTexturePiece("img/Tileset.png",
		130,320,30,30)
	TextureMap[Wall1] = NewTexturePiece("img/Tileset.png",
		256,544,32,32)
	//TextureMap[WormHole] = NewTexturePiece("img/Tileset.png",
	//	1664,352,32,32)
	TextureMap[WormHole] = NewTexturePiece("img/Tileset.png",
		577,1377,30,30)
	TextureMap[Strawberry] = NewTexturePiece("img/Tileset.png",
		1637,737,30,30)
	TextureMap[WormSpawn] = NewTexturePiece("img/Tileset.png",
		65,321,32,32)

	// load number textures
	TextureMap[NUM_0] = NewTexturePiece("img/Tileset.png",
		1355,9,9,14)
	TextureMap[NUM_1] = NewTexturePiece("img/Tileset.png",
		1387,9,9,14)
	TextureMap[NUM_2] = NewTexturePiece("img/Tileset.png",
		1419,9,9,14)
	TextureMap[NUM_3] = NewTexturePiece("img/Tileset.png",
		1451,9,9,14)
	TextureMap[NUM_4] = NewTexturePiece("img/Tileset.png",
		1484,9,9,14)
	TextureMap[NUM_5] = NewTexturePiece("img/Tileset.png",
		1515,9,9,14)
	TextureMap[NUM_6] = NewTexturePiece("img/Tileset.png",
		1547,9,9,14)
	TextureMap[NUM_7] = NewTexturePiece("img/Tileset.png",
		1579,9,9,14)
	TextureMap[NUM_8] = NewTexturePiece("img/Tileset.png",
		1611,9,9,14)
	TextureMap[NUM_9] = NewTexturePiece("img/Tileset.png",
		1643,9,9,14)
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

func NewTexturePiece(file string, x, y, w, h int) uint32 {
	imgFile, err := os.Open(file)
	if err != nil {
		log.Fatalf("texture %q not found on disk: %v\n", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(image.Rect(x,y,x+w,y+h))
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{x,y}, draw.Src)

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