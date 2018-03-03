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

// NewRTriangle creates a new right-triangle with the provided
// opposite and adjacent lengths.
func NewRTriangle(oLen, aLen, height float64) Builder {
	return Builder{
		SDF3: sdf.Extrude3D(sdf.Polygon2D([]sdf.V2{
			{},
			{X: oLen},
			{Y: aLen},
		}), height),
	}.Translate(-oLen/2, -aLen/2, height/2)
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

// Translate will move the object by the given amounts.
func (b Builder) Translate(x, y, z float64) Builder {
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Translate3d(
		sdf.V3{X: x, Y: y, Z: z},
	))
	return b
}

// Difference will subtract one or more objects from the receiver.
func (b Builder) Difference(other ...sdf.SDF3) Builder {
	for _, s := range other {
		b.SDF3 = sdf.Difference3D(b.SDF3, s)
	}
	return b
}

// Union will join multiple objects to the receiver.
func (b Builder) Union(other ...sdf.SDF3) Builder {
	other = append(other, b.SDF3)
	b.SDF3 = sdf.Union3D(other...)
	return b
}

// Intersection results in the overlapping parts of all objects.
func (b Builder) Intersection(other ...sdf.SDF3) Builder {
	for _, s := range other {
		b.SDF3 = sdf.Intersect3D(b.SDF3, s)
	}
	return b
}

// Mirror will mirror the object over the specified axis.
func (b Builder) Mirror(x, y, z bool) Builder {
	v := sdf.V3{
		X: 1, Y: 1, Z: 1,
	}
	if x {
		v.X = -1
	}
	if y {
		v.Y = -1
	}
	if z {
		v.Z = -1
	}
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Scale3d(v))
	return b
}
