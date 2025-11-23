import { generateRandomString, sha256base64url } from "./pkce";

export function createOIDCHandler({ issuer, clientId, redirectUri, scope }) {
  return {
    async createAuthUrl() {
      const state = generateRandomString(32);
      const codeVerifier = generateRandomString(64);
      const codeChallenge = await sha256base64url(codeVerifier);

      sessionStorage.setItem("pkce_verifier", codeVerifier);
      sessionStorage.setItem("state", state);

      const url =
        `${issuer}/oauth/v2/authorize?response_type=code` +
        `&client_id=${encodeURIComponent(clientId)}` +
        `&redirect_uri=${encodeURIComponent(redirectUri)}` +
        `&scope=${encodeURIComponent(scope)}` +
        `&code_challenge=${codeChallenge}` +
        `&code_challenge_method=S256` +
        `&state=${state}`;

      return url;
    },

    async processCallback() {
      const params = new URLSearchParams(window.location.search);
      const code = params.get("code");
      const returnedState = params.get("state");

      const storedState = sessionStorage.getItem("state");
      const codeVerifier = sessionStorage.getItem("pkce_verifier");

      if (!code || !returnedState || returnedState !== storedState) {
        throw new Error("Invalid state or code");
      }

      const body = new URLSearchParams({
        grant_type: "authorization_code",
        code,
        client_id: clientId,
        redirect_uri: redirectUri,
        code_verifier: codeVerifier
      });

      const tokenRes = await fetch(`${issuer}/oauth/v2/token`, {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body
      }).then((r) => r.json());

      const userinfo = await fetch(`${issuer}/oidc/v1/userinfo`, {
        headers: {
          Authorization: `Bearer ${tokenRes.access_token}`
        }
      }).then((r) => r.json());

      return {
        accessToken: tokenRes.access_token,
        idToken: tokenRes.id_token,
        user: userinfo
      };
    },

    async logout() {
      window.location.replace(`${issuer}/oidc/v1/end_session`);
    }
  };
}
