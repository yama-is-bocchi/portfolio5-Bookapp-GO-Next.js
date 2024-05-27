import { Inter } from "next/font/google";
import styles from "@/styles/Home.module.css";
import React, { useEffect, useState } from "react";
import {
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Typography,
  AppBar,
  Toolbar,
  IconButton,
  Container,
  Paper,
  Grid,
  TextField,
  Button,
} from "@mui/material";
import InboxIcon from "@mui/icons-material/Inbox";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import { Sidebar, GreenButton } from "@/component/Layout.tsx";
import { useRouter } from "next/router";
import {
  PasswordChecker,
  CautionComment,
  EscapeInput,
  CheckNormalLength,
  CheckMemoLength,
  CheckToken,
} from "@/methods/util.tsx";
import { UpdateBookapi } from "./../../api/Book.ts";

const inter = Inter({ subsets: ["latin"] });

const EditInfo = () => {
  const router = useRouter(); //ルーター
  const [name, setName] = useState(""); //UserName
  const [token, setToken] = useState(""); //Token
  const [PreTitle, setPreTitle] = useState(""); //編集前タイトル
  const [Title, setTitle] = useState(""); //タイトル
  const [Memo, setMemo] = useState(""); //メモ
  const [selectedItem, setSelectedItem] = useState(""); //種類
  const [showList, setShowList] = useState(false); //リスト状態

  //リスト展開ボタンクリック
  const handleButtonClick = () => {
    setShowList(!showList);
  };

  //リストアイテムボタンクリック
  const handleItemClick = async (item) => {
    await setSelectedItem(item);
    setShowList(false);
  };

  //送信ボタンクリック
  const SubmitClick = async () => {
    //入力に空欄が無いか確認
    if (Title === "" || Title === null) {
      CautionComment("caution", "タイトルを入力してください");
      return;
    }
    if (Memo === "" || Memo === null) {
      CautionComment("caution", "メモを入力してください");
      return;
    }
    if (selectedItem === "" || selectedItem === null) {
      CautionComment("caution", "本の状態を選択してください");
      return;
    }
    //長さチェック
    if (CheckNormalLength(Title) != true || CheckMemoLength(Memo) != true) {
      CautionComment("caution", "(タイトル20字以内メモ40字以内)");
      return;
    }

    //エスケープ処理
    let EscapedTitle = EscapeInput(Title);
    let EscapedMemo = EscapeInput(Memo);
    
    //サーバーに送信
    let ResultStatus = await UpdateBookapi(
      PreTitle,
      name,
      EscapedTitle,
      selectedItem,
      EscapedMemo,
      token
    );
    //タイトルが重複していないか確認
    if (ResultStatus === 403) {
      CautionComment("caution", "タイトルが重複しています");
      return;
    }
    if (ResultStatus === 200) {
      //成功
      //画面遷移
      router.push({
        pathname: "/MainMenu/BookEdit",
        query: { Name: name, Token: token },
      });
      return;
    } else {
      //失敗
      CautionComment("caution", "サーバーエラー");
      return;
    }
  };

  useEffect(() => {
    // クエリパラメータを取得
    const { Name, Token, Title, Kind, Memo } = router.query;
    //タイトルの内容を入力する
    if (Name && Token) {
      setName(Name);
      setToken(Token);

      //トークン確認
      if (CheckToken(Name, Token) === false) {
        //トークンが一致しない、または期限切れ
        router.push("/");
      } else {
        const EnterData = async () => {
          try {
            //正常
            await setPreTitle(Title);
            await setTitle(Title);
            await setSelectedItem(Kind);
            await setMemo(Memo);

            //各要素に値を代入しておく
            let ElementID = document.getElementById("TitleField");
            if (ElementID === null) return;
            ElementID.value = Title;
            ElementID = document.getElementById("MemoField");
            if (ElementID === null) return;
            ElementID.value = Memo;
            return;
          } catch (error) {
            console.error("データの取得に失敗しました:", error);
            return;
          }
        };
        EnterData();
      }
    } else {
      //Name,Tokenが無い
      router.push("/");
    }
  }, [router.query]);

  return (
    <main className={`${styles.main} ${inter.className}`}>
      <div>
        <Sidebar Name={name} Token={token} />
        <Container maxWidth="sm" sx={{ padding: 4 }}>
          <Paper sx={{ padding: 8, marginTop: 4, boxShadow: 8 }}>
            <Typography
              variant="h4"
              align="center"
              gutterBottom
              style={{ fontFamily: "Roboto, sans-serif" }}
            >
              本を編集
            </Typography>
            <Typography sx={{ color: "red", fontSize: "1.25em" }}>
              <span id="caution"></span>
            </Typography>
            <Grid container spacing={3}>
              <Grid item xs={12}>
                <TextField
                  id="TitleField"
                  fullWidth
                  label="タイトル"
                  variant="outlined"
                  style={{ fontFamily: "Arial, sans-serif" }}
                  onChange={(e) => {
                    setTitle(e.target.value);
                  }}
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  id="MemoField"
                  fullWidth
                  label="メモ"
                  variant="outlined"
                  style={{ fontFamily: "Arial, sans-serif" }}
                  onChange={(e) => {
                    setMemo(e.target.value);
                  }}
                />
              </Grid>
              <Grid item xs={12}>
                本の状態
                <Typography
                  onClick={handleButtonClick}
                  sx={{
                    cursor: "pointer",
                    border: "1px solid #ccc",
                    padding: "8px",
                    margin: "5px",
                    borderRadius: "3px",
                    fontSize: "1.1em",
                  }}
                >
                  {selectedItem}
                </Typography>
                {showList && (
                  <ul
                    style={{
                      listStyleType: "none",
                      background: "#f0f0f0",
                      padding: "10px",
                      borderRadius: "5px",
                      fontSize: "1.1em",
                    }}
                  >
                    <li
                      onClick={() => handleItemClick("購入済み")}
                      style={{
                        cursor: "pointer",
                        border: "1px solid #ccc",
                        padding: "8px",
                        margin: "5px",
                        borderRadius: "3px",
                      }}
                    >
                      購入済み
                    </li>
                    <li
                      onClick={() => handleItemClick("購入予定")}
                      style={{
                        cursor: "pointer",
                        border: "1px solid #ccc",
                        padding: "8px",
                        margin: "5px",
                        borderRadius: "3px",
                      }}
                    >
                      購入予定
                    </li>
                    <li
                      onClick={() => handleItemClick("完読済み")}
                      style={{
                        cursor: "pointer",
                        border: "1px solid #ccc",
                        padding: "8px",
                        margin: "5px",
                        borderRadius: "3px",
                      }}
                    >
                      完読済み
                    </li>
                    <li
                      onClick={() => handleItemClick("読書予定")}
                      style={{
                        cursor: "pointer",
                        border: "1px solid #ccc",
                        padding: "8px",
                        margin: "5px",
                        borderRadius: "3px",
                      }}
                    >
                      読書予定
                    </li>
                  </ul>
                )}
              </Grid>
              <Grid item xs={12}>
                <Grid container justifyContent="center">
                  <Button
                    variant="contained"
                    color="primary"
                    onClick={SubmitClick}
                  >
                    送信
                  </Button>
                </Grid>
              </Grid>
            </Grid>
          </Paper>
        </Container>
      </div>
    </main>
  );
};

export default EditInfo;
