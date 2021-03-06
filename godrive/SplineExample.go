package godrive

// TODO: Modify go2cs to properly pickup the following directives
//go2cs: inject-code[using System.Collections.Generic;]
//go2cs: inject-code[using UnityEngine;]
//go2cs: inject-code[using Vectrosity;]

import (
	"fmt"
)

// TODO: Directive that will call a post-build utility on specifed target, in this case
// wrapping converted SplineFollow3D structure with a class that inherits MonoBehaviour
//go2cs: post-build[unity-target.exe type=MonoBehaviour filename=SplineFollow3D.cs namespace=SplineExample]
type SplineFollow3D struct {
	Segments int // <- these public fields will be exposed to Unity editor
	DoLoop   bool
	Cube     Transform
	Speed    float32
}

type SplineIterator struct {
	source *SplineFollow3D
	line   VectorLine
	dist   float32
}

func (behaviour *SplineFollow3D) Awake() {
	// Set initial default values, Unity will serialize changes made in editor
	if behaviour.Segments == 0 {
		behaviour.Segments = 250
		behaviour.DoLoop = true
		behaviour.Speed = 0.05
	}
}

func (behaviour *SplineFollow3D) Start() IEnumerator {
	// Note that VectorLine expects a .NET List of Vector3 objects - this does not need
	// to exist here in Go (not compiling here), it just needs to be valid Go syntax for
	// the go2cs conversion process. Unity will compile translated C# version of code.
	// TODO: Fix Golang grammar to support new generics syntax, e.g.: List[Vector3] => List<Vector3>
	points := []Vector3{behaviour.Segments + 1}

	var splinePoints []Vector3
	i := 1

	obj := GameObject.Find(fmt.Sprintf("Sphere%d", i))
	i += 1

	for obj != nil {
		splinePoints = append(splinePoints, obj.transform.position)
		GameObject.Find(fmt.Sprintf("Sphere%d", i))
		i += 1
	}

	line := VectorLine{"Spline", points, 2.0, LineType.Continuous}
	line.MakeSpline(splinePoints, behaviour.Segments, behaviour.DoLoop)
	line.Draw3D()

	return &SplineIterator{behaviour, line, 0.0}
}

func (iterator *SplineIterator) MoveNext() bool {
	behaviour := iterator.source

	// Make the cube "ride" the spline at a constant speed
	if iterator.dist < 1.0 {
		iterator.dist += float32(Time.deltaTime) * behaviour.Speed
		behaviour.Cube.position = iterator.line.GetPoint3D01(iterator.dist)
	} else if behaviour.DoLoop {
		iterator.Reset()
	} else {
		return false
	}

	return true
}

func (iterator *SplineIterator) Reset() {
	iterator.dist = 0.0
}

func (iterator *SplineIterator) Current() object {
	return nil
}
