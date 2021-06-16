package lock

import "sync"

type RWMutex struct {
	rw   sync.RWMutex
	lock sync.Mutex
}

// Lock locks rw for writing.
// If the lock is already locked for reading or writing,
// Lock blocks until the lock is available.
func (rw *RWMutex) Lock() {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	rw.rw.Lock()
}

// Unlock unlocks rw for writing. It is a run-time error if rw is
// not locked for writing on entry to Unlock.
//
// As with Mutexes, a locked RWMutex is not associated with a particular
// goroutine. One goroutine may RLock (Lock) a RWMutex and then
// arrange for another goroutine to RUnlock (Unlock) it.
func (rw *RWMutex) Unlock() {
	rw.rw.Unlock()
}

// UnLockAndRLock 释放写锁并获取读锁
func (rw *RWMutex) UnLockAndRLock() {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	rw.rw.Unlock()
	rw.rw.RLock()
}

// RLock locks rw for reading.
//
// It should not be used for recursive read locking; a blocked Lock
// call excludes new readers from acquiring the lock. See the
// documentation on the RWMutex type.
func (rw *RWMutex) RLock() {
	rw.rw.RLock()
}

// RUnlock undoes a single RLock call;
// it does not affect other simultaneous readers.
// It is a run-time error if rw is not locked for reading
// on entry to RUnlock.
func (rw *RWMutex) RUnlock() {
	rw.rw.RUnlock()
}
