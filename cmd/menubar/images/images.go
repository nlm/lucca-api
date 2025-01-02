package images

import _ "embed"

var (
	//go:embed small-blue-diamond.png
	SmallBlueDiamond []byte
	//go:embed large-blue-diamond.png
	LargeBlueDiamond []byte
	//go:embed small-orange-diamond.png
	SmallOrangeDiamond []byte
	//go:embed large-orange-diamond.png
	LargeOrangeDiamond []byte
	//go:embed warning.png
	Warning []byte
)
