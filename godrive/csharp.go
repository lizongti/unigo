package godrive

type object interface{}

type VectorLine struct {
	a string
	b []Vector3
	c float64
	d int
}

func (*VectorLine) MakeSpline([]Vector3, int, bool) {

}

func (*VectorLine) Draw3D() {}

func (*VectorLine) GetPoint3D01(float32) int {
	return 0
}
