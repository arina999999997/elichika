// How login bonus is handled:
// There are 2 categories of login bonus:
// - IllustLoginBonus: 2d background
// - NaviLoginBonus: 3d with your navigator or the bday member or whatever
// There are 6 types of login bonus as listed in the enum

// Login bonus common info is stored in 3 tables:
// - s_login_bonus: store id, background id, start at, end at along with white_board_texture_asset,
//   Except for the id, start at, and end at, they all can be null
// - s_login_bonus_reward_day: store the day of the login bonus
// - s_login_bonus_reward_content: store the content of the login bonus days
// - s_login_bonus also additionally store server side parameters:
//   - handler: the name of the handle used to handle that login bonus
//   - handler config: a string config to be passed to the handler

// Login bonus user info is stored in u_login_bonus:
// - user_id: ...
// - login_bonus_id: ...
// - last_received_reward: the 0-indexed reward, default to -1 (have not received anything)
// - last_received_at: timestamp of the last time rewarded, default to 0

// The process is as follow:
// - User call bootstrap to get login bonus
// - The login bonuses are iterated through and are awarded if received.
// - Then the client send read requests that are not tracked
// - from recorded data, the present box count go up only after the read requests are sent, but the present count doesn't go up during the read request
// - this can interpreted as the server fill-in the present box data first, and then add present for the login bonus on top of that
// - although it can also be interpreted as the server only award item once read is called.

package login_bonus

import (
	"elichika/client"
	"elichika/gamedata"
	"elichika/userdata"
)

// handlers take the config, the relevant session, the relevant login bonus, and the bootstrap output

type HandlerType = func(string, *userdata.Session, *gamedata.LoginBonus, *client.BootstrapLoginBonus)

var handler = map[string]HandlerType{}

func init() {
	handler["beginner_login_bonus"] = beginnerLoginBonusHandler
	handler["normal_login_bonus"] = normalLoginBonusHandler
	handler["birthday_login_bonus"] = birthdayLoginBonusHandler
}
