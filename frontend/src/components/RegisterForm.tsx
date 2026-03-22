import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useRegister } from "../api/useRegister";

export default function RegisterForm() {
    const [firstname, setFirstname] = useState("")
    const [lastname, setLastname] = useState("")
    const [email, setEmail] = useState("")
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [dateOfBirth, setDateOfBirth] = useState("")
    const [gender, setGender] = useState<"male" | "female" | "other">("male")
    const navigate = useNavigate()
    const { mutate, isPending, isError, error } = useRegister()

    const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
        e.preventDefault()
        mutate(
            { firstname, lastname, email, username, password, dateOfBirth, gender },
            {
                onSuccess: () => navigate("/login")
            }
        )
    }

    return (
        <form onSubmit={handleSubmit} className="max-w-sm mx-auto p-6 bg-white shadow rounded">
            <h1 className="text-xl font-bold mb-4">Register</h1>
            {isError && <p className="text-red-500 mb-2">{(error as Error).message}</p>}
            <input
                type="text"
                placeholder="First Name"
                value={firstname}
                onChange={e => setFirstname(e.target.value)}
                className="w-full mb-3 p-2 border rounded"
            />
            <input
                type="text"
                placeholder="Last Name"
                value={lastname}
                onChange={e => setLastname(e.target.value)}
                className="w-full mb-3 p-2 border rounded"
            />
            <input
                type="email"
                placeholder="Email"
                value={email}
                onChange={e => setEmail(e.target.value)}
                className="w-full mb-3 p-2 border rounded"
            />
            <input
                type="text"
                placeholder="User Name"
                value={username}
                onChange={e => setUsername(e.target.value)}
                className="w-full mb-3 p-2 border rounded"
            />
            <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={e => setPassword(e.target.value)}
                className="w-full mb-3 p-2 border rounded"
            />
            <input
                type="date"
                placeholder="Date Of Birth"
                value={dateOfBirth}
                onChange={e => setDateOfBirth(e.target.value)}
                className="w-full mb-3 p-2 border rounded"
            />
            <select
                value={gender}
                onChange={e => setGender(e.target.value as "male" | "female" | "other")}
                className="w-full mb-3 p-2 border rounded"
            >
                <option value="male">Male</option>
                <option value="female">Female</option>
                <option value="other">Other</option>
            </select>
            <button
                type="submit"
                disabled={isPending}
                className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700"
            >
                {isPending ? "Registering Account..." : "Register"}
            </button>
            <button 
                type="button"
                onClick={() => navigate("/login")}
                className="w-full bg-gray-600 text-white py-2 rounded hover:bg-gray-700 mt-2"
            >
                I already have an account
            </button>
        </form>
    )
}