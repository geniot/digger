package utils

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

func AaellipseRGBA(renderer *sdl.Renderer, x int32, y int32, rx int32, ry int32, r uint8, g uint8, b uint8, a uint8) int {
	var result int
	var i int32
	var a2, b2, ds, dt, dxt, t, s, d int32
	var xp, yp, xs, ys, dyt, od, xx, yy, xc2, yc2 int32
	var cp float32
	var sab float64
	var weight, iWeight uint32

	/*
	 * Sanity check radii
	 */
	if (rx < 0) || (ry < 0) {
		return -1
	}

	/*
	 * Special cases for rx=0 and/or ry=0: draw a hline/vline/pixel
	 */
	if rx == 0 {
		if ry == 0 {
			return pixelRGBA(renderer, x, y, r, g, b, a)
		} else {
			return vlineRGBA(renderer, x, y-ry, y+ry, r, g, b, a)
		}
	} else {
		if ry == 0 {
			return hlineRGBA(renderer, x-rx, x+rx, y, r, g, b, a)
		}
	}

	/* Variable setup */
	a2 = rx * rx
	b2 = ry * ry

	ds = 2 * a2
	dt = 2 * b2

	xc2 = 2 * x
	yc2 = 2 * y

	sab = math.Sqrt(float64(a2 + b2))
	od = int32(math.Round(sab*0.01) + 1) /* introduce some overdraw */
	dxt = int32(math.Round(float64(a2)/sab)) + od

	t = 0
	s = -2 * a2 * ry
	d = 0

	xp = x
	yp = y - ry

	/* Draw */
	result = 0
	if a == 255 {
		err := renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
		if err != nil {
			result |= 1
		}
	} else {
		err := renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		if err != nil {
			result |= 1
		}
	}

	/* "End points" */
	result |= pixelRGBA(renderer, xp, yp, r, g, b, a)
	result |= pixelRGBA(renderer, xc2-xp, yp, r, g, b, a)
	result |= pixelRGBA(renderer, xp, yc2-yp, r, g, b, a)
	result |= pixelRGBA(renderer, xc2-xp, yc2-yp, r, g, b, a)

	for i = 1; i <= dxt; i++ {
		xp--
		d += t - b2

		if d >= 0 {
			ys = yp - 1
		} else if (d - s - a2) > 0 {
			if (2*d - s - a2) >= 0 {
				ys = yp + 1
			} else {
				ys = yp
				yp++
				d -= s + a2
				s += ds
			}
		} else {
			yp++
			ys = yp + 1
			d -= s + a2
			s += ds
		}

		t -= dt

		/* Calculate alpha */
		if s != 0 {
			cp = float32(math.Abs(float64(d)) / math.Abs(float64(s)))
			if cp > 1.0 {
				cp = 1.0
			}
		} else {
			cp = 1.0
		}

		/* Calculate weights */
		weight = uint32(cp * 255)
		iWeight = 255 - weight

		/* Upper half */
		xx = xc2 - xp
		result |= pixelRGBAWeight(renderer, xp, yp, r, g, b, a, iWeight)
		result |= pixelRGBAWeight(renderer, xx, yp, r, g, b, a, iWeight)

		result |= pixelRGBAWeight(renderer, xp, ys, r, g, b, a, weight)
		result |= pixelRGBAWeight(renderer, xx, ys, r, g, b, a, weight)

		/* Lower half */
		yy = yc2 - yp
		result |= pixelRGBAWeight(renderer, xp, yy, r, g, b, a, iWeight)
		result |= pixelRGBAWeight(renderer, xx, yy, r, g, b, a, iWeight)

		yy = yc2 - ys
		result |= pixelRGBAWeight(renderer, xp, yy, r, g, b, a, weight)
		result |= pixelRGBAWeight(renderer, xx, yy, r, g, b, a, weight)
	}

	/* Replaces original approximation code dyt = abs(yp - yc); */
	dyt = int32(math.Round(float64(b2)/sab)) + od

	for i = 1; i <= dyt; i++ {
		yp++
		d -= s + a2

		if d <= 0 {
			xs = xp + 1
		} else if (d + t - b2) < 0 {
			if (2*d + t - b2) <= 0 {
				xs = xp - 1
			} else {
				xs = xp
				xp--
				d += t - b2
				t -= dt
			}
		} else {
			xp--
			xs = xp - 1
			d += t - b2
			t -= dt
		}

		s += ds

		/* Calculate alpha */
		if t != 0 {
			cp = float32(math.Abs(float64(d)) / math.Abs(float64(t)))
			if cp > 1.0 {
				cp = 1.0
			}
		} else {
			cp = 1.0
		}

		/* Calculate weight */
		weight = uint32(cp * 255)
		iWeight = 255 - weight

		/* Left half */
		xx = xc2 - xp
		yy = yc2 - yp
		result |= pixelRGBAWeight(renderer, xp, yp, r, g, b, a, iWeight)
		result |= pixelRGBAWeight(renderer, xx, yp, r, g, b, a, iWeight)

		result |= pixelRGBAWeight(renderer, xp, yy, r, g, b, a, iWeight)
		result |= pixelRGBAWeight(renderer, xx, yy, r, g, b, a, iWeight)

		/* Right half */
		xx = xc2 - xs
		result |= pixelRGBAWeight(renderer, xs, yp, r, g, b, a, weight)
		result |= pixelRGBAWeight(renderer, xx, yp, r, g, b, a, weight)

		result |= pixelRGBAWeight(renderer, xs, yy, r, g, b, a, weight)
		result |= pixelRGBAWeight(renderer, xx, yy, r, g, b, a, weight)
	}

	return result
}

func pixelRGBAWeight(renderer *sdl.Renderer, x int32, y int32, r uint8, g uint8, b uint8, a uint8, weight uint32) int {
	var ax = uint32(a)
	ax = (ax * weight) >> 8
	if ax > 255 {
		a = 255
	} else {
		a = uint8(ax & 0x000000ff)
	}
	return pixelRGBA(renderer, x, y, r, g, b, a)
}

func pixelRGBA(renderer *sdl.Renderer, x int32, y int32, r uint8, g uint8, b uint8, a uint8) int {
	var result = 0

	if a == 255 {
		err := renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
		if err != nil {
			result |= 1
		}
	} else {
		err := renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		if err != nil {
			result |= 1
		}
	}
	err := renderer.SetDrawColor(r, g, b, a)
	if err != nil {
		result |= 1
	}
	err = renderer.DrawPoint(x, y)
	if err != nil {
		result |= 1
	}
	return result
}

func vlineRGBA(renderer *sdl.Renderer, x int32, y1 int32, y2 int32, r uint8, g uint8, b uint8, a uint8) int {
	var result = 0

	if a == 255 {
		err := renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
		if err != nil {
			result |= 1
		}
	} else {
		err := renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		if err != nil {
			result |= 1
		}
	}
	err := renderer.SetDrawColor(r, g, b, a)
	if err != nil {
		result |= 1
	}
	err = renderer.DrawLine(x, y1, x, y2)
	if err != nil {
		result |= 1
	}
	return result
}

func hlineRGBA(renderer *sdl.Renderer, x1 int32, x2 int32, y int32, r uint8, g uint8, b uint8, a uint8) int {
	var result = 0

	if a == 255 {
		err := renderer.SetDrawBlendMode(sdl.BLENDMODE_NONE)
		if err != nil {
			result |= 1
		}
	} else {
		err := renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
		if err != nil {
			result |= 1
		}
	}
	err := renderer.SetDrawColor(r, g, b, a)
	if err != nil {
		result |= 1
	}
	err = renderer.DrawLine(x1, y, x2, y)
	if err != nil {
		result |= 1
	}
	return result
}
