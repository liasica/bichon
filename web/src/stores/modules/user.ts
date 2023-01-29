import { defineStore } from "pinia";
import Cookies from "js-cookie";
import { useAccount, useWalletSignature } from "@/utils/ethereum";
import { useMemberNonce, reqSignin, useMemberInfo } from "@/api";

export default defineStore("user", {
  state: () => ({
    token: Cookies.get("token"),
    address: Cookies.get("address") || "",
    id: "",
    nickname: "",
    intro: "",
    loginLoading: false,
  }),
  actions: {
    async login() {
      try {
        if (this.loginLoading) return;
        this.loginLoading = true;
        const accounts = await useAccount();
        if (!accounts?.length) return;
        const address = accounts[0];
        this.address = address;
        const nonce = await useMemberNonce(address);
        if (!nonce) return;
        const signature = await useWalletSignature(address, nonce);
        if (!signature) return;
        const token = await reqSignin({
          address,
          nonce,
          signature,
        });
        if (!token) return;
        this.token = token;
        Cookies.set("address", address);
        Cookies.set("token", token);
      } catch (error) {
      } finally {
        this.loginLoading = false;
      }
      return this.token;
    },
    logout() {
      Cookies.set("token", "");
      this.token = "";
      // back to home
    },
    async getMemberInfo() {
      const res = await useMemberInfo(this.address);
      if (res) {
        this.id = res.id;
        this.nickname = res.nickname || "";
        this.intro = res.intro || "";
      }
    },
  },
  getters: {
    isLogin(state) {
      return state.token && state.address;
    },
  },
});
