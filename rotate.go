package builder

import (
	"math"

	"github.com/deadsy/sdfx/sdf"
)

// RotateXCopy will create n copies, evenly rotated about the X axis.
func (b Builder) RotateXCopy(n int) Builder {
	return b.RotateXCopyOrigin(n, 0, 0, 0)
}

// RotateYCopy will create n copies, evenly rotated about the Y axis.
func (b Builder) RotateYCopy(n int) Builder {
	return b.RotateYCopyOrigin(n, 0, 0, 0)
}

// RotateZCopy will create n copies, evenly rotated about the Z axis.
func (b Builder) RotateZCopy(n int) Builder {
	return b.RotateZCopyOrigin(n, 0, 0, 0)
}

func (b Builder) rotateCopyOrigin(fn func(float64, float64, float64, float64) Builder, n int, x, y, z float64) Builder {
	copies := make([]sdf.SDF3, n)
	for i := range copies {
		copies[i] = fn(math.Pi*2/float64(n)*float64(i), x, y, z)
	}

	b.SDF3 = sdf.Union3D(copies...)
	return b
}

// RotateXCopyOrigin will create n copies, evenly rotated using the coordinates as the origin.
func (b Builder) RotateXCopyOrigin(n int, x, y, z float64) Builder {
	return b.rotateCopyOrigin(b.RotateXOrigin, n, x, y, z)
}

// RotateYCopyOrigin will create n copies, evenly rotated using the coordinates as the origin.
func (b Builder) RotateYCopyOrigin(n int, x, y, z float64) Builder {
	return b.rotateCopyOrigin(b.RotateYOrigin, n, x, y, z)
}

// RotateZCopyOrigin will create n copies, evenly rotated using the coordinates as the origin.
func (b Builder) RotateZCopyOrigin(n int, x, y, z float64) Builder {
	return b.rotateCopyOrigin(b.RotateZOrigin, n, x, y, z)
}

// RotateX will rotate r radians about the X axis with origin (0, 0, 0).
func (b Builder) RotateX(r float64) Builder {
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.RotateX(r))
	return b
}

// RotateY will rotate r radians about the Y axis with origin (0, 0, 0).
func (b Builder) RotateY(radians float64) Builder {
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.RotateY(radians))
	return b
}

// RotateZ will rotate r radians about the Z axis with origin (0, 0, 0).
func (b Builder) RotateZ(radians float64) Builder {
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.RotateZ(radians))
	return b
}

func (b Builder) rotateOrigin(r sdf.M44, x, y, z float64) Builder {
	p := sdf.V3{X: x, Y: y, Z: z}
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Translate3d(p.Negate()))
	b.SDF3 = sdf.Transform3D(b.SDF3, r)
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Translate3d(p))
	return b
}

// RotateXOrigin will rotate r radians about the X axis using the
// provided coordinates as the origin.
func (b Builder) RotateXOrigin(r, x, y, z float64) Builder {
	return b.rotateOrigin(sdf.RotateX(r), x, y, z)
}

// RotateYOrigin will rotate r radians about the Y axis using the
// provided coordinates as the origin.
func (b Builder) RotateYOrigin(r, x, y, z float64) Builder {
	return b.rotateOrigin(sdf.RotateY(r), x, y, z)
}

// RotateZOrigin will rotate r radians about the Z axis using the
// provided coordinates as the origin.
func (b Builder) RotateZOrigin(r, x, y, z float64) Builder {
	return b.rotateOrigin(sdf.RotateZ(r), x, y, z)
}
