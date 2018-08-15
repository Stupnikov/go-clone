{
	img0 := image.NewGray(image.Rect(-1, -1, 3, 1))
	img0.Pix = []uint8{
		1, 2, 3, 4,
		5, 6, 7, 8,
	}
	img1Exp := image.NewGray(image.Rect(0, 0, 2, 4))
	img1Exp.Pix = []uint8{
		5, 1,
		6, 2,
		7, 3,
		8, 4,
	}

	f := Rotate270()
	img1 := image.NewGray(f.Bounds(img0.Bounds()))
	f.Draw(img1, img0, nil)

	if img1.Bounds().Size() != img1Exp.Bounds().Size() {
		t.Errorf("expected %v got %v", img1Exp.Bounds().Size(), img1.Bounds().Size())
	}
	if !comparePix(img1Exp.Pix, img1.Pix) {
		t.Errorf("expected %v got %v", img1Exp.Pix, img1.Pix)
	}
}