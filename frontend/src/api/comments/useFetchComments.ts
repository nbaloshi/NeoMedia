import { useInfiniteQuery } from "@tanstack/react-query"

interface Comment {
    id: string
    username: string
    content: string
    createdAt: string
}

export default function useFetchComments(postId: string) {
    const queryFn = async ({ pageParam = 0 }): Promise<Comment[]> => {
        const res = await fetch(`http://localhost:8080/comments?post_id=${postId}&page=${pageParam}&limit=10`, {
            credentials: "include",
        })
        
        if (!res.ok) throw new Error("Failed to fetch comments")
        return res.json()
    }

    return useInfiniteQuery({
        queryKey: ["comments", postId],
        queryFn,
        initialPageParam: 0,
        getNextPageParam: (lastPage, allPages) => {
            if (lastPage.length < 10) return undefined
            return allPages.length
        },
    })
}
