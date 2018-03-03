package builder

import (
	"github.com/deadsy/sdfx/sdf"
)

// Builder is used to create and manipulate 3D objects.
type Builder struct {
	sdf.SDF3
}

// NewHexagon creates a hexagon with the given size (diameter) and height.
func NewHexagon(size, height float64) Builder {
	boxWidth := size / 1.75
	return NewBox(size/2, boxWidth, height).SnapMinX(0).RotateZCopy(6)
}

// NewBox creates a new cubeoid with the given dimensions.
func NewBox(x, y, z float64) Builder {
	return Builder{SDF3: sdf.Box3D(sdf.V3{
		X: x, Y: y, Z: z,
	}, 0)}.Translate(0, 0, z/2)
}

// NewRoundedBox creates a new cubeoid with the given dimensions.
// Edges will be rounded by the provided r value.
func NewRoundedBox(x, y, z, r float64) Builder {
	return Builder{SDF3: sdf.Box3D(sdf.V3{
		X: x, Y: y, Z: z,
	}, r)}.Translate(0, 0, z/2)
}

// NewCylinder creates a new cylinder shape with the provided
// height and diameter.
func NewCylinder(height, diameter float64) Builder {
	return Builder{
		SDF3: sdf.Cylinder3D(height, diameter/2, 0),
	}.Translate(0, 0, height/2)
}

// NewCone creates a new cone shape with the provided start and end diameter.
func NewCone(height, d1, d2 float64) Builder {
	return Builder{
		SDF3: sdf.Cone3D(height, d1/2, d2/2, 0),
	}.Translate(0, 0, height/2)
}
