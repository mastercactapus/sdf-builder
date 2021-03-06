package builder

import "github.com/deadsy/sdfx/sdf"

// Translate will move the object by the given amounts.
func (b Builder) Translate(x, y, z float64) Builder {
	b.SDF3 = sdf.Transform3D(b.SDF3, sdf.Translate3d(
		sdf.V3{X: x, Y: y, Z: z},
	))
	return b
}

// TranslateCopy will move the object by the given amounts, leaving
// the original behind (as a copy/union).
func (b Builder) TranslateCopy(x, y, z float64) Builder {
	return b.TranslateCopyN(2, x, y, z)
}

// TranslateCopyN will move the object by the given amounts, leaving
// the original behind (as a copy/union), performing the operation n times.
func (b Builder) TranslateCopyN(n int, x, y, z float64) Builder {
	copies := make([]sdf.SDF3, n)
	copies[0] = b.SDF3
	for i, c := range copies[:n-1] {
		copies[i+1] = sdf.Transform3D(c, sdf.Translate3d(
			sdf.V3{X: x, Y: y, Z: z},
		))
	}
	b.SDF3 = sdf.Union3D(copies...)
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
