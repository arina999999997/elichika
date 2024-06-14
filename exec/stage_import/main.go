package main

import (
	"elichika/client"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type SimpleLiveNote struct {
	// Id                  int32 `json:"id"`
	CallTime     int32 `json:"call_time"`
	NoteType     int32 `json:"note_type"`
	NotePosition int32 `json:"note_position"`
	// GimmickId           int32 `json:"gimmick_id"`
	NoteAction int32 `json:"note_action"`
	WaveId     int32 `json:"wave_id"`
	// NoteRandomDropColor int32 `json:"note_random_drop_color"`
	// AutoJudgeType       int32 `json:"auto_judge_type"` // relevant because some note actually have different auto judge type (to force fully trigger stuff?)
}

type SimpleLiveStage struct {
	LiveDifficultyId int32                    `json:"live_difficulty_id"`
	LiveNotes        []SimpleLiveNote         `json:"live_notes"`
	LiveWaveSettings []client.LiveWaveSetting `json:"live_wave_settings"`

	Original *int32 `json:"original"`
}

type SimpleLiveStartResponse struct {
	Live struct {
		LiveStage SimpleLiveStage `json:"live_stage"`
	} `json:"live"`
}

type FullLiveStage = client.LiveStage
type FullLiveStartResponse struct {
	Live struct {
		LiveStage FullLiveStage `json:"live_stage"`
	} `json:"live"`
}

var fullStages map[int32]*FullLiveStage
var stages map[int32]*SimpleLiveStage
var stageOrigins map[int32]string

func isSame(a, b *SimpleLiveStage) bool {
	if len(a.LiveNotes) != len(b.LiveNotes) {
		return false
	}
	for i := range a.LiveNotes {
		if a.LiveNotes[i] != b.LiveNotes[i] {
			return false
		}
	}
	for i := range a.LiveWaveSettings {
		if a.LiveWaveSettings[i] != b.LiveWaveSettings[i] {
			return false
		}
	}
	return true
}

func parseStage(stage SimpleLiveStage, fullStage FullLiveStage, file string) {
	currentStage, exists := stages[stage.LiveDifficultyId]
	if !exists {
		stages[stage.LiveDifficultyId] = &stage
		fullStages[fullStage.LiveDifficultyId] = &fullStage
		stageOrigins[stage.LiveDifficultyId] = file
		return
	}
	// must be the same
	if len(stage.LiveNotes) != len(currentStage.LiveNotes) {
		panic(fmt.Sprint("Error: Inconsistent stage: ", stage.LiveDifficultyId, "\nDifferent Live Notes Lengths: ",
			len(currentStage.LiveNotes), " vs ", len(stage.LiveNotes), "\nFiles: ",
			stageOrigins[stage.LiveDifficultyId], " vs ", file))
	}
	for i := range stage.LiveNotes {
		if stage.LiveNotes[i] != currentStage.LiveNotes[i] {
			panic(fmt.Sprint("Error: Inconsistent stage: ", stage.LiveDifficultyId, "\nDifferent Notes: ",
				currentStage.LiveNotes[i], " vs ", stage.LiveNotes[i], "\nFiles: ",
				stageOrigins[stage.LiveDifficultyId], " vs ", file))
		}
	}
	for i := range stage.LiveWaveSettings {
		if stage.LiveWaveSettings[i] != currentStage.LiveWaveSettings[i] {
			panic(fmt.Sprint("Error: Inconsistent stage: ", stage.LiveDifficultyId, "\nDifferent Waves: ",
				currentStage.LiveWaveSettings[i], " vs ", stage.LiveWaveSettings[i], "\nFiles: ",
				stageOrigins[stage.LiveDifficultyId], " vs ", file))
		}
	}
}

func parseFile(file string) {
	text := utils.ReadAllText(file)
	var liveStage SimpleLiveStage
	var fullLiveStage FullLiveStage
	err := json.Unmarshal([]byte(text), &liveStage)
	if (err == nil) && (liveStage.LiveDifficultyId != 0) {
		err = json.Unmarshal([]byte(text), &fullLiveStage)
		utils.CheckErr(err)
		parseStage(liveStage, fullLiveStage, file)
		return
	}
	var live SimpleLiveStartResponse
	var fullLive FullLiveStartResponse
	err = json.Unmarshal([]byte(text), &live)
	if (err == nil) && (live.Live.LiveStage.LiveDifficultyId != 0) {
		err = json.Unmarshal([]byte(text), &fullLive)
		utils.CheckErr(err)
		parseStage(live.Live.LiveStage, fullLive.Live.LiveStage, file)
		return
	}
	fmt.Println("Warning: skipping file with unexpected format: ", file)

}

var files map[string]bool

func handleFile(path string) {
	files[path] = true
}

func handleDirectory(dir string) {
	if strings.HasSuffix(dir, ".json") {
		handleFile(dir)
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, file := range files {
		// fmt.Println(file.Name())
		handleDirectory(dir + "/" + file.Name())
	}

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: stage_import <one or more path/to/directory or path/to/file.json>.\nNote that this will recursively traverse everything.")
		return
	}
	files = make(map[string]bool)
	for i, path := range os.Args {
		if i == 0 {
			continue
		}
		handleDirectory(path)
	}
	fmt.Println("Discovered", len(files), "relevant files")
	stages = make(map[int32]*SimpleLiveStage)
	fullStages = make(map[int32]*FullLiveStage)
	stageOrigins = make(map[int32]string)
	for file := range files {
		parseFile(file)
	}
	fmt.Println("Found", len(stages), "different stages.\nRemoving similar maps")
	ids := []int32{}
	for id := range stages {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	for _, id := range ids {
		for _, other := range ids {
			if other == id {
				break
			}
			if isSame(stages[id], stages[other]) {
				fmt.Println("Found unoriginal map: ", id, " -> ", other)
				stages[id].Original = new(int32)
				*stages[id].Original = other
				break
			}
		}
		if stages[id].Original == nil {
			fmt.Println("Original map: ", id)
		}
	}

	for _, id := range ids {
		if stages[id].Original != nil {
			stages[id].LiveNotes = nil
		}
		output, err := json.Marshal(stages[id])
		utils.CheckErr(err)
		err = os.WriteFile(fmt.Sprint("assets/simple_stages/", id, ".json"), output, 0777)
		utils.CheckErr(err)
		output, err = json.Marshal(fullStages[id])
		utils.CheckErr(err)
		err = os.WriteFile(fmt.Sprint("assets/full_stages/", id, ".json"), output, 0777)
		utils.CheckErr(err)

	}

}
