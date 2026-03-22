import React, { createContext } from "react";

export interface User {
    id: string;
    email: string;
    username: string;
}

interface SessionContextType {
    user: User | null;
    setUser: React.Dispatch<React.SetStateAction<User | null>>;
    logout: () => Promise<void>;
    loading: boolean;
}

export const SessionContext = createContext<SessionContextType | null>(null)