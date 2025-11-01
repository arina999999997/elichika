package asset_manager

import (
	"elichika/log"
	"elichika/utils"

	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
)

func sortAndUnique(data []string) []string {
	temp := data
	data = nil
	sort.Strings(temp)
	size := 0
	for _, assetPath := range temp {
		if (size == 0) || (assetPath != data[size-1]) {
			data = append(data, assetPath)
			size++
		}
	}
	return data
}

// perform a cdn check on a list of pack to make sure they are all on the public cdn
// this means that the pack itself must be on the cdn or the metapack must be
// only do this for newly added pack
// for best effect, download and compare with files in /static
func performPackCheck(packs []string) string {
	publicCdn := "https://llsifas.imsofucking.gay/static/"
	localCdn := "static/"
	packs = sortAndUnique(packs)
	n := len(packs)
	log.Printf("Performing cdn pack check on %d packs\n", n)
	report := ""
	for i, pack := range packs {
		log.Printf("\tChecking pack %s (%d / %d)\n", pack, i+1, n)
		res, err := http.Get(publicCdn + pack)
		utils.CheckErr(err)
		defer res.Body.Close()
		if res.StatusCode != 200 {
			report += fmt.Sprintf("Failed to download pack from server: %s (status code: %d)\n", pack, res.StatusCode)
			continue
		}
		body, err := io.ReadAll(res.Body)
		utils.CheckErr(err)
		// compare to local file
		local, err := os.ReadFile(localCdn + pack)
		utils.CheckErr(err)
		if len(body) != len(local) {
			report += fmt.Sprintf("File difference for pack: %s (size isn't equal)\n", pack)
		}
		for i := range body {
			if body[i] != local[i] {
				report += fmt.Sprintf("File difference for pack: %s (byte %d isn't equal)\n", pack, i)
				break
			}
		}
	}
	return report
}
