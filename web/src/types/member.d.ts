declare namespace Model {
  interface NonceRes {
    nonce: string;
  }

  interface SigninReq {
    address: string;
    signature: string;
    nonce: string;
  }

  interface SigninRes {
    token: string;
  }

  interface Member {
    address: string;
    avatar: string;
    id: string;
    nickname: string;
  }

  interface MemberInfoReq {
    nickname: string;
    intro: string;
  }
}
