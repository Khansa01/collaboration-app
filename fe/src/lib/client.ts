import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { AuthService } from "../proto/auth_pb";
import { DocumentService } from "../proto/document_pb";
import useAuthStore from "../store/authStore";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const authInterceptor = (next: any) => async (req: any) => {
  const token = useAuthStore.getState().token;
  if (token) {
    req.header.set("Authorization", `Bearer ${token}`);
  }
  return await next(req);
};

const getTransport = () => createConnectTransport({
  baseUrl: import.meta.env.VITE_API_URL ?? "http://localhost:8080",
  interceptors: [authInterceptor],
});

export const authClient = createClient(AuthService, getTransport());
export const documentClient = createClient(DocumentService, getTransport());