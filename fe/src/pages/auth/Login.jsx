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
        <div className="min-h-screen flex" style={{ backgroundColor: "#F6F6F6" }}>
            {/* Left Panel */}
            <div className="hidden lg:flex w-1/2 flex-col justify-between p-12" style={{ backgroundColor: "#303030" }}>
                <div>
                    <h1 className="text-2xl font-bold" style={{ color: "#CBE86A" }}>Collabify</h1>
                </div>
                <div>
                    <p className="text-4xl font-bold leading-tight" style={{ color: "#F6F6F6" }}>
                        Collaborate in <br />real-time with <br />
                        <span style={{ color: "#CBE86A" }}>your team.</span>
                    </p>
                </div>
                <p className="text-sm" style={{ color: "#9E9E9E" }}>
                    Built with Go + gRPC + WebSocket
                </p>
            </div>

            {/* Right Panel */}
            <div className="flex flex-1 items-center justify-center p-8">
                <div className="w-full max-w-md">
                    <h2 className="text-3xl font-bold mb-2" style={{ color: "#303030" }}>Welcome back</h2>
                    <p className="mb-8 text-sm" style={{ color: "#9E9E9E" }}>Sign in to continue collaborating</p>

                    {error && (
                        <div className="mb-4 p-3 rounded text-sm" style={{ backgroundColor: "#fee2e2", color: "#dc2626" }}>
                            {error}
                        </div>
                    )}

                    <div className="mb-4">
                        <label className="block text-sm font-medium mb-2" style={{ color: "#303030" }}>Email</label>
                        <input
                            type="email"
                            placeholder="you@example.com"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            className="w-full p-3 rounded-lg border outline-none transition"
                            style={{ backgroundColor: "#fff", borderColor: "#E4E4E4", color: "#303030" }}
                        />
                    </div>

                    <div className="mb-6">
                        <label className="block text-sm font-medium mb-2" style={{ color: "#303030" }}>Password</label>
                        <input
                            type="password"
                            placeholder="••••••••"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            className="w-full p-3 rounded-lg border outline-none transition"
                            style={{ backgroundColor: "#fff", borderColor: "#E4E4E4", color: "#303030" }}
                        />
                    </div>

                    <button
                        onClick={handleLogin}
                        disabled={loading}
                        className="w-full p-3 rounded-lg font-semibold transition"
                        style={{ backgroundColor: "#CBE86A", color: "#303030" }}
                    >
                        {loading ? "Signing in..." : "Sign in"}
                    </button>

                    <p className="mt-6 text-center text-sm" style={{ color: "#9E9E9E" }}>
                        Don't have an account?{" "}
                        <a href="/register" className="font-semibold" style={{ color: "#303030" }}>
                            Register
                        </a>
                    </p>
                </div>
            </div>
        </div>
    );
};

export default Login;