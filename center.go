package builder

import "github.com/deadsy/sdfx/sdf"

// CenterX will translate the receiver so that the X center aligns
// with the provided value.
func (b Builder) CenterX(val float64) Builder {
	return b.Translate(
		val-b.BoundingBox().Center().X,
		0,
		0,
	)
}

// CenterY will translate the receiver so that the Y center aligns
// with the provided value.
func (b Builder) CenterY(val float64) Builder {
	return b.Translate(
		0,
		val-b.BoundingBox().Center().Y,
		0,
	)
}

// CenterZ will translate the receiver so that the Z center aligns
// with the provided value.
func (b Builder) CenterZ(val float64) Builder {
	return b.Translate(
		0,
		0,
		val-b.BoundingBox().Center().Z,
	)
}

// CenterXTo will translate the target so that the X center aligns
// with the X center of s. If s is nil, the X axis is used.
func (b Builder) CenterXTo(s sdf.SDF3) Builder { return b.CenterXToOffset(s, 0) }

// CenterXToOffset will translate the target so that the X center aligns
// with the X center of s, plus an offset. If s is nil, the X axis is used.
func (b Builder) CenterXToOffset(s sdf.SDF3, off float64) Builder {
	if s == nil {
		return b.CenterX(off)
	}
	return b.CenterX(s.BoundingBox().Center().X + off)
}

// CenterYTo will translate the target so that the Y center aligns
// with the Y center of s. If s is nil, the Y axis is used.
func (b Builder) CenterYTo(s sdf.SDF3) Builder { return b.CenterYToOffset(s, 0) }

// CenterYToOffset will translate the target so that the Y center aligns
// with the Y center of s, plus an offset. If s is nil, the Y axis is used.
func (b Builder) CenterYToOffset(s sdf.SDF3, off float64) Builder {
	if s == nil {
		return b.CenterY(off)
	}
	return b.CenterY(s.BoundingBox().Center().Y + off)
}

// CenterZTo will translate the target so that the Z center aligns
// with the Z center of s. If s is nil, the Z axis is used.
func (b Builder) CenterZTo(s sdf.SDF3) Builder { return b.CenterZToOffset(s, 0) }

// CenterZToOffset will translate the target so that the Z center aligns
// with the Z center of s, plus an offset. If s is nil, the Z axis is used.
func (b Builder) CenterZToOffset(s sdf.SDF3, off float64) Builder {
	if s == nil {
		return b.CenterZ(off)
	}
	return b.CenterZ(s.BoundingBox().Center().Z + off)
}
