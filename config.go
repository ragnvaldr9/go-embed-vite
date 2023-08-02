package vite

import (
	"html/template"
	"io/fs"
)

type ViteConfig struct {
	// file system of project root (required)
	RootFS fs.FS
	// pointer to html template (required)
	Template *template.Template
	// name of frontend source folder (need for development mode, default - src)
	SrcDir string
	// production or development, default - production
	Env string
	// react, vue, svelte, default - react
	Platform string
	// name of vite build output folder (default - dist)
	OutDir string
	// name of static assets folder (default - assets)
	AssetsDir string
	// AssetsURLPrefix (/assets/ for prod, /src/ for dev)
	AssetsURLPrefix string
	// host of dev server (default - localhost)
	DevServerHost string
	// port of dev server (default - 3000)
	DevServerPort string
	// frontend app main file (default - main.tsx)
	EntryPoint string
}

var defaults = map[string]string{
	"Env":                 "production",
	"Platform":            "react",
	"SrcDir":              "src",
	"AssetsURLPrefixProd": "/assets/",
	"AssetsURLPrefixDev":  "/src/",
	"OutDir":              "dist",
	"AssetsDir":           "assets",
	"DevServerHost":       "localhost",
	"DevServerPort":       "3000",
	"EntryPoint":          "main.tsx",
}

func (cfg *ViteConfig) setProdDefaults() {
	if cfg.AssetsURLPrefix == "" {
		cfg.AssetsURLPrefix = defaults["AssetsURLPrefixProd"]
	}
}

func (cfg *ViteConfig) setDevDefaults() {
	if cfg.DevServerHost == "" {
		cfg.DevServerHost = defaults["DevServerHost"]
	}

	if cfg.DevServerPort == "" {
		cfg.DevServerPort = defaults["DevServerPort"]
	}

	if cfg.AssetsURLPrefix == "" {
		cfg.AssetsURLPrefix = defaults["AssetsURLPrefixDev"]
	}

	if cfg.SrcDir == "" {
		cfg.SrcDir = defaults["SrcDir"]
	}

	if cfg.EntryPoint == "" {
		cfg.EntryPoint = defaults["EntryPoint"]
	}
}

func (cfg *ViteConfig) setDefaults() {
	if cfg.Env == "" {
		cfg.Env = defaults["Env"]
	}

	if cfg.Platform == "" {
		cfg.Platform = defaults["Platform"]
	}

	if cfg.OutDir == "" {
		cfg.OutDir = defaults["AssetsDir"]
	}

	if cfg.AssetsDir == "" {
		cfg.AssetsDir = defaults["AssetsDir"]
	}

	if cfg.Env == "production" {
		cfg.setProdDefaults()
	}

	if cfg.Env == "development" {
		cfg.setDevDefaults()
	}
}
