// moegirl.icu
// https://moegirl.icu/LoveLive!%E5%AD%A6%E5%9B%AD%E5%81%B6%E5%83%8F%E7%A5%ADALL_STARS/%E9%97%AA%E4%BA%AE%E7%81%AF%E6%B5%B7%E4%B9%8B%E5%A4%9C
//
// some of the event have complete list of reward, we can then use this to parse it

package event_marathon_dev

import (
	"elichika/utils"

	"fmt"
	"strconv"
	"strings"
)

var (
	// map from string to content_type,content_id
	MoegirlItem      = map[string]string{}
	MoegirlUnit      = map[string]bool{}
	MoegirlDataLines = []string{}
)

func init() {
	// skip ticket
	MoegirlItem["跳过券"] = "28,16001"
	// training ticket
	MoegirlItem["合宿券"] = "16,2200"

	// red macaron
	MoegirlItem["红色马卡龙★1"] = "12,1900"
	MoegirlItem["红色马卡龙★2"] = "12,1901"
	MoegirlItem["红色马卡龙★3"] = "12,1902"
	// blue macaron
	MoegirlItem["蓝色马卡龙★1"] = "12,1910"
	MoegirlItem["蓝色马卡龙★2"] = "12,1911"
	MoegirlItem["蓝色马卡龙★3"] = "12,1912"
	// green macaron
	MoegirlItem["绿色马卡龙★1"] = "12,1920"
	MoegirlItem["绿色马卡龙★2"] = "12,1921"
	MoegirlItem["绿色马卡龙★3"] = "12,1922"
	// yellow macaron
	MoegirlItem["黄色马卡龙★1"] = "12,1930"
	MoegirlItem["黄色马卡龙★2"] = "12,1931"
	MoegirlItem["黄色马卡龙★3"] = "12,1932"
	// purple macaron
	MoegirlItem["紫色马卡龙★1"] = "12,1940"
	MoegirlItem["紫色马卡龙★2"] = "12,1941"
	MoegirlItem["紫色马卡龙★3"] = "12,1942"
	// pink macaron
	MoegirlItem["桃色马卡龙★1"] = "12,1950"
	MoegirlItem["桃色马卡龙★2"] = "12,1951"
	MoegirlItem["桃色马卡龙★3"] = "12,1952"
	// silver macaron
	MoegirlItem["银色马卡龙★1"] = "12,1960"
	MoegirlItem["银色马卡龙★2"] = "12,1961"
	MoegirlItem["银色马卡龙★3"] = "12,1962"
	// gold macaron
	MoegirlItem["金色马卡龙★1"] = "12,1970"
	MoegirlItem["金色马卡龙★2"] = "12,1971"
	MoegirlItem["金色马卡龙★3"] = "12,1972"

	// Voltage type tree
	MoegirlItem["热度类型种子"] = "12,13101"
	MoegirlItem["热度类型幼苗"] = "12,13102"
	MoegirlItem["热度类型花苞"] = "12,13103"
	// SP type tree
	MoegirlItem["SP类型种子"] = "12,13201"
	MoegirlItem["SP类型幼苗"] = "12,13202"
	MoegirlItem["SP类型花苞"] = "12,13203"
	// Guard type tree
	MoegirlItem["防御类型种子"] = "12,13301"
	MoegirlItem["防御类型幼苗"] = "12,13302"
	MoegirlItem["防御类型花苞"] = "12,13303"
	// Skill type tree
	MoegirlItem["技能类型种子"] = "12,13401"
	MoegirlItem["技能类型幼苗"] = "12,13402"
	MoegirlItem["技能类型花苞"] = "12,13403"
	// Voltage type Book
	MoegirlItem["热度类型入门书"] = "12,12101"
	MoegirlItem["热度类型中阶书"] = "12,12102"
	MoegirlItem["热度类型高阶书"] = "12,12103"
	// SP type book
	MoegirlItem["SP类型入门书"] = "12,12201"
	MoegirlItem["SP类型中阶书"] = "12,12202"
	MoegirlItem["SP类型高阶书"] = "12,12203"
	// Guard type book
	MoegirlItem["防御类型入门书"] = "12,12301"
	MoegirlItem["防御类型中阶书"] = "12,12302"
	MoegirlItem["防御类型高阶书"] = "12,12303"
	// Skill type book
	MoegirlItem["技能类型入门书"] = "12,12401"
	MoegirlItem["技能类型中阶书"] = "12,12402"
	MoegirlItem["技能类型高阶书"] = "12,12403"

	// School idol badge
	MoegirlItem["学园偶像的凭证"] = "12,1700"

	// Insight pin
	MoegirlItem["灵感徽章★1"] = "6,1400"

	// practice charm
	MoegirlItem["幸运护身符★1"] = "6,1500"

	// EXP
	MoegirlItem["经验值"] = "4,0"
	// Gold
	MoegirlItem["金币"] = "10,0"
	// Star gem
	MoegirlItem["虹彩星石"] = "1,0"

	// Show Candy
	MoegirlItem["演唱会糖果（50）"] = "17,1300"
	MoegirlItem["演唱会糖果（50%）"] = "17,1300"
	MoegirlItem["演唱会糖果（100）"] = "17,1301"
	MoegirlItem["演唱会糖果（100%）"] = "17,1301"

	// School idol radiance
	MoegirlItem["学园偶像的光辉"] = "13,1800"
	// School idol wish
	MoegirlItem["学园偶像的心愿"] = "13,1805"

	// memento exchange ticket
	MoegirlItem["高级记忆碎片交换券"] = "21,21010"

	// pearl
	MoegirlItem["Smile珠子"] = "25,10101"
	MoegirlItem["Pure珠子"] = "25,10102"
	MoegirlItem["Cool珠子"] = "25,10103"
	MoegirlItem["Active珠子"] = "25,10104"
	MoegirlItem["Natural珠子"] = "25,10105"
	MoegirlItem["Elegant珠子"] = "25,10106"

	// counting units
	MoegirlUnit["个"] = true
	MoegirlUnit["张"] = true
	MoegirlUnit["本"] = true
	MoegirlUnit["名"] = true
	MoegirlUnit["張"] = true
}

func MoegirlSetItem(itemName, contentTypeContentId string) {
	MoegirlItem[itemName] = contentTypeContentId
}

func MoegirlParseHighestRank(data string) int {
	data = strings.Split(data, "～")[0]
	data = strings.TrimSpace(data)
	data = strings.ReplaceAll(data, "位", "")
	value, err := strconv.Atoi(data)
	utils.CheckErr(err)
	return value
}

func MoegirlParseCount(data string) int {
	data = strings.TrimSpace(data)
	data = strings.ReplaceAll(data, ",", "")
	for unit := range MoegirlUnit {
		data = strings.ReplaceAll(data, unit, "")
	}
	value, err := strconv.Atoi(data)
	utils.CheckErr(err)
	return value
}

func MoegirlMustGetItem(itemName string) string {
	item, exists := MoegirlItem[itemName]
	if !exists {
		panic("item doesn't exist: " + itemName)
	}
	return item
}

func MoegirlPointRewardBreakdown(line string) (string, string, string) {
	if line == "" {
		panic("line must not be empty")
	}
	tokens := strings.Split(line, "\t")
	fmt.Println(tokens)
	if len(tokens) != 2 {
		return "", "", ""
	}
	requiredPoint := strings.TrimSpace(tokens[0])
	reward := strings.TrimSpace(tokens[1])
	rewardTokens := strings.Split(reward, " ")
	tokenCnt := len(rewardTokens)
	rewardName := strings.TrimSpace(strings.Join(rewardTokens[:tokenCnt-1], " "))
	rewardAmount := strings.TrimSpace(rewardTokens[tokenCnt-1])
	rewardAmount = strings.ReplaceAll(rewardAmount, ",", "")
	for unit := range MoegirlUnit {
		rewardAmount = strings.ReplaceAll(rewardAmount, unit, "")
	}
	return requiredPoint, rewardName, rewardAmount
}

// return point_required,content_type,content_id,content_amount
func MoegirlGetPointReward(line string) string {
	if line == "" {
		return ""
	}
	requiredPoint, rewardName, rewardAmount := MoegirlPointRewardBreakdown(line)
	if requiredPoint == "" {
		fmt.Printf("Warning: Non empty line ignored: \n\t%s\n", line)
		return ""
	}
	contentTypeContentId, exists := MoegirlItem[rewardName]
	if exists {
		return fmt.Sprintf("%s,%s,%s", requiredPoint, contentTypeContentId, rewardAmount)
	} else {
		return "unknown_reward_name"
	}
}
