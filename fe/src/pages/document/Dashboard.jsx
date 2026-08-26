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
            setDocuments(res.documents);
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
        <div className="min-h-screen bg-gray-900 p-8">
            <div className="max-w-4xl mx-auto">
                <div className="flex justify-between items-center mb-8">
                    <h1 className="text-white text-3xl font-bold">My Documents</h1>
                    <button
                        onClick={logout}
                        className="text-gray-400 hover:text-white"
                    >
                        Logout
                    </button>
                </div>

                <div className="flex gap-4 mb-8">
                    <input
                        type="text"
                        placeholder="Nama dokumen baru..."
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        className="flex-1 p-3 rounded bg-gray-700 text-white"
                    />
                    <button
                        onClick={createDocument}
                        disabled={loading}
                        className="px-6 py-3 bg-purple-600 text-white rounded hover:bg-purple-700 disabled:opacity-50"
                    >
                        {loading ? "..." : "+ Buat Dokumen"}
                    </button>
                </div>

                <div className="grid gap-4">
                    {documents.length === 0 && (
                        <p className="text-gray-400 text-center py-8">
                            Belum ada dokumen — buat yang pertama!
                        </p>
                    )}
                    {documents.map((doc) => (
                        <div
                            key={doc.id}
                            onClick={() => navigate(`/document/${doc.id}`)}
                            className="bg-gray-800 p-6 rounded-lg cursor-pointer hover:bg-gray-700 transition"
                        >
                            <h2 className="text-white text-xl font-semibold">{doc.title}</h2>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Dashboard;