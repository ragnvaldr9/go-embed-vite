package vite

var validManifest string = `
	{
		"src/main.tsx": {
		  "file": "assets/main.647737a5.js",
		  "src": "src/main.tsx",
		  "isEntry": true,
		  "dynamicImports": [
			"src/components/screens/Home/dynamic/dynamic.js"
		  ],
		  "css": [
			"assets/main.ea45a986.css"
		  ],
		  "assets": [
			"assets/logo.ecc203fb.svg"
		  ]
		},
		"src/components/screens/Home/dynamic/dynamic.js": {
		  "file": "assets/dynamic.f001af7c.js",
		  "src": "src/components/screens/Home/dynamic/dynamic.js",
		  "isDynamicEntry": true
		}
  	}
`

var expectedMap AssetsData = AssetsData{
	"file":           "assets/main.647737a5.js",
	"css":            []any{"assets/main.ea45a986.css"},
	"assets":         []any{"assets/logo.ecc203fb.svg"},
	"imports":        nil,
	"dynamicImports": []any{"src/components/screens/Home/dynamic/dynamic.js"},
}

var expectedChuncks = []AssetsData{
	{
		"file":           "assets/main.647737a5.js",
		"src":            "src/main.tsx",
		"isEntry":        true,
		"css":            []any{"assets/main.ea45a986.css"},
		"assets":         []any{"assets/logo.ecc203fb.svg"},
		"dynamicImports": []any{"src/components/screens/Home/dynamic/dynamic.js"},
	},
	{
		"file":           "assets/dynamic.f001af7c.js",
		"src":            "src/components/screens/Home/dynamic/dynamic.js",
		"isDynamicEntry": true,
	},
}

var invalidManifest = `test`

var multipleEntryManifest = `
	{
		"src/main.tsx": {
		  "file": "assets/main.647737a5.js",
		  "src": "src/main.tsx",
		  "isEntry": true
		},
		"src/second.tsx": {
		  "file": "assets/second.f001af7c.js",
		  "src": "src/second.tsx",
		  "isEntry": true
		}
  	}
`
