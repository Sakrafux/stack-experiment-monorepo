install `nx` globally for ease of use like `npm install --global nx@latest`

ideally start out with `npm create nx-workspace` and create an integrated monorepo

 if repository already exists, run command to create child directory and copy-paste to parent

`npx nx@latest init` doesn't create all the boilerplate

in `nx.json` change the following default configuration to whatever you want:

```
"workspaceLayout": {
    "appsDir": "package",
    "libsDir": "package"
},
```

manually install `@nx/angular`, `@nx/react` or `@nx/node` depending on what applications you want to generate

`nx generate @nx/react:app <app>`

the target can also be path like `path/to/app` and will start from the specified `appsDir`

-----

the created projects are fully preconfigured and workable

the default configuration creates both `/dist` and `/coverage` at root level

we may not want this because we only want to work inside a single project

change `"outputPath": "dist/apps/realworld/frontend/react"` to `"outputPath": "apps/realworld/frontend/react/dist"` inside `project.json`

change to `coverageDirectory: './coverage'` in `jest.config.ts`


for convenience, add a `package.json` with the following scripts:
```
{
  "name": "realworld-frontend-react",
  "scripts": {
      "start": "nx serve realworld-frontend-react",
      "build": "nx build realworld-frontend-react",
      "test": "nx test realworld-frontend-react"
  },
  "nx": {
      "includedScripts": []
  }
}
```