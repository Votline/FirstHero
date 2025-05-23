package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"FirstHero/player"
	"FirstHero/shaders"
	"FirstHero/world"
)

const windowWidth = 800
const windowHeight = 500

func init() {
	runtime.LockOSThread()
}

func main() {
	pl := player.NewPlayer()
	gd := world.CreateGround()

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
	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mod glfw.ModifierKey) {	
	if action == glfw.Press {
			if key == glfw.KeySpace {
				if pl.CanJump {
					pl.SetTarget(1, pl.JumpHeight)
					pl.CanJump = false
				}
			}
		}
	})

	if err := gl.Init(); err != nil {
		log.Fatalln("OpenGL init error. \nErr: ", err)
	}

	program := gl.CreateProgram()
	shaders.CompileAndAttachShaders(program)
	gl.LinkProgram(program)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 7*4, nil)
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 7*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.LineWidth(3.0)
	glfw.SwapInterval(1)
	gl.UseProgram(program)

	for !window.ShouldClose() {
		processInput(window, pl)
		pl.UpdatePos(pl.RootLimb, gd)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		for _, limb := range pl.GetAllLimbs() {
			vertices, indices := limb.CreateLimb()
			gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
			gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

			gl.DrawElements(gl.TRIANGLES, int32(len(indices)), gl.UNSIGNED_INT, nil)
		}

		for _, block := range gd {
			vertices, indices := block.CreateQuad()
			gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
			gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

			gl.DrawElements(gl.TRIANGLES, int32(len(indices)), gl.UNSIGNED_INT, nil)

		}

		if err := gl.GetError(); err != gl.NO_ERROR {
			log.Printf("OpenGL error: 0x%x\n", err)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func processInput(w *glfw.Window, pl *player.Player) {
	if w.GetKey(glfw.KeyA) == glfw.Press {
		pl.SetTarget(0, -pl.Speed)
	}
	if w.GetKey(glfw.KeyD) == glfw.Press {
		pl.SetTarget(0, pl.Speed)
	}
}
