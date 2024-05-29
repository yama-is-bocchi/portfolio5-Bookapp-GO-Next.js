import { BaseUrl } from "./../common";
import { DeleteBook } from "@/structs/interface.tsx";

//本登録要求関数return200でOK
export async function AdminBookSubmitRegisterBook(
  Name: string,
  Title: string,
  Kind: string,
  Price: number,
  Memo: string,
  Token: string
) {
  const url = new URL("admin/book_resgister", BaseUrl);
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
        price:Price,
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