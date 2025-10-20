package asset_manager

// generate a report on the data
import (
	"fmt"
)

func GenerateReport() {
	for _, asset := range Assets {
		fmt.Printf("%s\n", asset.Info())
	}
}
