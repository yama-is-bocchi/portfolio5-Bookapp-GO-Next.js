//サインイン時に帰ってくるJSONデータ
interface SignInUser {
    Name: string;
    Token: string;
    error: string;
}

//削除するためにAPIサーバーに送信する本データ
interface DeleteBook{
    Name:string;
    Title:string;
    Kind:string;
    Memo:string;
    Token:string;
}