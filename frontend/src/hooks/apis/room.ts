import { useQuery } from "react-query";
import { getRoomMessages } from "../../api/room";

export const useRoomMessages = (roomId: string) => {

    return useQuery(
        ["room", roomId, "messages"],

        async () => {
            const messages = await getRoomMessages(roomId);
            return messages.data.data;
        },
        {
            initialData: null,
           
        }
    );
}