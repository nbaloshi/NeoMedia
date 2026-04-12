import formatDate from "../utils/formateDate";
import { CommentInput } from "./CommentInput";

export interface Post {
  id: string
  username: string
  content: string
  createdAt: string
}

export default function PostCard({ post }: { post: Post}) {
    return (
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
            <CommentInput postId={post.id} />
        </div>
    )
}