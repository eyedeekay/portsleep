# portsleep

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/.)
[![Go Report Card](https://goreportcard.com/badge/.)](https://goreportcard.com/report/.)

## Variables

DEFAULT_TIMEOUT is the default number of times to loop before deciding
that the port is not available. By default, it is set to 300 seconds.

```golang
var DEFAULT_TIMEOUT = 300
```

LOOP_TIMEOUT Set to true to loop forever instead of timing out

```golang
var LOOP_FOREVER = false
```

## Functions

### func [HostPortSleep](/sleep.go#L57)

`func HostPortSleep(host string, port int) bool`

HostPortSleep until host:port is available, or timeout after DEFAULT_TIMEOUT seconds.
return true if the port is available, false if the timeout is reached

### func [PortSleep](/sleep.go#L51)

`func PortSleep(port int) bool`

PortSleep until port is available, or timeout after DEFAULT_TIMEOUT seconds.
return true if the port is available, false if the timeout is reached

### func [SleepAddrLoop](/sleep.go#L37)

`func SleepAddrLoop(addr string)`

SleepAddrLoop until addr is available, or timeout after DEFAULT_TIMEOUT seconds
panic when the timeout is reached.

### func [SleepHostLoop](/sleep.go#L28)

`func SleepHostLoop(host string, port int)`

SleepHostLoop until host:port is available, or timeout after DEFAULT_TIMEOUT seconds
panic when the timeout is reached.

### func [SleepLoop](/sleep.go#L19)

`func SleepLoop(port int)`

SleepLoop until localhost:port is available, or timeout after DEFAULT_TIMEOUT seconds
panic when the timeout is reached.

### func [SleepUntilAvailable](/sleep.go#L64)

`func SleepUntilAvailable(scheme, host string, port int, timeout int) bool`

SleepUntilAvailable sleeps until scheme@host:port is available, or timeout
after timeout seconds.
return true if the port is available, false if the timeout is reached

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
