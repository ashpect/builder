# MetaCall Builder

Advanced builder based on Buildkit for selectively build compact Docker images selecting the only required languages.

## Build

```sh
go build cmd/main.go
```

## Run

```sh
./main py node rb | buildctl build --output type=docker,name=imagename | docker load
```

## Run with buildctl-daemonless

Requirements:

- [BuildKit](https://github.com/moby/buildkit/releases)
- [RootlessKit](https://github.com/rootless-containers/rootlesskit/releases)

MacOs:

For MacOs, you can use install buildkit using brew and lima for rootless containers, and run the script after the installation.

```console
$ brew install buildkit
$ brew install lima
```

```sh
./main py node rb | ./hack/buildctl.sh build --output type=docker,name=imagename | docker load
```

## Run with docker

Use the environment variable `BUILDER_ARGS` for passing the arguments.

With daemon:
```sh
BUILDER_ARGS="runtime py" docker compose up --exit-code-from client client
```

Rootless:
```sh
BUILDER_ARGS="runtime node" docker compose up --exit-code-from rootless rootless
```

You can also run the builder binary only:

```sh
BUILDER_ARGS="runtime rb" docker compose up --exit-code-from binary binary
```

## Linter

```sh
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.59.1 golangci-lint run -v --enable-all
```

## Useful Tools

[Dive](https://github.com/wagoodman/dive) can be used to analyze each layer of the generated image

```sh
dive <your-image-tag>
```
This opens up a window where in we can see changes in each layer of the image according to preferences.

**Ctrl + L** : To show only layer changes

**Tab** : To switch view from layers to current layer contents



