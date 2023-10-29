// Code generated by "gtigen"; DO NOT EDIT.

package units

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/girl/units.Value",
	ShortName: "units.Value",
	IDName:    "value",
	Doc:       "Value and units, and converted value into raw pixels (dots in DPI)",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Val", &gti.Field{Name: "Val", Type: "float32", Doc: "the value in terms of the specified unit", Directives: gti.Directives{}}},
		{"Un", &gti.Field{Name: "Un", Type: "Units", Doc: "the unit used for the value", Directives: gti.Directives{}}},
		{"Dots", &gti.Field{Name: "Dots", Type: "float32", Doc: "the computed value in raw pixels (dots in DPI)", Directives: gti.Directives{}}},
		{"DotsFunc", &gti.Field{Name: "DotsFunc", Type: "func(uc *Context) float32", Doc: "function to compute dots from units, using arbitrary expressions; if nil, standard ToDots is used", Directives: gti.Directives{}}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})