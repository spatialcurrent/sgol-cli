package sgol

import (
  "time"
)

import (
  "github.com/sirupsen/logrus"
)

type Command interface {
  GetName() string
  Parse(config *Config, args []string) error
  Run(log *logrus.Logger, start time.Time, version string) error
}
