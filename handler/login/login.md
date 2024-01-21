# login/Startup
Create new account
client send: 
- mask: random 32 bytes encrypted and then encoded in base64
- resemara_detection_identifier: reset marathon detection (to prevent/rewards rerolling)
- time_difference: second offset from utc based on timezone
- recaptcha_token: empty

server decode the 32 bytes, xor with the SessionKey, then response that as the authorization key along with an user id
