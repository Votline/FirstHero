package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"

	"FirstHero/player"
	"FirstHero/shaders"
)

const windowWidth = 800
const windowHeight = 500

func init() {
	runtime.LockOSThread()
}

func main() {
	rt := player.Limb{Name: "Root", Parent: nil,
		CurrentPos: [4]mgl32.Vec2{
			{-0.8, -0.4}, {-0.8, -0.55},
			{-0.73, -0.55}, {-0.73, -0.4},
		},
		TargetPos: [4]mgl32.Vec2{
			{-0.8, -0.4}, {-0.8, -0.55},
			{-0.73, -0.55}, {-0.73, -0.4},
		},
		Color: []mgl32.Vec4{
			{0.0, 0.0, 1.0, 1.0},
			{0.0, 0.0, 1.0, 1.0},
			{0.0, 0.0, 1.0, 1.0},
			{0.0, 0.0, 1.0, 1.0},
		},
	}
	rh := player.Limb{Name: "RightHand", Parent: &rt,
		CurrentPos: [4]mgl32.Vec2{
			{-0.05, 0.0}, {-0.05, 0.0},
			{-0.07, 0.0}, {-0.07, 0.0},
		},
		TargetPos: [4]mgl32.Vec2{
			{-0.05, 0.0}, {-0.05, 0.0},
			{-0.07, 0.0}, {-0.07, 0.0},
		},
		Color: []mgl32.Vec4{
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
		},
	}
	lh := player.Limb{Name: "LeftHand", Parent: &rt,
		CurrentPos: [4]mgl32.Vec2{
			{0.07, 0.0}, {0.07, 0.0},
			{0.05, 0.0}, {0.05, 0.0},
		},
		TargetPos: [4]mgl32.Vec2{
			{0.07, 0.0}, {0.07, 0.0},
			{0.05, 0.0}, {0.05, 0.0},
		},
		Color: []mgl32.Vec4{
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
			{1.0, 1.0, 0.0, 1.0},
		},
	}
	rl := player.Limb{Name: "RightLeg", Parent: &rt,
		CurrentPos: [4]mgl32.Vec2{
	    {0.0, -0.15}, {0.0, -0.15},
      {-0.035, -0.15}, {-0.035, -0.15},	
		},
		TargetPos: [4]mgl32.Vec2{
			{0.0, -0.15}, {0.0, -0.15},
			{-0.035, -0.15}, {-0.035, -0.15},
		},
		Color: []mgl32.Vec4{
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
			{0.49, 0.99, 0.0, 1.0},
		},
	}
	ll := player.Limb{Name: "LeftLeg", Parent: &rt,
    CurrentPos: [4]mgl32.Vec2{
      {0.025, -0.15}, {0.025, -0.15},
      {0.0, -0.15}, {0.0, -0.15},
		},
    TargetPos: [4]mgl32.Vec2{
      {0.025, -0.15}, {0.025, -0.15},
      {0.0, -0.15}, {0.0, -0.15},
    },
    Color: []mgl32.Vec4{
      {0.49, 0.99, 0.0, 1.0},
      {0.49, 0.99, 0.0, 1.0},
      {0.49, 0.99, 0.0, 1.0},
      {0.49, 0.99, 0.0, 1.0},
    },
  }
	head := player.Limb{Name: "Head", Parent: &rt,
    CurrentPos: [4]mgl32.Vec2{
      {-0.0, -0.0}, {0.0, -0.0},
      {0.0, 0.0}, {-0.0, 0.0},
    },
    TargetPos: [4]mgl32.Vec2{
      {0.0, 0.15}, {0.0, 0.0},
      {0.0, 0.0}, {0.05, 0.15},
		},
    Color: []mgl32.Vec4{
      {1.0, 1.0, 0.0, 1.0},
      {1.0, 1.0, 0.0, 1.0},
      {1.0, 1.0, 0.0, 1.0},
      {1.0, 1.0, 0.0, 1.0},
		},
  }
	allLimbs := make(map[string]*player.Limb)
	allLimbs["Head"] = &head
	allLimbs["LeftHand"] = &lh
	allLimbs["RightHand"] = &rh
	allLimbs["LeftLeg"] = &ll
	allLimbs["RightLeg"] = &rl
	pl := player.Player{Alpha: 0.1, Speed: 0.2, JumpHeight: 0.3,
		RootLimb: &rt, Limbs: allLimbs}

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
		if action == glfw.Press || action == glfw.Repeat {
			switch key {
			case glfw.KeyA:
				pl.SetTarget(0, -pl.Speed)
			case glfw.KeyD:
				pl.SetTarget(0, pl.Speed)
			case glfw.KeySpace:
				pl.SetTarget(1, pl.JumpHeight)
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
		pl.UpdatePos(pl.RootLimb)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		for _, limb := range pl.GetAllLimbs() {
			vertices, indices := limb.CreateLimb()

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
