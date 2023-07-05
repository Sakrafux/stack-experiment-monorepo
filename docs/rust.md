create `Cargo.toml` in root

declare workspace with members:

```
[workspace]
members = ["apps/.../rust"]
```

inherited keys and dependencies like this
```
[workspace.package]
authors = ["Andreas Hell"]

[workspace.dependencies]
axum = "0.6.18"
hyper = { version = "0.14.27", features = ["full"] }
tokio = { version = "1.29.1", features = ["full"] }
tower = "0.4.13"
```

-----

create child project with `cargo new` and fill like usual except inherited keys and dependencies like this:

```
authors.workspace = true

[dependencies]
axum.workspace = true
hyper = { workspace = true, features = ["full"] }
tokio = { workspace = true, features = ["full"] }
tower.workspace = true
```

