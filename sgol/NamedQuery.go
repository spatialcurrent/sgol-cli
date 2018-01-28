package sgol

type NamedQuery struct {
	Name     string   `json:"name" hcl:"name"`
	Sgol     string   `json:"sgol" hcl:"sgol"`
	Required []string `json:"required" hcl:"required"`
	Optional []string `json:"optional" hcl:"optional"`
}
