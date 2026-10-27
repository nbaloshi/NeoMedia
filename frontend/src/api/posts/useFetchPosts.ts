import { useInfiniteQuery } from "@tanstack/react-query"
import API_BASE_URL from "../../apiConfig";

interface Post {
    id: string
    username: string
    content: string
    createdAt: string
}

export default function useFetchPosts() {
    const queryFn = async ({ pageParam = 0 }): Promise<Post[]> => {
        const res = await fetch(`${API_BASE_URL}/posts?page=${pageParam}&limit=10`, {
            credentials: "include",
        })

        if (!res.ok) throw new Error("Failed to fetch posts")
        return res.json()
    }

    return useInfiniteQuery({
        queryKey: ["posts"],
        queryFn,
        initialPageParam: 0,
        getNextPageParam: (lastPage, allPages) => {
            if (lastPage.length < 10) return undefined
            return allPages.length
        },
    })
}
