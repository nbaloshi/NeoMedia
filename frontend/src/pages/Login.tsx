import LoginForm from "../components/LoginForm";

export default function LoginPage() {
    return (
        <div className="flex min-h-screen">
            <div className="w-3/5 flex flex-col">
                <div className="absolute top-1 left-6">
                    <img
                        src="/images/nm.png"
                        alt="NM illustration"
                        className="w-60 h-auto"
                    />
                </div>
                <div className="flex justify-center mt-10">
                    <img
                        src="/images/hero.png"
                        alt="Hero illustration"
                    />
                </div>
                <div>
                    <p className="ml-8 mb-2 text-6xl font-semibold">Share posts, chat instantly, and</p>
                    <p className="ml-8 text-6xl font-semibold text-indigo-600">stay connected.</p>
                </div>
            </div>
            <div className="w-[1px] bg-gray-400 my-12"></div>
            <div className="w-2/5 flex items-center justify-center">
                <LoginForm />
            </div>
        </div>
    )
}