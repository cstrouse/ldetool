/*
 This file was autogenerated via
 -----------------------------------------------------------
 gen-builtin --declarable --native --type Uint8 --name uint8
 -----------------------------------------------------------
 do not touch it with bare hands!
*/

package types

var _ Field = Uint8("")

// Uint8 represents field of type uint8
type Uint8 string

// Name returns field name
func (i Uint8) Name() string {
	return string(i)
}

// TypeName name of the type
func (i Uint8) TypeName() string {
	return "uint8"
}

// Register registers a field
func (i Uint8) Register(comment []string, registrator FieldRegistrator) {
	registrator.AddUint8(comment, i.Name())
}

// GoName returns Go's representation of this field's type
func (i Uint8) GoName() string {
	return "uint8"
}

func init() {
	if builtins == nil {
		builtins = map[string]func(name string) Field{}
	}
	builtins["uint8"] = func(fieldName string) Field {
		return Uint8(fieldName)
	}
	if natives == nil {
		natives = map[string]struct{}{}
	}
	natives["uint8"] = struct{}{}
	if declarables == nil {
		declarables = map[string]struct{}{}
	}
	declarables["uint8"] = struct{}{}

}
