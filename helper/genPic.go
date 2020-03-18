package helper

import "golang.org/x/tour/pic"

// Pic Generate a picture
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for iy := 0; iy < dx; iy++ {
		result[iy] = make([]uint8, dx)
		for ix := 0; ix < dx; ix++ {
			result[iy][ix] = uint8(ix) * uint8(iy)
		}

	}
	return result
}

// ShowPic prints Pic
func ShowPic() {
	pic.Show(Pic)
}
