import { create } from "zustand";

interface User {
  userId: string;
  name: string;
}

interface AuthState {
  token: string | null;
  user: User | null;
  setToken: (token: string) => void;
  logout: () => void;
}

const decodeJWT = (token: string): User | null => {
  try {
    const payload = JSON.parse(atob(token.split(".")[1]));
    return {
      userId: payload.user_id,
      name: payload.name ?? payload.email ?? "Anonymous",
    };
  } catch {
    return null;
  }
};

const storedToken = localStorage.getItem("token");

const useAuthStore = create<AuthState>((set) => ({
  token: storedToken,
  user: storedToken ? decodeJWT(storedToken) : null,

  setToken: (token: string) => {
    localStorage.setItem("token", token);
    set({ token, user: decodeJWT(token) });
  },

  logout: () => {
    localStorage.removeItem("token");
    set({ token: null, user: null });
  },
}));

export default useAuthStore;