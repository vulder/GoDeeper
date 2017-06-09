package main

import (
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"GoDeeper/gui"
)

type Square gui.Square
type P gui.Point


var window *glfw.Window


func init() {
	runtime.LockOSThread()
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

const (
 W = 42
 H = 42
)


func keyPress(w *glfw.Window, k glfw.Key, s int, act glfw.Action, mods glfw.ModifierKey) {
	// TODO: impl
}

func update() {

}

type pixel struct {
	r, g, b int
}

var slice []pixel


func main() {
	//iniGame()
	var err error
	err = glfw.Init()
	checkErr(err)
	defer glfw.Terminate()
	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err = glfw.CreateWindow(W, H, "GoDeeper", nil, nil)
	checkErr(err)
	window.MakeContextCurrent()
	window.SetKeyCallback(keyPress)
	err = gl.Init()
	checkErr(err)

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			update()
		}
	}()

	gl.Ortho(0, W, H, 0, -1 ,1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(255,255,255,0)
	gl.LineWidth(1)
	gl.Color3f(1,0,0)
	for !window.ShouldClose() {
		gui.DrawScene(window)
		glfw.PollEvents()
	}
}