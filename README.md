# Go Regex Benchmark

This repo is a benchmark for various Golang's regular expressions library. Based on benchmark by Rustem Kamalov [here](https://github.com/karust/regex-benchmark).

Unlike the original repository, here I only focus on Go language without caring about its performance compared to the other languages.

## Input Text

Like the original, the [input text](./input-text.txt) is a concatenation of [Learn X in Y minutes](https://github.com/adambard/learnxinyminutes-docs) repository.

## Regex Patterns

Like the original, the used patterns are:

- Email: `[\w\.+-]+@[\w\.-]+\.[\w\.-]+`
- URI: `[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`
- IPv4: `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`

## Measure

Unlike the original, measuring is done without including regex compilation. So the measurement only focused on pattern matching.

## Used Packages

There are 7 regular expression packages that used in this benchmark:

- **Go**: the `regexp` package from Go's standard library.
- **Grafana**: the package from `github.com/grafana/regexp` by Bryan Boreham that improve the standard library with several optimations. Actively maintained at the time this benchmark is written.
- **Modernc**: the package from `modernc.org/regexp` that implements experimental DFA support.
- **Regexp2**: the package from `github.com/dlclark/regexp2` that ports regex engine from .NET frameworks.
- **Re2**: the package from `github.com/wasilibs/go-re2` that binds `google/re2` regex engine using WebAssembly or cgo.
- **Hyperscan**: the package from `github.com/flier/gohs` that binds Intel's `hyperscan` regex engine using cgo.
- **PCRE**: the package from `github.com/GRbit/go-pcre` that binds PCRE regex engine using cgo.

## Result

The benchmark was run on Linux with Intel i7-8550U with RAM 16 GB.

|  Package  | Use CGO | Email (ms) | URI (ms) | IP (ms) | Total (ms) |  Times |
| :-------: | :-----: | ---------: | -------: | ------: | ---------: | -----: |
|  Re2 CGO  |   Yes   |      11.31 |    13.29 |   10.47 |      35.08 | 25.85x |
| Hyperscan |   Yes   |      27.16 |    27.71 |    0.99 |      55.86 | 16.23x |
|   PCRE    |   Yes   |      24.44 |    24.73 |    8.85 |      58.02 | 15.63x |
| Re2 WASM  |         |      52.09 |    58.98 |   49.18 |     160.26 |  5.66x |
|    Go     |         |     250.39 |   255.16 |  401.13 |     906.68 |  1.00x |
|  Modernc  |         |     282.63 |   261.73 |  407.65 |     952.01 |  0.95x |
|  Grafana  |         |     278.22 |   268.56 |  408.43 |     955.21 |  0.95x |
|  Regexp2  |         |    2298.71 |  2046.71 |   75.75 |    4421.17 |  0.21x |

- Re2 with cgo is the fastest, while Regexp2 is the slowest.
- Regex engines that binded using cgo are way faster than the ones without cgo.
- Re2 with WebAssembly is the fastest regex engine that doesn't use cgo.
- Regex engine by Modernc that implements DFA is actually a bit slower than the engine from standard library.
