// Code generated by "goki generate ./..."; DO NOT EDIT.

package abilities

import (
	"errors"
	"strconv"
	"strings"
	"sync/atomic"

	"goki.dev/enums"
)

var _AbilitiesValues = []Abilities{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

// AbilitiesN is the highest valid value
// for type Abilities, plus one.
const AbilitiesN Abilities = 15

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _AbilitiesNoOp() {
	var x [1]struct{}
	_ = x[Editable-(0)]
	_ = x[Selectable-(1)]
	_ = x[Activatable-(2)]
	_ = x[Pressable-(3)]
	_ = x[LongPressable-(4)]
	_ = x[DoubleClickable-(5)]
	_ = x[Draggable-(6)]
	_ = x[Droppable-(7)]
	_ = x[Slideable-(8)]
	_ = x[Checkable-(9)]
	_ = x[Scrollable-(10)]
	_ = x[Focusable-(11)]
	_ = x[FocusWithinable-(12)]
	_ = x[Hoverable-(13)]
	_ = x[LongHoverable-(14)]
}

var _AbilitiesNameToValueMap = map[string]Abilities{
	`Editable`:        0,
	`editable`:        0,
	`Selectable`:      1,
	`selectable`:      1,
	`Activatable`:     2,
	`activatable`:     2,
	`Pressable`:       3,
	`pressable`:       3,
	`LongPressable`:   4,
	`longpressable`:   4,
	`DoubleClickable`: 5,
	`doubleclickable`: 5,
	`Draggable`:       6,
	`draggable`:       6,
	`Droppable`:       7,
	`droppable`:       7,
	`Slideable`:       8,
	`slideable`:       8,
	`Checkable`:       9,
	`checkable`:       9,
	`Scrollable`:      10,
	`scrollable`:      10,
	`Focusable`:       11,
	`focusable`:       11,
	`FocusWithinable`: 12,
	`focuswithinable`: 12,
	`Hoverable`:       13,
	`hoverable`:       13,
	`LongHoverable`:   14,
	`longhoverable`:   14,
}

var _AbilitiesDescMap = map[Abilities]string{
	0:  `Editable means the element can be edited. Otherwise, it remains in ReadOnly mode.`,
	1:  `Selectable means it can be Selected`,
	2:  `Activatable means it can be made Active`,
	3:  `Pressable means it can be pressed but is not Activatable. Pressed items receive Click events, but do not get the automatic Active state otherwise associated with Activatable items.`,
	4:  `LongPressable indicates that an element can be LongPressed`,
	5:  `DoubleClickable indicates that an element does something different when it is clicked on twice in a row. If this is not set, DoubleClick events are processed in the same way as Click events.`,
	6:  `Draggable means it can be Dragged`,
	7:  `Droppable means it can receive DragEnter, DragLeave, and Drop events (not specific to current Drag item, just generally)`,
	8:  `Slideable means it has a slider element that can be dragged to change value. Cannot be both Draggable and Slideable.`,
	9:  `Checkable means it can be Checked`,
	10: `Scrollable means it can be Scrolled`,
	11: `Focusable means it can be Focused`,
	12: `FocusWithinable means it can be FocusedWithin`,
	13: `Hoverable means it can be Hovered`,
	14: `LongHoverable means it can be LongHovered`,
}

var _AbilitiesMap = map[Abilities]string{
	0:  `Editable`,
	1:  `Selectable`,
	2:  `Activatable`,
	3:  `Pressable`,
	4:  `LongPressable`,
	5:  `DoubleClickable`,
	6:  `Draggable`,
	7:  `Droppable`,
	8:  `Slideable`,
	9:  `Checkable`,
	10: `Scrollable`,
	11: `Focusable`,
	12: `FocusWithinable`,
	13: `Hoverable`,
	14: `LongHoverable`,
}

// String returns the string representation
// of this Abilities value.
func (i Abilities) String() string {
	str := ""
	for _, ie := range _AbilitiesValues {
		if i.HasFlag(ie) {
			ies := ie.BitIndexString()
			if str == "" {
				str = ies
			} else {
				str += "|" + ies
			}
		}
	}
	return str
}

// BitIndexString returns the string
// representation of this Abilities value
// if it is a bit index value
// (typically an enum constant), and
// not an actual bit flag value.
func (i Abilities) BitIndexString() string {
	if str, ok := _AbilitiesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Abilities value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Abilities) SetString(s string) error {
	*i = 0
	return i.SetStringOr(s)
}

// SetStringOr sets the Abilities value from its
// string representation while preserving any
// bit flags already set, and returns an
// error if the string is invalid.
func (i *Abilities) SetStringOr(s string) error {
	flgs := strings.Split(s, "|")
	for _, flg := range flgs {
		if val, ok := _AbilitiesNameToValueMap[flg]; ok {
			i.SetFlag(true, &val)
		} else if val, ok := _AbilitiesNameToValueMap[strings.ToLower(flg)]; ok {
			i.SetFlag(true, &val)
		} else {
			return errors.New(flg + " is not a valid value for type Abilities")
		}
	}
	return nil
}

// Int64 returns the Abilities value as an int64.
func (i Abilities) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Abilities value from an int64.
func (i *Abilities) SetInt64(in int64) {
	*i = Abilities(in)
}

// Desc returns the description of the Abilities value.
func (i Abilities) Desc() string {
	if str, ok := _AbilitiesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// AbilitiesValues returns all possible values
// for the type Abilities.
func AbilitiesValues() []Abilities {
	return _AbilitiesValues
}

// Values returns all possible values
// for the type Abilities.
func (i Abilities) Values() []enums.Enum {
	res := make([]enums.Enum, len(_AbilitiesValues))
	for i, d := range _AbilitiesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Abilities.
func (i Abilities) IsValid() bool {
	_, ok := _AbilitiesMap[i]
	return ok
}

// HasFlag returns whether these
// bit flags have the given bit flag set.
func (i Abilities) HasFlag(f enums.BitFlag) bool {
	return atomic.LoadInt64((*int64)(&i))&(1<<uint32(f.Int64())) != 0
}

// SetFlag sets the value of the given
// flags in these flags to the given value.
func (i *Abilities) SetFlag(on bool, f ...enums.BitFlag) {
	var mask int64
	for _, v := range f {
		mask |= 1 << v.Int64()
	}
	in := int64(*i)
	if on {
		in |= mask
		atomic.StoreInt64((*int64)(i), in)
	} else {
		in &^= mask
		atomic.StoreInt64((*int64)(i), in)
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Abilities) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Abilities) UnmarshalText(text []byte) error {
	return i.SetString(string(text))
}
