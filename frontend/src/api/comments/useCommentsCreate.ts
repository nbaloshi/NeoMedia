import { useMutation } from "@tanstack/react-query"

interface CommentCreateBody {
    post_id: string
    content: string
}

interface CommentCreateResponse {
    message: string
}

export default function useCommentCreate() {
    const mutationFn = async (body: CommentCreateBody): Promise<CommentCreateResponse> => {
        const res = await fetch("http://localhost:8080/comments", {
            method: "POST",
            headers: {"content-type": "application/json"},
            body: JSON.stringify(body),
            credentials: "include",
        })

        if (!res.ok) {
            const data = await res.json().catch(() => ({}))
            throw new Error(data.message || "Failed to create comment")
        }
        return res.json()
        
    }
    return useMutation({ mutationFn })
}