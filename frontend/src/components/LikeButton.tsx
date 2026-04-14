import { useQueryClient } from "@tanstack/react-query"
import useLikesPostsToggle from "../api/likes/useLikesPostsToggle"
import useFetchLikes from "../api/likes/useFetchLikesPosts"
import type React from "react"

export default function LikeButton({ postId }: { postId: string }) {
  const queryClient = useQueryClient()
  const { data, isLoading } = useFetchLikes(postId)
  const toggleLike = useLikesPostsToggle()

  if (isLoading) return <span>Loading...</span>

  const liked = data?.likedByMe ?? false
  const count = data?.count ?? 0

  const handleClick = (e: React.MouseEvent) => {
    e.stopPropagation()
    toggleLike.mutate(postId, {
      onSuccess: () => {
        // refetch likes after toggle to update count + state
        queryClient.invalidateQueries({ queryKey: ["likes", postId] })
      },
    })
  }

  return (
    <button
      onClick={handleClick}
      className="flex items-center space-x-1 hover:opacity-80 transition"
    >
      <svg
        className={`w-5 h-5 ${liked ? "text-red-500" : "text-gray-400"}`}
        fill="currentColor"
        viewBox="0 0 20 20"
      >
        <path d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 18.343l-6.828-6.829a4 4 0 010-5.656z" />
      </svg>
      <span className="text-sm">{count}</span>
    </button>
  )
}
