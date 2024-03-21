package module

type Crater struct {
	Id          string     `hcl:"id"`
	Name        string     `hcl:"name"`
	Description string     `hcl:"description"`
	Variants    []*Variant `hcl:"variant,block"`
	Provider    *Module
}
type Variant struct {
	Id          string `hcl:"id"`
	Name        string `hcl:"name"`
	Description string `hcl:"description"`
	Crater      *Crater
}
