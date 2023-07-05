# Guide to setup monorepo capabilities with Cargo for Rust

## Parent Cargo.toml

Manually create a `Cargo.toml` in the repo root and declare a workspace with members:

```
[workspace]
members = ["apps/.../app"]
```

Every new project needs to registered here with its full path.

All the dependencies used by the projects need to be defined for the workspace. Keys (limited variables) can also be defined, both togehter looking like this:
```
[workspace.package]
authors = ["Author"]

[workspace.dependencies]
dependency1 = "0.1.0"
dependency2 = { version = "0.1.0", features = ["full"] }
```

-----

## Child Cargo.toml

Create a new project with `cargo new` and use it like usual.

Keys and dependencies can now be inherited like this:

```
authors.workspace = true

[dependencies]
dependency1.workspace = true
dependency2 = { workspace = true, features = ["full"] }
```

-----

## More on this
https://doc.rust-lang.org/book/ch14-03-cargo-workspaces.html

https://doc.rust-lang.org/cargo/reference/workspaces.html