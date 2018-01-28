package sgol

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

import (
  "github.com/pkg/errors"
)

type HttpCommand struct {
  *BasicCommand
  server_url string
  auth_token string
}

func (cmd *HttpCommand) ParseServerUrl() error {
  if len(cmd.server_url) == 0 {
    cmd.server_url = cmd.config.Server.Url
    if len(cmd.server_url) == 0 {
      return errors.New("Error: missing server url")
    }
  }
  return nil
}

func (cmd *HttpCommand) ParseAuthToken() error {
  if len(cmd.auth_token) == 0 {
    if cmd.config.Authentication != nil {
      cmd.auth_token = cmd.config.Authentication.Token
    }
    if len(cmd.auth_token) == 0 {
      return errors.New("Error: missing authentication token")
    }
  }
  return nil
}

func (cmd *HttpCommand) MakeRequest(url string, auth_token string, verbose bool) (string, error) {

  client := &http.Client{}

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return "", err
  }

  if cmd.verbose {
    fmt.Println("Using auth token:", auth_token)
  }

  req.Header.Set("X-Auth-Token", auth_token)

  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  return string(body), nil
}
