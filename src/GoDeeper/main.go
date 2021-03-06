package main

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"GoDeeper/gui"
	"fmt"
	"GoDeeper/game"
)

var window *glfw.Window

func init() {
	runtime.LockOSThread()
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func keyPress(w *glfw.Window, k glfw.Key, s int, act glfw.Action, mods glfw.ModifierKey) {
	if act == glfw.Press || act == glfw.Repeat {
		switch k {
		case glfw.KeyDown:
			game.GoDown()
		case glfw.KeyRight:
			game.GoRight()
		case glfw.KeyUp:
			game.GoUp()
		case glfw.KeyLeft:
			game.GoLeft();
		}
	}
}

func update(dt time.Duration) {
	game.Update(dt)
}

type pixel struct {
	r, g, b int
}

var slice []pixel

func mouseClicked(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	x, y := w.GetCursorPos()
	fmt.Println("At Pos: ", x, y)
}

func main() {
	//iniGame()
	game.Init()
	var err error
	err = glfw.Init()
	checkErr(err)
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err = glfw.CreateWindow(gui.GetWidth(), gui.GetHigh(), "GoDeeper", nil, nil)
	checkErr(err)
	window.MakeContextCurrent()
	window.SetKeyCallback(keyPress)
	window.SetMouseButtonCallback(mouseClicked)

	err = gl.Init()
	checkErr(err)

	ticker := time.NewTicker(time.Millisecond * 1000)
	go func() {
		var lastTick = time.Now()

		for range ticker.C {
			currentTick := time.Now()
			update(currentTick.Sub(lastTick))
			lastTick = currentTick
		}
	}()

	gl.Ortho(0, float64(gui.GetWidth()), float64(gui.GetHigh()), 0, -1, 1)
	//gl.Enable(gl.DEPTH_TEST)
	gl.Disable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(255, 255, 255, 0)
	gl.LineWidth(1)
	gl.Color3f(1, 0, 0)
	gl.Enable(gl.DOUBLEBUFFER)

	gl.BlendFunc(gl.ONE, gl.ONE_MINUS_SRC_ALPHA)

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	})

	gui.LoadTextures()

	gl.Enable(gl.LIGHTING)
	ambient := []float32{0.5, 0.5, 0.5, 1}
	lightPosition := []float32{5, 5, 10, 0}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	gl.Enable(gl.LIGHT0)
	gl.Disable(gl.LIGHTING)

	game.Init()

	gui.InitDrawPieceStates()

	for !window.ShouldClose() {
		gui.DrawScene(window)
		glfw.PollEvents()
	}

	gui.FreeTextures()
}
