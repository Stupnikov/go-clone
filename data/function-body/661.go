{
	l := &table.Layouter{}
	{
		group := []view.View{}

		spacer := NewSpacer()
		l.Add(spacer, nil)

		switchView := view.NewSwitch()
		switchView.Value = v.app.Bluetooth.Enabled()
		switchView.OnSubmit = func(value bool) {
			v.app.Bluetooth.SetEnabled(!v.app.Bluetooth.Enabled())
		}

		cell1 := NewBasicCell()
		cell1.Title = "Bluetooth"
		cell1.AccessoryView = switchView
		group = append(group, cell1)

		for _, i := range AddSeparators(group, 60) {
			l.Add(i, nil)
		}
	}
	if v.app.Bluetooth.Enabled() {
		group := []view.View{}

		spacer := NewSpacerHeader()
		spacer.Title = "My Devices"
		l.Add(spacer, nil)

		for _, i := range v.app.Bluetooth.Devices() {
			device := i
			cell := NewBasicCell()
			cell.Title = device.SSID()
			if device.Connected() {
				cell.Subtitle = "Connected"
			} else {
				cell.Subtitle = "Not Connected"
			}
			cell.OnTap = func() {
				device.SetConnected(!device.Connected())
				v.Signal()
			}
			group = append(group, cell)
		}

		for _, i := range AddSeparators(group, 60) {
			l.Add(i, nil)
		}
	}

	scrollView := view.NewScrollView()
	scrollView.ContentLayouter = l
	scrollView.ContentChildren = l.Views()

	return view.Model{
		Children: []view.View{scrollView},
		Painter:  &paint.Style{BackgroundColor: backgroundColor},
		Options: []view.Option{
			&ios.StackBar{Title: "Settings"},
		},
	}
}