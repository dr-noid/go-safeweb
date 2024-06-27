package safehttp

import "fmt"

var Coverage = make(map[string]bool)

func InitializeCoverageMap() {
	// intialize the map with all branches as not taken, false
	Coverage["flight/write-1"] = false
	Coverage["flight/write-2"] = false
	Coverage["flight/write-3"] = false

	Coverage["flight/writeError-1"] = false
	Coverage["flight/writeError-2"] = false
	Coverage["flight/writeError-3"] = false

	Coverage["echo-1"] = false
	Coverage["echo-2"] = false
	Coverage["echo-3"] = false

	Coverage["uptime-1"] = false
	Coverage["uptime-2"] = false
	Coverage["uptime-3"] = false

	Coverage["cookie-1"] = false
	Coverage["cookie-2"] = false
	Coverage["cookie-3"] = false
	Coverage["Write-1"] = false
	Coverage["Write-2"] = false
	Coverage["Write-3"] = false

	Coverage["StripPrefix_1"] = false
	Coverage["StripPrefix_2"] = false
	Coverage["StripPrefix_3"] = false
	Coverage["StripPrefix_4"] = false

	Coverage["Header_addCookie_1"] = false
	Coverage["Header_addCookie_2"] = false
}

func PrintCoverage() {
	for key, value := range Coverage {
		if value {
			fmt.Println("Branch", key, "was taken")
		} else {
			fmt.Println("Branch", key, "was not taken")
		}
	}
}
