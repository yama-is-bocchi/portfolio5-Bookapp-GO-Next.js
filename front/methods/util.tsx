import { SubmitTokenCheck } from "@/pages/api/util.ts";

//定数
const NormalLengthLimit = 20;
const MemoLengthLimit = 40;
const specialChars: { [key: string]: string } = {
  '&': '',
  '<': '',
  '>': '',
  '"': '',
  "'": '',
  '/': '',
  '\\': '',
  '`': '',
  '=': '',
  '$': '',
  ';': '',
  ':': '',
  '(': '',
  ')': '',
  '{': '',
  '}': '',
  '[': '',
  ']': '',
  '|': '',
  '^': '',
  '%': '',
  '~': '',
  '#': '',
  '!': '',
  '@': '',
  '*': '',
  '?': '',
  '+': '',
};// 置き換える特殊文字のリストを定義


// パスワードの脆弱性確認
export function PasswordChecker(input: string): boolean {
  // 英字と数字が含まれているかチェック
  const hasLetter = /[a-zA-Z]/.test(input);
  const hasNumber = /[0-9]/.test(input);
  // 長さが5以上かチェック
  const lengthValid = input.length >= 5;

  // 英字と数字が含まれており、かつ長さが5以上の場合にTrueを返す
  return hasLetter && hasNumber && lengthValid;
}

//引数のIDを持つ要素に注意を促すテキスト挿入する
export function CautionComment(ID: string, comment: string): void {
  let element = document.getElementById(ID);
  if (element == null) return;
  element.textContent = comment;
  return;
}

//トークンチェックを非同期処理で行う
//trueで正常
export async function CheckToken(Name: string, Token: string): boolean {
  try {
    if ((await SubmitTokenCheck(Name, Token)) === 200) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    // 非同期処理中に発生したエラーをキャッチ
    return false;
  }
}

//ユーザーの通常の入力に対して長さ測る
//TrueでOK
export function CheckNormalLength(Input: string): boolean {
  if (Input.length <= NormalLengthLimit) {
    return true;
  } 
}


//ユーザーのメモの入力に対して長さ測る
//TrueでOK
export function CheckMemoLength(Input: string): boolean {
  if (Input.length <= MemoLengthLimit) {
    return true;
  } 
}

//ユーザーの入力の特殊文字をエスケープする
export function EscapeInput(Input: string): string {
  // 正規表現を使って特殊文字を置き換える
  return Input.replace(/[&<>"'\/\\`=$;:(){}[\]|^%~#@!*?+]/g, (char) => specialChars[char] || '');
}
