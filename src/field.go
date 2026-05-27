package main

type Field struct {
	app *Application
}

func NewField(app *Application) *Field {
	fld := &Field{}
	fld.app = app
	return fld
}
