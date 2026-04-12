import React, { useState } from "react";
import useCommentCreate from "../api/comments/useCommentsCreate";

export function CommentInput({ postId }: { postId: string }) {
    const [content, setContent] = useState("")
    const { mutate, isPending, isError, error } = useCommentCreate()

    const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (!content.trim()) {
            return
        }
        mutate({ post_id: postId, content })
        setContent("")
    }

    return (
        <form onSubmit={handleSubmit} className="flex flex-col">
            <div className="relative">
                <textarea
                    value={content}
                    onChange={e => setContent(e.target.value)}
                    placeholder="Add a comment..?"
                    className="w-full p-2 mt-3 border border-gray-300 rounded-lg resize-none focus:ring-2 focus:ring-cyan-600 pr-24"
                    rows={1}
                />
                <button
                    type="submit"
                    disabled={isPending || !content.trim()}
                    className="absolute bottom-2.5 right-2 bg-indigo-600 text-white px-4 py-1 rounded-lg hover:bg-indigo-700 disabled:opacity-50"
                >
                    {isPending ? "Commenting.." : "Enter"}
                </button>
            </div>
            {isError && (<p className="text-red-500 mb-2">{(error as Error).message}</p>)}
        </form>
    )
}