import { api } from 'api';
import { AxiosResponse } from 'axios';
import {
  MultipleArticlesResponse,
  NewArticleDto,
  NewArticleRequest,
  PaginationParams,
  SingleArticleResponse,
  UpdateArticleDto,
  UpdateArticleRequest,
} from 'models';

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

export const createArticleFavorite = async (slug: string): Promise<SingleArticleResponse> => {
  const result = await api.post<SingleArticleResponse, AxiosResponse<SingleArticleResponse>>(
    `/articles/${slug}/favorite`
  );

  result.data.article.author.image =
    result.data.article.author.image || 'https://api.realworld.io/images/demo-avatar.png';

  return result.data;
};

export const deleteArticleFavorite = async (slug: string): Promise<SingleArticleResponse> => {
  const result = await api.delete<SingleArticleResponse, AxiosResponse<SingleArticleResponse>>(
    `/articles/${slug}/favorite`
  );

  result.data.article.author.image =
    result.data.article.author.image || 'https://api.realworld.io/images/demo-avatar.png';

  return result.data;
};

export const getArticle = async (slug: string): Promise<SingleArticleResponse> => {
  const result = await api.get<SingleArticleResponse, AxiosResponse<SingleArticleResponse>>(`/articles/${slug}`);

  result.data.article.author.image =
    result.data.article.author.image || 'https://api.realworld.io/images/demo-avatar.png';

  return result.data;
};

export const createArticle = async (article: NewArticleDto): Promise<SingleArticleResponse> => {
  const result = await api.post<SingleArticleResponse, AxiosResponse<SingleArticleResponse>, NewArticleRequest>(
    '/articles',
    {
      article,
    }
  );

  result.data.article.author.image =
    result.data.article.author.image || 'https://api.realworld.io/images/demo-avatar.png';

  return result.data;
};

export const updateArticle = async (slug: string, article: UpdateArticleDto): Promise<SingleArticleResponse> => {
  const result = await api.put<SingleArticleResponse, AxiosResponse<SingleArticleResponse>, UpdateArticleRequest>(
    `/articles/${slug}`,
    {
      article,
    }
  );

  result.data.article.author.image =
    result.data.article.author.image || 'https://api.realworld.io/images/demo-avatar.png';

  return result.data;
};

export const deleteArticle = async (slug: string): Promise<void> => {
  await api.delete<void, AxiosResponse<void>>(`/articles/${slug}`);
};
