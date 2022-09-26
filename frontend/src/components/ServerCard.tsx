import react from "react";
import { Trash } from "@emotion-icons/fa-solid/Trash";

interface Props {
    name: string;
    host: string;
    onDelete?: () => void;
}

const ServerCard: react.FC<Props> = (props) => {
    return (
        <div className="bg-white p-4 rounded-md shadow-sm flex items-center justify-between group">
            <div>
                <h2 className="font-semibold text-base">{props.name}</h2>
                <p className="text-xs font-semibold text-neutral-400">{props.host}</p>
            </div>
            <div className="hidden group-hover:block">
                <Trash onClick={props.onDelete} className="text-neutral-400 cursor-pointer" size={12}/>
            </div>
        </div>
    )
}

export default ServerCard;