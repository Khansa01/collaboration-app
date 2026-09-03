import { useCallback } from "react";

const ToolbarButton = ({ onClick, active, children, title }) => (
    <button
        onClick={onClick}
        title={title}
        className="px-2 py-1 rounded text-sm font-medium transition"
        style={{
            backgroundColor: active ? "#CBE86A" : "transparent",
            color: active ? "#303030" : "#F6F6F6",
        }}
    >
        {children}
    </button>
);

const Toolbar = ({ editor }) => {
    if (!editor) return null;

    const addImage = useCallback(() => {
        const url = window.prompt("Masukkan URL gambar:");
        if (url) editor.chain().focus().setImage({ src: url }).run();
    }, [editor]);

    return (
        <div
            className="flex flex-wrap items-center gap-1 px-4 py-2 border-b"
            style={{ backgroundColor: "#303030", borderColor: "#444" }}
        >
            <ToolbarButton
                onClick={() => editor.chain().focus().toggleBold().run()}
                active={editor.isActive("bold")}
                title="Bold"
            >
                <b>B</b>
            </ToolbarButton>

            <ToolbarButton
                onClick={() => editor.chain().focus().toggleItalic().run()}
                active={editor.isActive("italic")}
                title="Italic"
            >
                <i>I</i>
            </ToolbarButton>

            <ToolbarButton
                onClick={() => editor.chain().focus().toggleUnderline().run()}
                active={editor.isActive("underline")}
                title="Underline"
            >
                <u>U</u>
            </ToolbarButton>

            <div style={{ width: 1, height: 20, backgroundColor: "#555", margin: "0 4px" }} />

            <ToolbarButton
                onClick={() => editor.chain().focus().toggleHeading({ level: 1 }).run()}
                active={editor.isActive("heading", { level: 1 })}
                title="Heading 1"
            >
                H1
            </ToolbarButton>

            <ToolbarButton
                onClick={() => editor.chain().focus().toggleHeading({ level: 2 }).run()}
                active={editor.isActive("heading", { level: 2 })}
                title="Heading 2"
            >
                H2
            </ToolbarButton>

            <ToolbarButton
                onClick={() => editor.chain().focus().toggleHeading({ level: 3 }).run()}
                active={editor.isActive("heading", { level: 3 })}
                title="Heading 3"
            >
                H3
            </ToolbarButton>

            <div style={{ width: 1, height: 20, backgroundColor: "#555", margin: "0 4px" }} />

            <ToolbarButton
                onClick={() => editor.chain().focus().toggleBulletList().run()}
                active={editor.isActive("bulletList")}
                title="Bullet List"
            >
                • List
            </ToolbarButton>

            <ToolbarButton
                onClick={() => editor.chain().focus().toggleOrderedList().run()}
                active={editor.isActive("orderedList")}
                title="Numbered List"
            >
                1. List
            </ToolbarButton>

        </div>
    );
};

export default Toolbar;