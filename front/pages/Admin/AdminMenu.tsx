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
} from "@mui/material";
import InboxIcon from "@mui/icons-material/Inbox";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import { AdminSidebar } from "@/component/Layout.tsx";
import { useRouter } from "next/router";
import{CautionComment,CheckToken}from "@/methods/util.tsx"
const inter = Inter({ subsets: ["latin"] });

const AdminMenu = () => {
  const router = useRouter();
  const [name, setName] = useState('');
  const [token, setToken] = useState('');

  useEffect(() => {
    // クエリパラメータを取得
    const { Name, Token } = router.query;
    if (Name && Token){
      setName(Name);
      setToken(Token);
      //トークン確認
      if(CheckToken(Name,Token)===false){
        //トークンが一致しない、または期限切れ
        router.push("/");
      }
      return
    }else{
      //Name,Tokenが無い
      router.push("/");
    }
  }, [router.query]);

  return (
    <main className={`${styles.main} ${inter.className}`}>
      <div>
        <AdminSidebar Name={name} Token={token}/>
        <Typography
          style={{
            fontFamily: "Arial, sans-serif",
            fontSize: "3rem",
            textAlign: "center",
            color:'red'
          }}
        >
          ***管理者画面***
        </Typography>
        
      </div>
    </main>
  );
};

export default AdminMenu;
