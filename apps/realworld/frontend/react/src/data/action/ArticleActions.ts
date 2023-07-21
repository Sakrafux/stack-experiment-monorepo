import { api } from 'api';
import { AxiosResponse } from 'axios';
import { MultipleArticlesResponse, PaginationParams } from 'models';

export type GetArticlesParams = PaginationParams & {
  tag?: string;
  author?: string;
  favorited?: string;
};

export const getArticles = async (params: GetArticlesParams): Promise<MultipleArticlesResponse> => {
  const result = await api.get<MultipleArticlesResponse, AxiosResponse<MultipleArticlesResponse>>('/articles', {
    params,
  });

  result.data.articles.forEach(
    article => (article.author.image = article.author.image || 'https://api.realworld.io/images/demo-avatar.png')
  );

  // If we want to keep less data in memory and use more requests instead, we can use this:
  // result.data.articles.forEach(article => (article.body = ''));

  return result.data;
};

export const getArticlesFeed = async (params: PaginationParams): Promise<MultipleArticlesResponse> => {
  const result = await api.get<MultipleArticlesResponse, AxiosResponse<MultipleArticlesResponse>>('/articles/feed', {
    params,
  });

  result.data.articles.forEach(
    article => (article.author.image = article.author.image || 'https://api.realworld.io/images/demo-avatar.png')
  );

  return result.data;
};
