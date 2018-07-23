package sgol

import (
  "fmt"
  "os"
  "net/url"
  "strings"
  "time"
)

import (
  "github.com/pkg/errors"
  "github.com/sirupsen/logrus"
)

type ExecCommand struct {
  *HttpCommand
  server_port int
}

func (cmd *ExecCommand) GetName() string {
  return "serve"
}

func (cmd *ExecCommand) Parse(args []string) error {

  fs := cmd.NewFlagSet(cmd.GetName())

  output_format_help := "Output format.  Options: "+strings.Join(cmd.config.GetFormatIds(), ", ")
  fs.StringVar(&cmd.output_format_text, "f", "", output_format_help)

  //c.sgol_config_path = flagSet.String("c", os.Getenv("SGOL_CONFIG_PATH"), "path to SGOL config.hcl")
  fs.StringVar(&cmd.backend_url, "u", os.Getenv("SGOL_BACKEND_URL"), "Backend url.")
  fs.BoolVar(&cmd.verbose, "verbose", false, "Provide verbose output")
  fs.BoolVar(&cmd.dry_run, "dry_run", false, "Connect to destination, but don't import any data.")
  fs.BoolVar(&cmd.version, "version", false, "Version")
  fs.BoolVar(&cmd.help, "help", false, "Print help")

  err := fs.Parse(args)
  if err != nil {
    return err
  }

  err = cmd.ParseBackendUrl()
  if err != nil {
    return err
  }

  return nil
}

func (cmd *ExecCommand) Run(log *logrus.Logger, start time.Time, version string) error {

  if cmd.help {
    cmd.PrintHelp(cmd.GetName(), version)
  }

	if len(cmd.named_query_name) > 0 {
		named_query_object := &NamedQuery{}
		for _, named_query := range cmd.config.NamedQueries {
			if named_query.Name == cmd.named_query_name {
				named_query_object = named_query
				break
			}
		}
		if len(named_query_object.Name) == 0 {
			return errors.New("Error: Could not find named query with name " + cmd.named_query_name + ".")
		}
    if cmd.verbose {
      fmt.Println("Using query template "+named_query_object.Sgol)
    }
		ctx, err := BuildContext(cmd.flagSet.Args(), named_query_object.Required, named_query_object.Optional)
		if err != nil {
			return err
		}
    query, err := RenderTemplate(named_query_object.Sgol, ctx)
    if err != nil {
			return err
		}
    cmd.query = query
	}

	url := cmd.backend_url + "/exec." + cmd.output_format.Extension + "?q=" + url.QueryEscape(cmd.query)
	if cmd.verbose {
		fmt.Println("url: " + url)
	}

  output_text, err := cmd.MakeRequest(url, cmd.auth_token, cmd.verbose)
  if err != nil {
    return err
  }

  err = cmd.WriteOutput(cmd.output_uri, output_text)
  if err != nil {
    return err
  }

  elapsed := time.Since(start)
  if cmd.verbose {
    log.Info("Done in " + elapsed.String())
  }

  return nil
}

func NewServeCommand(config *Config) *ServeCommand {
  return &ServeCommand{
    HttpCommand: &HttpCommand{
      BasicCommand: &BasicCommand{
        config: config,
      },
    },
  }
}
