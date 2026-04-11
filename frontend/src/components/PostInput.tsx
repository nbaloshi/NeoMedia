import React, { useState } from "react";
import usePostsCreate from "../api/posts/usePostsCreate";

export default function PostInput() {
    const [content, setContent] = useState("")
    const { mutate, isPending, isError, error } = usePostsCreate()

    const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (!content.trim()) {
            return
        }
        mutate({ content })
        setContent("")
    }

    return (
        <form onSubmit={handleSubmit} className="flex flex-col">
            <div className="relative">
                <textarea
                    value={content}
                    onChange={e => setContent(e.target.value)}
                    placeholder="Whats on your mind..?"
                    className="w-full p-3 pr-20 border border-gray-300 rounded-lg resize-none focus:ring-2 focus:ring-cyan-600"
                    rows={3}
                />
                <button
                    type="submit"
                    disabled={isPending || !content.trim()}
                    className="absolute bottom-5 right-4 bg-indigo-600 text-white px-4 py-1 rounded-lg hover:bg-indigo-700 disabled:opacity-50"
                >
                    {isPending ? "Posting.." : "Post"}
                </button>
            </div>
            {isError && (<p className="text-red-500 mb-2">{(error as Error).message}</p>)}
        </form>
    )
}