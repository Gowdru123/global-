// (c) Jisin0

package utils

import "github.com/PaulSonOfLars/gotgbot/v2"

var TEXT map[string]string = map[string]string{

	"START": `
<b>Hᴇʏ %v ɪᴍ %v ᴀɴ 𝐈 𝐚𝐦 𝐠𝐥𝐨𝐛𝐚𝐥 𝐭𝐡𝐞𝐚𝐭𝐞𝐫 𝐩𝐫𝐢𝐧𝐭 𝐛𝐨𝐭 𝐈 𝐰𝐢𝐥𝐥 𝐬𝐞𝐧𝐝 𝐚𝐥𝐥 𝐥𝐚𝐧𝐠𝐮𝐚𝐠𝐞𝐬 𝐧𝐞𝐰 𝐭𝐡𝐞𝐚𝐭𝐞𝐫 𝐩𝐫𝐢𝐧𝐭 𝐦𝐨𝐯𝐢𝐞'𝐬</b>

<i>𝐔 𝐜𝐚𝐧 𝐚𝐥𝐬𝐨 𝐚𝐝𝐝 𝐦𝐞 𝐢𝐧 𝐠𝐫𝐨𝐮𝐩 𝐈 𝐰𝐢𝐥𝐥 𝐩𝐫𝐨𝐯𝐢𝐝𝐞 𝐭𝐡𝐞𝐚𝐭𝐞𝐫 𝐩𝐫𝐢𝐧𝐭 𝐦𝐨𝐯𝐢𝐞'𝐬 𝐢𝐧 𝐮𝐫 𝐠𝐫𝐨𝐮𝐩 𝐌𝐲 𝐂𝐫𝐞𝐚𝐭𝐞𝐫 @jnaneshgowdru</i>
	`,
	"ABOUT": `
<b>Լαɳցᥙαցҽ</b> : <a href='https://ROCKERSBACKUP'>ಕನ್ನಡ</a>
<b>Ƒɾα𝓶ҽɯσƙ</b> : <a href='https://jnaneshgowdru'>jnanesh</a>
<b>Sҽɾʋҽɾ</b> : <a href='heroku.com'>𝗛𝗲𝗿𝗼𝗸𝘂</a>
<b>Ɗα𝜏αẞαടҽ</b> : <a href='mongodb.org'>𝗠𝗼𝗻𝗴𝗼𝗗𝗕</a>
<b>Ɗҽʋҽɬσρҽɾ</b> : <a href='t.me/jnaneshgowdru'>jnanesh</a>
<b>Sᥙρρσɾ𝜏</b> : <a href='t.me/ROCKERSBACKUP'>𝗛𝗲𝗿𝗲</a>
	`,

	"MF": `
<b>Mᴀɴᴜᴀʟ ғɪʟᴛᴇʀs ᴀʟʟᴏᴡ ʏᴏᴜ ᴛᴏ sᴀᴠᴇ ᴄᴜsᴛᴏᴍ ғɪʟᴛᴇʀs ᴏᴛʜᴇʀ ᴛʜᴀɴ ᴛʜᴇ ᴀᴜᴛᴏᴍᴀᴛɪᴄ ᴏɴᴇs. Fɪʟᴛᴇʀs ᴄᴀɴ ʙᴇ ᴏғ ᴛᴇxᴛ/ᴘʜᴏᴛᴏ/ᴅᴏᴄᴜᴍᴇɴᴛ/ᴀᴜᴅɪᴏ/ᴀɴɪᴍᴀᴛɪᴏɴ/ᴠɪᴅᴇᴏ .</b>

<b><u>Nᴇᴡ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/filter "keyword" text</code> or
Rᴇᴘʟʏ ᴛᴏ ᴀ ᴍᴇssᴀɢᴇ -><code>/filter "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/filter "hi" hello</code>
[<code>hello</code>] -> Reply -> <code>/filter hi</code>

<b><u>Sᴛᴏᴘ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/stop "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/stop "hi"</code>

<b><u>Vɪᴇᴡ ғɪʟᴛᴇʀs :</u></b>
<code>/filters</code>
`,

	"GF": `
<b>Gʟᴏʙᴀʟ ғɪʟᴛᴇʀs ᴀʀᴇ ᴍᴀɴᴜᴀʟ ғɪʟᴛᴇʀs sᴀᴠᴇᴅ ʙʏ ʙᴏᴛ ᴀᴅᴍɪɴs ᴛʜᴀᴛ ᴡᴏʀᴋ ɪɴ ᴀʟʟ ᴄʜᴀᴛs. Tʜᴇʏ ᴘʀᴏᴠɪᴅᴇ ʟᴀᴛᴇsᴛ ᴍᴏᴠɪᴇs ɪɴ ᴀ ᴇᴀsʏ ᴛᴏ ᴜsᴇ ғᴏʀᴍᴀᴛ.</b>

<b><u>Sᴛᴏᴘ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/stop "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/stop "et"</code>

<b><u>Vɪᴇᴡ ғɪʟᴛᴇʀs :</u></b>
/gfilters
`,
	"CONNECT": `
<b>Cᴏɴɴᴇᴄᴛɪᴏɴs ᴀʟʟᴏᴡ ʏᴏᴜ ᴛᴏ ᴍᴀɴᴀɢᴇ ʏᴏᴜʀ ɢʀᴏᴜᴘ ʜᴇʀᴇ ɪɴ ᴘᴍ ɪɴsᴛᴇᴀᴅ ᴏғ sᴇɴᴅɪɴɢ ᴛʜᴏsᴇ ᴄᴏᴍᴍᴀɴᴅs ᴘᴜʙʟɪᴄʟʏ ɪɴ ᴛʜᴇ ɢʀᴏᴜᴘ ⠘⁾</b>

<b><u>Cᴏɴɴᴇᴄᴛ :</u></b>
-> Fɪʀsᴛ ɢᴇᴛ ʏᴏᴜʀ ɢʀᴏᴜᴘ's ɪᴅ ʙʏ sᴇɴᴅɪɴɢ /id ɪɴ ʏᴏᴜʀ ɢʀᴏᴜᴘ
-> <code>/connect [group_id]</code>

<b><u>Dɪsᴄᴏɴɴᴇᴄᴛ :</u></b>
<code>/disconnect</code>
`,

	"BROADCAST": `
<b>The broadcast feature allows bot admins to broadcast a message to all of the bot's users.</b>

<I>Broadcasts are of two types :</i>
 ~ Full Broadcast - Broadcast to all of the bot users <code>/broadcast</code>
 ~ Concast - Broadcast to only users who are connected to a chat <code>/concast</code>
`,

	"HELP": `
<b>To know how to use my modules use the buttons below to get all my commands with usage examples !</b>
`,
}

var BUTTONS map[string][][]gotgbot.InlineKeyboardButton = map[string][][]gotgbot.InlineKeyboardButton{
	"START": {
		{
			{Text: "☂ Aʙᴏᴜᴛ ☂", CallbackData: "edit(ABOUT)"},
			{Text: "🧭 Help 🧭", CallbackData: "edit(HELP)"},
			{Text: "🫂 Sᴜᴘᴘᴏʀᴛ 🫂", Url: "t.me/jnaneshgowdru"},
		},
	},
	"ABOUT": {
		{
			{Text: "𝙷𝙾𝙼𝙴", CallbackData: "edit(START)"},
			{Text: "𝚂𝚃𝙰𝚃𝚂", CallbackData: "stats"},
		},
	},
	"STATS": {
		{
			{Text: "𝙱𝙰𝙲𝙺", CallbackData: "edit(ABOUT)"},
			{Text: "𝚁𝙴𝙵𝚁𝙴𝚂𝙷", CallbackData: "stats"},
		},
	},
	"HELP": {
		{{Text: "Fɪʟᴛᴇʀ", CallbackData: "edit(MF)"},
			{Text: "Gʟᴏʙᴀʟ", CallbackData: "edit(GF)"},
		}, {
			{Text: "Cᴏɴɴᴇᴄᴛ", CallbackData: "edit(CONNECT)"}, {Text: "Broadcast", CallbackData: "edit(BROADCAST)"},
		},
		{{Text: "Bᴀᴄᴋ ➔", CallbackData: "edit(START)"}},
	},
}
