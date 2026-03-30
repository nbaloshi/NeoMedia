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
        <form onSubmit={handleSubmit} className="w-9/12">
            <h1 className="text-lg font-semibold mb-7">Log in to your account</h1>
            {isError && <p className="text-red-500 mb-2">{(error as Error).message}</p>}
            <input 
                type="email"
                placeholder="Email"
                value={email}
                onChange={e => setEmail(e.target.value)}
                className="w-full mb-4 p-4 border border-gray-300 rounded-2xl
                             focus:border-cyan-600 focus:ring-cyan-600 focus:ring-2"
            />
            <input 
                type="password"
                placeholder="Password"
                value={password}
                onChange={e => setPassword(e.target.value)}
                className="w-full mb-6 p-4 border border-gray-300 rounded-2xl
                            focus:border-cyan-600 focus:ring-cyan-600 focus:ring-2"
            />
            
            <button 
                type="submit"
                disabled={isPending}
                className="w-full mb-3 font-semibold bg-indigo-600 text-white py-2 rounded-full hover:bg-indigo-700"
            >
                {isPending ? "Logging in..." : "Log in"}
            </button>
            <p className="flex justify-center p-5 mb-8 font-semibold">Forgot password? Too bad.</p>
            <button 
                type="button"
                onClick={() => navigate("/register")}
                className="w-full font-semibold bg-white text-cyan-600 border border-gray-300 py-2 rounded-full hover:bg-gray-300 mb-2"
            >
                Create an account
            </button>
            <p className="flex justify-center font-semibold text-indigo-600">NeoMedia</p>
        </form>
    )
}