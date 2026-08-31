import { create } from "@bufbuild/protobuf";
import { useEffect, useRef, useState } from "react";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { PresenceRequestSchema, PresenceService, UserPresenceSchema } from "../proto/presence_pb";
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

const usePresence = (docId: string) => {
  const { token, user } = useAuthStore();
  const [users, setUsers] = useState<User[]>([]);
  const streamRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  useEffect(() => {
    if (!docId || !token || !user) return;

    const transport = createConnectTransport({
      baseUrl: (import.meta as any).env?.VITE_API_URL ?? "http://localhost:8080",
    });

    const client = createClient(PresenceService, transport);

    let cancelled = false;

    const start = async () => {
      try {
        const stream = client.syncPresence(
        (async function* () {
            yield create(PresenceRequestSchema, {
            docId,
            presence: create(UserPresenceSchema, {
                userId: user.userId,
                docId,
                name: user.name,
                isOnline: true,
            }),
            });
            await new Promise((_, reject) => {
            streamRef.current = setInterval(() => {
                if (cancelled) reject(new Error("cancelled"));
            }, 5000) as unknown as ReturnType<typeof setTimeout>;
            });
        })(),
        { headers: { Authorization: `Bearer ${token}` } }
        );

        for await (const res of stream) {
          if (cancelled) break;
          setUsers(res.users as User[]);
        }
      } catch {
        if (!cancelled) setTimeout(start, 3000); // reconnect
      }
    };

    start();

    return () => {
      cancelled = true;
      if (streamRef.current) clearInterval(streamRef.current as unknown as number);
    };
  }, [docId, token, user]);

  return { users, getColor };
};

export default usePresence;