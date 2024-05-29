package admin_methods

import (
	"Bookapp/ent"
	"Bookapp/ent/admin"
	"Bookapp/ent/miss"
	"Bookapp/ent/user"
	"Bookapp/lock"
	"Bookapp/miss_methods"
	"Bookapp/structs"
	"Bookapp/token"
	"Bookapp/util"
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"sync"
)

// サインインのHTTPハンドラ関数
func AdminSignInHandler(c *gin.Context) {

	// 変数reqをSignRequestで定義
	var req structs.SignRequest

	//reqに取得したデータを格納、変換でエラーが起きた場合はエラーを返して終了
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(401, gin.H{"error": "err"})
		return
	}

	//エスケープ処理
	e_Name := util.EscapeInput(req.Name)
	e_Password := util.EscapeInput(req.Password)

	// ユーザの検索、パスワードの照合を行う
	sign_in_user, err := structs.Client.User.Query().
		Where(user.NameEQ(e_Name), user.PasswordEQ(e_Password)).
		First(context.Background())

	//パスワード,ユーザー名が不一致またはユーザー名が存在しない
	if err != nil {

		//失敗したNameがMissテーブルに存在するなら更新,存在しないなら追加する
		//まずUserNameが実在するものか確認する
		miss_user, err := util.GetUserEnt(e_Name)

		if err != nil {
			//存在しないので
			c.JSON(401, gin.H{"error": "err"})
			return
		}

		//アカウントがロックされているか確認

		//ユーザーIDがロックテーブルに無いか確認する
		//あればこの場でreturnする
		existlocklist := lock.CheckLocklist(structs.Client)

		if structs.ExistCheckfunc(miss_user.ID, existlocklist) {
			c.JSON(403, gin.H{"error": "lock"})
			return
		}

		//UserNameがMissテーブルに存在するか確認
		existmisslist := miss_methods.CheckMisslist(structs.Client)

		if structs.ExistCheckfunc(miss_user.ID, existmisslist) {
			//存在する

			target_user, err := structs.Client.Miss.Query().
				Where(miss.UserIDEQ(miss_user.ID)).
				First(context.Background())

			if err != nil {
				c.JSON(401, gin.H{"error": "err"})
				return
			}
			//インクリメント処理
			//5回超えたらロックテーブルに追加する

			_, err = miss_methods.IncrementCountatMissList(structs.Client, target_user.ID, target_user.UserID, target_user.Count)

			if err != nil {
				c.JSON(401, gin.H{"error": "err"})
				return
			}

			c.JSON(401, gin.H{"error": "err"})
			return
		}

		//Missテーブルには存在しないので追加処理
		_, err = miss_methods.InsertMissList(structs.Client, miss_user.ID)
		if err != nil {
			c.JSON(401, gin.H{"error": "err"})
		}
		c.JSON(401, gin.H{"error": "err"})
		return
	}

	//ユーザーIDがロックテーブルに無いか確認する
	//あればこの場でreturnする
	existlocklist := lock.CheckLocklist(structs.Client)

	if structs.ExistCheckfunc(sign_in_user.ID, existlocklist) {
		c.JSON(403, gin.H{"error": "lock"})
		return
	}

	//Adminテーブルに存在するIDか確認する。
	//存在しなければミステーブルに追加、インクリメント、ロックを行う
	existadminlist := CheckAdmin(structs.Client)

	if !structs.ExistCheckfunc(sign_in_user.ID, existadminlist) {
		//存在しないので上記処理を行う
		//UserNameがMissテーブルに存在するか確認
		existmisslist := miss_methods.CheckMisslist(structs.Client)

		if structs.ExistCheckfunc(sign_in_user.ID, existmisslist) {
			//存在する

			target_user, err := structs.Client.Miss.Query().
				Where(miss.UserIDEQ(sign_in_user.ID)).
				First(context.Background())

			if err != nil {
				c.JSON(401, gin.H{"error": "err"})
				return
			}
			//インクリメント処理
			//5回超えたらロックテーブルに追加する

			_, err = miss_methods.IncrementCountatMissList(structs.Client, target_user.ID, target_user.UserID, target_user.Count)

			if err != nil {
				c.JSON(401, gin.H{"error": "err"})
				return
			}

			c.JSON(401, gin.H{"error": "err"})
			return
		}

		//Missテーブルには存在しないので追加処理
		_, err = miss_methods.InsertMissList(structs.Client, sign_in_user.ID)
		if err != nil {
			c.JSON(401, gin.H{"error": "err"})
		}
		c.JSON(401, gin.H{"error": "err"})
		return
	}

	//ゴルーチン定義
	var WaitGroup sync.WaitGroup
	WaitGroup.Add(2)
	ReverseChan := make(chan bool, 1)
	TokenChan := make(chan string, 1)

	//ミスリストに存在するなら0回にアップデートする
	go func() {
		defer WaitGroup.Done()
		existmisslist := miss_methods.CheckMisslist(structs.Client)

		if structs.ExistCheckfunc(sign_in_user.ID, existmisslist) {
			//存在するので0に戻す
			_, err := miss_methods.ReverseZeroMissCount(structs.Client, sign_in_user.ID)

			if err != nil {
				c.JSON(401, gin.H{"error": "err"})
				return
			}
		}

		ReverseChan <- true

	}()

	//トークンを生成して返す
	go func() {
		defer WaitGroup.Done()
		generatedToken, err := token.SaveToken(structs.Client, sign_in_user.ID)

		if err != nil {
			c.JSON(401, gin.H{"error": "err"})
			return
		}
		TokenChan <- generatedToken

	}()

	//ゴルーチンの処理を待つ
	WaitGroup.Wait()
	close(ReverseChan)
	close(TokenChan)

	newToken := <-TokenChan

	c.JSON(200, gin.H{"user": sign_in_user, "token": newToken, "error": "OK"})

}

func CheckAdmin(client *ent.Client) structs.ExitsChecktype {
	return func(user_id int) bool {
		// 新規発行か確認
		_, err := client.Admin.Query().
			Where(admin.UserIDEQ(user_id)).
			First(context.Background())
		return err == nil
	}

}
