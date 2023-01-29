import { useUserStore } from "@/stores";
import { GroupKeyDexie } from "./key";

export const useKeyDB = (): GroupKeyDexie | undefined => {
  const user = useUserStore();

  if (user.address) {
    const db = new GroupKeyDexie(user.address);
    return db;
  }
  return undefined;
};
