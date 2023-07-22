import { ProfileDto } from './ProfileDto';

export type CommentDto = {
  id: number;
  createdAt: Date;
  updatedAt: Date;
  body: string;
  author: ProfileDto;
};
