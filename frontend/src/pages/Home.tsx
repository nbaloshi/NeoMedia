import Feed from "../components/Feed";
import Header from "../components/Header";
import PostInput from "../components/PostInput";

export default function HomePage() {
    return (
        <div className="flex flex-col min-h-screen">
            <Header />
            <div className="flex flex-row flex-1">

                <div className="w-[50%] p-4 flex flex-col">
                    <div className="mb-4">
                        <PostInput />
                    </div>
                    <div className=" bg-gray-200 shadow rounded p-4 overflow-y-auto h-[754px]">
                        <Feed />
                    </div>
                </div>

                <div className="w-[30%] p-4 flex flex-col">
                    <div className="flex-1 bg-gray-200 shadow rounded p-4 overflow-y-auto">
                        <p className="text-gray-500">Chat area...</p>
                    </div>
                </div>

                <div className="w-[20%] p-4 flex flex-col">
                    <div className="flex-1 bg-gray-200 shadow rounded p-4 overflow-y-auto">
                        <p className="text-gray-500">User list...</p>
                    </div>
                </div>
                
            </div>
        </div>
    )
}