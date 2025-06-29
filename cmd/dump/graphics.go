package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"strings"
)

type Region interface {
	BoundingBox() []float64
}

// single value Raster.
//
//	TBD if we need RasterRGB
type Raster [][]float64

func (r Raster) BoundingBox() []float64 {
	return []float64{0.0, 0.0, float64(len(r[0])), float64(len(r))}
}
func (r Raster) String() string {
	return fmt.Sprintf("Raster( %d x %d )", len(r[0]), len(r))
}

type Graphics []any

func (c Graphics) String() string {
	out := []string{}
	for _, a := range c {
		out = append(out, a.(fmt.Stringer).String())
	}
	return strings.Join(out, ",")
}

func (c Graphics) BoundingBox() []float64 {
	bb := []float64{math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64}
	for _, a := range c {
		switch r := a.(type) {
		case Region:
			b := r.BoundingBox()

			// xmin
			if b[0] < bb[0] {
				bb[0] = b[0]
			}

			// ymin
			if b[1] < bb[1] {
				bb[1] = b[1]
			}

			// xmax
			if b[2] > bb[2] {
				bb[2] = b[2]
			}

			// ymax
			if b[3] > bb[3] {
				bb[3] = b[3]
			}
		}
	}
	return bb
}

type Text struct {
	x    float64
	y    float64
	body string
}

func (c Text) String() string {
	return fmt.Sprintf("Text(%g %g %q)", c.x, c.y, c.body)
}

func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

func colorPercent(val float64) int {
	return int(math.Round(100 * clamp01(val)))
}

type GrayScale struct {
	val float64
}

type RGBColor struct {
	r float64
	g float64
	b float64
}

func bbPoints(c []float64) []float64 {
	bb := []float64{math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64}
	for i := 0; i < len(c); i += 2 {
		if c[i] < bb[0] {
			bb[0] = c[i]
		}
		if c[i] > bb[2] {
			bb[2] = c[i]
		}
		if c[i+1] < bb[1] {
			bb[1] = c[i+1]
		}
		if c[i+1] > bb[3] {
			bb[3] = c[i+1]
		}
	}
	return bb
}

type Rectangle struct {
	xmin float64
	ymin float64
	xmax float64
	ymax float64
}

func (c Rectangle) BoundingBox() []float64 {
	return []float64{c.xmin, c.ymin, c.xmax, c.ymax}
}
func (c Rectangle) String() string {
	return fmt.Sprintf("rectangle(%g,%g %g,%g)", c.xmin, c.ymin, c.xmax, c.ymax)
}

type RoundedRectangle struct {
	xmin float64
	ymin float64
	xmax float64
	ymax float64
	rx   float64
	ry   float64
}

func (c RoundedRectangle) BoundingBox() []float64 {
	return []float64{c.xmin, c.ymin, c.xmax, c.ymax}
}
func (c RoundedRectangle) String() string {
	return fmt.Sprintf("rectangle(%g,%g %g,%g %g,%g)", c.xmin, c.ymin, c.xmax, c.ymax, c.rx, c.ry)
}

type Point []float64

func (c Point) BoundingBox() []float64 {
	return bbPoints(c)
}

func (c Point) String() string {
	// can't use %v here
	out := make([]string, len(c), len(c))
	for i := 0; i < len(c); i += 2 {
		out = append(out, fmt.Sprintf("%g,%g", c[i], c[i+1]))
	}
	return "Point(" + strings.Join(out, " ") + ")"
}

type Polygon []float64

func (c Polygon) BoundingBox() []float64 {
	return bbPoints(c)
}

func (c Polygon) String() string {
	// can't use %v here
	out := make([]string, len(c), len(c))
	for i := 0; i < len(c); i += 2 {
		out = append(out, fmt.Sprintf("%g,%g", c[i], c[i+1]))
	}
	return "Polygon(" + strings.Join(out, " ") + ")"
}

type Line []float64

func (c Line) BoundingBox() []float64 {
	return bbPoints(c)
}

func (c Line) String() string {
	// can't use %v here
	out := make([]string, len(c), len(c))
	for i := 0; i < len(c); i += 2 {
		out = append(out, fmt.Sprintf("%g,%g", c[i], c[i+1]))
	}
	return "Line2d(" + strings.Join(out, " ") + ")"
}

func (c RGBColor) String() string {
	return fmt.Sprintf("RGBColor(%g %g %g)", c.r, c.g, c.b)
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(%g %g %g)", c.x, c.y, c.r)
}

// satifies Region
func (c Circle) BoundingBox() []float64 {
	return []float64{c.x - c.r, c.y - c.r, c.x + c.r, c.y + c.r}
}

type Ellipse struct {
	x  float64
	y  float64
	rx float64
	ry float64
}

func (c Ellipse) String() string {
	return fmt.Sprintf("Circle(%g %g %g %g)", c.x, c.y, c.rx, c.ry)
}

// satifies Region
func (c Ellipse) BoundingBox() []float64 {
	return []float64{c.x - c.rx, c.y - c.ry, c.x + c.rx, c.y + c.ry}
}

type Disk struct {
	x  float64
	y  float64
	rx float64
	ry float64
}

func (c Disk) String() string {
	rx := c.rx
	ry := c.ry
	if ry == 0 {
		return fmt.Sprintf("Disk(%g %g %g)", c.x, c.y, rx)
	}
	return fmt.Sprintf("Disk(%g %g %g %g)", c.x, c.y, rx, ry)
}

// satifies Region
func (c Disk) BoundingBox() []float64 {
	rx := c.rx
	ry := c.ry
	if ry == 0 {
		ry = rx
	}
	return []float64{c.x - c.rx, c.y - c.ry, c.x + c.rx, c.y + c.ry}
}

func CSSColor(c any) string {
	switch col := c.(type) {
	case GrayScale:
		return fmt.Sprintf("hsl(0 0 %d)", colorPercent(col.val))
	case RGBColor:
		return fmt.Sprintf("rgb(%d %d %d)", colorPercent(col.r), colorPercent(col.g), colorPercent(col.b))
	default:
		return "NOPE"
	}
}

func toSVGlist(ctx *svgContext, args []any) {
	for _, a := range args {
		switch shape := a.(type) {
		case Text:
			// TODO XML ESCAPE BODY
			x, y := ctx.remap(shape.x, shape.y)
			ctx.Write(fmt.Sprintf("<text x=\"%g\" y=\"%g\" text-anchor=\"middle\">%s</text>", x, y, shape.body))
		case Circle:
			x, y := ctx.remap(shape.x, shape.y)
			r := ctx.remapScalar(shape.r)
			ctx.Write(fmt.Sprintf("<circle cx=\"%g\" cy=\"%g\" r=\"%g\"/>", x, y, r))
		case Ellipse:
			x, y := ctx.remap(shape.x, shape.y)
			rx := ctx.remapScalar(shape.rx)
			ry := ctx.remapScalar(shape.ry)
			ctx.Write(fmt.Sprintf("<ellipse cx=\"%g\" cy=\"%g\" rx=\"%g\" ry=\"%g\"/>", x, y, rx, ry))
		case Disk:
			//TBD
		case Rectangle:
			// ATTENTION
			// using upper-left corneri (html), NOT the lower-left (normal)
			xmin, ymin := ctx.remap(shape.xmin, shape.ymax)
			h := ctx.remapScalar(shape.xmax - shape.xmin)
			w := ctx.remapScalar(shape.ymax - shape.ymin)
			ctx.Write(fmt.Sprintf("<rect x=\"%g\" y=\"%g\" width=\"%g\" height=\"%g\"/>", xmin, ymin, w, h))
		case RoundedRectangle:

			xmin, ymin := ctx.remap(shape.xmin, shape.ymax)
			h := ctx.remapScalar(shape.xmax - shape.xmin)
			w := ctx.remapScalar(shape.ymax - shape.ymin)
			rx := ctx.remapScalar(shape.rx)
			ry := ctx.remapScalar(shape.ry)
			ctx.Write(fmt.Sprintf("<rect x=\"%g\" y=\"%g\" width=\"%g\" height=\"%g\" rx=\"%g\" ry=\"%g\"/>", xmin, ymin, w, h, rx, ry))
		case Line:
			ctx.Write(fmt.Sprintf("<polyline points=%q fill=%q />", ctx.makeLinePoints(shape), "none"))

		case Point:
			// TBD
		case Polygon:
			ctx.Write(fmt.Sprintf("<polygson points=%q />", ctx.makeLinePoints(shape)))
		case Raster:
			rwidth := ctx.scalex
			rheight := ctx.scaley
			fmt.Printf("rwidth=%g, rheight=%g\n", rwidth, rheight)
			for i := 0; i < len(shape); i++ {
				for j := 0; j < len(shape[i]); j++ {
					xmin := float64(i) * rwidth
					ymin := float64(j+1) * rheight
					//let [xmin, ymin] = remapPoint(ctx.tx, ctx.boxHeight, [i, j + 1]);

					fill := CSSColor(GrayScale{shape[i][j]})
					ctx.Write(fmt.Sprintf("<rect x=\"%g\" y=\"%g\" width=\"%g\" height=\"%g\" fill=%q stroke=\"none\"/>", xmin, ymin, rwidth, rheight, fill))
				}
			}

		default:
			// NOP.. skip it
		}
	}
}

type svgContext struct {
	out       []string
	scalex    float64
	scaley    float64
	tx        float64
	ty        float64
	boxwidth  float64
	boxheight float64
}

func (r *svgContext) Write(s string) {
	r.out = append(r.out, s)
}

func (r *svgContext) String() string {
	return strings.Join(r.out, "")
}

// remape or rescale a single value scalar
func (r *svgContext) remapScalar(val float64) float64 {
	return r.scalex * val
}

// remap a point to SVG space... tricky!
func (r *svgContext) remap(x, y float64) (float64, float64) {
	nx := x*r.scalex + r.tx
	ny := y*r.scaley + r.ty
	ny = r.boxheight - ny
	return nx, ny
}
func (r *svgContext) makeLinePoints(pts []float64) string {
	out := make([]string, 0, len(pts)/2)
	for i := 0; i < len(pts); i += 2 {
		x, y := r.remap(pts[i], pts[i+1])
		out = append(out, fmt.Sprintf("%g,%g", x, y))
	}
	return strings.Join(out, " ")
}

func ToSVG(g Graphics) string {
	bb := g.BoundingBox()

	boxWidth := float64(320)
	boxHeight := float64(320)
	svgWidth := bb[2] - bb[0]
	svgHeight := bb[3] - bb[1]

	xScale := boxWidth / svgWidth
	yScale := boxHeight / svgHeight
	scale := math.Min(xScale, yScale)
	tx := -bb[0] * scale
	ty := -bb[1] * scale
	viewbox := fmt.Sprintf("0 0 %d %d", int(boxWidth), int(boxHeight))

	var ctx = &svgContext{
		out:       []string{},
		scalex:    scale,
		scaley:    scale,
		tx:        tx,
		ty:        ty,
		boxwidth:  boxWidth,
		boxheight: boxHeight,
	}
	ctx.Write(fmt.Sprintf("<svg version=\"1.1\" role=\"img\" viewBox=%q width=\"%d\" height=\"%d\" xmlns=\"http://www.w3.org/2000/svg\">", viewbox, int(boxWidth), int(boxHeight)))

	toSVGlist(ctx, g)

	ctx.Write("</svg>")
	return ctx.String()
}

func RandomBit() int {
	return rand.IntN(2)
}

func WriteRaster(table [][]float64, name string) error {
	g := Graphics{ Raster(table) }
	svg := ToSVG(g)
	return os.WriteFile(name, []byte(svg), 0644)
}

func main2() {
	xmax := 64
	ymax := 64
	table := make([][]float64, xmax)
	for i := 0; i < ymax; i++ {
		table[i] = make([]float64, xmax)
		for j := 0; j < xmax; j++ {
			table[i][j] = 0.0// float64(RandomBit())
		}
	}
	table[0][0] = 1.0
	WriteRaster(table, "crap.svg")
}
