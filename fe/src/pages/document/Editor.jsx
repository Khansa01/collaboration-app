import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useEditor, EditorContent } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Collaboration from "@tiptap/extension-collaboration";
import * as Y from "yjs";
import { documentClient } from "../../lib/client";

const Editor = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [doc] = useState(() => new Y.Doc());
    const [title, setTitle] = useState("");
    const [saving, setSaving] = useState(false);

    const editor = useEditor({
        extensions: [
            StarterKit.configure({ history: false }),
            Collaboration.configure({ document: doc }),
        ],
        editorProps: {
            attributes: {
                class: "prose prose-invert max-w-none focus:outline-none min-h-screen p-8",
            },
        },
    });

    useEffect(() => {
        const fetchDoc = async () => {
            try {
                const res = await documentClient.getDocument({ id });
                setTitle(res.document?.title || "");
            } catch (err) {
                console.error(err);
            }
        };
        fetchDoc();
    }, [id]);

    const handleSave = async () => {
        if (!editor) return;
        try {
            setSaving(true);
            const content = JSON.stringify(editor.getJSON());
            await documentClient.updateDocument({
                id,
                content,
            });
            console.log("Saved!");
        } catch (err) {
            console.error("Error saving:", err);
        } finally {
            setSaving(false);
        }
    };

    if (!editor) return null;

    return (
        <div className="min-h-screen bg-gray-900">
            {/* Navbar */}
            <div className="flex items-center justify-between px-8 py-4 bg-gray-800 border-b border-gray-700">
                <div className="flex items-center gap-4">
                    <button
                        onClick={() => navigate("/dashboard")}
                        className="text-gray-400 hover:text-white"
                    >
                        ← Back
                    </button>
                    <h1 className="text-white font-semibold">{title}</h1>
                </div>
                <button
                    onClick={handleSave}
                    disabled={saving}
                    className="px-4 py-2 bg-purple-600 text-white rounded hover:bg-purple-700 disabled:opacity-50"
                >
                    {saving ? "Saving..." : "Save"}
                </button>
            </div>

            {/* Editor */}
            <div className="max-w-4xl mx-auto mt-8">
                <EditorContent editor={editor} className="text-white" />
            </div>
        </div>
    );
};

export default Editor;