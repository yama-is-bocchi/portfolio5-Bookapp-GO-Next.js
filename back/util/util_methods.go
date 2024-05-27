package util 

import(
	"Bookapp/ent"
	"Bookapp/ent/user"
	"context"
	"Bookapp/structs"
	"strings"
)

//名前からユーザーの情報を取得
func GetUserEnt(Name string)(*ent.User,error){
	user, err := structs.Client.User.Query().
	Where(user.NameEQ(Name)).
	First(context.Background())
	return user,err
}

//エスケープ処理関数
func EscapeInput(input string) string {
	// 置き換える特殊文字のリストを定義
	specialChars := []string{
		"&", "<", ">", "\"", "'", "/", "\\", "`", "=", "$", ";", ":",
		"(", ")", "{", "}", "[", "]", "|", "^", "%", "~", "#", "!",
		"@", "*", "?", "+",
	}

	// 特殊文字を空文字列に置き換える
	for _, char := range specialChars {
		input = strings.ReplaceAll(input, char, "")
	}

	return input
}

