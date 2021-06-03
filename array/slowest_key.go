package array

func slowestKey(releaseTimes []int, keysPressed string) byte {
	maxDuration := releaseTimes[0]
	slowest := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		duration := releaseTimes[i] - releaseTimes[i-1]
		if duration > maxDuration {
			maxDuration = duration
			slowest = keysPressed[i]
		} else if maxDuration == duration && keysPressed[i] > slowest {
			maxDuration = duration
			slowest = keysPressed[i]
		}

	}

	return slowest
}
