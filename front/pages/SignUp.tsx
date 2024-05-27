import React, { useState } from "react";
import {
  Container,
  Paper,
  Grid,
  TextField,
  Button,
  Typography,
  IconButton,
  InputAdornment,
} from "@mui/material";
import { useRouter } from "next/router";
import { Visibility, VisibilityOff } from "@mui/icons-material";
import { Inter } from "next/font/google";
import styles from "@/styles/Home.module.css";
import{PasswordChecker,CautionComment,EscapeInput,CheckNormalLength}from "@/methods/util.tsx"
import {SubmitSignUp} from "./api/Sign.ts";
const inter = Inter({ subsets: ["latin"] });

export default function SignuP() {
 
  const [showPassword, setShowPassword] = useState(false);//パスワード可視性フック
  const [Password, setPassword] = useState("");//新規パスワード
  const [CheckPassword, setCheckPassword] = useState("");//新規パスワード
  const [UserName, setUserName] = useState("");//新規名前
  const CautionID="caution";//注意ID
  const router = useRouter();//ルーター

  //パスワード可視性
  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };
  //パスワード可視性
  const handleMouseDownPassword = (event) => {
    event.preventDefault();
  };


  //送信ボタンクリック
  const SubmitClick=async()=>{
    //パスワードが確認と同じか
    if (Password!==CheckPassword){
      //注意を促す
      CautionComment(CautionID,"確認パスワードが違います");
      return;
    }

    //長さチェック
    if(CheckNormalLength(UserName)!=true||CheckNormalLength(Password)!=true){
      CautionComment(CautionID,"入力が長すぎます(20字以内)");
      return;
    }

    //エスケープ処理
    let EscapedName=EscapeInput(UserName);
    let EscapedPassword=EscapeInput(Password);

    //パスワードが脆弱じゃないか
    if(PasswordChecker(EscapedPassword)===false){
      CautionComment(CautionID,"パスワードが脆弱です");
      return;
    }

    //サーバーの返答を待つ
    let ServerResult=await SubmitSignUp(EscapedName,EscapedPassword);

    //Nameが重複していないか&サーバーからの返答が200か?
    if (ServerResult===403){
      CautionComment(CautionID,"名前が重複しています。");
      return;
    }
    //Nameが重複していないか&サーバーからの返答が200か?
    if (ServerResult!==201){
      CautionComment(CautionID,"サーバー処理エラー");
      return;
    }
    //サインアップ成功
    //メニュー画面へ遷移
    CautionComment(CautionID,"サインアップ成功!!少々お待ちください。");
    router.push("/SignIn");
  }

  return (
    <main className={`${styles.main} ${inter.className}`}>
    <div>
      <Container maxWidth="md" sx={{ padding: 4 }}>
        <Paper sx={{ padding: 8, marginTop: 4, boxShadow: 8 }}>
          <Typography
            variant="h4"
            align="center"
            gutterBottom
            style={{ fontFamily: "Roboto, sans-serif" }}
          >
            サインアップ
          </Typography>
          <Typography sx={{ color: "red", fontSize: "1.25em" }}>
          <span id="caution">
          </span>
        </Typography>
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="新規名前"
                variant="outlined"
                style={{ fontFamily: "Arial, sans-serif" }}
                onChange={(e) => {
                  setUserName(e.target.value);
                }}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="新規パスワード"
                type={showPassword ? "text" : "password"}
                variant="outlined"
                style={{ fontFamily: "Arial, sans-serif" }}
                autoComplete="new-password"
                onChange={(e) => {
                  setPassword(e.target.value);
                }}
                InputProps={{
                  endAdornment: (
                    <InputAdornment position="end">
                      <IconButton
                        aria-label="パスワードの表示を切り替える"
                        onClick={handleClickShowPassword}
                        onMouseDown={handleMouseDownPassword}
                        edge="end"
                      >
                        {showPassword ? <VisibilityOff /> : <Visibility />}
                      </IconButton>
                    </InputAdornment>
                  ),
                }}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                fullWidth
                label="新規パスワード(確認)"
                type={showPassword ? "text" : "password"}
                variant="outlined"
                style={{ fontFamily: "Arial, sans-serif" }}
                autoComplete="new-password"
                onChange={(e) => {
                  setCheckPassword(e.target.value);
                }}
                InputProps={{
                  endAdornment: (
                    <InputAdornment position="end">
                      <IconButton
                        aria-label="パスワードの表示を切り替える"
                        onClick={handleClickShowPassword}
                        onMouseDown={handleMouseDownPassword}
                        edge="end"
                      >
                        {showPassword ? <VisibilityOff /> : <Visibility />}
                      </IconButton>
                    </InputAdornment>
                  ),
                }}
              />
            </Grid>
            <Grid item xs={12}>
              <Grid container justifyContent="center">
                <Button variant="contained" color="primary" onClick={SubmitClick}>
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
}
