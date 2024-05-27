import { Inter } from "next/font/google";
import styles from "@/styles/Home.module.css";
import React, { useEffect, useState } from "react";
import {
  Card,
  CardContent,
  Typography,
  Container,
  List,
  ListItem,
  Checkbox,
  Button,
} from "@mui/material";
import InboxIcon from "@mui/icons-material/Inbox";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import { Sidebar, RedButton } from "@/component/Layout.tsx";
import { useRouter } from "next/router";
import { CautionComment, CheckToken } from "@/methods/util.tsx";
import { GetBookList,DeleteBookapi } from "./../api/Book.ts";
import {DeleteBook} from "@/structs/interface.tsx"
const inter = Inter({ subsets: ["latin"] });

const BookDelete = () => {
  const router = useRouter(); //ルーター
  const [name, setName] = useState(""); //UserName
  const [token, setToken] = useState(""); //Token
  const [ApiDatas, setApiDatas] = useState([]); //APIからの受信データ
  const [checkedItems, setCheckedItems] = useState({});//チェックされた本のBool配列
  const deleteBooks: DeleteBook[] = [];//削除したい本のデータ


  //削除ボタンクリック
  const DeleteBtnClick = async () => {
    //配列を確認する(checkedItems)
    const checkItems = (): boolean => {
      //チェックされているか
      const ItemarrayLen=Object.keys(checkedItems).length;
      if (ItemarrayLen == 0) {
        //チェックがない
        return false;
      } 
      else 
      {
        //チェックが一つはある

        for(const item in checkedItems) {
          //どこの配列にチェックされているか確認
          //item=number
          if(checkedItems[item]==true){
            //チェックされている
            //対象のApiDatasの要素をNeedDeleteBooksに追加する
            const TargetBook: DeleteBook = {
              Name: name,
              Title: ApiDatas[item].title,
              Kind: ApiDatas[item].kind,
              Memo: ApiDatas[item].memo,
              Token:token
            };
            deleteBooks.push(TargetBook);
           }
        }
        return true;
      }
    };
    //アイテムから対象の本データを配列に挿入
    if(await checkItems()===true){
      //選択が一つ以上あり全て配列に挿入した
      //APIを呼び出し結果を待つ
      const ResultStatus=await DeleteBookapi(deleteBooks);

      if(ResultStatus!==200){
        
        if(ResultStatus===403){
          console.log(deleteBooks)
          CautionComment("caution","サーバーエラー")
          return
        }

       
      }
      else{
        //成功
        //画面再描画
        router.push({
          pathname: "/MainMenu/BookDelete",
          query: { Name: name, Token: token},
        });
        return
      }
    }    
  };

  //カードチェック
  const handleToggle = (index) => {
    setCheckedItems((prevCheckedItems) => ({
      ...prevCheckedItems,
      [index]: !prevCheckedItems[index],
    }));
  };

  useEffect(() => {
    // クエリパラメータを取得
    const { Name, Token } = router.query;
    if (Name && Token) {
      setName(Name);
      setToken(Token);
      //トークン確認
      if (CheckToken(Name, Token) === false) {
        //トークンが一致しない、または期限切れ
        router.push("/");
      } else {
        //トークン一致
        //リスト取得
        const fetchData = async () => {
          try {
            //API送信
            setApiDatas(await GetBookList(Name, Token));
          } catch (error) {
            console.error("データの取得に失敗しました:", error);
          }
        };
        fetchData();
      }
      return;
    } else {
      //Name,Tokenが無い
      router.push("/");
    }
  }, [router.query]);

  return (
    <main className={`${styles.main} ${inter.className}`}>
      <div>
        <Sidebar Name={name} Token={token} />
        <Container>
          <Typography variant="h4" component="h1" gutterBottom>
            削除したい本を選択してください。
            <RedButton variant="contained" color="primary" onClick={DeleteBtnClick}>
              削除
            </RedButton>
          </Typography>
          <Typography sx={{ color: "red", fontSize: "1.25em" }}>
              <span id="caution"></span>
            </Typography>
          <List>
            {ApiDatas.map((item, index) => (
              <ListItem
                key={index}
                style={{ marginBottom: "16px", padding: 0 }}
              >
                <Card
                  style={{
                    width: "100%",
                    display: "flex",
                    alignItems: "center",
                  }}
                >
                  <Checkbox
                    checked={!!checkedItems[index]}
                    onChange={() => handleToggle(index)}
                  />
                  <CardContent>
                    <Typography variant="h5" component="h2">
                      {item.title}
                    </Typography>
                    <Typography color="textSecondary">
                      種類: {item.kind}
                    </Typography>
                    <Typography variant="body2" component="p">
                      メモ: {item.memo}
                    </Typography>
                  </CardContent>
                </Card>
              </ListItem>
            ))}
          </List>
        </Container>
      </div>
    </main>
  );
};

export default BookDelete;
