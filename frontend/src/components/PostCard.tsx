import { useState } from "react";
import formatDate from "../utils/formateDate";
import { CommentInput } from "./CommentInput";
import CommentsList from "./CommentsList";
import LikeButton from "./LikeButton";

interface Post {
  id: string
  username: string
  content: string
  createdAt: string
}

export default function PostCard({ post }: { post: Post}) {
    const [expanded, setExpanded] = useState(false)

    return (
        <div key={post.id}
            className="p-4 bg-indigo-100 shadow rounded m-4 hover:bg-indigo-200 hover:cursor-pointer transition"
            onClick={() => setExpanded(prev => !prev)}
            >
            <div className="flex justify-between items-start mb-2">
                <p className="font-semibold">{post.username}</p>
                <span className="text-xs text-gray-500">{formatDate(post.createdAt)}</span>
            </div>
            <p className="break-words">{post.content}</p>
            <CommentInput postId={post.id} />

            {expanded && <CommentsList postId={post.id} />}

            <LikeButton postId={post.id} />

            
        </div>
    )
}