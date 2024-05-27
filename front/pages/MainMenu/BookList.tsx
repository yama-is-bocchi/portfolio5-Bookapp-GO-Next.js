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
  ListItemText,
} from "@mui/material";
import InboxIcon from "@mui/icons-material/Inbox";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import { Sidebar } from "@/component/Layout.tsx";
import { useRouter } from "next/router";
import { CautionComment, CheckToken } from "@/methods/util.tsx";
import { GetBookList } from "./../api/Book.ts";
const inter = Inter({ subsets: ["latin"] });

const BookList = () => {
  const router = useRouter(); //ルーター
  const [name, setName] = useState(""); //UserName
  const [token, setToken] = useState(""); //Token
  const [ApiDatas, setApiDatas] = useState([]); //APIからの受信データ

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
            登録本データ一覧
          </Typography>
          <List>
            {ApiDatas.map((item, index) => (
              <ListItem
                key={index}
                style={{ marginBottom: "16px", padding: 0 }}
              >
                <Card style={{ width: "100%" }}>
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

export default BookList;
