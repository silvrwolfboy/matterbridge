// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WatermarkLayout undocumented
type WatermarkLayout int

const (
	// WatermarkLayoutVHorizontal undocumented
	WatermarkLayoutVHorizontal WatermarkLayout = 0
	// WatermarkLayoutVDiagonal undocumented
	WatermarkLayoutVDiagonal WatermarkLayout = 1
)

// WatermarkLayoutPHorizontal returns a pointer to WatermarkLayoutVHorizontal
func WatermarkLayoutPHorizontal() *WatermarkLayout {
	v := WatermarkLayoutVHorizontal
	return &v
}

// WatermarkLayoutPDiagonal returns a pointer to WatermarkLayoutVDiagonal
func WatermarkLayoutPDiagonal() *WatermarkLayout {
	v := WatermarkLayoutVDiagonal
	return &v
}
