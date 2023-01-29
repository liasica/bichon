import { createFetch } from "@vueuse/core";
import crypto from "crypto-js";
import { useUserStore, useAppStore } from "@/stores";
import router from "@/router";
import { useWalletSignature } from "./ethereum";
import { useToast } from "vue-toastification";
import { ResponseError } from "@/types/errors";

const toast = useToast();

type ApiResponse<T> = {
  code: number;
  message: string;
  data: T;
};

export const useApiFetch = (url: string, sign?: boolean) => {
  const app = useAppStore();
  return createFetch({
    baseUrl: url.startsWith("http") ? "" : app.baseUrl,
    options: {
      async beforeFetch({ options }) {
        const user = useUserStore();
        if (user.token) {
          options.headers = {
            ...options.headers,
            Authorization: `Bearer ${user.token}`,
            ["X-Member-Address"]: user.address,
          };
        }

        const body = options.body?.toString();
        if (options.method === "POST" && sign && body) {
          const ts = Math.round(new Date().getTime() / 1000).toString();
          const str = crypto.MD5(body).toString() + ts;
          const signature = await useWalletSignature(user.address, str);
          console.info(`str: ${str}, signature: ${signature}`);
          options.headers = {
            ...options.headers,
            ["X-Signature"]: signature,
            ["X-Timestamp"]: ts,
          };
        }

        return { options };
      },
      afterFetch({ data, response }) {
        return { data, response };
      },
      onFetchError({ data, error }) {
        return { data, error };
      },
    },
    fetchOptions: {
      mode: "cors",
    },
  })(url);
};

// useGet usePost usePut useDelete

export const useGet = async <T>(url: string): Promise<any> => {
  const { data } = await useApiFetch(url).get().json<ApiResponse<T>>();
  const res = data.value;
  if (res?.code !== 200) {
    const user = useUserStore();

    switch (res?.code) {
      case 401:
        // redirect group list
        toast.clear();
        toast("Auth failed!");

        router.replace({ path: "/" });
        await user.login();

      default:
        toast.clear();
        toast(res?.message || "Error Encountered!");
        break;
    }
    const error = new ResponseError(res!.message, res!.code, res!.data);

    return Promise.reject(error);
  }
  return res!.data;
};

export const usePost = async <T, P = unknown>(
  url: string,
  payload: P,
  sign?: boolean
): Promise<any> => {
  const { data } = await useApiFetch(url, sign)
    .post(payload)
    .json<ApiResponse<T>>();
  const res = data.value;

  if (res?.code !== 200) {
    const user = useUserStore();

    switch (res?.code) {
      case 401:
        toast.clear();
        toast("Auth failed!");

        router.replace({ path: "/" });
        await user.login();

      default:
        toast.clear();
        toast(res?.message || "Error Encountered!");
        break;
    }
    const error = new ResponseError(res!.message, res!.code, res!.data);

    return Promise.reject(error);
  }
  return res!.data;
};
