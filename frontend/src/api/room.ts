import { BaseMessage, SuccessResponse } from "../@types/api";
import { protectedApiClient } from "../utils/apiClient";


export const getRoomMessages = (roomId: string) => {
    return protectedApiClient.get<SuccessResponse<BaseMessage[]>>(`/room/${roomId}/messages`);
};