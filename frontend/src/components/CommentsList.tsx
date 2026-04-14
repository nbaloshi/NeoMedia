import { useEffect, useRef } from "react"
import useFetchComments from "../api/comments/useFetchComments"
import formatDate from "../utils/formateDate"

export default function CommentsList({ postId }: { postId: string }) {
    const {
        data,
        isLoading,
        isError,
        error,
        fetchNextPage,
        hasNextPage,
        isFetchingNextPage,
    } = useFetchComments(postId)
    
    const loadMoreRef = useRef<HTMLDivElement | null>(null)
    
    useEffect(() => {
        if (!hasNextPage || isFetchingNextPage) return
        const observer = new IntersectionObserver(
            entries => {
                if (entries[0].isIntersecting) {
                    fetchNextPage()
                }
            },
            { threshold: 1.0 }
        )
        if (loadMoreRef.current) {
            observer.observe(loadMoreRef.current)
        }
        return () => {
            if (loadMoreRef.current) {
                observer.unobserve(loadMoreRef.current)
            }
        }
    }, [hasNextPage, isFetchingNextPage, fetchNextPage])
    
    if (isLoading) return <p>Loading comments...</p>
    if (isError) return <p className="text-red-500">{(error as Error).message}</p>

    return (
        <div className="mt-4 border-t pt-2 space-y-2">
            {data?.pages.map((page, i) => (
                <div key={i}>
                    {page.map(comment => (
                        <div key={comment.id} className="p-4 bg-indigo-100 shadow rounded m-4">
                            <div className="flex justify-between items-start mb-2">
                                <p className="font-semibold">{comment.username}</p>
                                <span className="text-xs text-gray-500">{formatDate(comment.createdAt)}</span>
                            </div>
                            <p className="break-words">{comment.content}</p>
                        </div>
                    ))}
                </div>
            ))}
            {hasNextPage && (
                <div ref={loadMoreRef} className="h-10 flex justify-center items-center">
                    {isFetchingNextPage && <p>Loading more...</p>}
                </div>
            )}
        </div>
    )
}