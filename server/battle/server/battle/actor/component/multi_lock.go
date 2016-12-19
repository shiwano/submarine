package component

// MultiLock represents a lock that is locked in duplicate.
type MultiLock struct {
	count int
}

// IsLocked determines whether the lock is locked.
func (l *MultiLock) IsLocked() bool {
	return l.count > 0
}

// Lock the lock.
func (l *MultiLock) Lock() {
	l.count++
}

// Unlock the lock.
func (l *MultiLock) Unlock() {
	l.count--
	if l.count < 0 {
		l.count = 0
	}
}
