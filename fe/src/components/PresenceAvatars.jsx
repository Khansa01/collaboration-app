import usePresence from "../hooks/usePresence.tsx";


const PresenceAvatars = ({ docId }) => {
    const { users, getColor } = usePresence(docId);

    const online = users.filter((u) => u.isOnline);

    if (online.length === 0) return null;

    return (
        <div className="flex items-center gap-1">
            {online.slice(0, 5).map((u) => (
                <div
                    key={u.userId}
                    title={u.name}
                    className="w-8 h-8 rounded-full flex items-center justify-center text-xs font-bold"
                    style={{
                        backgroundColor: getColor(u.userId),
                        color: "#303030",
                        border: "2px solid #303030",
                    }}
                >
                    {u.name?.[0]?.toUpperCase() ?? "?"}
                </div>
            ))}
            {online.length > 5 && (
                <div
                    className="w-8 h-8 rounded-full flex items-center justify-center text-xs font-bold"
                    style={{ backgroundColor: "#9E9E9E", color: "#F6F6F6", border: "2px solid #303030" }}
                >
                    +{online.length - 5}
                </div>
            )}
        </div>
    );
};

export default PresenceAvatars;