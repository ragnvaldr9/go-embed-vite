package vite

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

func (v *Vite) Favicon() (template.HTML, error) {
	var href string = "/assets/favicon.svg"

	tag := fmt.Sprintf(`<link rel="icon" type="image/svg+xml" href="%v" />`, href)
	tmpl, err := template.New("favicon").Parse(tag)

	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	tmpl.Execute(&buffer, v)

	return template.HTML(buffer.String()), nil
}

func (v *Vite) Tags() (template.HTML, error) {
	b := strings.Builder{}

	if v.Env == "production" {
		b.WriteString(`
			{{ $data := .Data }}
			<script type="module" crossorigin src="/{{ $data.file }}"></script>
			
			{{ range $href := $data.imports }}
				<link rel="modulepreload" href="/{{ $href }}">
			{{ end }}
			
			{{ range $href := $data.css }}
				<link rel="stylesheet" href="/{{ $href }}">
			{{ end }}
		`)
	}

	if v.Env == "development" {
		if v.Platform == "react" {
			b.WriteString(`
				<script type="module">
					import RefreshRuntime from '{{ .DevServerURL }}/@react-refresh'
					RefreshRuntime.injectIntoGlobalHook(window)
					window.$RefreshReg$ = () => {}
					window.$RefreshSig$ = () => (type) => type
					window.__vite_plugin_react_preamble_installed__ = true
				</script>
			`)
		}

		b.WriteString(`
			<script type="module" src="{{ .DevServerURL }}/@vite/client"></script>
        	<script type="module" src="{{ .DevServerURL }}/{{ .SrcDir }}/main.tsx"></script>
		`)
	}

	tmpl, err := template.New("tags").Parse(b.String())

	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	tmpl.Execute(&buffer, v)

	return template.HTML(buffer.String()), nil
}
