package safehttp

var Coverage = make(map[string]bool)

func InitializeCoverageMap() {
	// intialize the map with all branches as not taken, false
	Coverage["echo-1"] = false
	Coverage["echo-2"] = false
	Coverage["echo-3"] = false

	Coverage["uptime-1"] = false
	Coverage["uptime-2"] = false
	Coverage["uptime-3"] = false
}

func PrintCoverage() {
	for key, value := range Coverage {
		if value {
			println("Branch", key, "was taken")
		} else {
			println("Branch", key, "was not taken")
		}
	}
}
