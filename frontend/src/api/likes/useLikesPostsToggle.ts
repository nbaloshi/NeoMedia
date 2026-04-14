import { useMutation, useQueryClient } from "@tanstack/react-query"

interface Response {
    liked: boolean
}

export default function useLikesPostsToggle() {
    const queryClient = useQueryClient()
    const mutationFn = async (postId: string): Promise<Response> => {
        const res = await fetch(`http://localhost:8080/likes-post?post_id=${postId}`, {
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
        onSuccess: (_, postId) => {
            queryClient.invalidateQueries({ queryKey: ["likes", postId ]})
        }
    })
}