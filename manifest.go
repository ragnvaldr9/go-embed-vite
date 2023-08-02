package vite

import (
	"encoding/json"
	"errors"
	"io/fs"
)

func read(fsys fs.FS, path string) ([]byte, error) {
	content, err := fs.ReadFile(fsys, path)

	return content, err
}

func mapChunck(c map[string]any, dist AssetsData) {
	for k, v := range c {
		dist[k] = v
	}
}

func checkBool(v any) bool {
	vv, ok := v.(bool)

	if ok && vv {
		return vv
	}

	return false
}

func processChunck(chunck, data AssetsData) []AssetsData {
	list := []AssetsData{}

	if checkBool(chunck["isEntry"]) {
		mapChunck(chunck, data)
		list = append(list, chunck)
	} else {
		var node = make(AssetsData)
		mapChunck(chunck, node)
		list = append(list, node)
	}

	return list
}

func mapManifest(m any) (AssetsData, []AssetsData, error) {
	manifest, ok := m.(map[string]any)

	if !ok {
		return nil, nil, errors.New(INVALID_MANIFEST_STRUCT)
	}

	raw := AssetsData{}
	var chuncks []AssetsData

	for _, chunck := range manifest {
		m, ok := chunck.(map[string]any)
		if ok {
			chuncks = processChunck(m, raw)
		} else {
			return nil, nil, errors.New(INVALID_MANIFEST_STRUCT)
		}
	}

	target := map[string]any{}
	target["file"] = raw["file"]
	target["css"] = raw["css"]
	target["assets"] = raw["assets"]
	target["imports"] = raw["imports"]
	target["dynamicImports"] = raw["dynamicImports"]

	return target, chuncks, nil
}

func (v *Vite) parseManifest(dist *fs.FS, path string) error {
	bytes, err := read(*dist, path)

	if err != nil {
		return err
	}

	var jsonData any

	json.Unmarshal(bytes, &jsonData)

	data, chuncks, err := mapManifest(jsonData)

	if err != nil {
		return err
	}

	v.setArgs(data, false)

	v.chuncks = &chuncks

	return nil
}
