import { api } from 'api';
import { AxiosResponse } from 'axios';
import { MultipleCommentsResponse, NewCommentDto, NewCommentRequest, SingleCommentResponse } from 'models';

export const getArticleComments = async (slug: string): Promise<MultipleCommentsResponse> => {
  const response = await api.get<MultipleCommentsResponse>(`/articles/${slug}/comments`);

  response.data.comments.forEach(comment => {
    comment.author.image = comment.author.image || 'https://api.realworld.io/images/demo-avatar.png';
  });

  return response.data;
};

export const createArticleComment = async (slug: string, comment: NewCommentDto): Promise<SingleCommentResponse> => {
  const response = await api.post<SingleCommentResponse, AxiosResponse<SingleCommentResponse>, NewCommentRequest>(
    `/articles/${slug}/comments`,
    { comment }
  );

  response.data.comment.author.image =
    response.data.comment.author.image || 'https://api.realworld.io/images/demo-avatar.png';

  return response.data;
};

export const deleteArticleComment = async (slug: string, id: number): Promise<void> => {
  await api.delete(`/articles/${slug}/comments/${id}`);
};
