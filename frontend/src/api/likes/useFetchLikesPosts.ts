import { useQuery } from "@tanstack/react-query"

interface LikesResponse {
  count: number
  likedByMe: boolean
}

export default function useFetchLikes(postId: string) {
  const queryFn = async (): Promise<LikesResponse> => {
    const res = await fetch(`http://localhost:8080/likes-post?post_id=${postId}`,{
        credentials: "include"
    })
    if (!res.ok) {
      const data = await res.json().catch(() => ({}))
      throw new Error(data.message || "Failed to fetch likes")
    }
    return res.json()
  }

  return useQuery<LikesResponse>({
    queryKey: ["likes", postId],
    queryFn,
  })
}