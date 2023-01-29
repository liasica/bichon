import { Buffer } from "buffer";

export const useWalletSignature = async (
  address: string,
  data: string
): Promise<string> => {
  const buff = Buffer.from(data, "utf-8");
  const signature = await window.ethereum.request<string>({
    method: "personal_sign",
    params: [buff.toString("hex"), address],
  });
  return signature ?? "";
};

export const useAccount = async (): Promise<string[]> => {
  const accounts = (await window.ethereum.request<string[]>({
    method: "eth_requestAccounts",
  })) as string[];
  return accounts;
};
