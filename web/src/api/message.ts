import { useGet, usePost } from "@/utils/request";

/**
 * 创建消息
 * @param req 消息信息(content需要使用shared key进行aes加密)
 * @returns 消息id和创建时间
 */
export const useMessageCreate = async (
  req: Model.ChatMessage
): Promise<Model.ChatMessage | undefined> =>
  usePost<Model.ChatMessageCreatRes, Model.ChatMessage>("/message", req);

/**
 * 获取群消息
 */
export const useMessageList = async (
  req: Model.ChatMessageListReq
): Promise<Model.ChatMessage[]> => {
  let query = "";
  if (req) {
    const searchParams = new URLSearchParams();
    const keys = Object.keys(req) as (keyof typeof req)[];
    keys.forEach((key) => searchParams.append(key, `${req[key]}`));
    query = `?${searchParams.toString()}`;
  }

  return useGet(`/message${query}`);
};
