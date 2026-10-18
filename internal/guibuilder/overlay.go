package guibuilder

import (
	"image/color"

	"github.com/Uzer007-admin/granio/v2"
	"github.com/Uzer007-admin/granio/v2/canvas"
	"github.com/Uzer007-admin/granio/v2/container"
	"github.com/Uzer007-admin/granio/v2/layout"
	"github.com/Uzer007-admin/granio/v2/theme"
	"github.com/Uzer007-admin/granio/v2/widget"

	"github.com/Uzer007-admin/granio-ui"
)

type overlay struct {
	widget.BaseWidget

	b         *Builder
	indicator *canvas.Rectangle
}

func newOverlay(b *Builder) *overlay {
	o := &overlay{b: b}
	o.ExtendBaseWidget(o)
	return o
}

func (o *overlay) CreateRenderer() granio.WidgetRenderer {
	r := canvas.NewRectangle(color.Transparent)
	r.StrokeColor = color.Transparent
	r.StrokeWidth = 4
	r.Resize(granio.NewSize(20, 10))
	r.Move(granio.NewPos(10, 10))

	o.indicator = r

	return widget.NewSimpleRenderer(container.NewWithoutLayout(r))
}

func (o *overlay) Tapped(pe *granio.PointEvent) {
	rootPos := granio.CurrentApp().Driver().AbsolutePositionForObject(o.b.root)
	pos := pe.AbsolutePosition.Subtract(rootPos)
	obj := findObject(o.b.root, pos)

	// TODO update when an item is removed, inserted, or if the UI resizes
	o.indicator.StrokeColor = theme.Color(theme.ColorNamePrimary)
	objAbsPos := granio.CurrentApp().Driver().AbsolutePositionForObject(obj)
	objPos := objAbsPos.Subtract(rootPos)
	o.indicator.Move(objPos)
	o.indicator.Resize(obj.Size())

	o.b.choose(obj)
}

func findObject(o granio.CanvasObject, p granio.Position) granio.CanvasObject {
	switch w := o.(type) {
	case *granio.Container:
		for _, child := range w.Objects {
			if !child.Visible() || !insideObject(child, p) {
				continue
			}

			match := findObject(child, p.Subtract(child.Position()))
			if match != nil && isContainerOrWidget(match) {
				return match
			}
			if isContainerOrWidget(child) {
				return child
			}
		}

		// also see if we have any raw canvas objects
		for _, child := range w.Objects {
			if !insideObject(child, p) {
				continue
			}

			return child
		}

		return w
	case granio.Widget:
		for _, child := range refyne.DropZonesForObject(o) {
			if !child.Visible() || !insideObject(child, p) {
				continue
			}

			match := findObject(child, p.Subtract(child.Position()))
			if match != nil && isContainerOrWidget(match) {
				return match
			}
			if isContainerOrWidget(child) {
				return child
			}
		}
	}

	return nil
}

func insideObject(o granio.CanvasObject, p granio.Position) bool {
	pos := o.Position()
	if p.X < pos.X || p.Y < pos.Y {
		return false
	}

	size := o.Size()
	return p.X < pos.X+size.Width && p.Y < pos.Y+size.Height
}

func isContainerOrWidget(o granio.CanvasObject) bool {
	switch o.(type) {
	case granio.Widget, *granio.Container, *layout.Spacer:
		return true
	default:
		return false
	}
}
