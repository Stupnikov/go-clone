{
	testData := []struct {
		desc           string
		r, g, b        float32
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{

		{
			"color balance 0, 0, 0, 0x0",
			0, 0, 0,
			image.Rect(0, 0, 0, 0),
			image.Rect(0, 0, 0, 0),
			[]uint8{},
			[]uint8{},
		},
		{
			"color balance 0, 0, 0",
			0, 0, 0,
			image.Rect(-1, -1, 1, 3),
			image.Rect(0, 0, 2, 4),
			[]uint8{
				0x00, 0x10, 0x20, 0x30, 0x99, 0x66, 0x33, 0xFF,
				0xF0, 0xE0, 0xD0, 0xC0, 0x11, 0x66, 0xBB, 0x00,
				0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
				0xE0, 0x50, 0xA0, 0x77, 0xEE, 0xFE, 0xEE, 0xFD,
			},
			[]uint8{
				0x00, 0x10, 0x20, 0x30, 0x99, 0x66, 0x33, 0xFF,
				0xF0, 0xE0, 0xD0, 0xC0, 0x11, 0x66, 0xBB, 0x00,
				0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
				0xE0, 0x50, 0xA0, 0x77, 0xEE, 0xFE, 0xEE, 0xFD,
			},
		},

		{
			"color balance 10, -20, 200",
			10, -20, 200,
			image.Rect(-1, -1, 1, 3),
			image.Rect(0, 0, 2, 4),
			[]uint8{
				0x00, 0x10, 0x20, 0x30, 0x99, 0x66, 0x33, 0xFF,
				0xF0, 0xE0, 0xD0, 0xC0, 0x11, 0x66, 0xBB, 0x00,
				0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF,
				0xE0, 0x50, 0xA0, 0x77, 0xEE, 0xFE, 0xEE, 0xFD,
			},
			[]uint8{
				0x00, 0x0D, 0x60, 0x30, 0xA8, 0x52, 0x99, 0xFF,
				0xFF, 0xB3, 0xFF, 0xC0, 0x13, 0x52, 0xFF, 0x00,
				0x00, 0x00, 0x00, 0x00, 0xFF, 0xCC, 0xFF, 0xFF,
				0xF6, 0x40, 0xFF, 0x77, 0xFF, 0xCB, 0xFF, 0xFD,
			},
		},
	}

	for _, d := range testData {
		src := image.NewNRGBA(d.srcb)
		src.Pix = d.srcPix

		f := ColorBalance(d.r, d.g, d.b)
		dst := image.NewNRGBA(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}
}