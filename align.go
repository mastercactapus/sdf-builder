package builder

import "github.com/deadsy/sdfx/sdf"

// AlignXTo will cause the higher X edge of the receiver to align
// with the higher X edge of s. If s is nil, the X axis is used.
func (b Builder) AlignXTo(s sdf.SDF3) Builder { return b.AlignXToOffset(s, 0) }

// AlignXToOffset will cause the higher X edge of the receiver to align
// with the higher X edge of s plus an offset. If s is nil, the X axis is used.
func (b Builder) AlignXToOffset(s sdf.SDF3, off float64) Builder {
	if s == nil {
		return b.AlignX(off)
	}
	return b.AlignX(s.BoundingBox().Max.X + off)
}

// AlignYTo will cause the higher Y edge of the receiver to align
// with the higher Y edge of s. If s is nil, the X axis is used.
func (b Builder) AlignYTo(s sdf.SDF3) Builder { return b.AlignYToOffset(s, 0) }

// AlignYToOffset will cause the higher Y edge of the receiver to align
// with the higher Y edge of s plus an offset. If s is nil, the Y axis is used.
func (b Builder) AlignYToOffset(s sdf.SDF3, off float64) Builder {
	if s == nil {
		return b.AlignY(off)
	}
	return b.AlignY(s.BoundingBox().Max.Y + off)
}

// AlignZTo will cause the higher Z edge of the receiver to align
// with the higher Z edge of s. If s is nil, the X axis is used.
func (b Builder) AlignZTo(s sdf.SDF3) Builder { return b.AlignZToOffset(s, 0) }

// AlignZToOffset will cause the higher Z edge of the receiver to align
// with the higher Z edge of s plus an offset. If s is nil, the Z axis is used.
func (b Builder) AlignZToOffset(s sdf.SDF3, off float64) Builder {
	if s == nil {
		return b.AlignZ(off)
	}
	return b.AlignZ(s.BoundingBox().Max.Z + off)
}

// AlignX will cause the higher X edge of the receiver to match
// the provided value.
func (b Builder) AlignX(val float64) Builder {
	return b.Translate(
		val-b.BoundingBox().Max.X,
		0,
		0,
	)
}

// AlignY will cause the higher Y edge of the receiver to match
// the provided value.
func (b Builder) AlignY(val float64) Builder {
	return b.Translate(
		0,
		val-b.BoundingBox().Max.Y,
		0,
	)
}

// AlignZ will cause the higher Z edge of the receiver to match
// the provided value.
func (b Builder) AlignZ(val float64) Builder {
	return b.Translate(
		0,
		0,
		val-b.BoundingBox().Max.Z,
	)
}
