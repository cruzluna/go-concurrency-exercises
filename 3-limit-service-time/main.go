//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync/atomic"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

const KILL_TIME = 10

func (u *User) AddTime(sec int64) int64 {
	return atomic.AddInt64(&u.TimeUsed, sec)
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	done := make(chan bool)
	if u.IsPremium {
		process()
		return true
	}

	if atomic.LoadInt64(&u.TimeUsed) >= KILL_TIME {
		return false
	}

	go func() {
		process()
		done <- true
	}()

	tick := time.Tick(time.Second)

	for {
		select {
		case <-done:
			return true
		case <-tick:
			if i := u.AddTime(1); i >= KILL_TIME {
				return false
			}

		}
	}
}

func main() {
	RunMockServer()
}
