# FirstHero

FirstHero is a minimalist 2D game built from scratch using OpenGL and GLFW.

## Description

The game features completely code-based graphics with no external assets - no textures, no imported animations, all rendered procedurally.

### Key Features:
- **Pure code graphics**: All visuals generated algorithmically
- **Custom engine**: Built from ground up without using existing game engines
- **Lightweight**: Minimal dependencies and resource consumption
- **Procedural content**: Dynamic game elements created at runtime

## Technologies
- **Go** (1.24.1) - primary programming language
- **OpenGL** (v4.1) — rendering (Compat Profile)
- **GLFW** (v3.3) - window management and input handling
- **MathGL** - mathematical operations for game physics and transformations

## Installation
1. Clone the repository:
   - `git clone https://github.com/Votline/FirstHero`
2. Install dependencies:
   - `go mod download`
3. Build:
   - `go build`
4. Run:
   - `./main`

## License
This project is licensed under [GNU AGPL v3](LICENSE).
