package sgol

type Config struct {
	Server         *Server         `hcl:"server"`
	Formats        []*Format       `hcl:"formats"`
	Authentication *Authentication `hcl:"authentication"`
	NamedQueries   []*NamedQuery   `hcl:"queries"`
}

func (config *Config) GetFormatIds() []string {
	formatIds := make([]string, len(config.Formats))
	for i := 0; i < len(formatIds); i++ {
		formatIds[i] = config.Formats[i].Id
	}
	return formatIds
}
