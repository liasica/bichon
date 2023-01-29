declare namespace Model {
  type ChatMessage = {
    id?: string;
    memberId: string;
    groupId: string;
    content: string;
    createdAt?: string;
    member?: Model.Member;
    quote?: ChatMessage;
    index?: number;
    isFirst?: boolean;
    isLast?: boolean;
    isSelf?: boolean;
    isQuote?: boolean;
    parentId?: string;
  };

  type ChatMessageCreatRes = {
    id: string;
    createdAt: string;
  };

  type ChatMessageListReq = {
    groupId: string;
    time?: string;
  };
}
