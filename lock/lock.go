package lock

import "sync"

type RWMutex struct {
	sync.RWMutex
	lock sync.Mutex
}

// Lock locks rw for writing.
// If the lock is already locked for reading or writing,
// Lock blocks until the lock is available.
func (rw *RWMutex) Lock() {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	rw.RWMutex.Lock()
}

// UnLockAndRLock 释放写锁并获取读锁
func (rw *RWMutex) UnLockAndRLock() {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	rw.Unlock()
	rw.RLock()
}
