import { IUserResponse } from '@app/Types/Responses/User';

export interface ICommentResponse {
    id: number;
    comment: string;
    timestamp: string;
	user: IUserResponse;
}
