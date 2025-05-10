package main

import (
	"log"
	"runtime"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"

	"FirstHero/shaders"
	"FirstHero/primShapes"
)

const windowWidth = 800
const windowHeight = 500

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("GLFW init error. \nErr: ", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "FirstHero", nil, nil)
	if err != nil {
		log.Fatalln("GLFW create window error. \nErr: ", err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("OpenGL init error. \nErr: ", err)
	}

	program := gl.CreateProgram()
	shaders.CompileAndAttachShaders(program)
	gl.LinkProgram(program)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	t := primShapes.Triangular{StartPos: mgl32.Vec2{-0.8, -0.4}, Width: 0.1, Height: 0.3}
	vertices, verticesQuan := t.CreateTriangular()

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.LineWidth(3.0)
	glfw.SwapInterval(1)
	gl.UseProgram(program)
	for !window.ShouldClose(){
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.DrawArrays(gl.LINE_LOOP, 0, verticesQuan)

		if err := gl.GetError(); err != gl.NO_ERROR {
			log.Println("OpenGL error. \nErr: ", err)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
