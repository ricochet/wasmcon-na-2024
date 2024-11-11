# WasmCon NA 2024 - Choose your own adventure

## Getting Started

This workshop will walkthrough building a WebAssembly component with [WASIP2](https://github.com/WebAssembly/WASI/blob/main/wasip2/README.md),
how to assemble (compose) a declarative application, and finally deploy (run)
across diverse environments.

## Chapter 1: Build

WebAssembly (Wasm) components are a portable, language agnostic, and sandbox-able compilation target.

Many language toolchains support building WebAssembly component AKA a `.wasm` where the first bytes of the binary are `(component)`, by setting `wasm32-wasip2` as the target.

Some language toolchains benefit from an SDK
geared towards producing components like `componentize-py` and `componentize-dotnet`. For this workshop, we will focus on Go with the TinyGo compiler and Rust.

Once you have built a Wasm component, it is portable to any runtime that supports components with `wasip2`.
The WebAssembly runtime that we will use is [wasmtime](wasmtime.dev) from the Bytecode Alliance foundation.

The application platform that we will
use to run our Wasm applications is CNCF [wasmCloud](https://wasmcloud.com).

[wasmCloud](https://wasmcloud.com) has a CLI called [wash](https://wasmcloud.com/docs/installation)(wasmCloud shell) that simplifies the build, develop, and distribute loop of Wasm components. Install wash now. Note that all of the Wasm components that we will build
can also be built with direct tooling. That will take more flags and commands than `wash build`.

### Install wash

Follow the install instructions for installing wash for your system.

Now you must make your first major decision of this journey.
What type of application will you create and for which language?

What language will you use?

```bash
export WASMCON_LANG="go"
```

```bash
export WASMCON_LANG="rust"
```

```bash
WASMCON_APP="password-checker"
wash new component --git ricochet/wasmcon-na-2024 --subfolder ${WASMCON_LANG}/${WASMCON_APP}
```

Decide which application to create:
```bash
WASMCON_APP="dog-fetcher"
```

Or

```bash
WASMCON_APP="password-checker"
```

### Go with TinyGo

```bash
wash new component --git ricochet/wasmcon-na-2024 --subfolder go/dog-fetcher
cd component
wash dev
```

Checkout the cutest dogs on the planet by continuing to download random dogs. Who will get Fallout's dogmeat?

```bash
curl localhost:8000
```

or

```bash
wash new component --git ricochet/wasmcon-na-2024 --subfolder go/password-checker
cd component
./setup.sh # little pre-populate 
wash dev
```

For password-checker, check out the possible inputs:

```bash
curl localhost:8000/api/v1/check -d '{"value": "wasmconisawesome"}'
{"valid":true,"length":16}

curl localhost:8000/api/v1/check -d '{"value": "veryshort"}'
{"valid":false,"message":"insecure password, try including more special characters, using uppercase letters, using numbers or using a longer password"}

curl localhost:8000/api/v1/check -d '{"value": "1234"}'
{"valid":false,"message":"password is in the list of 500 worst passwords"}
```

For the curious, this is the equivalent command to build a go component:

```bash
go generate
tinygo build --target=wasip2 --wit-package ./wit --wit-world password-checker
```

### Rust

```bash
wash new component --git ricochet/wasmcon-na-2024 --subfolder rust/dog-fetcher
cd component
wash dev
```

```bash
wash new component --git ricochet/wasmcon-na-2024 --subfolder rust/password-checker
cd component
wash dev
```

### Test

Were you successful?

```bash
curl localhost:8000
```

How large is your binary?

```bash
du -h build/component.wasm
```

Let's get WIT-ty

```bash
wasm-tools component wit build/component.wasm
```

### Extend

Let's see what it's like to add a new dependency. 

Search for and uncomment the code for `PART 2:`

1. First uncomment blobstore in the `world.wit`
2. Then uncomment all of the `PART 2:` sources

```bash
# fetch new WIT deps
# both `wash dev` and `wash build` call `wash wit deps`
# to populate deps based on world.wit
wash build

# P.S. to manually do this with Bytecode Alliance tooling, run:
# wkg wit fetch

# regenerate bindings if necessary
# go generate
# rebuild
```

`wash build` detected a change in the world definition and automatically fetched WIT dependencies.

## Chapter 2: Compose

Applications are composed of components and optionally stateful providers.

In cloud native nomenclature, an application is the combination of microservices and necessary services like stateful storage and hardware enabled capabilities to run your app. Components are equivalent to stateless 12-factor apps. A provider is a type of host plugin that is used to extend the host to stateful services. Providers can be builtin to the host or run remotely.

You might be wondering, why would I need an application platform if wasmtime (with plugins) should be able to run my component? For the same reason we need kubernetes to orchestrate containers at scale, we gain many benefits of Wasm native orchestration. WebAssembly components are a fundamentally new unit of compute whose benefits are best realized via native orchestration.

Components offer a fundamentally finer-grained abstraction than containers, like Kubernetes for WebAssembly, so wasmCloud provides a Wasm-native orchestrator to best take advantage of the unique properties that WebAssembly components can provide. Wasm-native works with cloud-native and runs seamlessly on Kubernetes or any other container execution engine like AWS Fargate, Microsoft AKS, or Google Cloud Run.

wasmtime is like the kubelet, and wasmCloud is the entire platform like kubernetes including a scheduler and runtime abstraction.

Unlike microservices where we need to layer on API Gateways and service meshes to understand the API provided by a microservice, with components, they are themselves declarative and introspectable.

Remember where we looked at and modified the WIT? Let's take that a step further by **linking** our component
interface with other interfaces declaratively.

The declarative application manifest that is scheduled by `wadm`.

This is the stage where we most often operate. Writing applications has never been the problem. Maintaining them
is the real challenge in software development today.

Today we have platform engineering teams that provide core platform services, think databases, authentication, observability services, etc. The tension between getting teams to upgrade their dependencies and being able to use the latest features of a platform, can now be decoupled from the component that contains only business logic, and
the platform services who may now update at *runtime*.

We've been running with the filesystem. That's really convenient for local development but that's not what we
should use for reliability, scalability, and distributability. Blob storage is great at this! Let's use it.

This compose step is the bridge between developer teams and platform engineering.

Let's create a declarative manifest from our dev loop:

```bash
wash dev --manifest-output-dir .
```

Now modify that manifest to use `blobstore-s3`!

```bash
# find and uncomment part 2 code
# PART 2:
```

## Chapter 3: Run

This part is often scary for developers. Running reliably at scale in environments that are very different
from local development machines, developers are used to experiencing significant friction to deploy into
environments with controls.

With CNCF [wasmCloud](wasmcloud.com), whether the deployment is running on your local machine, in the cloud,
or on the edge, the experience for developers and platform engineers is the same.

Let's pretend for a second that your laptop is a kind of edge.

```bash
wash up -d
wash app deploy wadm.yaml
```

Let's scale out the apps to another host.

```bash
wash up -d
```

```bash
wash ui
```

```bash
curl localhost:8080
```

We need to get rid of the local file if we want to deploy remotely.
Let's push our `.wasm` as an OCI artifact to GHCR.

```bash
wash push ghcr.io/ricochet/components/${WASMCON_APP}-${WASMCON_LANG}:0.2.0 build/*_s.wasm

# P.S. BA tooling for this is wkg
```

Do we need to care what language this app was written in? Nope. It's portable. But is it cloud native? Let's find out.

Shutdown local hosts.

```bash
wash down --all
```

Modify local file path to use your OCI ref.

```bash
wash app deploy wadm.yaml
```

### Deploying to Kubernetes

If you have access to another k8s cluster, deploy the wasmCloud platform helm chart.

For running k8s locally, use kind.

```bash
./deploy/setup-kind.sh
```

```bash
wash app deploy wadm.yaml
# or
kubectl app deploy wadm.yaml
```

### Your princess is in another cluster

```bash
# start another kind cluster
# deploy to that cluster as well
```

### Users are on the edge

Change the label to the host on my lattice to my edge host.

```bash
wash app deploy wadm.yaml
```

Noticing a trend? The experience is the *same*. Unify software development lifecycle steps of build, compose, and run.

With a unified control plane, it is now possible to spread compute across clouds, regions, and edges including your own.
