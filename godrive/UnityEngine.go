package godrive

type Transform struct {
	position int
}

type Collision interface{}

type GameObjectClass struct {
	transform *Transform
}

var GameObject *GameObjectClass

func (*GameObjectClass) Find(string) *GameObjectClass {
	return &GameObjectClass{}
}

type Vector3 interface {
}

var Time struct {
	deltaTime int
}

var LineType struct {
	Continuous int
}
