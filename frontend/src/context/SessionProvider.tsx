import { useEffect, useState, type ReactNode } from "react";
import type { User } from "./SessionContext";
import{ SessionContext } from "./SessionContext";

interface SessionProviderProps {
    children: ReactNode
}

export function SessionProvider({ children }: SessionProviderProps) {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetch("http://localhost:8080/me", { credentials: "include" })
            .then(res => res.ok ? res.json() : null)
            .then(data => {
                setUser(data);
                setLoading(false);
            })
            .catch(() => setLoading(false));
    }, []);

    const logout = async () => {
        await fetch("http://localhost:8080/logout", {
            method: "POST",
            credentials: "include",
        });
        setUser(null);
    };

    return (
        <SessionContext.Provider value={{ user, setUser, logout, loading }}>
            {children}
        </SessionContext.Provider>
    )
}