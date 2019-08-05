// Code generated by gogen.Generate; DO NOT EDIT.

package pkg1

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
)

var cuegenvalMyStruct = cuegenMake("MyStruct", &MyStruct{})

// Validate validates x.
func (x *MyStruct) Validate() error {
	return cuegenCodec.Validate(cuegenvalMyStruct, x)
}

// Complete completes x.
func (x *MyStruct) Complete() error {
	return cuegenCodec.Complete(cuegenvalMyStruct, x)
}

var cuegenvalOtherStruct = cuegenMake("OtherStruct", &OtherStruct{})

// Validate validates x.
func (x *OtherStruct) Validate() error {
	return cuegenCodec.Validate(cuegenvalOtherStruct, x)
}

var cuegenvalString = cuegenMake("String", nil)

// ValidateCUE validates x.
func (x String) ValidateCUE() error {
	return cuegenCodec.Validate(cuegenvalString, x)
}

var cuegenvalSpecialString = cuegenMake("SpecialString", nil)

// ValidateSpecialString validates x.
func ValidateSpecialString(x string) error {
	return cuegenCodec.Validate(cuegenvalSpecialString, x)
}

var cuegenvalPtr = cuegenMake("Ptr", Ptr(nil))

// ValidatePtr validates x.
func ValidatePtr(x Ptr) error {
	return cuegenCodec.Validate(cuegenvalPtr, x)
}

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	v := cuegenInstance.Lookup(name)
	if !v.Exists() {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 518 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\x94R]k\xd4@\x14\x9d\x9b\xae`.U\xf0\a\b\xe3<m\xa4\xcd~\x80\b\xa1Qk\xad\u0407v\x17\x8b\"\x88\x0f\xe3\xeclv\xd8l&$w\u0165\xb6\xa0\xd6\xda_\xe7\xab\xff\xa6\x91|\xb5\xabo\xcd\xcb\\n\xe6\x9c{\xe6\xdcs\xaf\xf8\xe5\x80S\\2(\xbe3\xf6\xb4\xf8\xb6\x01\xb0i\x92\x9cd\xa2\xf4+I\xb2\xec\xc3\x06t\xdeXK\xe00\xe8\x8c%\xcd`\x93\xc1\x9d\xd7&\xd69\x14\x17\x8c\xb1\x87\xc5O\a\xe0\xfe\x87\x8fj\xa9\xfd\xa9\x89\x1b\xe4\x05\x83\u2731n\xf1c\x03\xe0\xeeM\xff\x9c\x81\x03\x9d#\xb9\xd0%Q\xa7j\"c\xec\xca\xf9]\\2\a\x00\xb6\xd4R\xc72\x89|\x9bE\xbd\xc8\xf6t\xa2\xec\xc4$e\xad\xecD\xf7H\xe74\x91${\xe9<\x1a\x00\xc0\x83\xf2\ucd7a}\xb5\xd4p\x05\u007fR\xa9\xe62\u04bc\xfc\x89h\x16\xa9\u0348w\xd1\x15\xb7`\x1f\nt\xc5B\u04ac<s\xcaL\x12\xe5\x02=\xc4\xc3\xd51eKE\x01?Aw7\xe0|'\x1c\xf4\xd1}\x19p\x1e\x9e\t%I\xf0\xaf\xfc\xb1\x98\xd8H\xa0;z\x1e\xf0\x11\xcdtVc\xd0=\bx)k\xe8\x1fT\xaa\x0e5\x9e\xf2\x17\x91\xedn)\xbbHcM:\xdck\n\x0f\u05c0\xed\xb0F\x88\xbfg\x13\x92&\xc9w\x93UW\xbc\x17\x1e\xba\xe3\xa0\xe6\x1d\x1b5/Y\xf1\xb8\xba\x1a\xf0\xe6{\x14\n\xd1\xd6\xd5\xc0\xcf26\x13I:|\xd7\x14{o\xf7=<N\xb522n\xc1\xe1\x99\xc8\ub3a8Q\xb4JuX\xab\xf0p\xb40t=\x80s\x93\x10_\x1f\xb1\xed\xe1\x91M\xf6\xbf\x98\x9c*\xb2\x93\u02a5\x1a\xdc<{\xdb\u00e9\xb5A\t\xc51e\xed3K\xdf\xfd\xc3eL&\x8d\xf5h\xda\x1d\xf4=<E\u019c\xdb\xe4c\xd8\xe4c\xf8o>\xe4Z:\x86\xd7\xe9\xb8Y1\xb6\x9bi\xc5\xec\f\xfa\xfd5\xe5\xff\x99/?)Q\x8a\xab}\x0f\xf8\xb3'\xc8\xd8\xdf\x00\x00\x00\xff\xff\xa9a\xee\x8d^\x03\x00\x00")
