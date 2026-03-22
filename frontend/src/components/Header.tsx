import { useContext } from "react"
import { SessionContext } from "../context/SessionContext"

export default function Header() {
    const session = useContext(SessionContext);
    if (!session) throw new Error("SessionContext not found");

    const { user, logout } =  session;

    return (
        <header className="flex items-center justify-between px-6 py-4 bg-white shadow">
            <div className="text-sm font-medium">
                {user ? `Welcome, ${user.username}` : null}
            </div>
            <div className="text-lg font-bold">NeoMedia</div>
            <button
                onClick={logout}
                className="text-sm text-red-500 hover:underline"
            >
                Logout
            </button>
        </header>
    )
}