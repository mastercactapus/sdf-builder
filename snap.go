package builder

// SizeX returns the bounding box size for the X dimension.
func (b Builder) SizeX() float64 { return b.BoundingBox().Size().X }

// SizeY returns the bounding box size for the Y dimension.
func (b Builder) SizeY() float64 { return b.BoundingBox().Size().Y }

// SizeZ returns the bounding box size for the Z dimension.
func (b Builder) SizeZ() float64 { return b.BoundingBox().Size().Z }

// MinX will return the low X-coordinate of the bounding box.
func (b Builder) MinX() float64 { return b.BoundingBox().Min.X }

// MidX will return the center X-coordinate of the bounding box.
func (b Builder) MidX() float64 { return b.BoundingBox().Center().X }

// MaxX will return the high X-coordinate of the bounding box.
func (b Builder) MaxX() float64 { return b.BoundingBox().Max.X }

// SnapMinX will snap the low X-edge to the provided value.
func (b Builder) SnapMinX(val float64) Builder { return b.Translate(val-b.MinX(), 0, 0) }

// SnapMidX will snap the X-center to the provided value.
func (b Builder) SnapMidX(val float64) Builder { return b.Translate(val-b.MidX(), 0, 0) }

// SnapMaxX will snap the high X-edge to the provided value.
func (b Builder) SnapMaxX(val float64) Builder { return b.Translate(val-b.MaxX(), 0, 0) }

// MinY will return the low Y-coordinate of the bounding box.
func (b Builder) MinY() float64 { return b.BoundingBox().Min.Y }

// MidY will return the center Y-coordinate of the bounding box.
func (b Builder) MidY() float64 { return b.BoundingBox().Center().Y }

// MaxY will return the high Y-coordinate of the bounding box.
func (b Builder) MaxY() float64 { return b.BoundingBox().Max.Y }

// SnapMinY will snap the low Y-edge to the provided value.
func (b Builder) SnapMinY(val float64) Builder { return b.Translate(0, val-b.MinY(), 0) }

// SnapMidY will snap the Y-center to the provided value.
func (b Builder) SnapMidY(val float64) Builder { return b.Translate(0, val-b.MidY(), 0) }

// SnapMaxY will snap the high Y-edge to the provided value.
func (b Builder) SnapMaxY(val float64) Builder { return b.Translate(0, val-b.MaxY(), 0) }

// MinZ will return the low Z-coordinate of the bounding box.
func (b Builder) MinZ() float64 { return b.BoundingBox().Min.Z }

// MidZ will return the center Z-coordinate of the bounding box.
func (b Builder) MidZ() float64 { return b.BoundingBox().Center().Z }

// MaxZ will return the high Z-coordinate of the bounding box.
func (b Builder) MaxZ() float64 { return b.BoundingBox().Max.Z }

// SnapMinZ will snap the low Z-edge to the provided value.
func (b Builder) SnapMinZ(val float64) Builder { return b.Translate(0, 0, val-b.MinZ()) }

// SnapMidZ will snap the Z-center to the provided value.
func (b Builder) SnapMidZ(val float64) Builder { return b.Translate(0, 0, val-b.MidZ()) }

// SnapMaxZ will snap the high Z-edge to the provided value.
func (b Builder) SnapMaxZ(val float64) Builder { return b.Translate(0, 0, val-b.MaxZ()) }
