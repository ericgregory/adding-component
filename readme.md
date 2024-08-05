---
title: "README"
---

# Component example

This example implements a [Go-based adding component example from the Component Model Book](https://component-model.bytecodealliance.org/language-support/go.html) with the **washCloud Shell (`wash`)** CLI. The repo includes:

* **`add.go`** - The Go code for the component
* **`wit/world.wit`** - The WebAssembly Interface Type (WIT) file defining the `add` interface
* **`wasmcloud.toml`** - A TOML file defining metadata and configuring the build process for `wash build`

Clone the repo and `wash build` from the project repo:

```shell
wash build
```

This will generate bindings in the `gen` directory and create a Wasm binary in the `build` directory. (Actually, it will create two binaries, but for our purposes we're only interested in the final output, `adder_s.wasm`.)

Once the component is built, use `wash inspect` to view the component's imports and exports:

```shell
wash inspect --wit ./build/adder_s.wasm
```
```shell
package root:component;

world root {
  import wasi:io/error@0.2.0;
  import wasi:io/streams@0.2.0;
  import wasi:cli/stdin@0.2.0;
  import wasi:cli/stdout@0.2.0;
  import wasi:cli/stderr@0.2.0;
  import wasi:clocks/wall-clock@0.2.0;
  import wasi:filesystem/types@0.2.0;
  import wasi:filesystem/preopens@0.2.0;

  export docs:adder/add@0.1.0;
}
```

We're exporting the add interface. 