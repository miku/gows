# Go Workshop Notes

> time and place: TBD

## Outline

> 14/28

* [Introduction](Intro.md)
    * Motivation and Landscape
    * Hello World
* [History](History.md)
    * Timeline of events
* [More First Programs](MoreFirstPrograms.md)
    * Entry point
    * Importing Code
    * Visibility (public, private)
* [Language Overview](Language.md)
    * Basic Types
    * Control Structures: if, for, switch
* [More Types](MoreTypes.md)
    * Arrays
    * Slices
    * Maps
    * Structs
    * Methods
    * Pointers
    * Gotchas
* [Interfaces](Interfaces.md)
    * Structural Typing
    * Small interfaces
    * Variants (basic, embedding, general, [ref/spec](https://go.dev/ref/spec#Interface_types))
* [Go OOP?](OO.md)
    * Is Go object oriented? [FAQ](https://go.dev/doc/faq#Is_Go_an_object-oriented_language)
* [Error Handling](Errors.md)
    * Custom Error Types
    * Wrapping and unwrapping errors
* [Project Layout](Projects.md)
    * typical structure
    * naming recommendations
    * import path and resolution
    * Go modules
    * mixing public and private code
    * versioning libraries
* [IO](IO.md)
    * working with files
    * readers and writers
* [Concurrency](Concurrency.md)
    * classic and CSP style
    * goroutines
    * channels
    * select
    * the `sync` package
    * error handling
    * helpers: `errgroup`
* [Testing Go Code](Testing.md)
    * Unit Test
    * Subtests
    * Benchmarks
    * Testcontainers
* [HTTP clients](HTTP.md)
    * clients and transport
    * standard clients, third party clients
* [HTTP servers](Servers.md)
    * handlers
    * router, e.g. gorilla/mux
    * testing
* [database access](Databases.md)
    * package db, db/sql
    * database drivers
    * sqlx helper
