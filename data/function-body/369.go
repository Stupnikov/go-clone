{
	return func(start, end, now int32) float32 {
		// A Cubic-Bezier curve restricted to starting at (0,0) and
		// ending at (1,1) is defined as
		//
		//  B(t) = 3*(1-t)^2*t*P0 + 3*(1-t)*t^2*P1 + t^3
		//
		// with derivative
		//
		//  B'(t) = 3*(1-t)^2*P0 + 6*(1-t)*t*(P1-P0) + 3*t^2*(1-P1)
		//
		// Given a value x ∈ [0,1], we solve for t using Newton's
		// method and solve for y using t.

		x := linear(start, end, now)

		// Solve for t using x.
		t := x
		for i := 0; i < 5; i++ {
			t2 := t * t
			t3 := t2 * t
			d := 1 - t
			d2 := d * d

			nx := 3*d2*t*x0 + 3*d*t2*x1 + t3
			dxdt := 3*d2*x0 + 6*d*t*(x1-x0) + 3*t2*(1-x1)
			if dxdt == 0 {
				break
			}

			t -= (nx - x) / dxdt
			if t <= 0 || t >= 1 {
				break
			}
		}
		if t < 0 {
			t = 0
		}
		if t > 1 {
			t = 1
		}

		// Solve for y using t.
		t2 := t * t
		t3 := t2 * t
		d := 1 - t
		d2 := d * d
		y := 3*d2*t*y0 + 3*d*t2*y1 + t3

		return y
	}
}