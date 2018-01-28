package sgol

import (
  "bytes"
  "text/template"
)

func RenderTemplate(template_text string, ctx map[string]string) (string, error) {
  tmpl, err := template.New("tmpl").Parse(template_text)
  if err != nil {
    return "", err
  }
  buf := new(bytes.Buffer)
  err = tmpl.ExecuteTemplate(buf, "tmpl", ctx)
  if err != nil {
    return "", err
  }
  return buf.String(), nil
}
