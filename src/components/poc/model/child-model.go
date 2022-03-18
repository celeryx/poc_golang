package poc

type ChildModel struct {
	Rut  string `json:"rut"   validate:"required,numeric,min=7"`
	Name string `json:"name"  validate:"required,alpha"`
}
