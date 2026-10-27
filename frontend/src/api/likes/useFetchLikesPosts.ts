import { useQuery } from "@tanstack/react-query"
import API_BASE_URL from "../../apiConfig";

interface LikesResponse {
  count: number
  likedByMe: boolean
}

export default function useFetchLikesPosts(postId: string) {
  const queryFn = async (): Promise<LikesResponse> => {
    const res = await fetch(`${API_BASE_URL}/likes-post?post_id=${postId}`,{
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