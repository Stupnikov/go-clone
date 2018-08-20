{
	l := &constraint.Layouter{}
	l.Solve(func(s *constraint.Solver) {
		if v.Axis == layout.AxisY {
			s.Height(200)
		} else {
			s.Width(200)
		}
	})

	label := view.NewTextView()
	label.String = strconv.Itoa(v.Index)
	l.Add(label, func(s *constraint.Solver) {
	})

	border := view.NewBasicView()
	border.Painter = &paint.Style{BackgroundColor: colornames.Gray}
	l.Add(border, func(s *constraint.Solver) {
		s.Height(1)
		s.LeftEqual(l.Left())
		s.RightEqual(l.Right())
		s.BottomEqual(l.Bottom())
	})

	return view.Model{
		Children: l.Views(),
		Layouter: l,
		Painter:  &paint.Style{BackgroundColor: colornames.White},
	}
}