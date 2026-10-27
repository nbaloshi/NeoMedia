import { useMutation, useQueryClient } from "@tanstack/react-query"
import API_BASE_URL from "../../apiConfig";

interface Response {
    liked: boolean
}

export default function useLikesCommentsToggle() {
    const queryClient = useQueryClient()
    const mutationFn = async (commentId: string): Promise<Response> => {
        const res = await fetch(`${API_BASE_URL}/likes-comment?comment_id=${commentId}`, {
            method: "POST",
            credentials: "include",
        })
        if (!res.ok) {
            const data = await res.json().catch(() => ({}))
            throw new Error(data.message || "Failed to toggle like")
        }
        return res.json()
    }
    return useMutation({
        mutationFn ,
        onSuccess: (_, commentId) => {
            queryClient.invalidateQueries({ queryKey: ["likes", commentId ]})
        }
    })
}