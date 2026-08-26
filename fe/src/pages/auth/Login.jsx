import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { authClient } from "../../lib/client";
import useAuthStore from "../../store/authStore";

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);
    const { setToken } = useAuthStore();
    const navigate = useNavigate();

    const handleLogin = async () => {
        try {
            setLoading(true);
            setError("");
            const res = await authClient.login({ email, password });
            setToken(res.token);
            navigate("/dashboard");
        } catch (err) {
            setError("Email atau password salah!");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-900">
            <div className="bg-gray-800 p-8 rounded-lg w-96">
                <h1 className="text-white text-2xl font-bold mb-6">Login</h1>

                {error && <p className="text-red-400 mb-4">{error}</p>}

                <input
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    className="w-full p-3 rounded bg-gray-700 text-white mb-4"
                />

                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    className="w-full p-3 rounded bg-gray-700 text-white mb-6"
                />

                <button
                    onClick={handleLogin}
                    disabled={loading}
                    className="w-full p-3 bg-purple-600 text-white rounded hover:bg-purple-700 disabled:opacity-50"
                >
                    {loading ? "Loading..." : "Login"}
                </button>

                <p className="text-gray-400 mt-4 text-center">
                    Belum punya akun?{" "}
                    <a href="/register" className="text-purple-400">
                        Register
                    </a>
                </p>
            </div>
        </div>
    );
};

export default Login;