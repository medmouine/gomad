[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/medmouine/gomad.svg)](https://github.com/medmouine/gomad)
[![GitHub go.mod Go version of a Go module](https://github.com/medmouine/gomad/workflows/Go/badge.svg)](https://github.com/medmouine/gomad/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/medmouine/gomad/maybe.svg)](https://pkg.go.dev/github.com/medmouine/gomad/maybe)
[![codecov](https://codecov.io/gh/medmouine/gomad/branch/main/graph/badge.svg?token=3DJBNCU1NG)](https://codecov.io/gh/medmouine/gomad)

### Functional Patterns in Golang
# GOMAD (Early stage)
### This package is still in an early stage of development. Feel free to open a PR and contribute or just open an issue to help me priorities features.

Following is a list of future modules coming up in no specific order:

- [X] Maybe
- [X] Either
- [X] Result
- [ ] List
- [ ] Pipe
- [ ] HKT
- [ ] Monad
- [ ] IO
- [ ] Task
- [ ] Reader

### Prerequisites
All these modules use the newly added features of Golang v1.18 (Still in beta as of today) notably type and function generics.

## Install

```
go get github.com/medmouine/gomad/<Desired module>

i.e
go get github.com/medmouine/gomad/maybe
```

## Modules
### Maybe
`Maybe` is a monadic pattern allowing for data manipulation while abstracting whether or not the value actually exists or is `nil`. For example, if we fetch data from an external API that could be `nil`, we can still perform manipulation on it while disregarding its actual state. The `Maybe` struct will take care of managing the value itself. This is similar to the Maybe interface in [Elm](https://package.elm-lang.org/packages/elm/core/latest/Maybe) or [Haskell](https://wiki.haskell.org/Maybe) or [Optional in Java](https://docs.oracle.com/javase/8/docs/api/java/util/Optional.html). This is helpful for CRUD operations by simplifying the code and allowing for seamless manipulation of nullable data.

You can use the functions `Just`, `None` or `Nillable` to instanciate a `Maybe` struct. The type parameter will be determined by the passed argument or by specifying it. For example:
```
maybeNilInteger := maybe.Nillable[int](nil)

nilInteger := maybe.None[int]()

someInteger := maybe.Just[int](1)
```


<!-- #### Just
`Just` referes to a value we know is not `nil`.
```
TODO
```
##### None
`Just` referes to a value we know is `nil`.
```
TODO
``` -->
TODO
### Either
Allows to manipulate pairs of mutually exclusive data. For example, if we would want to fall back to a value B if A answers to a specific predicate. This interface allows integrating this behavior seamlessly by abstracting all the underlying logic of managing both values. A common use case for this is form validation for front-end applications.

TODO
### Result
This interface aim at abstracting all logic related to operations susceptible to failures, such as external API calls, etc. It offers constructors and methods to safely manipulate the result in case of success and handle errors gracefully in case of failure.

TODO
