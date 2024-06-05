package staging

import (
	helpers "github.com/metacall/builder/pkg/helper"
	"github.com/moby/buildkit/client/llb"
)

// Contains the commands in steps specific to doing operations on any llb for MetaCall

func DownloadPackages(state *llb.State) {
	state.Run(llb.Shlex("ls -l /bin")).
		Run(llb.Shlex("apt-get update")).
		Run(llb.Shlex("apt-get install -y --no-install-recommends build-essential git cmake libgtest-dev wget apt-utils apt-transport-https gnupg dirmngr ca-certificates"))
}

func MetaCallClone(state *llb.State, branch string) {
	state.Run(llb.Shlexf("git -c http.sslVerify=false clone --depth 1 --single-branch --branch=%v https://github.com/metacall/core.git", branch)).
		Run(llb.Shlex("mkdir core/build"))
}

func MetaCallEnvBase(state *llb.State, arg string) {
	state.Run(llb.Shlexf("bash tools/metacall-environment.sh base backtrace %v", arg))
}

func MetaCallConfigure(state *llb.State, arg string) {
	state.File(llb.Mkdir("build", 0777)).
		Run(llb.Shlexf("bash tools/metacall-configure.sh tests scripts ports install %v", arg))
}

func MetaCallBuild(state *llb.State) {
	state.Run(llb.Shlex("bash tools/metacall-build.sh"))
}

func MetacallRuntime(state *llb.State, arg string) {
	state.Run(llb.Shlexf("bash tools/metacall-runtime.sh backtrace ports %v", arg))
}

func RemoveBuild(state *llb.State) llb.State {
	return state.File(llb.Rm("build"))
}

func AddCli(src, dst llb.State) llb.State {
	return dst.With(helpers.CopyFrom(src, "/usr/local/bin/metacallcli*", "/usr/local/bin/metacall"))
}
