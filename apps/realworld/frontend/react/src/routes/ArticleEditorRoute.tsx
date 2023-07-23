import { ArticleEditor } from 'components';
import { getArticle } from 'data';
import { ArticleDto } from 'models';
import { useEffect, useState } from 'react';
import { useLocation, useParams } from 'react-router-dom';

const ArticleEditorRoute = () => {
  const [article, setArticle] = useState<ArticleDto>({
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
  });

  const { slug } = useParams();
  const { article: articleFromRoute } = (useLocation()?.state as { article: ArticleDto }) ?? {};

  useEffect(() => {
    if (slug) {
      if (articleFromRoute) {
        setArticle(articleFromRoute);
      } else {
        getArticle(slug).then(result => setArticle(result.article));
      }
    }
  }, [articleFromRoute, slug]);

  return <ArticleEditor article={article} slug={slug} />;
};

export default ArticleEditorRoute;
