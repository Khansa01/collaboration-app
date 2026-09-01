import { useEffect, useRef, useState } from "react";
import useAuthStore from "../store/authStore";

interface User {
  userId: string;
  name: string;
  isOnline: boolean;
}

const COLORS = ["#CBE86A", "#F6A86A", "#6AB4F6", "#F66A8A", "#A86AF6"];

const getColor = (userId: string) => {
  const hash = userId.split("").reduce((acc, c) => acc + c.charCodeAt(0), 0);
  return COLORS[hash % COLORS.length];
};

const WS_URL = (import.meta as any).env?.VITE_WS_URL ?? "ws://localhost:8080";

const usePresence = (docId: string) => {
  const { token, user } = useAuthStore();
  const [users, setUsers] = useState<User[]>([]);
  const wsRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    if (!docId || !token || !user) return;

    const connect = () => {
      const ws = new WebSocket(`${WS_URL}/presence/${docId}?token=${token}`);
      wsRef.current = ws;

      ws.onopen = () => {
        ws.send(JSON.stringify({
          type: "join",
          userId: user.userId,
          name: user.name,
        }));
      };

      ws.onmessage = (e) => {
        try {
          const data = JSON.parse(e.data);
          if (data.type === "presence") {
            setUsers(data.users);
          }
        } catch {}
      };

      ws.onclose = () => {
        setTimeout(connect, 3000);
      };
    };

    connect();

    return () => {
      wsRef.current?.close();
    };
  }, [docId, token, user]);

  return { users, getColor };
};

export default usePresence;