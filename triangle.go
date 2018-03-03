package builder

import (
	"math"

	"github.com/deadsy/sdfx/sdf"
)

// NewSlice will create a triangle with the given vertex length
// and origin angle.
func NewSlice(height, vertexLen, angle float64) Builder {
	o := math.Tan(angle/2) * vertexLen
	return Builder{
		SDF3: sdf.Extrude3D(sdf.Polygon2D([]sdf.V2{
			{},
			{X: vertexLen, Y: o},
			{X: vertexLen, Y: -o},
		}), height),
	}.Translate(0, 0, height/2)
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
