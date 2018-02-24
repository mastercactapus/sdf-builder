package builder

import "github.com/deadsy/sdfx/sdf"

// RotateZCopy will create n copies, evenly rotated about the Z axis.
func (b Builder) RotateZCopy(n int) Builder {
	b.SDF3 = sdf.RotateCopy3D(b.SDF3, n)
	return b
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

// RotateXOrigin will rotate r radians about the X axis using the
// provided coordinates as the origin.
func (b Builder) RotateXOrigin(r, x, y, z float64) Builder {
	p := sdf.V3{X: x, Y: y, Z: z}
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Translate3d(p.Negate()))
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.RotateX(r))
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Translate3d(p))
	return b
}

// RotateYOrigin will rotate r radians about the Y axis using the
// provided coordinates as the origin.
func (b Builder) RotateYOrigin(r, x, y, z float64) Builder {
	p := sdf.V3{X: x, Y: y, Z: z}
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Translate3d(p.Negate()))
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.RotateY(r))
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Translate3d(p))
	return b
}

// RotateZOrigin will rotate r radians about the Z axis using the
// provided coordinates as the origin.
func (b Builder) RotateZOrigin(r, x, y, z float64) Builder {
	p := sdf.V3{X: x, Y: y, Z: z}
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Translate3d(p.Negate()))
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.RotateZ(r))
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Translate3d(p))
	return b
}
