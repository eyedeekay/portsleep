package portsleep

import (
	"fmt"
	"net"
	"strconv"
	time "time"
)

// DEFAULT_TIMEOUT is the default number of times to loop before deciding
// that the port is not available. By default, it is set to 300 seconds.
var DEFAULT_TIMEOUT = 300

// LOOP_TIMEOUT Set to true to loop forever instead of timing out
var LOOP_FOREVER = false

// SleepLoop until localhost:port is available, or timeout after DEFAULT_TIMEOUT seconds
// panic when the timeout is reached.
func SleepLoop(port int) {
	if !PortSleep(port) {
		err := fmt.Errorf("port %d is not available", port)
		panic(err)
	}
}

// SleepHostLoop until host:port is available, or timeout after DEFAULT_TIMEOUT seconds
// panic when the timeout is reached.
func SleepHostLoop(host string, port int) {
	if !SleepUntilAvailable("tcp", host, port, DEFAULT_TIMEOUT) {
		err := fmt.Errorf("port %d on host %s is not available", port, host)
		panic(err)
	}
}

// SleepAddrLoop until addr is available, or timeout after DEFAULT_TIMEOUT seconds
// panic when the timeout is reached.
func SleepAddrLoop(addr string) {
	host, sport, err := net.SplitHostPort(addr)
	if err != nil {
		panic(err)
	}
	port, _ := strconv.Atoi(sport)
	if !SleepUntilAvailable("tcp", host, port, DEFAULT_TIMEOUT) {
		err := fmt.Errorf("addr %s is not available", addr)
		panic(err)
	}
}

// PortSleep until port is available, or timeout after DEFAULT_TIMEOUT seconds.
// return true if the port is available, false if the timeout is reached
func PortSleep(port int) bool {
	return SleepUntilAvailable("tcp", "localhost", port, DEFAULT_TIMEOUT)
}

// HostPortSleep until host:port is available, or timeout after DEFAULT_TIMEOUT seconds.
// return true if the port is available, false if the timeout is reached
func HostPortSleep(host string, port int) bool {
	return SleepUntilAvailable("tcp", host, port, DEFAULT_TIMEOUT)
}

// SleepUntilAvailable sleeps until scheme@host:port is available, or timeout
// after timeout seconds.
// return true if the port is available, false if the timeout is reached
func SleepUntilAvailable(scheme, host string, port int, timeout int) bool {
	x := 0
	for x < timeout {
		addr := net.JoinHostPort(host, fmt.Sprintf("%d", port))
		conn, err := net.Listen(scheme, addr)
		if err == nil {
			conn.Close()
			return true
		}
		time.Sleep(time.Duration(1) * time.Second)
		if x%5 == 0 {
			fmt.Println("Waiting for " + scheme + "://" + addr + " to become available")
		}
		if timeout > 0 {
			x = increment(x)
		}
	}
	return false
}

func increment(count int) int {
	if LOOP_FOREVER {
		return count
	}
	count++
	return count
}
