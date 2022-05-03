# Go Workshop Notes

> time and place: TBD

## Outline

> 14/28

* [Introduction](01-Intro.md)
    * Motivation and Landscape
    * Hello World
    * Task: "helloworld"
* [History](02-History.md)
    * Timeline of events
* [More First Programs](03-MoreFirstPrograms.md)
    * Entry point
    * Importing Code
    * Visibility (public, private)
* [Language Overview](04-Language.md)
    * Basic Types
    * Variable Declarations
    * Control Structures: if, for, switch
* [More Types](05-MoreTypes.md)
    * Arrays
    * Slices
    * Maps
    * Structs
    * Methods
    * Pointers
    * Gotchas
* [Interfaces](06-Interfaces.md)
    * Structural Typing
    * Small interfaces
    * Variants (basic, embedding, general, [ref/spec](https://go.dev/ref/spec#Interface_types))
* [Go OOP?](07-OO.md)
    * Is Go object oriented? [FAQ](https://go.dev/doc/faq#Is_Go_an_object-oriented_language)
* [Error Handling](08-Errors.md)
    * Custom Error Types
    * Wrapping and unwrapping errors
* [Project Layout](09-Projects.md)
    * typical structure
    * naming recommendations
    * import path and resolution
    * Go modules
    * mixing public and private code
    * versioning libraries
* [IO](10-IO.md)
    * working with files
    * readers and writers
* [Serialization](11-Serialization.md)
    * struct tags
    * JSON
    * XML
* [Concurrency](12-Concurrency.md)
    * classic and CSP style
    * goroutines
    * channels
    * select
    * the `sync` package
    * error handling
    * helpers: `errgroup`
* [Testing Go Code](13-Testing.md)
    * Unit Test
    * Subtests
    * Benchmarks
    * Testcontainers
* [HTTP clients](14-HTTP.md)
    * clients and transport
    * standard clients, third party clients
* [HTTP servers](15-Servers.md)
    * handlers
    * router, e.g. gorilla/mux
    * testing
* [Database access](16-Databases.md)
    * package db, db/sql
    * database drivers
    * sqlx helper
