import { BaseUrl } from "./common";
import { DeleteBook } from "@/structs/interface.tsx";

//本登録要求関数return200でOK
export async function SubmitRegisterBook(
  Name: string,
  Title: string,
  Kind: string,
  Memo: string,
  Token: string
) {
  const url = new URL("books/resgister", BaseUrl);
  return new Promise((resolve) => {
    fetch(url.href, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: Name,
        title: Title,
        kind: Kind,
        memo: Memo,
        token: Token,
      }),
    }).then((res) => resolve(res.status));
  });
}

//本一覧関数return200でOK
export async function GetBookList(Name: string, Token: string) {
  const url = new URL("books/get_list", BaseUrl);
  return new Promise((resolve) => {
    fetch(url.href, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: Name,
        token: Token,
      }),
    })
      .then((res) => res.json())
      .then((data) => resolve(data));
  });
}

//本の更新処理関数
export async function UpdateBookapi(
  Pretitle: string,
  Name: string,
  Title: string,
  Kind: string,
  Memo: string,
  Token: string
) {
  const url = new URL("books/update", BaseUrl);
  return new Promise((resolve) => {
    fetch(url.href, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        pretitle:Pretitle,
        name:Name,
        title:Title,
        kind:Kind,
        memo:Memo,
        token:Token
      }),
    }).then((res) => resolve(res.status));
  });
}

//本の選択削除
export async function DeleteBookapi(DeleteBooks: DeleteBook[]) {
  const url = new URL("books/delete", BaseUrl);
  return new Promise((resolve) => {
    fetch(url.href, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify(DeleteBooks),
    }).then((res) => resolve(res.status));
  });
}
