{
	testData := []struct {
		desc           string
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{

		{
			"sobel 0x0",
			image.Rect(0, 0, 0, 0),
			image.Rect(0, 0, 0, 0),
			[]uint8{},
			[]uint8{},
		},
		{
			"sobel 6x6",
			image.Rect(-1, -1, 5, 5),
			image.Rect(0, 0, 6, 6),
			[]uint8{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x99, 0x99, 0x99, 0x99,
				0x00, 0x00, 0x99, 0x99, 0x99, 0x99,
				0x00, 0x00, 0x99, 0x99, 0x99, 0x99,
				0x00, 0x00, 0x99, 0x99, 0x99, 0x99,
			},
			[]uint8{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0xd8, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xff, 0xff, 0xff, 0xff, 0xff,
				0x00, 0xff, 0xff, 0x00, 0x00, 0x00,
				0x00, 0xff, 0xff, 0x00, 0x00, 0x00,
				0x00, 0xff, 0xff, 0x00, 0x00, 0x00,
			},
		},
	}

	for _, d := range testData {
		src := image.NewGray(d.srcb)
		src.Pix = d.srcPix

		f := Sobel()
		dst := image.NewGray(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}
}