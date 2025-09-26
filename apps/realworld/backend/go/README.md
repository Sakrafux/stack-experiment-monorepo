# ![RealWorld Example App](../../logo.png)

> ### Go codebase containing real world examples (CRUD, auth, advanced patterns, etc) that adheres to the [RealWorld](https://github.com/gothinkster/realworld) spec and API.


### [Demo](https://demo.realworld.io/)&nbsp;&nbsp;&nbsp;&nbsp;[RealWorld](https://github.com/gothinkster/realworld)


This codebase was created to demonstrate a fully fledged fullstack application built with **Go** including CRUD operations, authentication, routing, pagination, and more.

For more information on how to this works with other frontends/backends, head over to the [RealWorld](https://github.com/gothinkster/realworld) repo.


# How it works

It's a simple Go application.

The openapi specification has been faithfully translated to dtos and endpoints.

Routing is handled via Chi router.

Database access is done via SQLX and a postgres driver.

Otherwise, the implementation is largely reliant on the standard library.

# Experiences

Go feels very verbose to use, which is largely driven by the explicit handling of errors as values. While in other languages you often just
re-throw exception implicitly until you want to deal with them, in Go you either have to explicitly handle them, regardless of whether you
want to re-throw or deal with them.

Another issue is the lack of advanced language features, especially annotations for code generation, which make it cumbersome to implement 
standard features like mappers without relying on excessive reflection during runtime, which just degrades performance. It makes it cumbersome
to statically generate code, which is very typical for strictly defined operations such as mapping types.

The packaging itself is very weird and opinionated. Basically, all files in the same package are treated as a single file for all intents
and purposes, which isn't that much of an issue at first. However, simultaneously limiting this to the same directory level severely limits
one's options in structuring the code, especially considering that cyclical imports are completely forbidden. You can't just import 
a single member you need, like a type definition, which means you would have to extract all common members into another package, but this
often means we are separating code that semantically belongs together. 
This leads to such situations such as, while I would like all APIs to be in the same package, I want to separate them by domain. Then maybe 
I have so many endpoint methods that I would like to split them into different files simply for code organization. It is easy to see how this
could lead to either very large files or many files in the same directory with even moderately sized applications.

The issue of cyclical dependencies makes layered architectures very difficult to use without explicitly defining the interfaces and types 
everywhere. At that point, you can already use hexagonal architecture instead, which admittedly feels more natural for the way Go works. 
However, it is often overkill for small-to-medium projects. For small projects, a very flat structure without greater architecture feels
great to use. Being forced into certain architecture paradigms can then make certain operations very cumbersome, i.e. cross-boundary operations.

While the idea of using implicit interfaces certainly has its merits and can be used to great success, the lack of implicit documentation
usually found by using a defined interface is apparent. At least optional syntax would be helpful.

The code is not optimal and could potentially be optimized, reducing code by ~20% and making some iterated database accesses single access.

As the result is a compiled binary, the program executes expectedly fast. 

----

# Statistics

Experience Level (Language): 2/5

Time to Complete: 20h

LoC: 2406