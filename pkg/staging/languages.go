package staging

import "errors"

// Contains the list of languages used to validate the args
// Will also contain the any other command line arguments we want to give to set up env

type Match struct {
	str  string
	flag bool
}

var languageValidate = map[string]Match{

	"base":       {"base", false},
	"py":         {"python", false},
	"rb":         {"ruby", false},
	"nodejs":     {"nodejs", false},
	"typescript": {"typescript", false},
	"wasm":       {"wasm", false},
	"java":       {"java", false},
	"c":          {"c", false},
	"cobol":      {"cobol", false},
	"go":         {"go", false},
	"rust":       {"rust", false},

	// "cache": "cache",
	// "netcore":   "netcore",
	// "netcore2":  "netcore2",
	// "netcore5":  "netcore5",
	// "netcore7":  "netcore7",
	// "rapidjson": "rapidjson",
	// "funchook":  "funchook",
	// "v8":        "v8",
	// "v8rep54":   "v8rep54",
	// "v8rep57":   "v8rep57",
	// "v8rep58":   "v8rep58",
	// "v8rep52":   "v8rep52",
	// "v8rep51":   "v8rep51",
	// "file":       "file",
	// "rpc":        "rpc",
	// "swig":  "swig",
	// "pack":     "pack",
	// "coverage": "coverage",
	// "clangformat": "clangformat",
	// "backtrace"	: "backtrace",
	// "sandbox"	: "sandbox",
	// "scripts":    "scripts",
	// "examples":   "examples",
	// "tests":      "tests",
	// "benchmarks": "benchmarks",
	// "ports":      "ports",
}

func (mb *metaBuilder) ValidateLanguages(args []string) error {
	lang := make(map[string]bool)
	for _, arg := range args {
		if _, ok := languageValidate[arg]; !ok {
			return errors.New("Invalid language")
		}
		lang[arg] = true
	}

	mb.meta.languages = lang
	return nil
}
