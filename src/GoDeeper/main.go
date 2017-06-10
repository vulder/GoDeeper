package main

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"GoDeeper/gui"
	"fmt"
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
	// TODO: impl
}

func update() {

}

type pixel struct {
	r, g, b int
}

var slice []pixel

func mouseClicked(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	x, y := w.GetCursorPos()
	fmt.Println("At Pos: ",x ,y)
}

func main() {
	//iniGame()
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

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			update()
		}
	}()

	gl.Ortho(0, float64(gui.GetWidth()), float64(gui.GetHigh()), 0, -1 ,1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(255,255,255,0)
	gl.LineWidth(1)
	gl.Color3f(1,0,0)
	gl.Enable(gl.DOUBLEBUFFER)

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		gl.Viewport(0,0, int32(width), int32(height))
	})

	w := int32(gui.GetWidth())
	h := int32(gui.GetHigh())


	for !window.ShouldClose() {
		gui.DrawScene(window, w, h)
		glfw.PollEvents()
	}
}