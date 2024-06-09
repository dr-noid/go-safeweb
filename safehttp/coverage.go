package safehttp
import "fmt"

var Coverage = make(map[string]bool)

func InitializeCoverageMap() {
	Coverage["cookie-1"] = false
	Coverage["cookie-2"] = false
	Coverage["cookie-3"] = false
	Coverage["Write-1"] = false
	Coverage["Write-2"] = false
	Coverage["Write-3"] = false

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