import { useGet, usePost } from "@/utils/request";

export const reqSignin = async (req: Model.SigninReq): Promise<string> => {
  const res = await usePost("/member", req);
  return res.token;
};

export const useMemberNonce = async (address: string): Promise<string> => {
  const res = await useGet(`/member/nonce/${address}`);
  return res.nonce;
};

export const useMemberInfo = async (address: string): Promise<any> => {
  const res = await useGet(`/member/${address}`);
  return res;
};

export const updateMemberInfo = async (
  req: Model.MemberInfoReq
): Promise<any> => {
  const res = await usePost("/member/update", req);
  return res;
};

export const useNodes = async (): Promise<string[]> => {
  const res = await useGet(`${import.meta.env.VITE_BASE_API_URL}/nodes`);
  return res;
};
