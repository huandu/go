package runtime

// GetGoroutineId returns current goroutine id.
func GetGoroutineId() int64 {
        gp := getg()
        return gp.goid
}

