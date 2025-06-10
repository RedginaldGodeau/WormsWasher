package fluit

type Vector2[V float64 | int] struct {
	X V
	Y V
}

func NewVector2[V float64 | int](x V, y V) Vector2[V] {
	return Vector2[V]{x, y}
}

func VectorAdd[V float64 | int](v1 Vector2[V], v2 Vector2[V]) Vector2[V] {
	return Vector2[V]{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func VectorSum[V float64 | int](v1 Vector2[V], mult V) Vector2[V] {
	return Vector2[V]{X: v1.X + mult, Y: v1.Y + mult}
}
