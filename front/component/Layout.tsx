import React from "react";
import styles from "@/styles/Layout.module.css";
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
  Button,
} from "@mui/material";
import ListAltIcon from "@mui/icons-material/ListAlt";
import AddIcon from "@mui/icons-material/Add";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";
import LogoutIcon from "@mui/icons-material/Logout";
import { useRouter } from "next/router";
import { styled } from "@mui/system";
import { Inter } from "next/font/google";
const inter = Inter({ subsets: ["latin"] });

//赤ボタン
export const RedButton = styled(Button)(({ theme }) => ({
  backgroundColor: "red",
  "&:hover": {
    backgroundColor: "darkred",
  },
}));

//緑ボタン
export const GreenButton = styled(Button)(({ theme }) => ({
  backgroundColor: "green",
  "&:hover": {
    backgroundColor: "darkgreen",
  },
}));

//サイドバー
export const Sidebar = (UserDatas: string) => {
  const router = useRouter(); //ルーター

  //一覧ボタンクリック
  const ListClick = () => {
    router.push({
      pathname: "/MainMenu/BookList",
      query: { Name: UserDatas.Name, Token: UserDatas.Token },
    });
  };
  //追加ボタンクリック
  const AddClick = () => {
    router.push({
      pathname: "/MainMenu/Add",
      query: { Name: UserDatas.Name, Token: UserDatas.Token },
    });
  };

  //編集ボタンクリック
  const EditClick = () => {
    router.push({
      pathname: "/MainMenu/BookEdit",
      query: { Name: UserDatas.Name, Token: UserDatas.Token },
    });
  };

  //削除ボタンクリック
  const DeleteClick = () => {
    router.push({
      pathname: "/MainMenu/BookDelete",
      query: { Name: UserDatas.Name, Token: UserDatas.Token },
    });
  };
  //サインアウトクリック
  const ExitClick = () => {
    router.push("/");
  };
  return (
    <Drawer variant="permanent">
      <List>
        <Typography
          style={{
            fontFamily: "Arial, sans-serif",
            fontSize: "2.25rem",
            textAlign: "center",
          }}
        >
          BookManager
        </Typography>
        <Typography
          style={{
            fontFamily: "Arial, sans-serif",
            fontSize: "1.75rem",
            textAlign: "center",
          }}
        >
          {UserDatas.Name}様
        </Typography>
        <ListItem />
        <ListItem />
        <ListItem />
        <ListItem button onClick={ListClick}>
          <ListItemIcon>
            <ListAltIcon />
          </ListItemIcon>
          <Typography
            style={{
              fontFamily: "Arial, sans-serif",
              fontSize: "1.75rem",
            }}
          >
            一覧 (List)
          </Typography>
        </ListItem>
        <ListItem button onClick={AddClick}>
          <ListItemIcon>
            <AddIcon />
          </ListItemIcon>
          <Typography
            style={{
              fontFamily: "Arial, sans-serif",
              fontSize: "1.75rem",
            }}
          >
            追加 (Add)
          </Typography>
        </ListItem>
        <ListItem button onClick={EditClick}>
          <ListItemIcon>
            <EditIcon />
          </ListItemIcon>
          <Typography
            style={{
              fontFamily: "Arial, sans-serif",
              fontSize: "1.75rem",
            }}
          >
            編集 (Edit)
          </Typography>
        </ListItem>
        <ListItem button onClick={DeleteClick}>
          <ListItemIcon>
            <DeleteIcon />
          </ListItemIcon>
          <Typography
            style={{
              fontFamily: "Arial, sans-serif",
              fontSize: "1.75rem",
            }}
          >
            削除 (Delete)
          </Typography>
        </ListItem>
      </List>

      <ListItem button onClick={ExitClick}>
        <ListItemIcon>
          <LogoutIcon />
        </ListItemIcon>
        <Typography
          style={{
            fontFamily: "Arial, sans-serif",
            fontSize: "2rem",
          }}
        >
          サインアウト
        </Typography>
      </ListItem>
    </Drawer>
  );
};
