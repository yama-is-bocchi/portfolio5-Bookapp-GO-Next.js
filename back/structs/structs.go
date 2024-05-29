package structs

import (
	"Bookapp/ent"
)

//クライアント(Book_db)データベースのグローバル変数
var Client *ent.Client

// サインアップで送られてくるリクエストを型定義
type SignRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 通信毎に送られてくるトークンチェックの型定義
type TokenCheckRequest struct {
	Name  string `json:"name" binding:"required"`
	Token string `json:"token" binding:"required"`
}

// 本登録する際に送られてくるリクエストの型定義
type BookRegisterRequest struct {
	Name  string `json:"name" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Kind string `json:"kind" binding:"required"`
	Memo string `json:"memo" binding:"required"`
	Token string `json:"token" binding:"required"`
}

// 本更新する際に送られてくるリクエストの型定義
type BookUpdateRequest struct {
	Pretitle string `json:"Pretitle" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Kind string `json:"kind" binding:"required"`
	Memo string `json:"memo" binding:"required"`
	Token string `json:"token" binding:"required"`
}


//テーブルにユーザーIDが存在するか確認する関数の型定義
type ExitsChecktype func(int) bool

//渡された関数で確認する関数
//true存在する,false存在しない
func ExistCheckfunc(user_id int,existchecker ExitsChecktype)bool{
	return existchecker(user_id)
}