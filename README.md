# Go Regex Benchmark

This repo is a benchmark for various Golang's regular expressions library. Based on benchmark by Rustem Kamalov [here][original-benchmark].

Unlike the original repository, here I only focus on Go language without caring about its performance compared to the other languages.

## Table of Contents

- [Input Text](#input-text)
- [Regex Patterns](#regex-patterns)
  - [Short Regex](#short-regex)
  - [Long Regex](#long-regex)
- [Used Packages](#used-packages)
- [Measure](#measure)
- [Result](#result)
  - [Short Regex](#short-regex-1)
  - [Long Regex](#long-regex-1)

## Input Text

Like the original, the [input text](./input-text.txt) is a concatenation of [Learn X in Y minutes][x-in-y] repository.

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

There are 8 regular expression packages that used in this benchmark:

- **Go**: the `regexp` package from Go's standard library.
- **Grafana**: the package from [`github.com/grafana/regexp@speedup`][grafana] by Bryan Boreham that improve the standard library with several optimations. Actively maintained at the time this benchmark is written.
- **Modernc**: the package from [`modernc.org/regexp`][modernc] that implements experimental DFA support.
- **Regexp2**: the package from [`github.com/dlclark/regexp2`][regexp2] that ports regex engine from .NET frameworks.
- **RE2**: the package from [`github.com/wasilibs/go-re2`][go-re2] that binds [`google/re2`][google-re2] regex engine using WebAssembly or cgo.
- **Hyperscan**: the package from [`github.com/flier/gohs`][go-hyperscan] that binds Intel's [`hyperscan`][hyperscan] regex engine using cgo.
- **PCRE**: the package from [`github.com/GRbit/go-pcre`][go-pcre] that binds PCRE regex engine using cgo.
- **re2go**: unlike the others, this one is not using any regex package. Instead, it use [`re2go`][re2go] to compile regular expressions to Go code using TDFA algorithm.
- **Code Search**: the package from [`github.com/google/codesearch/regexp`][codesearch]. Rather than a regex engine, it's more like a grep engine for searching huge amount of codes. As a grep engine, it's not really comparable to the other packages. It's not thread safe as well. However it's still included here to see what Go actually can do.

## Measure

Unlike the original, measuring is done without including regex compilation. So the measurement only focused on pattern matching.

The measurement are done 10 times, then the smallest durations are used as the final durations.

## Result

The benchmark was run on Linux with Intel i7-8550U with RAM 16 GB.

### Short Regex

|   Package   | Use CGO | Email (ms) | URI (ms) | IP (ms) | Total (ms) |  Times |
| :---------: | :-----: | ---------: | -------: | ------: | ---------: | -----: |
|   RE2 CGO   |   Yes   |       9.51 |    11.93 |    9.57 |      31.01 | 29.05x |
| Code Search |         |      12.16 |    13.33 |   12.34 |      37.83 | 23.81x |
|  Hyperscan  |   Yes   |      26.65 |    25.28 |    0.86 |      52.79 | 17.06x |
|    PCRE     |   Yes   |      22.59 |    23.72 |    8.59 |      54.89 | 16.41x |
|    re2go    |         |      75.26 |    45.69 |   22.46 |     143.40 |  6.28x |
|  RE2 WASM   |         |      52.57 |    58.11 |   49.18 |     159.86 |  5.64x |
| Go std lib  |         |     249.04 |   254.24 |  397.59 |     900.88 |  1.00x |
|   Modernc   |         |     260.28 |   252.38 |  391.30 |     903.95 |  1.00x |
|   Grafana   |         |     278.33 |   264.12 |  423.10 |     965.54 |  0.93x |
|   Regexp2   |         |    2259.23 |  2040.55 |   77.92 |    4377.71 |  0.21x |

- For short regex, RE2 with cgo is the fastest while Regexp2 is the slowest.
- Regex engines that utilize cgo are way faster than the ones without cgo (excluding Code Search since that one is not a regex engine).
- For code without cgo, re2go has the best performance. However some Go's regex syntaxes are not supported by re2go, so there are needs to modify the regex patterns.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance albeit a bit slower than re2go.
- Regex engine by Modernc that implements DFA is actually a bit slower than the engine from standard library.
- It's amazing to see how fast Code Search is. It's almost as fast as RE2 with cgo.

### Long Regex

|   Package   | Use CGO | Long Date (ms) |     Times |
| :---------: | :-----: | -------------: | --------: |
|  Hyperscan  |   Yes   |           1.13 | 11395.46x |
|   RE2 CGO   |   Yes   |          13.21 |   972.69x |
|    re2go    |         |          48.11 |   267.09x |
| Code Search |         |          48.95 |   262.49x |
|  RE2 WASM   |         |          63.38 |   202.71x |
|    PCRE     |   Yes   |         174.54 |    73.61x |
|   Grafana   |         |        3364.24 |     3.82x |
|   Regexp2   |         |        4330.83 |     2.97x |
|   Modernc   |         |       12411.65 |     1.04x |
| Go std lib  |         |       12848.43 |     1.00x |

- For long regex, Hyperscan with cgo is the fastest while Modernc is the slowest.
- For code without cgo, re2go has the best performance. It's as fast as Code Search.
- For code without cgo but with full regex compatibility, RE2 WASM has the best performance (since Code Search is not a regex engine).
- It's interesting to see how fast Hyperscan handling long regex pattern. It's even faster than when it's used for short regex.
- It's interesting to see PCRE's performance become slower for long regex compared to RE2 and Hyperscan. It's even slower than RE2 WASM.
- It's also interesting to see how fast Grafana compared to the Go's standard library.

[original-benchmark]: https://github.com/karust/regex-benchmark
[x-in-y]: https://github.com/adambard/learnxinyminutes-docs
[grafana]: https://github.com/grafana/regexp/tree/speedup?tab=readme-ov-file
[modernc]: https://gitlab.com/cznic/regexp
[regexp2]: https://github.com/dlclark/regexp2
[go-re2]: https://github.com/wasilibs/go-re2
[google-re2]: https://github.com/google/re2
[go-hyperscan]: https://github.com/flier/gohs
[hyperscan]: https://www.intel.com/content/www/us/en/developer/articles/technical/introduction-to-hyperscan.html
[go-pcre]: https://github.com/GRbit/go-pcre
[re2go]: https://re2c.org/manual/manual_go.html
[codesearch]: https://github.com/google/codesearch
