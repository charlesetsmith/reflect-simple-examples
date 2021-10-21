package main

import (
	"fmt"
	"reflect"
)

type Info interface {
	NewFrame(interface{})
	PrintFrame()
}

// Interface Handlers
func NewFrame(f Info, info interface{}) {
	f.NewFrame(info)
}

func PrintFrame(f Info) {
	f.PrintFrame()
}

// ***************************************************************************************

// Binfo handlers

// Simple Beacon Structure to load via interface to New
type Binfo struct {
	Session uint32
}

func (b *Binfo) NewFrame(info interface{}) {
	e := reflect.ValueOf(info).Elem()
	// Uint() only returns uint64 so cast it to what Session is
	b.Session = uint32(e.FieldByName("Session").Uint())
}

func (b *Binfo) PrintFrame() {
	fmt.Println("Beacon Session:", b.Session)
}

// ***************************************************************************************

// Finfo handlers

// More Complex Frame Structure to load via interface to New
type Finfo struct {
	Session  uint32
	Progress uint64
	Inrespto uint64
}

func (f *Finfo) NewFrame(info interface{}) {
	// Get the Element
	e := reflect.ValueOf(info).Elem()
	// Uint() only returns uint64 so cast it to what Session is
	f.Session = uint32(e.FieldByName("Session").Uint())
	f.Progress = e.FieldByName("Progress").Uint()
	f.Inrespto = e.FieldByName("Inrespto").Uint()
}

func (f *Finfo) PrintFrame() {
	fmt.Println("Metadata Session:", f.Session, "Progress:", f.Progress, "Inrespto:", f.Inrespto)
}

// ***************************************************************************************

// Sinfo Handlers
type Sinfo struct {
	Session  uint32
	Progress uint64
	Inrespto uint64
	Start    []uint64
	End      []uint64
}

func (s *Sinfo) NewFrame(info interface{}) {
	// Uint() only returns uint64 so cast it to what Session is
	e := reflect.ValueOf(info).Elem()
	s.Session = uint32(e.FieldByName("Session").Uint())
	s.Progress = e.FieldByName("Progress").Uint()
	s.Inrespto = e.FieldByName("Inrespto").Uint()

	// Are we a Start slice
	if e.FieldByName("Start").Kind() == reflect.Slice {
		// How long is the Start slice
		len := e.FieldByName("Start").Len()
		// Loop through the entries in the Start slice
		for i := 0; i < len; i++ {
			s.Start = append(s.Start,
				e.FieldByName("Start").Index(i).Uint())
		}
	}
	// Are we a End slice
	if e.FieldByName("End").Kind() == reflect.Slice {
		// How long is the End slice
		len := e.FieldByName("End").Len()
		// Loop through the entries in the End slice
		for i := 0; i < len; i++ {
			s.End = append(s.End,
				e.FieldByName("End").Index(i).Uint())
		}
	}
}

func (s *Sinfo) PrintFrame() {
	fmt.Println("Metadata Session:", s.Session, "Progress:", s.Progress, "Inrespto:", s.Inrespto)
	fmt.Println(" Start:", s.Start, " End:", s.End)
}

// ***************************************************************************************

// Xinfo handlers

// Used in Xinfo
type Hole struct {
	Start uint32
	End   uint32
}

// Really Complex Frame Structure to load via interface to New
type Xinfo struct {
	Session  uint32
	Progress uint64
	Inrespto uint64
	Desc     string
	Holes    []Hole // This is the complex bit an array of Hole structure
}

func (x *Xinfo) NewFrame(info interface{}) {
	e := reflect.ValueOf(info).Elem()
	// Uint() only returns uint64 so cast it to what Session is
	x.Session = uint32(e.FieldByName("Session").Uint())
	x.Progress = e.FieldByName("Progress").Uint()
	x.Inrespto = e.FieldByName("Inrespto").Uint()
	// Are we a Slice of Holes Structures
	if e.FieldByName("Holes").Kind() == reflect.Slice {
		// Loop through the entries in the Holes slice
		for i := 0; i < e.FieldByName("Holes").Len(); i++ {
			var h Hole
			// Get the Start and End from within the Holes Structure
			h.Start = uint32(e.FieldByName("Holes").Index(i).FieldByName("Start").Uint())
			h.End = uint32(e.FieldByName("Holes").Index(i).FieldByName("End").Uint())
			x.Holes = append(x.Holes, h)
		}
	}
	x.Desc = e.FieldByName("Desc").String()
}

func (x *Xinfo) PrintFrame() {
	fmt.Println("Metadata Session:", x.Session, "Progress:",
		"Desc:", x.Desc, x.Progress, "Inrespto:", x.Inrespto)
	fmt.Println(" Holes:", x.Holes)
	//    fmt.Println("Start:", m.Start)
	//    fmt.Println("End:", m.End)
}

// ***************************************************************************************

func main() {
	// var f Info

	var s Sinfo
	sinfo := Sinfo{Session: 0, Progress: 1, Inrespto: 2,
		Start: []uint64{500, 600, 700}, End: []uint64{700, 800, 900}}
	NewFrame(&s, &sinfo)
	PrintFrame(&s)

	var b Binfo
	binfo := Binfo{Session: 88}
	NewFrame(&b, &binfo)
	PrintFrame(&b)

	var x Xinfo
	ho := []Hole{{Start: 11, End: 12}, {Start: 13, End: 14}}
	xinfo := Xinfo{Session: 99, Progress: 100, Inrespto: 101, Desc: "Hello",
		Holes: ho}
	NewFrame(&x, &xinfo)
	PrintFrame(&x)
}
