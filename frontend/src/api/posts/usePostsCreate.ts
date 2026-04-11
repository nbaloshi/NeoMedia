import { useMutation } from "@tanstack/react-query"

interface PostsCreateBody {
    content: string
}

interface PostsCreateResponse {
    message: string
}

export default function usePostsCreate() {
    const mutationFn = async (body: PostsCreateBody): Promise<PostsCreateResponse> => {
        const res = await fetch("http://localhost:8080/posts", {
            method: "POST",
            headers: {"content-type": "application/json"},
            body: JSON.stringify(body),
            credentials: "include",
        })

        if (!res.ok) {
            const data = await res.json().catch(() => ({}))
            throw new Error(data.message || "Failed to create post")
        }
        return res.json()
    }
    return useMutation({ mutationFn })
}