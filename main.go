package main

import "github.com/andlabs/ui"

func main() {
	err := ui.Main(func() {
		window := ui.NewWindow("Hello", 600, 400, false)
		window.SetMargined(true)
		window.OnClosing(func(*ui.Window) bool {
			return true
		})

		list := ui.NewVerticalBox()
		list.Append(ui.NewLabel("email1"), false)
		list.Append(ui.NewLabel("email2"), false)
		inbox := ui.NewGroup("Inbox")
		inbox.SetChild(list)

		subject := ui.NewLabel("subject")
		content := ui.NewLabel("content")
		labels := ui.NewVerticalBox()
		labels.Append(ui.NewLabel("From "), false)
		labels.Append(ui.NewLabel("To "), false)
		labels.Append(ui.NewLabel("Date "), false)

		values := ui.NewVerticalBox()
		from := ui.NewLabel("email")
		values.Append(from, false)
		to := ui.NewLabel("email")
		values.Append(to, false)
		date := ui.NewLabel("date")
		values.Append(date, false)

		meta := ui.NewHorizontalBox()
		meta.SetPadded(true)
		meta.Append(labels, false)
		meta.Append(values, true)

		detail := ui.NewVerticalBox()
		detail.SetPadded(true)
		detail.Append(subject, false)
		detail.Append(meta, false)
		detail.Append(ui.NewHorizontalSeparator(), false)
		detail.Append(content, true)

		content1 := ui.NewHorizontalBox()
		content1.SetPadded(true)
		content1.Append(inbox, false)
		content1.Append(ui.NewVerticalSeparator(), false)
		content1.Append(detail, true)

		window.SetChild(content1)
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func NewEmail() {
	window := ui.NewWindow("New GoMail", 400, 320, false)
	window.SetMargined(true)
	window.OnClosing(func(*ui.Window) bool {
		return true
	})
	subject := ui.NewEntry()
	subject.SetText("subject")
	toBox := ui.NewHorizontalBox()
	toBox.SetPadded(true)
	toBox.Append(ui.NewLabel("To"), false)
	to := ui.NewEntry()
	to.SetText("email")
	toBox.Append(to, true)
	content := ui.NewEntry()
	content.SetText("email content")
	buttonBox := ui.NewHorizontalBox()
	buttonBox.SetPadded(true)
	buttonBox.Append(ui.NewLabel(""), true)
	buttonBox.Append(ui.NewButton("Cancel"), false)
	buttonBox.Append(ui.NewButton("Send"), false)
	layout := ui.NewVerticalBox()
	layout.SetPadded(true)
	layout.Append(subject, false)
	layout.Append(toBox, false)
	layout.Append(content, true)
	layout.Append(buttonBox, false)
	window.SetChild(layout)
	window.Show()
}

func ToolbarAndMenu() *ui.Box {
	toolbar := ui.NewHorizontalBox()
	toolbar.Append(ui.NewButton("New"), false)
	toolbar.Append(ui.NewButton("Reply"), false)
	toolbar.Append(ui.NewButton("Reply All"), false)
	toolbar.Append(ui.NewLabel(" "), false)
	toolbar.Append(ui.NewVerticalSeparator(), false)
	toolbar.Append(ui.NewLabel(" "), false)
	toolbar.Append(ui.NewButton("Delete"), false)
	toolbar.Append(ui.NewLabel(" "), false)
	toolbar.Append(ui.NewVerticalSeparator(), false)
	toolbar.Append(ui.NewLabel(" "), false)
	toolbar.Append(ui.NewButton("Cut"), false)
	toolbar.Append(ui.NewButton("Copy"), false)
	toolbar.Append(ui.NewButton("Paste"), false)

	return toolbar
}
