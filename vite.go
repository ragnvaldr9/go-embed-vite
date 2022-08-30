package vite

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"path"
)

type AssetsData = map[string]any
type Vite struct {
	RootFS          fs.FS
	DistFS          fs.FS
	Env             string
	Platform        string
	SrcDir          string
	AssetsPath      string
	AssetsDir       string
	AssetsURLPrefix string
	Template        *template.Template
	DevServerURL    string
	data            AssetsData
	chuncks         *[]AssetsData
}

var v *Vite

func NewVite(cfg *ViteConfig) (*Vite, error) {
	if cfg.RootFS == nil {
		return nil, errors.New(NO_ROOT_FS)
	}

	if cfg.Template == nil {
		return nil, errors.New(NO_TEMPLATE)
	}

	cfg.setDefaults()

	v = &Vite{
		data: AssetsData{},
	}

	v.Env = cfg.Env
	v.Platform = cfg.Platform
	v.Template = cfg.Template

	distFs, err := fs.Sub(cfg.RootFS, "static")

	if err != nil {
		return nil, err
	}

	v.DistFS = distFs

	if v.Env == "production" {
		v.AssetsDir = cfg.AssetsDir

		err := v.parseManifest(&v.DistFS, "manifest.json")

		if err != nil {
			return nil, err
		}

		v.AssetsPath = path.Join("static", v.AssetsDir)
	}

	if v.Env == "development" {
		v.SrcDir = cfg.SrcDir
		v.DevServerURL = fmt.Sprintf("http://%v:%v", cfg.DevServerHost, cfg.DevServerPort)
	}

	v.AssetsURLPrefix = cfg.AssetsURLPrefix

	return v, nil
}
