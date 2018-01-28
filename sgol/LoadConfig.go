package sgol

import (
	"fmt"
	"io/ioutil"
)

import (
	"github.com/hashicorp/hcl"
	"github.com/pkg/errors"
)

func LoadConfig(path string, verbose bool) (*Config, error) {

	result := &Config{}

	d, err := ioutil.ReadFile(path)
	if err != nil {
		return result, errors.New(fmt.Sprintf("Error reading %s: %s", path, err))
	}

	obj, err := hcl.Parse(string(d))
	if err != nil {
		return result, errors.New(fmt.Sprintf("Error parsing %s: %s", path, err))
	}

	if err := hcl.DecodeObject(&result, obj); err != nil {
		return result, errors.New(fmt.Sprintf("Error parsing %s: %s", path, err))
	}

	return result, nil

}
