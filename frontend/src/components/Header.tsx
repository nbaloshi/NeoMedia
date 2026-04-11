import { useContext } from "react"
import { SessionContext } from "../context/SessionContext"

export default function Header() {
    const session = useContext(SessionContext);
    if (!session) throw new Error("SessionContext not found");

    const { user, logout } =  session;

    return (
        <header className="flex items-center justify-between px-6 py-4 bg-indigo-700">
            <div className="text-xl font-semibold text-white">
                {user ? `Welcome, ${user.username}` : null}
            </div>
            {/* <div className="text-lg font-bold">NeoMedia</div> */}
            <div>
                <img
                    src="/images/nm2.png"
                    alt="NM illustration"
                    className="w-60 h-auto"
                />
            </div>
            <button
                onClick={logout}
                className="w-28 font-semibold bg-white text-cyan-600 border border-gray-300 py-2 rounded-full hover:bg-gray-300 mb-2"
            >
                Logout
            </button>
        </header>
    )
}