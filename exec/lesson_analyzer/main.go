package main

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type JsonDumpFile struct {
	Path       string
	Data       []byte
	Time       int64
	IsResponse bool
}

func (jdf *JsonDumpFile) GetActualData() json.RawMessage {
	items := []json.RawMessage{}
	err := json.Unmarshal(jdf.Data, &items)
	utils.CheckErr(err)
	return items[len(items)-2]
}

var files []JsonDumpFile

func handleFile(path string) {
	file := JsonDumpFile{
		Path: path,
	}
	// executeLesson%3fp=a&mv=2d61e7b4e89961c7&id=33&u=773494794&t=1688062407673&l=en(1)
	tStart := strings.Index(path, "&t=") + 3
	for i := range path {
		if i < tStart {
			continue
		}

		if (path[i] > '9') || (path[i] < '0') {
			break
		}
		file.Time = file.Time*10 + int64(path[i]-'0')
	}
	file.IsResponse = strings.HasSuffix(path, "(1)")
	file.Data = []byte(utils.ReadAllText(path))
	files = append(files, file)
}

func handleDirectory(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		// not a dir
		handleFile(dir)
		return
	}
	for _, file := range files {
		// fmt.Println(file.Name())
		handleDirectory(dir + "/" + file.Name())
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: lesson_analyser path/to/lesson/dump/directory\nNote that this will recursively traverse everything.")
		return
	}

	handleDirectory(os.Args[1])
	sort.Slice(files, func(i, j int) bool {
		if files[i].Time != files[j].Time {
			return files[i].Time < files[j].Time
		}
		return files[j].IsResponse
	})

	n := len(files)
	if n%5 != 0 {
		panic("wrong number of file. expected /executeLesson request response, then /resultLesson response, then /skillEditResult request and response!")
	}
	for i := 0; i < n; i += 5 {
		var executeReq request.ExecuteLessonRequest
		var executeResp response.ExecuteLessonResponse
		var resultResp response.LessonResultResponse
		var err error
		err = json.Unmarshal(files[i].GetActualData(), &executeReq)
		utils.CheckErr(err)
		err = json.Unmarshal(files[i+1].GetActualData(), &executeResp)
		utils.CheckErr(err)
		err = json.Unmarshal(files[i+2].GetActualData(), &resultResp)
		utils.CheckErr(err)

		smallDropCount := []int{}
		isSmallDrop := true

		currentRun := 0
		dropCount := []int{}

		currentMegaphoneRun := 0
		megaphoneDropCount := []int{}

		// fmt.Println(files[i].Path)
		// fmt.Println(files[i+1].Path)
		// fmt.Println(files[i+2].Path)
		// fmt.Println(executeReq)
		for _, item := range resultResp.DropItemList.Slice {
			if item.IsSubscription {
				if currentRun != 0 {
					dropCount = append(dropCount, currentRun)
					megaphoneDropCount = append(megaphoneDropCount, currentMegaphoneRun)
					currentMegaphoneRun = 0
					currentRun = 0
					isSmallDrop = true
				}
			} else {
				if item.ContentType == enum.ContentTypeMemberGuildSupport {
					currentMegaphoneRun++
					if isSmallDrop == true {
						isSmallDrop = false
						smallDropCount = append(smallDropCount, currentRun)
					}
				} else {
					if currentMegaphoneRun > 0 {
						megaphoneDropCount = append(megaphoneDropCount, currentMegaphoneRun)
						currentMegaphoneRun = 0
					}
				}
				currentRun++
			}
		}
		// for i := range dropCount {
		// 	fmt.Println(dropCount[i],"\t", megaphoneDropCount[i])
		// }
		for _, c := range megaphoneDropCount {
			fmt.Println(c)
		}
	}
}
