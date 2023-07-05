# Guide to setup a JavaScript/TypeScript Monorepo with Nx

## Root

Install `nx` globally for ease of use with `npm install --global nx@latest`

Ideally start out with `npm create nx-workspace` and create an integrated monorepo, since this creates all the boilerplate required. If the repository already exists, running the command to create a new (temporary) repository and copy all the required files from there.

`npx nx@latest init` doesn't create all the boilerplate and needs some work afterwards.

### Structure

In `nx.json` change the following default configuration to whatever target locations you require:

```
"workspaceLayout": {
    "appsDir": "package",
    "libsDir": "package"
},
```

You need to manually install `@nx/angular`, `@nx/react` or `@nx/node` depending on what applications you want to generate with `nx`. Alternatively, you can setup projects manually as well.

The command `nx generate @nx/react:app <app>` then creates a fully runnable application.

The target for the generate command can also be path like `path/to/app` and will be placed in the specified `appsDir`.

-----

## Projects

The created projects are fully preconfigured and workable. Any `nx`-specific (and some others, e.g. ESLint, tsconfig) configurations reference either the app by its fully qualified path starting from `appsDir` (or `libsDir`) or relative paths like `../../..`

The default configuration creates both `/dist` and `/coverage` at root level, but since we may not want this, because we want to place the output next to its sources, we need to change `"outputPath": "dist/<appsDir>/.../app"` to `"outputPath": "<appsDir>/.../app/dist"` inside `project.json`.

For the coverage, change to `coverageDirectory: './coverage'` in `jest.config.ts`

For convenience scripts when using the specific project as working directory, add a `package.json` with the following scripts:
```
{
  "name": "<app>",
  "scripts": {
      "start": "nx serve <app>",
      "build": "nx build <app>",
      "test": "nx test <app>"
  },
  "nx": {
      "includedScripts": []
  }
}
```

-----

## More on this
https://nx.dev/getting-started/tutorials/integrated-repo-tutorial