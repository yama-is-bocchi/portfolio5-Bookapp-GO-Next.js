import { BaseUrl} from './../common';

//サインイン要求関数return200でOK
export async function AdminSubmitSignIn(Name: string,Password:string) {
    const url = new URL("admin/sign_in", BaseUrl);
    return new Promise((resolve) => {
      fetch(url.href, {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: Name,
          password: Password
        }),
      }).then( res =>res.json())
      .then(data=>resolve(data));
    });
  }
  
  