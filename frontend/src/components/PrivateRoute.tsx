import { useContext, type ReactNode } from "react";
import { SessionContext } from "../context/SessionContext";
import { Navigate } from "react-router-dom";

interface PrivateRouteProps {
    children: ReactNode
}

export function PrivateRoute({ children }: PrivateRouteProps) {
    const session = useContext(SessionContext);

    if (!session) return null;
    const {user, loading } = session;

    if(loading) return <p>Loading...</p>;
    if (!user) return <Navigate to="/login" replace />;
    return children;
}