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
import MenuBookIcon from "@mui/icons-material/MenuBook";
import MenuIcon from "@mui/icons-material/Menu";
import AddIcon from "@mui/icons-material/Add";
import LibraryBooksIcon from "@mui/icons-material/LibraryBooks";
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

  //メニュー画面
  const MenuClick = () => {
    router.push({
      pathname: "/MainMenu",
      query: { Name: UserDatas.Name, Token: UserDatas.Token },
    });
  };

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
        <ListItem button onClick={MenuClick}>
          <ListItemIcon>
            <MenuIcon />
          </ListItemIcon>
          <Typography
            style={{
              fontFamily: "Arial, sans-serif",
              fontSize: "1.75rem",
            }}
          >
            メニュー画面
          </Typography>
        </ListItem>
        <ListItem button onClick={MenuClick}>
          <ListItemIcon>
            <LibraryBooksIcon />
          </ListItemIcon>
          <Typography
            style={{
              fontFamily: "Arial, sans-serif",
              fontSize: "1.75rem",
            }}
          >
            おすすめ技術書
          </Typography>
        </ListItem>
        <ListItem button onClick={ListClick}>
          <ListItemIcon>
            <MenuBookIcon />
          </ListItemIcon>
          <Typography
            style={{
              fontFamily: "Arial, sans-serif",
              fontSize: "1.75rem",
            }}
          >
            登録一覧 (List)
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
            登録 (Add)
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




//管理者サイドバー
export const AdminSidebar = (UserDatas: string) => {
  const router = useRouter(); //ルーター


  //一覧ボタンクリック
  const ListClick = () => {
    router.push({
      pathname: "/AdminMenu/AdminBookList",
      query: { Name: UserDatas.Name, Token: UserDatas.Token },
    });
  };
  //追加ボタンクリック
  const AddClick = () => {
    router.push({
      pathname: "/Admin/AdminMenu/AdminAddBook",
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
    <Drawer variant="permanent" PaperProps={{
      style: { backgroundColor: 'lightblue' }, // 背景色を設定
    }}>
      <List>
        <Typography
          style={{
            fontFamily: "Arial, sans-serif",
            fontSize: "2.25rem",
            textAlign: "center",
          }}
        >
          ***管理者画面***
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
            <MenuBookIcon />
          </ListItemIcon>
          <Typography
            style={{
              fontFamily: "Arial, sans-serif",
              fontSize: "1.75rem",
            }}
          >
            管理者本一覧
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
            管理者本登録
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
            管理者本編集
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
            管理者本削除
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