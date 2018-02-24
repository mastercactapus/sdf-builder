package builder

import "github.com/deadsy/sdfx/sdf"

// SnapX will cause the lower X edge of the receiver to match
// the provided value.
func (b Builder) SnapX(val float64) Builder {
	return b.Translate(
		val-b.BoundingBox().Min.X,
		0,
		0,
	)
}

// SnapY will cause the lower Y edge of the receiver to match
// the provided value.
func (b Builder) SnapY(val float64) Builder {
	return b.Translate(
		0,
		val-b.BoundingBox().Min.Y,
		0,
	)
}

// SnapZ will cause the lower Z edge of the receiver to match
// the provided value.
func (b Builder) SnapZ(val float64) Builder {
	return b.Translate(
		0,
		0,
		val-b.BoundingBox().Min.Z,
	)
}

// SnapXTo will cause the lower X edge of the receiver to match
// the higher X edge of s. If s is nil, the X axis is used.
func (b Builder) SnapXTo(s sdf.SDF3) Builder { return b.SnapXToOffset(s, 0) }

// SnapXToOffset will cause the lower X edge of the receiver to match
// the higher X edge of s, plus an offset. If s is nil, the X axis is used.
func (b Builder) SnapXToOffset(s sdf.SDF3, off float64) Builder {
	if s == nil {
		return b.SnapX(off)
	}
	return b.SnapX(s.BoundingBox().Max.X + off)
}

// SnapYTo will cause the lower Y edge of the receiver to match
// the higher Y edge of s. If s is nil, the Y axis is used.
func (b Builder) SnapYTo(s sdf.SDF3) Builder { return b.SnapYToOffset(s, 0) }

// SnapYToOffset will cause the lower Y edge of the receiver to match
// the higher Y edge of s, plus an offset. If s is nil, the Y axis is used.
func (b Builder) SnapYToOffset(s sdf.SDF3, off float64) Builder {
	if s == nil {
		return b.SnapY(off)
	}
	return b.SnapY(s.BoundingBox().Max.Y + off)
}

// SnapZTo will cause the lower Z edge of the receiver to match
// the higher Z edge of s. If s is nil, the Z axis is used.
func (b Builder) SnapZTo(s sdf.SDF3) Builder { return b.SnapZToOffset(s, 0) }

// SnapZToOffset will cause the lower Z edge of the receiver to match
// the higher Z edge of s, plus an offset. If s is nil, the Z axis is used.
func (b Builder) SnapZToOffset(s sdf.SDF3, off float64) Builder {
	if s == nil {
		return b.SnapZ(off)
	}
	return b.SnapZ(s.BoundingBox().Max.Z + off)
}
