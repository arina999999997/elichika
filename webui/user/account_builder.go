package user

import (
	"elichika/client"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/item"
	"elichika/locale"
	"elichika/router"
	"elichika/subsystem/user_beginner_challenge"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_member"
	"elichika/subsystem/user_mission"
	"elichika/subsystem/user_status"
	"elichika/subsystem/user_story_event_history"
	"elichika/subsystem/user_story_main"
	"elichika/subsystem/user_story_member"
	"elichika/subsystem/user_suit"
	"elichika/subsystem/user_training_tree"
	"elichika/subsystem/user_voice"
	"elichika/userdata"
	"elichika/webui/webui_utils"

	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

// builder form is a page with features on lines and a button at the end of line to send it

type BuilderFeature struct {
	FeatureName string // used as the main label of the feature
	FeaturePath string // the path to post request to, relative to the builder page
	Input       string // the content of a form
	ButtonText  string // the text on the button
	Handler     func(ctx *gin.Context)
}

var builderFeatures []BuilderFeature

func AddBuilderFeature(featureName, featurePath, inputHTML, buttonText string, handler func(ctx *gin.Context)) {
	builderFeatures = append(builderFeatures,
		BuilderFeature{
			FeatureName: featureName,
			FeaturePath: featurePath,
			Input:       inputHTML,
			ButtonText:  buttonText,
			Handler:     handler,
		})
	router.AddHandler("/webui/user", "POST", "/builder/"+featurePath, handler)
}

func (bf *BuilderFeature) ToHTML() string {
	html := fmt.Sprintf(`<div>
	<form id="%s" method="POST" enctype="multipart/form-data">
		<label>%s</label>
		%s
		<input type="button" onclick="submit_form('%s', './builder/%s')" value="%s">
	</form>
</div>`, bf.FeaturePath, bf.FeatureName, bf.Input, bf.FeaturePath, bf.FeaturePath, bf.ButtonText)
	return html
}

func builderForm(ctx *gin.Context) {
	form :=
		`<div><label>Perform actions in the game quickly, just like you actually did the things in the client!</label></div>
<div><label>Note that this feature perform "simulated client actions" on your account, if you just want some help with the resource then you should use the Resource Helper instead.</label></div>
<div><label>To use, just click on the relevant button after setting the input, if there's any.</label></div>
<div><label>Some request will take a while to finish, be patience and wait for it.</label></div>
<div><label>Also, don't be too greedy or your account might become unplayable due to overflow.</label></div>
<div><label>Finally, if you just want to have a maxed out account then click the button on the right</label>
<button onclick="if (confirm('This will DESTROY your existing progress, proceed?')) get_maxed_account()">Get Maxed out Account</button>
</div>
<br>

<script>
function get_maxed_account() {
	var forms = document.getElementsByTagName("form")
	for (let form of forms) {
		form.reset()
	}
	var inputs = document.getElementsByTagName("input")
	for (let input of inputs) {
		if (input.type == "checkbox") {
			input.setAttribute("checked", null)
		} else if (input.type == "number") {
			input.setAttribute("value", input.getAttribute("max"))
		}
	}
	var ogAlert = window.alert
	window.alert = function(data) {
		console.log(data)
	}
	finished = function() {
		window.alert = ogAlert
		alert("Done")
	}
	var funcs = [finished]
	for (let i = forms.length - 1; i >= 0; i--) {
		const form = forms[i].id
		const func = funcs[funcs.length - 1]
		funcs.push(function() {
			submit_form(form, "./builder/" + form, false, func)
		})
	}
	funcs[funcs.length-1]()
}
</script>
`
	for _, feature := range builderFeatures {
		form += feature.ToHTML()
		form += "\n"
	}
	ctx.HTML(http.StatusOK, "logged_in_user.html", gin.H{
		"body": form,
	})
}

func init() {
	addFeature("Account builder", "builder")
	router.AddHandler("/webui/user", "GET", "/builder", builderForm)
	AddBuilderFeature("Add user EXP", "add_user_exp",
		`<input type="number" name="exp" min="1" max="10000000" value="0">`, "Add user exp",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			exp := webui_utils.GetFormInt32(ctx, "exp")
			user_status.AddUserExp(session, exp)
			session.Finalize()
			webui_utils.CommonResponse(ctx, fmt.Sprintf("Added user EXP: %d", exp), "")
		})
	AddBuilderFeature("Add LP", "add_lp",
		`<input type="number" name="lp" min="1" max="10000" value="0">`, "Add LP",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			lp := webui_utils.GetFormInt32(ctx, "lp")
			user_status.AddUserLp(session, lp)
			session.Finalize()
			webui_utils.CommonResponse(ctx, fmt.Sprintf("Added LP: %d", lp), "")
		})
	AddBuilderFeature("Add Star Gem", "add_sns_coin",
		`<input type="number" name="sns_coin" min="1" max="10000" value="0">`, "Add Star Gem",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			snsCoin := webui_utils.GetFormInt32(ctx, "sns_coin")
			user_content.AddContent(session, item.StarGem.Amount(snsCoin))
			session.Finalize()
			webui_utils.CommonResponse(ctx, fmt.Sprintf("Added star gem: %d", snsCoin), "")
		})
	AddBuilderFeature("Add member coin", "add_subscription_coin",
		`<input type="number" name="subscription_coin" min="1" max="10000" value="0">`, "Add member coin",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			subscriptionCoin := webui_utils.GetFormInt32(ctx, "subscription_coin")
			user_content.AddContent(session, item.MemberCoin.Amount(subscriptionCoin))
			session.Finalize()
			webui_utils.CommonResponse(ctx, fmt.Sprintf("Added member coin: %d", subscriptionCoin), "")
		})
	AddBuilderFeature("Add login day", "add_login_day",
		`<input type="number" name="login_day" min="1" max="10000" value="0">`, "Add login day",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			loginDay := webui_utils.GetFormInt32(ctx, "login_day")
			session.UserStatus.LoginDays += loginDay
			session.Finalize()
			webui_utils.CommonResponse(ctx, fmt.Sprintf("Added login day: %d", loginDay), "")
		})

	AddBuilderFeature("Add costume", "add_suit",
		`<select name="suit_option" required>
			<option value="all">All costumes</option>
			<option value="paid">Paid costumes</option>
		</select>`, "Add costume",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			required := webui_utils.GetFormString(ctx, "suit_option")

			for _, suit := range session.Gamedata.Suit {
				if (required == "all") || (suit.SuitReleaseRoute == enum.SuitReleaseRouteSecret) {
					user_suit.InsertUserSuit(session, suit.Id)
				}
			}
			session.Finalize()
			webui_utils.CommonResponse(ctx, "Added costumes", "")
		})

	AddBuilderFeature("Clear ALL missions", "clear_mission",
		`<input type="checkbox" name="confirm"><label>I want to clear ALL missions and lose track of the current progress</label>`, "Clear all missions",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			confirm := webui_utils.GetFormBool(ctx, "confirm")
			if !confirm {
				webui_utils.CommonResponse(ctx, "Check the confirm box if you really want to complete all missions", "")
				return
			}
			user_mission.FetchMission(session)

			for {
				ids := []int32{}
				for _, mission := range session.UserModel.UserMissionByMissionId.Map {
					if mission.IsReceivedReward {
						continue
					}
					mission.IsCleared = true
					mission.MissionCount = session.Gamedata.Mission[mission.MissionMId].MissionClearConditionCount
					ids = append(ids, mission.MissionMId)
				}
				if len(ids) == 0 {
					break
				}
				user_mission.ReceiveReward(session, ids)
			}
			session.Finalize()
			webui_utils.CommonResponse(ctx, "Cleared all mission, all reward delivered", "")
		})

	AddBuilderFeature("Clear ALL beginner challenge", "clear_beginner_challenge",
		`<input type="checkbox" name="confirm"><label>I want to clear ALL  beginner challenges and lose track of the current progress</label>`, "Clear all challenges",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			confirm := webui_utils.GetFormBool(ctx, "confirm")
			if !confirm {
				webui_utils.CommonResponse(ctx, "Check the confirm box if you really want to complete all challenges", "")
				return
			}
			challengeCells := user_beginner_challenge.GetBeginnerChallengeCells(session)

			for _, cell := range challengeCells {
				if cell.IsRewardReceived {
					continue
				}
				cell.IsRewardReceived = true
				user_beginner_challenge.UpdateChallengeCell(session, *cell)
			}
			session.Finalize()
			webui_utils.CommonResponse(ctx, "Cleared all beginner challenges", "")
		})

	{
		memberOptions := ""
		memberPattern := `<option value="%d">%s</option>` + "\n"

		gd := locale.Locales["en"].Gamedata
		members := []*gamedata.Member{}
		for _, member := range gd.Member {
			members = append(members, member)
		}
		sort.Slice(members, func(i, j int) bool {
			return members[i].Id < members[j].Id
		})
		// all filter
		memberIdToBitIndex := func(id int32) int32 {
			if id <= 9 {
				return 1 << (id - 1)
			} else if id <= 109 {
				return 1 << (id - 101 + 9)
			} else {
				return 1 << (id - 201 + 18)
			}
		}

		// everyone
		{
			everyone := int32(0)
			for _, member := range members {
				everyone |= memberIdToBitIndex(member.Id)
			}
			memberOptions += fmt.Sprintf(memberPattern, everyone, "Everyone")
		}
		// group
		{
			group := [4]int32{}
			for _, member := range members {
				group[member.MemberGroup] |= memberIdToBitIndex(member.Id)
			}
			for i := int32(1); i <= 3; i++ {
				memberOptions += fmt.Sprintf(memberPattern, group[i], gd.MemberGroup[i].GroupName)
			}
		}
		// member
		{
			for _, member := range members {
				memberOptions += fmt.Sprintf(memberPattern, memberIdToBitIndex(member.Id), member.Name)
			}
		}

		AddBuilderFeature("Add 1 copy of all cards of ", "add_card",
			fmt.Sprintf(`<select name="filter">
				%s
			</select>`, memberOptions), "Add card",
			func(ctx *gin.Context) {
				session := ctx.MustGet("session").(*userdata.Session)
				filter := webui_utils.GetFormInt32(ctx, "filter")
				for _, card := range session.Gamedata.Card {
					if (filter & memberIdToBitIndex(card.Member.Id)) == 0 {
						continue
					}
					user_card.AddUserCardByCardMasterId(session, card.Id)
				}
				session.Finalize()
				webui_utils.CommonResponse(ctx, "Added cards", "")
			})

		AddBuilderFeature("Increase to max limit break all owned cards of ", "mlb",
			fmt.Sprintf(`<select name="filter">
				%s
			</select>`, memberOptions), "Increase Limit Break",
			func(ctx *gin.Context) {
				session := ctx.MustGet("session").(*userdata.Session)
				filter := webui_utils.GetFormInt32(ctx, "filter")
				session.PopulateUserModelField("UserCardByCardId")
				cardIds := []int32{}
				for _, card := range session.UserModel.UserCardByCardId.Map {
					if (filter & memberIdToBitIndex(session.Gamedata.Card[card.CardMasterId].Member.Id)) == 0 {
						continue
					}
					for i := card.Grade + 1; i <= 5; i++ {
						cardIds = append(cardIds, card.CardMasterId)
					}
				}
				for _, cardId := range cardIds {
					user_training_tree.GradeUpCard(session, cardId, 0)
				}
				session.Finalize()
				webui_utils.CommonResponse(ctx, "Max limit breaked cards", "")
			})

		AddBuilderFeature("Increase bond point of", "increase_love",
			fmt.Sprintf(`<select name="filter">
				%s
			</select> by <input type="number" min="1" max="100000000" name="amount">`, memberOptions), "Increase Bond",
			func(ctx *gin.Context) {
				session := ctx.MustGet("session").(*userdata.Session)
				filter := webui_utils.GetFormInt32(ctx, "filter")
				amount := webui_utils.GetFormInt32(ctx, "amount")
				for _, member := range session.Gamedata.Member {
					if (filter & memberIdToBitIndex(member.Id)) == 0 {
						continue
					}
					user_member.AddMemberLovePoint(session, member.Id, amount)
				}
				session.Finalize()
				webui_utils.CommonResponse(ctx, "Increased bond points", "")
			})

		AddBuilderFeature("Unlock all available bond board boards of ", "unlock_love_panel",
			fmt.Sprintf(`<select name="filter">
				%s
			</select>`, memberOptions), "Unlock boards",
			func(ctx *gin.Context) {
				session := ctx.MustGet("session").(*userdata.Session)
				filter := webui_utils.GetFormInt32(ctx, "filter")
				for _, member := range session.Gamedata.Member {
					if (filter & memberIdToBitIndex(member.Id)) == 0 {
						continue
					}
					userMember := user_member.GetMember(session, member.Id)
					userMemberStartPanelId := member.Id + 1000
					userLovePanel := user_member.GetMemberLovePanel(session, member.Id)
					panel := session.Gamedata.MemberLovePanel[userMemberStartPanelId]
					level := int32(0)
					for panel != nil {
						if panel.LoveLevelMasterLoveLevel > userMember.LoveLevel {
							break
						} else {
							panel = panel.NextPanel
							level++
						}
					}
					userLovePanel.MemberLovePanelCellIds.Slice = []int32{}
					for i := int32(0); i < level; i++ {
						for j := int32(1); j <= 5; j++ {
							userLovePanel.MemberLovePanelCellIds.Append(i*10000 + j*1000 + member.Id)
						}
					}
					user_member.UpdateMemberLovePanel(session, userLovePanel)
				}
				session.Finalize()
				webui_utils.CommonResponse(ctx, "Unlocked bond board(s)", "")
			})

		AddBuilderFeature("Practice (tiles and level) all owned cards of", "practice_cards",
			fmt.Sprintf(`<select name="filter">
			%s
		</select>`, memberOptions), "Practice",
			func(ctx *gin.Context) {
				session := ctx.MustGet("session").(*userdata.Session)
				filter := webui_utils.GetFormInt32(ctx, "filter")
				session.PopulateUserModelField("UserCardByCardId")
				maxCardLevel := map[int32]map[int32]int32{}

				prepare := func(memberId int32) {
					_, exist := maxCardLevel[memberId]
					if exist {
						return
					}
					maxCardLevel[memberId] = map[int32]int32{}
					maxCardLevel[memberId][enum.CardRarityTypeRare] = session.Gamedata.CardRarity[enum.CardRarityTypeRare].MaxLevel
					maxCardLevel[memberId][enum.CardRarityTypeSRare] = session.Gamedata.CardRarity[enum.CardRarityTypeSRare].MaxLevel
					maxCardLevel[memberId][enum.CardRarityTypeURare] = session.Gamedata.CardRarity[enum.CardRarityTypeURare].MaxLevel
					userLovePanel := user_member.GetMemberLovePanel(session, memberId)
					userLovePanel.Fix()
					for i, cellId := range userLovePanel.MemberLovePanelCellIds.Slice {
						cellContent := session.Gamedata.MemberLovePanelCell[cellId]
						switch cellContent.BonusType {
						case enum.MemberLovePanelEffectTypeRLevel:
							maxCardLevel[memberId][enum.CardRarityTypeRare] += cellContent.BonusValue
						case enum.MemberLovePanelEffectTypeSrLevel:
							maxCardLevel[memberId][enum.CardRarityTypeSRare] += cellContent.BonusValue
						case enum.MemberLovePanelEffectTypeUrLevel:
							maxCardLevel[memberId][enum.CardRarityTypeURare] += cellContent.BonusValue
						default:
						}
						if i%5 == 4 { // full level
							panel := session.Gamedata.MemberLovePanel[*cellContent.MemberLovePanelMasterId]
							for _, bonus := range panel.Bonuses {
								switch bonus.BonusType {
								case enum.MemberLovePanelEffectTypeRLevel:
									maxCardLevel[memberId][enum.CardRarityTypeRare] += bonus.BonusValue
								case enum.MemberLovePanelEffectTypeSrLevel:
									maxCardLevel[memberId][enum.CardRarityTypeSRare] += bonus.BonusValue
								case enum.MemberLovePanelEffectTypeUrLevel:
									maxCardLevel[memberId][enum.CardRarityTypeURare] += bonus.BonusValue
								default:
								}
							}

						}
					}
				}

				for _, card := range session.UserModel.UserCardByCardId.Map {
					masterCard := session.Gamedata.Card[card.CardMasterId]
					memberId := masterCard.Member.Id
					if (filter & memberIdToBitIndex(memberId)) == 0 {
						continue
					}
					prepare(memberId)
					card.Level = maxCardLevel[memberId][masterCard.CardRarityType]
					if card.IsAllTrainingActivated {
						continue
					}
					trainingTree := user_training_tree.GetUserTrainingTree(session, card.CardMasterId)
					hasCellId := map[int32]bool{}
					for _, cell := range trainingTree.Slice {
						hasCellId[cell.CellId] = true
					}
					cellContents := masterCard.TrainingTree.TrainingTreeMapping.TrainingTreeCellContents
					addedCells := []int32{}
					for i, cell := range cellContents {
						if (i == 0) || (cell.RequiredGrade > card.Grade) || hasCellId[cell.CellId] {
							continue
						}
						addedCells = append(addedCells, cell.CellId)
					}
					if len(addedCells) > 0 {
						user_training_tree.ActivateTrainingTreeCells(session, card.CardMasterId, addedCells)
					}
				}
				session.Finalize()
				webui_utils.CommonResponse(ctx, "Practiced(s)", "")
			})

		AddBuilderFeature("Finish owned bond episodes of", "finish_member_story",
			fmt.Sprintf(`<select name="filter">
				%s`, memberOptions), "Finish",
			func(ctx *gin.Context) {
				session := ctx.MustGet("session").(*userdata.Session)
				filter := webui_utils.GetFormInt32(ctx, "filter")
				for _, storyMember := range session.Gamedata.StoryMember {
					if (filter & memberIdToBitIndex(storyMember.MemberMId)) == 0 {
						continue
					}
					if user_member.GetMember(session, storyMember.MemberMId).LoveLevel < storyMember.LoveLevel {
						continue
					}
					user_story_member.FinishStoryMember(session, storyMember.Id)
				}
				session.UserModel.UserInfoTriggerBasicByTriggerId.Clear()
				session.Finalize()
				webui_utils.CommonResponse(ctx, "Finished bond episodes", "")
			})
	}

	AddBuilderFeature("Finish main story", "unlock_main_story",
		`<input type="checkbox" name="confirm"><label>I want to finish the whole main stories and lose track of the current progress</label>`, "Unlock",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			confirm := webui_utils.GetFormBool(ctx, "confirm")
			if !confirm {
				webui_utils.CommonResponse(ctx, "Check the confirm box if you really want to finish main story", "")
				return
			}
			for _, story := range session.Gamedata.StoryMainChapter {
				for _, cell := range story.Cells {
					user_story_main.InsertUserStoryMain(session, cell)
				}
			}
			session.Finalize()
			webui_utils.CommonResponse(ctx, "Finish main story!", "")
		})
	AddBuilderFeature("Unlock voice lines", "unlock_voice",
		"", "Unlock",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)

			for _, voice := range session.Gamedata.NaviVoice {
				if voice.ListType == enum.NaviVoiceListTypeNo {
					continue
				}
				user_voice.UpdateUserVoice(session, voice.Id, false)
			}
			session.Finalize()
			webui_utils.CommonResponse(ctx, "Unlocked all voice lines!", "")
		})
	AddBuilderFeature("Unlock all backgrounds", "unlock_background",
		"", "Unlock",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			for _, bg := range session.Gamedata.CustomBackground {
				user_content.AddContent(session, client.Content{
					ContentType:   enum.ContentTypeCustomBackground,
					ContentId:     bg.Id,
					ContentAmount: 1,
				})
			}
			session.Finalize()
			webui_utils.CommonResponse(ctx, "Unlocked all backgrounds!", "")
		})
	AddBuilderFeature("Unlock all event stories", "unlock_event_story",
		"", "Unlock",
		func(ctx *gin.Context) {
			session := ctx.MustGet("session").(*userdata.Session)
			for _, story := range session.Gamedata.StoryEventHistory {
				user_story_event_history.UnlockEventStory(session, story.StoryEventId)
			}
			session.Finalize()
			webui_utils.CommonResponse(ctx, "Unlocked all event stories!", "")
		})
}
