package core

type Vector2 struct {
	x float64
	y float64
}

func NewVector(x, y float64) Vector2 {
	return Vector2{x, y}
}

func (v Vector2) Get() (float64, float64) {
	return v.x, v.y
}

func (v Vector2) Equal(vector Vector2) bool {
	return v.x == vector.x && v.y == vector.y
}

func (v *Vector2) Add(vector Vector2) {
	v.x += vector.x
	v.y += vector.y
}

func (v *Vector2) Mult(n float64) Vector2 {
	return NewVector(v.x*n, v.y*n)
}

func (v *Vector2) X() float64 {
	return v.x
}
func (v *Vector2) Y() float64 {
	return v.y
}

func (v *Vector2) XInt() int {
	return int(v.x)
}
func (v *Vector2) YInt() int {
	return int(v.y)
}
