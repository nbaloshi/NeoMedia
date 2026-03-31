import { useMutation } from "@tanstack/react-query"

interface RegisterBody {
    firstname: string
    lastname: string
    email: string
    username: string
    password: string
    dateOfBirth: string
    gender: "male" | "female" | "other"
    profilePic?: string
}

interface RegisterResponse {
    message: string
}

export function useRegister() {
    const mutationFn = async (body: RegisterBody): Promise<RegisterResponse> => {
        const res = await fetch("http://localhost:8080/register", {
            method: "POST",
            headers: { "content-type": "application/json" },
            body: JSON.stringify(body),
        })

        if (!res.ok) {
            const data = await res.json().catch(() => ({}))
            throw new Error(data.message || "Registeration failed")
        }

        return res.json()
    }
    return useMutation({ mutationFn })
}