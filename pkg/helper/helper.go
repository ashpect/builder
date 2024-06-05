package helpers

import (
	"github.com/moby/buildkit/client/llb"
)

func CopyForStates(src llb.State, dst llb.State, srcpaths []string, dstpath string) llb.State {
	return dst.With(
		CopyMultiple(src, srcpaths, dstpath),
	)
}

func CopyMultiple(src llb.State, srcPaths []string, destPath string) llb.StateOption {
	var stateOptions []llb.StateOption
	for _, srcPath := range srcPaths {
		stateOptions = append(stateOptions, CopyFrom(src, srcPath, destPath))
	}

	return func(s llb.State) llb.State {
		for _, stateOption := range stateOptions {
			s = stateOption(s)
		}
		return s
	}
}

func CopyFrom(src llb.State, srcPath, destPath string) llb.StateOption {
	return func(s llb.State) llb.State {
		return Copy(src, srcPath, s, destPath)
	}
}

func Copy(src llb.State, srcPath string, dest llb.State, destPath string) llb.State {
	return dest.File(llb.Copy(src, srcPath, destPath, &llb.CopyInfo{
		AllowWildcard:  true,
		AttemptUnpack:  true,
		CreateDestPath: true,
	}))
}

func ParseIntoString(args map[string]bool) (string, error) {

	var languages string
	for lang, _ := range args {
		languages += lang + " "
	}
	return languages, nil

}
