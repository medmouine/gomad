[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/medmouine/gomad.svg)](https://github.com/medmouine/gomad)
[![GitHub go.mod Go version of a Go module](https://github.com/medmouine/gomad/workflows/Go/badge.svg)](https://github.com/medmouine/gomad/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/medmouine/gomad/maybe.svg)](https://pkg.go.dev/github.com/medmouine/gomad/maybe)
[![codecov](https://codecov.io/gh/medmouine/gomad/branch/main/graph/badge.svg?token=3DJBNCU1NG)](https://codecov.io/gh/medmouine/gomad)

### Functional Patterns in Golang
# GOMAD (Early stage)
### This package is still in an early stage of development. Feel free to open a PR and contribute or just open an issue to help me priorities features.

Following is a list of future modules coming up in no specific order:

- [X] Maybe
- [ ] Either
- [ ] Result
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
`Maybe` is a monadic pattern allowing for data manipulation while abstracting whether or not the value actually exists or is `nil`. For example, if we fetch data from an external API that could be `nil`, we can still perform manipulation on it while disregarding its actual state. The `Maybe` struct will take care of managing the value itself. This is similar to the Maybe interface in [Elm](https://package.elm-lang.org/packages/elm/core/latest/Maybe) or [Haskell](https://wiki.haskell.org/Maybe) or [Optional in Java](https://docs.oracle.com/javase/8/docs/api/java/util/Optional.html).

You can use the functions `Just`, `None` or `Nillable` to instanciate a `Maybe` struct. The type parameter will be determined by the passed argument or by specifying it. For example:
```
maybeNilInteger := maybe.Nillable[int](nil)

nilInteger := maybe.None[int]()

someInteger := maybe.Just[int](1)
```

#### Usage
TODO
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
