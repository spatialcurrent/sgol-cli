package sgol

type Config struct {
	Server         *Server         `hcl:"server"`
	Formats        []*Format       `hcl:"formats"`
	Authentication *Authentication `hcl:"authentication"`
	NamedQueries   []*NamedQuery   `hcl:"queries"`
}
