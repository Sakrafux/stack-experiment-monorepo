parent maven pom

`<packaging>pom</packaging>`

```
<modules>
    <module>apps/.../module1</module>
    <module>apps/.../module2</module>
</modules>
```

all dependencies with versions

------

in child poms

**No** `<packaging>pom</packaging>`

reference parent like:

```
<parent>
    <groupId>com.sakrafux</groupId>
    <artifactId>stack-experiment-monorepo</artifactId>
    <version>${revision}</version>
    <relativePath>../../../../pom.xml</relativePath>
</parent>
```

maven variables like `${...}` will be resolved by parent

compilation targets need to be set by child

```
<properties>
    <maven.compiler.source>${java.version}</maven.compiler.source>
    <maven.compiler.target>${java.version}</maven.compiler.target>
</properties>
```

dependencies without version number -> resolved by parent

build plugins in child, for necessary version references -> variables resolved by parent

config files like `checkstyle.xml` can be referenced relatively like `<configLocation>../../../../checkstyle.xml</configLocation>`

now application can be run from monorepo root or in the individual project