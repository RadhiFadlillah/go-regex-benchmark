This repository is a simple benchmark to compare performance between regexp in standard library and [github.com/dlclark/regexp2][0].

```
STANDARD LIBRARY
EMAIL: GOT 92 MATCHES IN 270ms
URI  : GOT 5301 MATCHES IN 294ms
IP   : GOT 5 MATCHES IN 465ms

REGEXP2 LIBRARY
EMAIL: GOT 92 MATCHES IN 2742ms
URI  : GOT 5301 MATCHES IN 2246ms
IP   : GOT 5 MATCHES IN 113ms
```

[0]: https://github.com/dlclark/regexp2
