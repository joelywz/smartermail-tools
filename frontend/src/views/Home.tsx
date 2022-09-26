import ServerCard from "../components/ServerCard";

export default function Home() {
    return (
        <main className="m-8">
            <div className="flex justify-between items-center mb-4">
                <h1 className="font-bold text-2xl">Server List</h1>
                <button className="bg-blue-500 px-2 py-1.5 text-xs text-white rounded-sm font-semibold ">Add Server</button>
            </div>

            <ul className="flex flex-col gap-2">
                <li>
                    <ServerCard name="A" host="https://mail.htecloud.com"/>
                </li>

            </ul>
        </main>
    )
}