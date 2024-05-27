package book_methods

import (
	"Bookapp/ent/book"
	"Bookapp/lock"
	"Bookapp/structs"
	"Bookapp/token"
	"Bookapp/util"
	"context"
	"github.com/gin-gonic/gin"
	"sync"
)

// 本登録ハンドラ関数
func RegisterHandler(c *gin.Context) {

	// 変数reqをBookRegisterRequestで定義
	var req structs.BookRegisterRequest

	//reqに取得したデータを格納、変換でエラーが起きた場合はエラーを返して終了
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(401, gin.H{"error": "変換できません"})
		return
	}

	//エスケープ処理
	e_Name := util.EscapeInput(req.Name)
	e_Title := util.EscapeInput(req.Title)
	e_Kind := util.EscapeInput(req.Kind)
	e_Memo := util.EscapeInput(req.Memo)
	e_Token := util.EscapeInput(req.Token)

	//ユーザーNameからIDを求める
	subimiteduser, err := util.GetUserEnt(e_Name)

	//ユーザー名が存在しない
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	//アカウントがロックされているか確認
	//あればこの場でreturnする
	existlocklist := lock.CheckLocklist(structs.Client)

	if structs.ExistCheckfunc(subimiteduser.ID, existlocklist) {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	//トークンを確認する
	if !token.CheckNameToken(e_Name, e_Token) {
		c.JSON(401, gin.H{"error": "token"})
		return
	}

	//ユーザーIDとe_titleが一致するデータが登録されていたらエラー
	_, err = structs.Client.Book.Query().
		Where(book.UserIDEQ(subimiteduser.ID), book.TitleEQ(e_Title)).
		First(context.Background())

	if err == nil {
		//存在する
		c.JSON(403, gin.H{"error": "invalid credentials"})
		return

	}
	//DBに登録する
	newBook, err := structs.Client.Book.
		Create().
		SetTitle(e_Title).
		SetKind(e_Kind).
		SetUserID(subimiteduser.ID).
		SetMemo(e_Memo).
		Save(context.Background())

	//データベース登録失敗
	if err != nil {
		c.JSON(401, gin.H{"error": "データベース登録処理", "book": newBook})
		return
	}

	//正常終了
	c.JSON(200, gin.H{"book": newBook})
}

// 本参照ハンドラ関数
func GetListHandler(c *gin.Context) {

	// 変数reqをTokenCheckRequestで定義
	var req structs.TokenCheckRequest

	//reqに取得したデータを格納、変換でエラーが起きた場合はエラーを返して終了
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "変換できません"})
		return
	}
	//エスケープ処理
	e_Name := util.EscapeInput(req.Name)
	e_Token := util.EscapeInput(req.Token)

	//ユーザーNameからIDを求める
	subimiteduser, err := util.GetUserEnt(e_Name)

	//ユーザー名が存在しない
	if err != nil {
		c.JSON(403, gin.H{"error": "invalid credentials"})
		return
	}

	//アカウントがロックされているか確認
	//あればこの場でreturnする
	existlocklist := lock.CheckLocklist(structs.Client)

	if structs.ExistCheckfunc(subimiteduser.ID, existlocklist) {
		c.JSON(403, gin.H{"error": "invalid credentials"})
		return
	}

	//トークンを確認
	if !token.CheckNameToken(e_Name, e_Token) {
		c.JSON(403, gin.H{"error": "invalid credentials"})
		return
	}

	//対象のユーザーIDのBookテーブルのデータを全て送信
	Datas, err := structs.Client.Book.Query().
		Where(book.UserIDEQ(subimiteduser.ID)).
		All(context.Background())

	//存在しないまたはエラー
	if err != nil {
		c.JSON(403, gin.H{"error": "not exist"})
		return
	}

	//正常終了
	c.JSON(200, Datas)
}

// 本更新アップデートハンドラ関数
func UpdateHandler(c *gin.Context) {
	//ユーザーIDとタイトルが一致するBookテーブルのデータを取得してReqの内容に更新して返す
	// 変数reqをBookUpdateRequestで定義
	var req structs.BookUpdateRequest

	//reqに取得したデータを格納、変換でエラーが起きた場合はエラーを返して終了
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "変換できません"})
		return
	}

	//エスケープ処理
	e_Pretitle := util.EscapeInput(req.Pretitle)
	e_Name := util.EscapeInput(req.Name)
	e_Title := util.EscapeInput(req.Title)
	e_Kind := util.EscapeInput(req.Kind)
	e_Memo := util.EscapeInput(req.Memo)
	e_Token := util.EscapeInput(req.Token)
	//ユーザーNameからIDを求める
	subimiteduser, err := util.GetUserEnt(e_Name)

	//ユーザー名が存在しない
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid credentials"})
		return
	}

	//アカウントがロックされているか確認
	//あればこの場でreturnする
	existlocklist := lock.CheckLocklist(structs.Client)

	if structs.ExistCheckfunc(subimiteduser.ID, existlocklist) {
		c.JSON(400, gin.H{"error": "invalid credentials"})
		return
	}

	//トークンを確認する
	if !token.CheckNameToken(e_Name, e_Token) {
		c.JSON(401, gin.H{"error": "token"})
		return
	}

	//reqのpretitleとuserIDのデータが存在するか確認
	target_book, err := structs.Client.Book.Query().
		Where(book.UserIDEQ(subimiteduser.ID), book.TitleEQ(e_Pretitle)).
		First(context.Background())

	if err != nil {
		//存在しない
		c.JSON(400, gin.H{"error": "invalid credentials"})
		return
	}

	//更新後のデータが既に登録されているデータに重複しないか確認
	if e_Pretitle != e_Title {
		duplicatedBook, err := structs.Client.Book.Query().
			Where(book.UserIDEQ(subimiteduser.ID), book.TitleEQ(e_Title)).
			First(context.Background())

		if err == nil {
			//重複している
			c.JSON(403, gin.H{"duplicated": duplicatedBook})
			return
		}
	}
	//アップデートする
	updatedBook, err := structs.Client.Book.
		UpdateOneID(target_book.ID).
		SetKind(e_Kind).SetUserID(subimiteduser.ID).SetTitle(e_Title).SetMemo(e_Memo).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(200, gin.H{"book": updatedBook})
}

// 選択本削除ハンドラ関数
func DeleteHandler(c *gin.Context) {
	// 変数reqをBookRegisterRequestで定義
	var req []structs.BookRegisterRequest

	//reqに取得したデータを格納、変換でエラーが起きた場合はエラーを返して終了
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(403, gin.H{"error": "変換できません"})
		return
	}
	//ゴルーチン定義
	var WaitGroup sync.WaitGroup
	//要素を1つずつ削除する
	for _, item := range req {
		WaitGroup.Add(1)
		go func(item structs.BookRegisterRequest) {
			defer WaitGroup.Done()

			//エスケープ処理
			e_Name := util.EscapeInput(item.Name)
			e_Title := util.EscapeInput(item.Title)
			e_Token := util.EscapeInput(item.Token)
			//titleとuserIDからBookIDを取得
			//ユーザーNameからIDを求める
			subimiteduser, err := util.GetUserEnt(e_Name)

			//ユーザー名が存在しない
			if err != nil {
				c.JSON(403, gin.H{"error": "invalid credentials"})
				return
			}

			//アカウントがロックされているか確認
			//あればこの場でreturnする
			existlocklist := lock.CheckLocklist(structs.Client)

			if structs.ExistCheckfunc(subimiteduser.ID, existlocklist) {
				c.JSON(403, gin.H{"error": "invalid credentials"})
				return
			}

			//トークンを確認する
			if !token.CheckNameToken(e_Name, e_Token) {
				c.JSON(401, gin.H{"error": "token"})
				return
			}
			//BookIDを取得する
			target_book, err := structs.Client.Book.Query().
				Where(book.UserIDEQ(subimiteduser.ID), book.TitleEQ(e_Title)).
				First(context.Background())

			if err != nil {
				//存在しない
				c.JSON(403, gin.H{"error": "invalid credentials"})
				return
			}

			//BookIDを削除
			err = structs.Client.Book.
				DeleteOneID(target_book.ID).
				Exec(context.Background())

			if err != nil {
				c.JSON(404, gin.H{"error": "削除に失敗しました。"})
				return
			}

		}(item)
	}

	//ゴルーチンの処理を待つ
	WaitGroup.Wait()
	c.JSON(200, gin.H{"book": "ok"})
}
