import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { documentClient } from "../../lib/client";
import useAuthStore from "../../store/authStore";

const Dashboard = () => {
    const [documents, setDocuments] = useState([]);
    const [title, setTitle] = useState("");
    const [loading, setLoading] = useState(false);
    const { token, logout } = useAuthStore();
    const navigate = useNavigate();

    const fetchDocuments = async () => {
        try {
            const res = await documentClient.listDocuments({});
            setDocuments(res.documents || []);
        } catch (err) {
            console.error(err);
        }
    };

    const createDocument = async () => {
        if (!title.trim()) return;
        try {
            setLoading(true);
            const res = await documentClient.createDocument({ title });
            navigate(`/document/${res.document.id}`);
        } catch (err) {
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        if (!token) {
            navigate("/login");
            return;
        }
        fetchDocuments();
    }, [token]);

    return (
        <div className="min-h-screen" style={{ backgroundColor: "#F6F6F6" }}>
            {/* Navbar */}
            <div className="px-4 md:px-8 py-4 flex items-center justify-between" style={{ backgroundColor: "#303030" }}>
                <h1 className="text-xl font-bold" style={{ color: "#CBE86A" }}>
                    Collabify
                </h1>
                <button
                    onClick={logout}
                    className="text-sm font-medium px-4 py-2 rounded-lg transition"
                    style={{ color: "#CBE86A" }}
                >
                    Sign out
                </button>
            </div>

            {/* Content */}
            <div className="max-w-4xl mx-auto px-4 md:px-8 py-8 md:py-12">
                <h2 className="text-2xl md:text-3xl font-bold mb-2" style={{ color: "#303030" }}>My Documents</h2>
                <p className="mb-8 text-sm" style={{ color: "#9E9E9E" }}>Create and collaborate on documents in real-time</p>

                {/* Create Document */}
                <div className="flex flex-col sm:flex-row gap-3 mb-10">
                    <input
                        type="text"
                        placeholder="New document name..."
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        onKeyDown={(e) => e.key === "Enter" && createDocument()}
                        className="flex-1 p-3 rounded-lg border outline-none w-full"
                        style={{ backgroundColor: "#fff", borderColor: "#E4E4E4", color: "#303030" }}
                    />
                    <button
                        onClick={createDocument}
                        disabled={loading}
                        className="w-full sm:w-auto px-6 py-3 rounded-lg font-semibold transition disabled:opacity-50"
                        style={{ backgroundColor: "#CBE86A", color: "#303030" }}
                    >
                        {loading ? "..." : "+ New"}
                    </button>
                </div>

                {/* Document List */}
                {documents.length === 0 ? (
                    <div className="text-center py-20">
                        <p className="text-4xl mb-4">📄</p>
                        <p className="font-medium mb-1" style={{ color: "#303030" }}>No documents yet</p>
                        <p className="text-sm" style={{ color: "#9E9E9E" }}>Create your first document above</p>
                    </div>
                ) : (
                    <div className="grid gap-3">
                        {documents.map((doc) => (
                            <div
                                key={doc.id}
                                onClick={() => navigate(`/document/${doc.id}`)}
                                className="p-5 rounded-xl border cursor-pointer transition hover:shadow-md"
                                style={{ backgroundColor: "#fff", borderColor: "#E4E4E4", borderLeft: "4px solid #CBE86A" }}
                            >
                                <h3 className="font-semibold" style={{ color: "#303030" }}>{doc.title}</h3>
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    );
};

export default Dashboard;