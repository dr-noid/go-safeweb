package safehttp
import "fmt"

var Coverage = make(map[string]bool)

func InitializeCoverageMap() {
	// intialize the map with all branches as not taken, false

	// coverage for handler/StripPrefix

	Coverage["StripPrefix_1"] = false
	Coverage["StripPrefix_2"] = false
	Coverage["StripPrefix_3"] = false
	Coverage["StripPrefix_4"] = false

	// coverage for header/addCookie
	Coverage["Header_addCookie_1"] = false
	Coverage["Header_addCookie_2"] = false

}

func PrintCoverage() {
	for k, v := range Coverage {
		if v {
			fmt.Println("Branch", k, "was taken")
		} else {
			fmt.Println("Branch", k, "was not taken")
		}
	}
}