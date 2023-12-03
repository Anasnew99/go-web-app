import { SuccessResponse } from "../@types/api";
import { apiClient } from "../utils/apiClient"

interface LoginResponse {
    token: string;
}

export const login = (username: string, password: string) => {

    return apiClient.post<SuccessResponse<LoginResponse>>("/auth/login", { username, password });
}

export interface RegisterRequest {
    username: string;
    password: string;
    email: string;
}

export const register = (body: RegisterRequest)=>{
    return apiClient.post<SuccessResponse>("/auth/register", body);
}

