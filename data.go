package vite

var reservedKeys = map[string]bool{
	"file":           true,
	"src":            true,
	"isEntry":        true,
	"isDynamicEntry": true,
	"dynamicImports": true,
	"css":            true,
	"assets":         true,
}

func (v *Vite) setArgPrivate(k string, val any) error {
	v.data[k] = val

	return nil
}

func (v *Vite) setArgPublic(k string, val any) error {
	if !reservedKeys[k] {
		v.data[k] = val
	} else {
		return createReadOnlyError(k)
	}

	return nil
}

func (v *Vite) setArg(k string, val any, isPublic bool) error {
	if isPublic {
		return v.setArgPrivate(k, v)
	} else {
		return v.setArgPublic(k, v)
	}
}

func (v *Vite) updateArgs(data AssetsData, setter func(k string, v any) error) error {
	var err error

	for k, val := range data {
		err = setter(k, val)
	}

	return err
}

func (v *Vite) setArgs(data AssetsData, isPublic bool) error {
	if isPublic {
		return v.updateArgs(data, v.setArgPublic)
	} else {
		return v.updateArgs(data, v.setArgPrivate)
	}
}

func (v *Vite) Data() AssetsData {
	return v.data
}

func (v *Vite) SetArgs(vars AssetsData) error {
	return v.setArgs(vars, true)
}

func (v *Vite) SetArg(k string, val any) error {
	return v.setArgPublic(k, v)
}
