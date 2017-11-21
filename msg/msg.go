package msg

//go:generate msgp
//go:generate ffjson $GOFILE
type Sample3 struct {
	Foo       int
	Bar       int
	Age       int
	FirstName string
	LastName  string
}

type Sample3List struct {
	samples []Sample3
}

const Eight = 8

type MyInt int
type Data []byte

//go:generate msgp
type MyStruct struct {
	Which map[string]*MyInt `msg:"which"`
	Other Data              `msg:"other"`
	Nums  [Eight]float64    `msg:"nums"`
}
