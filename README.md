# Nudge

Keeps your computer awake.

A simple Go program inspired by [Caffeine][caffeine] to occasionally press a key
to keep your screen on/unlocked.

```sh
go get github.com/travis-g/nudge
```

The automated keypress functionality is provided by [go-vgo/robotgo][robotgo]
([Apache 2.0/MIT][robotgo-license]).

## Usage

```console
$ nudge --help
Usage of nudge:
  -disable
        start disabled
  -interval duration
        interval between nudges (default 59s)
  -key string
        key to press (default "f16")
```

If <kbd>F16</kbd> gives trouble try `-key shift`.

[caffeine]: http://www.zhornsoftware.co.uk/caffeine/
[robotgo]: https://github.com/go-vgo/robotgo
[robotgo-license]: https://github.com/go-vgo/robotgo#license
