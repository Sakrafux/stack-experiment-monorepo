# Guide to setup monorepo capabilities with Maven for Java

## Parent Pom

This `pom.xml` lies in the repo root and will manage all the dependency versions to ensure consistency. Furthermore, it enables local libraries to depend on without the need to upload artifacts.

It is important to mark this pom with the correct packaging mode like this:
`<packaging>pom</packaging>`

Any new project needs to be registered in the parent pom in the `<modules>`-section:

```
<modules>
    <module>apps/.../module1</module>
    <module>apps/.../module2</module>
</modules>
```

All dependencies that are required by any project will need to be specified here as usual.

Since some versions are still required to be annotated in the child poms, handling all versions with `<properties>` is adviced.

------

## Child Pom

Ensure that **no** `<packaging>pom</packaging>` is present, as this prevents the building of any artifacts.

The parent pom needs to be referenced like this:

```
<parent>
    <groupId>com.example</groupId>
    <artifactId>demo</artifactId>
    <version>${revision}</version>
    <relativePath>../../../../pom.xml</relativePath>
</parent>
```

The properties defined in the parent are available here as well with simple `${...}` syntax.

The compilation targets for the compilere need to be set by child pom like this:

```
<properties>
    <maven.compiler.source>${java.version}</maven.compiler.source>
    <maven.compiler.target>${java.version}</maven.compiler.target>
</properties>
```

Although all the required dependencies still need to be specified, their version numbers will be automatically inferred by the parent.

Any required build plugins need to be placed in the child and as they may need to reference dependencies with specific versions, the properties defined by the parent can once again be used to ensure consistency.

Config files e.g. `checkstyle.xml` can be referenced relatively like `<configLocation>../../../../checkstyle.xml</configLocation>`.

The application can now be run from the monorepo root or in the individual project by itself.

-----

## More on this
https://maven.apache.org/guides/mini/guide-multiple-modules-4.html

https://www.baeldung.com/maven-multi-module
