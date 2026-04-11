import { useEffect, useRef } from "react";
import useFetchPosts from "../api/posts/useFetchPosts";
import formatDate from "../utils/formateDate";

export default function Feed() {
    const {
        data,
        isLoading,
        isError,
        error,
        fetchNextPage,
        hasNextPage,
        isFetchingNextPage,
    } = useFetchPosts()

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

    if (isLoading) return <p>Loading posts...</p>
    if (isError) return <p className="text-red-500">{(error as Error).message}</p>

    return (
        <div className="space-y-4">
            {data?.pages.map((page, i) => (
                <div key={i}>
                    {page.map(post => (
                        <div key={post.id} className="p-4 bg-indigo-100 shadow rounded m-4">
                            <div className="flex justify-between items-start mb-2">
                                <p className="font-semibold">{post.username}</p>
                                <span className="text-xs text-gray-500">
                                    {formatDate(post.createdAt)}
                                </span>
                            </div>
                            <p className="break-words">
                                {post.content}
                            </p>
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