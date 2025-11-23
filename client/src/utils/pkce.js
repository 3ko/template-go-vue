export function generateRandomString(length) {
  const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~";
  let result = "";
  const randomArray = new Uint8Array(length);
  crypto.getRandomValues(randomArray);
  randomArray.forEach((c) => (result += chars[c % chars.length]));
  return result;
}

export async function sha256base64url(input) {
  const enc = new TextEncoder();
  const bytes = await crypto.subtle.digest("SHA-256", enc.encode(input));
  let base64 = btoa(String.fromCharCode(...new Uint8Array(bytes)));
  return base64.replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/, "");
}
