package poc

type PocModel struct {
	Rut            string        `json:"rut"            validate:"required,numeric,min=7"`
	Name           string        `json:"name"           validate:"required,alpha"`
	Addresses      []string      `json:"addresses"      validate:"required,gt=0,dive,required"`
	Age            string        `json:"age"            validate:"required,numeric"`
	BirthDate      string        `json:"birthDate"      validate:"required"`
	Email          string        `json:"email"          validate:"required,email"`
	ProfilePicture string        `json:"profilePicture" validate:"required,base64"`
	Childrens      []*ChildModel `json:"childrens"      validate:"required,gt=0,dive,required"`
	CreditCard     string        `json:"creditCard"     validate:"required,numeric"`
	PostalCode     string        `json:"postalCode"     validate:"required"`
	MobilePhone    string        `json:"mobilePhone"    validate:"required,e164"`
}
