import { useMutation } from "@tanstack/react-query"

interface LoginBody {
    email: string
    password: string
}

interface LoginResponse {
    message: string
}

export function useLogin() {
    const mutationFn =  async (body: LoginBody): Promise<LoginResponse> => {
        const res = await fetch("http://localhost:8080/login", {
            method: "POST",
            headers: { "content-type": "application/json" },
            body: JSON.stringify(body),
            credentials: "include",
        })

        if(!res.ok) {
            const data = await res.json().catch(() => ({}))
            throw new Error(data.message || "Login failed")
        }
        return res.json()
    }
    return useMutation({ mutationFn })
}