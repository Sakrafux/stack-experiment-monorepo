import { Article } from 'components';
import { getArticle } from 'data';
import { ArticleDto } from 'models';
import { useEffect, useState } from 'react';
import { useLocation, useParams } from 'react-router-dom';

const ArticleRoute = () => {
  const [article, setArticle] = useState<ArticleDto>();

  const { slug } = useParams();
  const { article: articleFromRoute } = (useLocation()?.state as { article: ArticleDto }) ?? {};

  useEffect(() => {
    if (articleFromRoute) {
      setArticle(articleFromRoute);
    } else {
      getArticle(slug!).then(result => setArticle(result.article));
    }
  }, [articleFromRoute, slug]);

  useEffect(() => {
    document.title = `${article?.title ?? 'Article'} — Conduit`;
  }, [article?.title]);

  const articleOrDefault = article ?? {
    slug: '',
    title: '',
    description: '',
    body: '',
    tagList: [],
    author: { username: '', bio: '', image: '', following: false },
    createdAt: new Date(),
    updatedAt: new Date(),
    favorited: false,
    favoritesCount: 0,
  };

  return <Article article={articleOrDefault} setArticle={setArticle} />;
};

export default ArticleRoute;
