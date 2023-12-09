export interface ErrorResponse {
    error_code: string;
    message: string;
    status: boolean;
}

export interface SuccessResponse<T=unknown> {
    data: T;
    status: boolean;
    message: string;
}

export interface BaseUser {
	id: string;
	username: string;
	email: string;
}

export interface BaseRoom {
	id: string;
	description: string;
	password: string;
	room_owner: BaseUser;
	created_at: number;
}

export interface BaseMessage {
	id: string;
	username: string;
	message: string;
	timestamp: number;
}





