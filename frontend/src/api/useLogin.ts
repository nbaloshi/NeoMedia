export async function Login(email: string, password: string) {
    const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({ email, password }),
        credentials: "include",
    })

    if(!res.ok) {
        const data = await res.json().catch(() => ({}))
        throw new Error(data.message || "Login failed")
    }

    return res
}