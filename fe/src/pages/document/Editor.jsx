import { useEffect, useState, useRef } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useEditor, EditorContent } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Collaboration from "@tiptap/extension-collaboration";
import * as Y from "yjs";
import { WebsocketProvider } from "y-websocket";
import { documentClient } from "../../lib/client";
import useAuthStore from "../../store/authStore";
import PresenceAvatars from "../../components/PresenceAvatars";

const WS_URL = import.meta.env.VITE_WS_URL ?? "ws://localhost:8080";

const Editor = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const { token } = useAuthStore();
    const [title, setTitle] = useState("");
    const [saving, setSaving] = useState(false);
    const [connected, setConnected] = useState(false);
    const ydocRef = useRef(new Y.Doc());
    const providerRef = useRef(null);

    useEffect(() => {
        const ydoc = ydocRef.current;

        const provider = new WebsocketProvider(
            `${WS_URL}/ws`,
            id,
            ydoc,
            {
                params: { token: token || "" },
                connect: true,
                resyncInterval: 5000,   // resync tiap 5 detik kalau ada gap
                maxBackoffTime: 10000,  // max reconnect delay 10 detik
            }
        );

        providerRef.current = provider;

        provider.on("status", ({ status }) => {
            setConnected(status === "connected");
        });

        const fetchDoc = async () => {
            try {
                const res = await documentClient.getDocument({ id });
                setTitle(res.document?.title || "");
            } catch (err) {
                console.error(err);
            }
        };

        fetchDoc();

        return () => {
            provider.destroy();
        };
    }, [id]);

    const editor = useEditor({
        extensions: [
            StarterKit.configure({ history: false }),
            Collaboration.configure({ document: ydocRef.current }),
        ],
        editorProps: {
            attributes: {
                class: "outline-none min-h-screen p-12 max-w-3xl mx-auto",
                style: "color: #303030; font-size: 16px; line-height: 1.8;",
            },
        },
    });

    useEffect(() => {
        if (!editor || !id) return;

        const loadContent = async () => {
            try {
                const res = await documentClient.getDocument({ id });
                setTitle(res.document?.title || "");
                if (res.document?.content) {
                    const content = JSON.parse(res.document.content);
                    editor.commands.setContent(content);
                }
            } catch (err) {
                console.error(err);
            }
        };

        loadContent();
    }, [editor, id]);

    if (!editor) return null;

    const handleSave = async () => {
        if (!editor) return;
        try {
            setSaving(true);
            const content = JSON.stringify(editor.getJSON());
            await documentClient.updateDocument({ id, content });
        } catch (err) {
            console.error("Error saving:", err);
        } finally {
            setSaving(false);
        }
    };

    if (!editor) return null;

    return (
        <div className="min-h-screen" style={{ backgroundColor: "#F6F6F6" }}>
            {/* Navbar */}
            <div className="px-8 py-4 flex items-center justify-between sticky top-0 z-10" style={{ backgroundColor: "#303030" }}>
                <div className="flex items-center gap-4">
                    <button
                        onClick={() => navigate("/dashboard")}
                        className="text-sm transition"
                        style={{ color: "#9E9E9E" }}
                    >
                        ← Back
                    </button>
                    <span style={{ color: "#9E9E9E" }}>|</span>
                    <h1 className="font-semibold" style={{ color: "#F6F6F6" }}>{title}</h1>
                </div>

                <div className="flex items-center gap-4">
                    <PresenceAvatars docId={id} />
                    <span
                        className="text-xs px-3 py-1 rounded-full font-medium"
                        style={{
                            backgroundColor: connected ? "#CBE86A" : "#E4E4E4",
                            color: connected ? "#303030" : "#9E9E9E"
                        }}
                    >
                        {connected ? "● Live" : "○ Offline"}
                    </span>

                    <button
                        onClick={handleSave}
                        disabled={saving}
                        className="px-4 py-2 rounded-lg text-sm font-semibold transition disabled:opacity-50"
                        style={{ backgroundColor: "#CBE86A", color: "#303030" }}
                    >
                        {saving ? "Saving..." : "Save"}
                    </button>
                </div>
            </div>

            {/* Editor */}
            <div style={{ backgroundColor: "#fff", minHeight: "calc(100vh - 60px)" }}>
                <EditorContent editor={editor} />
            </div>
        </div>
    );
};

export default Editor;