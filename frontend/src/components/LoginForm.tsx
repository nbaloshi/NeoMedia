import React, { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useLogin } from "../api/useLogin";
import { SessionContext } from "../context/SessionContext";

export default function LoginForm() {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const navigate = useNavigate()
    const { mutate, isPending, isError, error } = useLogin()

    const session = useContext(SessionContext);
    if (!session) throw new Error("SessionContext not found");

    const { setUser } = session;

    const handleSubmit =  async (e: React.SyntheticEvent<HTMLFormElement>) => {
        e.preventDefault()
        mutate(
            { email, password },
            {
                onSuccess: async () => {
                    const res = await fetch("http://localhost:8080/me", { credentials: "include" });
                    console.log(res)
                    const data = await res.json();
                    console.log(data)
                    setUser(data);
                    navigate("/home");
                }
            }
        )
    }

    return (
        <form onSubmit={handleSubmit} className="max-w-sm mx-auto p-6 bg-white shadow rounded">
            <h1 className="text-xl font-bold mb-4">Login</h1>
            {isError && <p className="text-red-500 mb-2">{(error as Error).message}</p>}
            <input 
                type="email"
                placeholder="Email"
                value={email}
                onChange={e => setEmail(e.target.value)}
                className="w-full mb-3 p-2 border rounded"
            />
            <input 
                type="password"
                placeholder="Password"
                value={password}
                onChange={e => setPassword(e.target.value)}
                className="w-full mb-3 p-2 border rounded"
            />
            <button 
                type="submit"
                disabled={isPending}
                className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700"
            >
                {isPending ? "Logging in..." : "Log in"}
            </button>
            <button 
                type="button"
                onClick={() => navigate("/register")}
                className="w-full bg-gray-600 text-white py-2 rounded hover:bg-gray-700 mt-2"
            >
                Create an account
            </button>
        </form>
    )
}