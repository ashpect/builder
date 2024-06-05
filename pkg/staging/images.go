package staging

import (
	helpers "github.com/metacall/builder/pkg/helper"
	"github.com/moby/buildkit/client/llb"
)

type Meta struct {
	branch       string
	languages    map[string]bool
	baseImage    llb.State
	depsImage    llb.State
	devImage     llb.State
	runtimeImage llb.State
}

type MetaBuilder interface {
	SetBranch(branch string) MetaBuilder
	GetDepsImage() llb.State
	GetDevImage() llb.State
	GetRuntimeImage() llb.State
	BuildBaseImage(base llb.State) MetaBuilder
	BuildDepsImage() MetaBuilder
	BuildDevImage() MetaBuilder
	BuildRuntimeImage() MetaBuilder
	ConstructMetaImage(imgtype string, branch string, base llb.State, args []string) *Meta

	Build() *Meta
}

func NewMetaBuilder() MetaBuilder {
	return &metaBuilder{
		meta: &Meta{},
	}
}

type metaBuilder struct {
	meta *Meta
}

func (mb *metaBuilder) SetBranch(branch string) MetaBuilder {
	mb.meta.branch = branch
	return mb
}

func (mb *metaBuilder) GetDepsImage() llb.State {
	return mb.meta.depsImage
}

func (mb *metaBuilder) GetDevImage() llb.State {
	return mb.meta.devImage
}

func (mb *metaBuilder) GetRuntimeImage() llb.State {
	return mb.meta.runtimeImage
}

func (mb *metaBuilder) BuildBaseImage(baseImage llb.State) MetaBuilder {
	mb.meta.baseImage = baseImage
	DownloadPackages(&mb.meta.baseImage)
	MetaCallClone(&mb.meta.baseImage, mb.meta.branch)

	return mb
}

func (mb *metaBuilder) BuildDepsImage() MetaBuilder {

	args, err := helpers.ParseIntoString(mb.meta.languages)
	if err != nil {
		panic(err)
	}
	mb.meta.depsImage = mb.meta.baseImage
	MetaCallEnvBase(&mb.meta.depsImage, args)

	return mb
}

func (mb *metaBuilder) BuildDevImage() MetaBuilder {
	args, err := helpers.ParseIntoString(mb.meta.languages)
	if err != nil {
		panic(err)
	}

	mb.BuildDepsImage()
	mb.meta.devImage = mb.meta.depsImage
	MetaCallConfigure(&mb.meta.devImage, args)
	MetaCallBuild(&mb.meta.devImage)

	return mb
}

func (mb *metaBuilder) BuildRuntimeImage() MetaBuilder {
	args, err := helpers.ParseIntoString(mb.meta.languages)
	if err != nil {
		panic(err)
	}
	mb.meta.runtimeImage = mb.meta.baseImage
	MetacallRuntime(&mb.meta.runtimeImage, args)

	return mb
}

func (mb *metaBuilder) ConstructMetaImage(imgtype string, branch string, base llb.State, args []string) *Meta {

	mb.SetBranch(branch)
	mb.BuildBaseImage(base)
	if imgtype == "deps" {
		mb.BuildDepsImage()
	} else if imgtype == "dev" {
		mb.BuildDevImage()
	} else if imgtype == "runtime" {
		mb.BuildRuntimeImage()
	}

	err := mb.ValidateLanguages(args)
	if err != nil {
		panic(err)
	}

	return mb.Build()
}

func (mb *metaBuilder) Build() *Meta {
	return mb.meta
}
