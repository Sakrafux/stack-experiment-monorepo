# Stack Experiment Monorepo

Like the name of repo may suggest, this repo contains code for multiple example applications in multiple languages.

To not only make use of code colocation for easy comparison between languages and holistic development of both front- and backend, but also to test out the concept, I designed this in a Monorepo structure.

Even though some tools like e.g. **Bazel** seem to be ideal for such a use case, the lack of proper tooling, IDE integration and straight-up bugs led me to instead mix and match tools for muliple languages in the same repo.

So far this includes **Maven** using *Maven Modules* for Java, **Nx** for JavaScript/TypeScript and **Cargo** using *Cargo Workspace* for Rust.

Though I may lose some of the power of a single master build tool, like incremental builds, caching, etc., this doesn't really matter, since the scale of this repo will be too small to warrant such features and the contained projects will be largely independent. More specifically, they will be independent from other language modules and importing modules of the same language is usually handled in their language specific build toolchain.

