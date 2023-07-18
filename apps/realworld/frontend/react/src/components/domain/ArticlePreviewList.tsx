import { Pagination } from 'components';
import { ArticlePages } from 'data';
import { useEffect, useState } from 'react';
import ArticlePreview from './ArticlePreview';

export type ArticlePreviewListProps = {
  articlePages: ArticlePages;
  pageSize: number;
  getArticlesForPage: (page: number) => Promise<void>;
};

const ArticlePreviewList = ({ articlePages, pageSize, getArticlesForPage }: ArticlePreviewListProps) => {
  const [page, setPage] = useState(0);
  const [activePage, setActivePage] = useState(0);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setIsLoading(true);
    getArticlesForPage(page).then(() => {
      setActivePage(page);
      setIsLoading(false);
    });
  }, [getArticlesForPage, page]);

  if (articlePages.articlesCount === 0) return <div className="article-preview">No articles are here... yet.</div>;

  return (
    <>
      {articlePages.articles[activePage].map(article => (
        <ArticlePreview key={article.slug} article={article} />
      ))}
      {isLoading && <div className="article-preview">Loading articles...</div>}
      <Pagination page={page} onPageChange={setPage} totalPages={Math.ceil(articlePages.articlesCount / pageSize)} />
    </>
  );
};

export default ArticlePreviewList;
