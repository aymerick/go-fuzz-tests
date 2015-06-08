# go-fuzz-tests

Various tests using [go-fuzz](https://github.com/dvyukov/go-fuzz)

Example:

    $ go-fuzz-build github.com/aymerick/go-fuzz-tests/raymond
    $ go-fuzz -bin=./raymond-fuzz.zip -corpus=raymond/corpus -workdir=output
