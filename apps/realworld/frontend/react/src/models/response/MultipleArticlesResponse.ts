import { ArticleDto } from '../ArticleDto';

export type MultipleArticlesResponse = {
  articles: ArticleDto[];
  articlesCount: number;
};
