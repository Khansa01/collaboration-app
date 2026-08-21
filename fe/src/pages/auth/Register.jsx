import { useState } from "react";
import { authClient } from "../../lib/client";

const Register = () => {
    const [name, setName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");

    const handleRegister = async () => {
        try {
            await authClient.register({ name, email, password });
            window.location.href = "/login";
        } catch (err) {
            setError("Registrasi gagal, coba lagi!");
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-900">
            <div className="bg-gray-800 p-8 rounded-lg w-96">
                <h1 className="text-white text-2xl font-bold mb-6">Register</h1>

                {error && <p className="text-red-400 mb-4">{error}</p>}

                <input
                    type="text"
                    placeholder="Nama"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    className="w-full p-3 rounded bg-gray-700 text-white mb-4"
                />

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
                    onClick={handleRegister}
                    className="w-full p-3 bg-purple-600 text-white rounded hover:bg-purple-700"
                >
                    Register
                </button>

                <p className="text-gray-400 mt-4 text-center">
                    Udah punya akun?{" "}
                    <a href="/login" className="text-purple-400">
                        Login
                    </a>
                </p>
            </div>
        </div>
    );
};

export default Register;