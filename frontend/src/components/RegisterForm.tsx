import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useRegister } from "../api/auth/useRegister";

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
        <form onSubmit={handleSubmit} className="max-w-2xl mx-auto p-6">
            <h1 className="text-2xl font-semibold mb-1 flex justify-center">Get started on NeoNedia</h1>
            <p className="flex justify-center p-5 mb-5"
            >Create an account to connect with friends, family and communities of people who share your interests.</p>
            {isError && <p className="text-red-500 mb-2">{(error as Error).message}</p>}
            <input
                type="text"
                placeholder="First Name"
                value={firstname}
                onChange={e => setFirstname(e.target.value)}
                className="w-full mb-4 p-4 border border-gray-300 rounded-2xl
                            focus:border-cyan-600 focus:ring-cyan-600 focus:ring-2"
            />
            <input
                type="text"
                placeholder="Last Name"
                value={lastname}
                onChange={e => setLastname(e.target.value)}
                className="w-full mb-4 p-4 border border-gray-300 rounded-2xl
                            focus:border-cyan-600 focus:ring-cyan-600 focus:ring-2"
            />
            <input
                type="email"
                placeholder="Email"
                value={email}
                onChange={e => setEmail(e.target.value)}
                className="w-full mb-4 p-4 border border-gray-300 rounded-2xl
                            focus:border-cyan-600 focus:ring-cyan-600 focus:ring-2"
            />
            <input
                type="text"
                placeholder="User Name"
                value={username}
                onChange={e => setUsername(e.target.value)}
                className="w-full mb-4 p-4 border border-gray-300 rounded-2xl
                            focus:border-cyan-600 focus:ring-cyan-600 focus:ring-2"
            />
            <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={e => setPassword(e.target.value)}
                className="w-full mb-4 p-4 border border-gray-300 rounded-2xl
                            focus:border-cyan-600 focus:ring-cyan-600 focus:ring-2"
            />
            <input
                type="date"
                placeholder="Date Of Birth"
                value={dateOfBirth}
                onChange={e => setDateOfBirth(e.target.value)}
                className="w-full mb-4 p-4 border border-gray-300 rounded-2xl
                            focus:border-cyan-600 focus:ring-cyan-600 focus:ring-2"
            />
            <select
                value={gender}
                onChange={e => setGender(e.target.value as "male" | "female" | "other")}
                className="w-full mb-4 p-4 border border-gray-300 rounded-2xl
                            focus:border-cyan-600 focus:ring-cyan-600 focus:ring-2"
            >
                <option value="male">Male</option>
                <option value="female">Female</option>
                <option value="other">Other</option>
            </select>
            <p 
                className="flex justify-center p-5 mb-8"
            >By tapping Register, you agree to create an account and acccept NeoMedia's terms and conditions.</p>
            <button
                type="submit"
                disabled={isPending}
                className="w-full mb-3 font-semibold bg-indigo-600 text-white py-2 rounded-full hover:bg-indigo-700"
            >
                {isPending ? "Registering Account..." : "Register"}
            </button>
            <button 
                type="button"
                onClick={() => navigate("/login")}
                className="w-full font-semibold bg-white text-cyan-600 border border-gray-300 py-2 rounded-full hover:bg-gray-300 mb-2"
            >
                I already have an account
            </button>
        </form>
    )
}