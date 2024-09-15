# Go Regex Benchmark

This repo is a benchmark for various Golang's regular expressions library. Based on benchmark by Rustem Kamalov [here][original-benchmark].

For its input, this benchmark use [text](./input-text.txt) from concatenation of [Learn X in Y minutes][x-in-y] repository. As the original benchmark pointed out, it's not the best representative text for real world benchmark. So, as additional comparison I also made [another benchmark][benchmark-web] that uses around 1000 web pages for input, so feel free to check it out.

## Table of Contents

- [How to Run](#how-to-run)
- [Measurement](#measurement)
- [Regex Patterns](#regex-patterns)
  - [Short Regex](#short-regex)
  - [Long Regex](#long-regex)
- [Used Packages](#used-packages)
  - [Native Go Packages](#native-go-packages)
  - [Regex with CGO Binding](#regex-with-cgo-binding)
  - [Regex with Web Assembly Binding](#regex-with-web-assembly-binding)
  - [Regex Compiler](#regex-compiler)
- [Result](#result)
  - [Short Regex](#short-regex-1)
  - [Long Regex](#long-regex-1)
- [License](#license)

## How to Run

If you are using GNU Make, you can simply run `make` in the root directory of this benchmark:

```
make clean-all    # clean the old build
make build-all    # rebuild the benchmark executable
make              # run the benchmark
```

## Measurement

Unlike the original, measurement is done without including regex compilation, so it's only focused on pattern matching. The measurement are done 10 times, then the smallest durations are used as the final durations.

## Regex Patterns

In this benchmark, there are 2 kinds of regexes that will be tested: short regex and long regex.

### Short Regex

For the short regex, we use 3 patterns from the original benchmark:

- Email: `[\w\.+-]+@[\w\.-]+\.[\w\.-]+`
- URI: `[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`
- IPv4: `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`

### Long Regex

For the long regex, we use pattern for detecting multilingual long date texts. Given the following patterns:

- day:

  ```
  [0-3]?[0-9]
  ```

- year:

  ```
  199[0-9]|20[0-3][0-9]
  ```

- month:

  ```
  January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara
  ```

The final pattern is defined as:

```
(?i)({month})\s({day})(?:st|nd|rd|th)?,?\s({year})|({day})(?:st|nd|rd|th|\.)?\s(?:of\s)?({month})[,.]?\s({year})
```

## Used Packages

There are 12 regular expressions that used in this benchmark:

- 5 are regex packages in native Go code.
- 3 are regex packages with CGO binding.
- 1 is regex packages with WASM binding.
- 3 are regex compiler to compile regular expressions to native Go code.

### Native Go Packages

1. **Go** is the `regexp` package from Go's standard library.
2. **Grafana** is the package from [`github.com/grafana/regexp@speedup`][grafana] by Bryan Boreham that improve the standard library with several optimations. Actively maintained at the time this benchmark is written.
3. **Modernc** is the package from [`modernc.org/regexp`][modernc] that implements experimental DFA support. However, it doesn't use DFA implementation from `google/re2` and instead uses its own implementation.
4. **Regexp2** is the package from [`github.com/dlclark/regexp2`][regexp2] that ports regex engine from .NET frameworks.
5. **Matloob** is the package from [`github.com/matloob/regexp`][matloob] that implements RE2 DFA. It's still not complete and its development is halted, but it's still benchmarked here to see how fast native Go's regex once DFA implemented.
6. **Code Search** is the package from [`github.com/google/codesearch/regexp`][codesearch]. It uses DFA algorithm with trigram indexing and was used by Google Code Search.

   Since it's used for search engine, currently its API only supports grep-like matching as in for every matching patterns it will returns the entire line instead of only returning the matching phrase.

   It's not thread safe, so every regex can only be used one goroutine at a time. And since its API only supports grep matching, currently it's not suitable for daily use. However it's still benchmarked here to glimpse the possible performance that can be achieved by native Go's regex engine in the future.

### Regex with CGO Binding

1. **RE2 CGO** is the package from [`github.com/wasilibs/go-re2`][go-re2] that binds [`google/re2`][google-re2] regex engine using cgo. Since Go's regex also use `google/re2` syntax, it can be used as drop in replacement for Go's native regex.
2. **Hyperscan** is the package from [`github.com/flier/gohs`][go-hyperscan] that binds Intel's [`hyperscan`][hyperscan] regex engine using cgo.
3. **PCRE** is the package from [`github.com/GRbit/go-pcre`][go-pcre] that binds PCRE regex engine using cgo.

### Regex with Web Assembly Binding

1. **RE2 WASM** is the package from [`github.com/wasilibs/go-re2`][go-re2] that binds [`google/re2`][google-re2] regex engine using Web Assembly. Like its cgo counterpart, it can be used as drop in replacement for Go's native regex.

### Regex Compiler

1. **re2go** is a regex compiler from [`re2c.org`][re2c] that compiles regular expressions into a native Go codes. Despite its name, it's not related with `google/re2` and uses its own lookahead TDFA algorithm.

   Since originally it only supports C, `re2go` use its own regex syntax which is not really compatible with Go's regex syntax. Fortunately it also supports [Flex regex syntax][flex] which is kinda similar with Go's syntax, so modifying the existing pattern is pretty easy.

   It's actively maintained and has been used in production and many open source projects, e.g. PHP and Ninja.

2. **Regexp2go** is a regex compiler from [`github.com/CAFxX/regexp2go`][regexp2go]. Despite its name, it's not related with `github.com/dlclark/regexp2` package mentioned above.

   It's similar in spirit to `re2go`, but aiming for compatibility with Go's regex syntax. At the time this benchmark written, this compiler hasn't been updated for 2 years and its documentation mentioned that it's not recommended to use in production.

   However for basic regex and with enough testing I reckon it should be good enough to use.

3. **Regexp2cg** is a regex compiler from [`github.com/dlclark/regexp2cg`][regexp2cg]. It's related with `github.com/dlclark/regexp2` package mentioned above.

   It will compile regular expressions into Go codes which can be used by the `dlclark/regexp2` package.

## Result

As reminder, as mentioned above this benchmark might not be the best representative text for real world. So make sure to check my [other benchmark][benchmark-web] that uses around 1000 web pages for its input.

Another note, re2go as a compiler has its own way to handle regular expressions. Thanks to this, in this benchmark its result for email and URI is a bit "unfair" compared to the other packages. This is because for those cases re2go templates in this benchmark is optimized with multiple rules to handle regex with quantifier groups. For more details, please check out [this issue][re2go-issue] where I discuss re2go performance with its maintainer.

With that out of the way, here is the benchmark result. The benchmark was run on my Linux PC with Intel i7-8550U and RAM 16 GB.

### Short Regex

|   Package   |     Type      | Email (ms) | URI (ms) | IP (ms) | Total (ms) |  Times |
| :---------: | :-----------: | ---------: | -------: | ------: | ---------: | -----: |
|   RE2 CGO   |      CGO      |       8.98 |    11.14 |    8.95 |      29.07 | 29.41x |
| Code Search | Native (Grep) |      11.61 |    12.04 |   11.10 |      34.75 | 24.61x |
|    re2go    |   Compiler    |      19.50 |    18.57 |    7.78 |      45.85 | 18.65x |
|  Hyperscan  |      CGO      |      25.07 |    24.65 |    0.83 |      50.55 | 16.91x |
|    PCRE     |      CGO      |      22.00 |    22.93 |    8.23 |      53.16 | 16.08x |
|   Matloob   |    Native     |      45.97 |    51.01 |   45.66 |     142.64 |  5.99x |
|  RE2 WASM   |     WASM      |      49.68 |    56.69 |   49.50 |     155.88 |  5.49x |
|  Regexp2Go  |   Compiler    |      67.32 |   517.95 |  165.11 |     750.38 |  1.14x |
|   Modernc   |    Native     |     233.35 |   232.46 |  377.78 |     843.59 |  1.01x |
|     Go      |    Native     |     235.59 |   238.13 |  381.30 |     855.02 |  1.00x |
|   Grafana   |    Native     |     253.65 |   235.67 |  378.36 |     867.68 |  0.99x |
|  Regexp2cg  |   Compiler    |    1914.50 |  1734.85 |   65.36 |    3714.71 |  0.23x |
|   Regexp2   |    Native     |    2127.81 |  1924.57 |   74.60 |    4126.98 |  0.21x |

Some interesting points:

- It's amazing to see how fast Code Search is. It's almost as fast as RE2 with cgo.
- Matloob is very fast for short regexes. It's even faster than RE2 WASM. Unfortunately its development is halted, so it's not suitable for daily use.
- The optimized re2go template generates a very fast native Go code. In the case of IP pattern, it's as fast as RE2 with cgo. However since some Go's regex syntaxes are not supported by re2go, there are needs to modify the regex patterns before using it.
- For code without cgo, has full regex compatibility and actively maintained, RE2 WASM has the best performance albeit a bit slower than re2go.

### Long Regex

|   Package   |     Type      | Long Date (ms) |     Times |
| :---------: | :-----------: | -------------: | --------: |
|  Hyperscan  |      CGO      |           1.13 | 10963.19x |
|   RE2 CGO   |      CGO      |          12.67 |   978.67x |
| Code Search | Native (Grep) |          42.46 |   291.97x |
|    re2go    |   Compiler    |          42.62 |   290.93x |
|  RE2 WASM   |     WASM      |          59.46 |   208.51x |
|    PCRE     |      CGO      |         162.66 |    76.22x |
|   Grafana   |    Native     |        3221.90 |     3.85x |
|   Regexp2   |    Native     |        4113.02 |     3.01x |
|  Regexp2cg  |   Compiler    |        5004.37 |     2.48x |
|  Regexp2Go  |   Compiler    |        7140.02 |     1.74x |
|   Modernc   |    Native     |       12228.11 |     1.01x |
|     Go      |    Native     |       12398.38 |     1.00x |
|   Matloob   |    Native     |       12938.24 |     0.96x |

Some interesting points:

- Hyperscan is really fast at handling long regex pattern. It's even faster than when it's used for short regex.
- Eventhough Matloob is very fast for short regexes, it's still slow for long regex.
- PCRE's performance become a lot slower for long regex compared to RE2 and Hyperscan. It's even slower than RE2 WASM.
- For native Go code, Grafana is pretty fast at handling long regex pattern.
- For code without cgo, regex that compiled by re2go has the best performance. It has similar performance as Code Search.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance (since Code Search currently can't be used as daily regex engine).

## License

Like the original benchmark, this benchmark is also released under MIT license.

[original-benchmark]: https://github.com/karust/regex-benchmark
[benchmark-web]: https://github.com/RadhiFadlillah/go-regex-benchmark-web
[x-in-y]: https://github.com/adambard/learnxinyminutes-docs
[grafana]: https://github.com/grafana/regexp/tree/speedup?tab=readme-ov-file
[modernc]: https://gitlab.com/cznic/regexp
[matloob]: https://github.com/matloob/regexp
[regexp2]: https://github.com/dlclark/regexp2
[codesearch]: https://github.com/google/codesearch
[go-re2]: https://github.com/wasilibs/go-re2
[google-re2]: https://github.com/google/re2
[go-hyperscan]: https://github.com/flier/gohs
[hyperscan]: https://www.intel.com/content/www/us/en/developer/articles/technical/introduction-to-hyperscan.html
[go-pcre]: https://github.com/GRbit/go-pcre
[re2c]: https://re2c.org/manual/manual_go.html
[flex]: https://github.com/westes/flex
[regexp2go]: https://github.com/CAFxX/regexp2go
[regexp2cg]: https://github.com/dlclark/regexp2cg
[re2go-issue]: https://github.com/skvadrik/re2c/issues/487
