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
import { Sidebar,GreenButton } from "@/component/Layout.tsx";
import { useRouter } from "next/router";
import { CautionComment, CheckToken } from "@/methods/util.tsx";
import { GetBookList, DeleteBookapi } from "./../api/Book.ts";
import { DeleteBook } from "@/structs/interface.tsx";
const inter = Inter({ subsets: ["latin"] });

const BookEdit = () => {
  const router = useRouter(); //ルーター
  const [name, setName] = useState(""); //UserName
  const [token, setToken] = useState(""); //Token
  const [ApiDatas, setApiDatas] = useState([]); //APIからの受信データ
  const [checkedItems, setCheckedItems] = useState({}); //チェックされた本のBool配列
  const deleteBooks: DeleteBook[] = []; //削除したい本のデータ

  //編集ボタンクリック
  const EditBtnClick = async(e) => {
    if (await e.target.id === null) return;
    const element=ApiDatas.find(el=>el.title===e.target.id)
    if (await element === null) return;
    //e.target.id=タイトル
    //編集ページへ遷移
    router.push({
      pathname: "/MainMenu/BookEdit/EditInfo",
      query: {
        Name: name,
        Token: token,
        Title: e.target.id,
        Kind:element.kind,
        Memo:element.memo,
      },
    });
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
            編集したい本を選択してください。
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
                  <CardContent style={{ flex: 1 }}>
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
                  <GreenButton
                    variant="contained"
                    color="primary"
                    style={{ marginRight: "16px" }}
                    id={item.title}
                    onClick={EditBtnClick}
                  >
                    <Typography
                      onClick={EditBtnClick}
                      id={item.title}
                      sx={{ fontSize: "1.25em" }}
                    >
                      編集
                    </Typography>
                  </GreenButton>
                </Card>
              </ListItem>
            ))}
          </List>
        </Container>
      </div>
    </main>
  );
};

export default BookEdit;
