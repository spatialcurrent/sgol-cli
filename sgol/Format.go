package sgol

type Format struct {
	Id        string   `json:"id", hcl:"id"`
	Extension string   `json:"extension" hcl:"extension"`
	Aliases   []string `json:"aliases" hcl:"aliases"`
}
