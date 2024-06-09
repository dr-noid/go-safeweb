package safehttp
import "fmt"
var Coverage = make(map[string]bool)

func InitializeCoverageMap() {
	// intialize the map with all branches as not taken, false
	Coverage["flight/write-1"] = false
	Coverage["flight/write-1"] = false
	Coverage["flight/write-1"] = false

	Coverage["flight/writeError-1"] = false
	Coverage["flight/writeError-2"] = false
	Coverage["flight/writeError-3"] = false


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