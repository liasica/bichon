declare namespace Model {
  interface GroupCreateReq {
    name: string;
    category: string;
    maxMembers: number;
    intro?: string;
    sharedPublic: string; // 用户公钥
  }

  interface GroupDetail {
    id: string;
    address: string;
    category: string;
    intro: string;
    membersMax: number;
    name: string;
    owner: boolean;
    public: boolean;
    members?: Model.Member[];
    createdAt: string;
    unreadCount?: number;
    unreadId?: string;
    unreadTime?: string;
    membersCount?: number;
    inviteCode?: string;
  }

  interface GroupKeyReq {
    groupId: string;
    sharedPublic: string; // 用户公钥
  }

  type GroupShareKey = {
    groupPublicKey: string;
    keyId: string; // group key 存储标识
  };

  type GroupDetailWithPublicKey = GroupDetail & GroupShareKey;

  type GroupListReq = ApiPaginationReq & {
    name?: string;
    category?: string;
  };

  type GroupListRes = GroupMeta & {
    joined: boolean;
  };

  type GroupKeyUsedReq = {
    keys: GroupKeyUsed[];
  };

  interface UpdateGroupReq {
    category?: string;
    groupId: string;
    intro?: string;
    maxMembers?: number;
    name?: string;
    public?: boolean;
  }
}
