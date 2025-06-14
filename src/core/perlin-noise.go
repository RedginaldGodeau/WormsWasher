package core

import (
	"math"
	"math/rand"
)

type Perlin struct {
	permutation []int
	p           []int
	size        int
	randMax     int
}

func NewPerlin() *Perlin {
	p := &Perlin{
		permutation: []int{
			151, 160, 137, 91, 90, 15,
			131, 13, 201, 95, 96, 53, 194, 233, 7, 225, 140, 36, 103, 30, 69, 142, 8, 99, 37, 240, 21, 10, 23,
			190, 6, 148, 247, 120, 234, 75, 0, 26, 197, 62, 94, 252, 219, 203, 117, 35, 11, 32, 57, 177, 33,
			88, 237, 149, 56, 87, 174, 20, 125, 136, 171, 168, 68, 175, 74, 165, 71, 134, 139, 48, 27, 166,
			77, 146, 158, 231, 83, 111, 229, 122, 60, 211, 133, 230, 220, 105, 92, 41, 55, 46, 245, 40, 244,
			102, 143, 54, 65, 25, 63, 161, 1, 216, 80, 73, 209, 76, 132, 187, 208, 89, 18, 169, 200, 196,
			135, 130, 116, 188, 159, 86, 164, 100, 109, 198, 173, 186, 3, 64, 52, 217, 226, 250, 124, 123,
			5, 202, 38, 147, 118, 126, 255, 82, 85, 212, 207, 206, 59, 227, 47, 16, 58, 17, 182, 189, 28, 42,
			223, 183, 170, 213, 119, 248, 152, 2, 44, 154, 163, 70, 221, 153, 101, 155, 167, 43, 172, 9,
			129, 22, 39, 253, 19, 98, 108, 110, 79, 113, 224, 232, 178, 185, 112, 104, 218, 246, 97, 228,
			251, 34, 242, 193, 238, 210, 144, 12, 191, 179, 162, 241, 81, 51, 145, 235, 249, 14, 239, 107,
			49, 192, 214, 31, 181, 199, 106, 157, 184, 84, 204, 176, 115, 121, 50, 45, 127, 4, 150, 254,
			138, 236, 205, 93, 222, 114, 67, 29, 24, 72, 243, 141, 128, 195, 78, 66, 215, 61, 156, 180,
		},
		size:    256,
		randMax: 256,
	}
	p.load()
	return p
}

func (p *Perlin) load() {
	p.p = make([]int, 2*p.size)
	for i := 0; i < p.size; i++ {
		p.p[i] = p.permutation[i]
		p.p[p.size+i] = p.p[i]
	}
}

func (p *Perlin) Seed(seed int64) {
	rand.Seed(seed)
	p.permutation = make([]int, p.size)
	for i := 0; i < p.size; i++ {
		p.permutation[i] = i
	}
	n := len(p.permutation)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		p.permutation[i], p.permutation[j] = p.permutation[j], p.permutation[i]
	}
	p.load()
}

func fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}

func lerp(t, a, b float64) float64 {
	return a + t*(b-a)
}

func grad(hash int, x, y, z float64) float64 {
	h := hash & 15
	u := x
	if h < 8 {
		u = x
	} else {
		u = y
	}
	v := y
	if h < 4 {
		v = y
	} else if h == 12 || h == 14 {
		v = x
	} else {
		v = z
	}
	if (h & 1) == 0 {
		// Do nothing, u is positive
	} else {
		u = -u
	}
	if (h & 2) == 0 {
		// Do nothing, v is positive
	} else {
		v = -v
	}
	return u + v
}

func (p *Perlin) Noise(x, y, z float64) float64 {
	X := int(math.Floor(x)) & 255
	Y := int(math.Floor(y)) & 255
	Z := int(math.Floor(z)) & 255
	xf := x - math.Floor(x)
	yf := y - math.Floor(y)
	zf := z - math.Floor(z)
	u := fade(xf)
	v := fade(yf)
	w := fade(zf)
	A := p.p[X] + Y
	AA := p.p[A] + Z
	AB := p.p[A+1] + Z
	B := p.p[X+1] + Y
	BA := p.p[B] + Z
	BB := p.p[B+1] + Z

	return lerp(w, lerp(v, lerp(u, grad(p.p[AA], xf, yf, zf),
		grad(p.p[BA], xf-1, yf, zf)),
		lerp(u, grad(p.p[AB], xf, yf-1, zf),
			grad(p.p[BB], xf-1, yf-1, zf))),
		lerp(v, lerp(u, grad(p.p[AA+1], xf, yf, zf-1),
			grad(p.p[BA+1], xf-1, yf, zf-1)),
			lerp(u, grad(p.p[AB+1], xf, yf-1, zf-1),
				grad(p.p[BB+1], xf-1, yf-1, zf-1))))
}
