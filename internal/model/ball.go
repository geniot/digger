package model

import (
	"geniot.com/geniot/digger/internal/utils"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Ball struct {
	x         float32
	y         float32
	width     float32
	height    float32
	xVelocity float32
	yVelocity float32

	born       uint32
	lastUpdate uint32
}

func NewBall() *Ball {
	return &Ball{}
}

func (ball Ball) Render(renderer *sdl.Renderer) {
	wnd, _ := renderer.GetWindow()
	width, height := wnd.GetSize()
	alpha := uint8(150)
	radius := width / 3
	if height/3 < radius {
		radius = height / 3
	}

	renderer.SetDrawColor(0, 0, 0, alpha)

	colourInside := [3]uint8{0, 0, 0}
	colourAround := [3]uint8{192, 192, 192}

	//DrawSimpleCircle(renderer, 100.0, 100.0, 100.0, 100.0, 100.0)
	//DrawCircle(renderer, 100.0, 100.0, 100.0, float64(width/2), float64(height/2))
	utils.AaellipseRGBA(renderer, width/2, height/2, radius+1, radius, colourAround[0], colourAround[1], colourAround[2], alpha)
	utils.AaellipseRGBA(renderer, width/2, height/2, radius, radius, colourInside[0], colourInside[1], colourInside[2], alpha)
	utils.AaellipseRGBA(renderer, width/2, height/2, radius-1, radius-1, colourAround[0], colourAround[1], colourAround[2], alpha)
	//ball.drawCircle(renderer, width/2, height/2, radius)
}

func DrawSimpleCircle(r *sdl.Renderer, x float64, y float64, cx float64, cy float64, radius float64) {
	px := cx + radius*math.Cos(0)
	py := cy + radius*math.Sin(0)

	for teta := float64(1); teta <= 360; teta++ {
		px = cx + radius*math.Cos(teta)
		py = cy + radius*math.Sin(teta)
		r.DrawLine(int32(px), int32(py), int32(x), int32(y))
		px = x
		py = y
	}
}

/*
*
int radius = 200;
int new_x = 0;
int new_y = 0;
int old_x =  SCR_CEN_X + 200;
int old_y = SCR_CEN_Y;
float step = (M_PI * 2) / 50;
*/
func DrawCircle(r *sdl.Renderer, radius float64, new_x float64, new_y float64, scr_cen_x float64, scr_cen_y float64) {
	old_x := scr_cen_x + 200
	old_y := scr_cen_y
	step := (math.Pi * 2) / 50

	r.SetDrawColor(0, 0, 0, 255)

	for theta := float64(0); theta <= (math.Pi * 2); theta += step {
		new_x = scr_cen_x + radius*math.Cos(theta)
		new_y = scr_cen_y - radius*math.Sin(theta)

		r.DrawLine(int32(old_x), int32(old_y), int32(new_x), int32(new_y))

		old_x = new_x
		old_y = new_y
	}

	new_x = scr_cen_x + (radius * math.Cos(0))
	new_y = scr_cen_y - (radius * math.Sin(0))
	r.DrawLine(int32(old_x), int32(old_y), int32(new_x), int32(new_y))

}

func (ball Ball) drawCircle(renderer *sdl.Renderer, x0 int32, y0 int32, radius int32) {
	x := radius - 1
	y := int32(0)
	dx := int32(1)
	dy := int32(1)
	err := dx - (radius << 1)
	for x >= y {
		renderer.DrawPoint(x0+x, y0+y)
		renderer.DrawPoint(x0+y, y0+x)
		renderer.DrawPoint(x0-y, y0+x)
		renderer.DrawPoint(x0-x, y0+y)
		renderer.DrawPoint(x0-x, y0-y)
		renderer.DrawPoint(x0-y, y0-x)
		renderer.DrawPoint(x0+y, y0-x)
		renderer.DrawPoint(x0+x, y0-y)
		if err <= 0 {
			y++
			err += dy
			dy += 2
		} else {
			x--
			dx += 2
			err += dx - (radius << 1)
		}
	}
}
