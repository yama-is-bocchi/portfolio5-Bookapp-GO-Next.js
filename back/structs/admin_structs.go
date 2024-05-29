package structs

// おすすめ本登録する際に送られてくるリクエストの型定義
type SuggestBookRegisterRequest struct {
	Name  string `json:"name" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Kind string `json:"kind" binding:"required"`
	Price int `json:"price" binding:"required"`
	Memo string `json:"memo" binding:"required"`
	Token string `json:"token" binding:"required"`
}

// おすすめ本更新する際に送られてくるリクエストの型定義
type SuggestBookUpdateRequest struct {
	Pretitle string `json:"Pretitle" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Kind string `json:"kind" binding:"required"`
	Price int `json:"price" binding:"required"`
	Memo string `json:"memo" binding:"required"`
	Token string `json:"token" binding:"required"`
}