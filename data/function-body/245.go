{
	colors := []color.NRGBA{
		{0, 0, 0, 0},
		{255, 255, 255, 255},
		{50, 100, 150, 255},
		{150, 100, 50, 200},
	}
	pixels := []pixel{
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0.196, 0.392, 0.588, 1},
		{0.588, 0.392, 0.196, 0.784},
	}

	img := image.NewNRGBA(image.Rect(-1, -2, 3, 2))
	pg := newPixelGetter(img)
	var row []pixel
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			img.Set(x, y, colors[x-img.Bounds().Min.X])
		}
		pg.getPixelRow(y, &row)
		if !comparePixelSlices(row, pixels, 0.005) {
			t.Errorf("getPixelRow y=%d %v %v", y, row, pixels)
		}
	}
}