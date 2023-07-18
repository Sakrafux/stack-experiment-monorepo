import { ArticlePreviewList } from 'components';
import { getArticles, GetArticlesParams, useArticleContext } from 'data';
import { environment } from 'environments/environment';
import { useEffect } from 'react';
import { useParams } from 'react-router-dom';

export type ProfileArticlesRouteProps = {
  favorites?: boolean;
};

const ProfileArticlesRoute = ({ favorites }: ProfileArticlesRouteProps) => {
  const { state, dispatch } = useArticleContext();

  const { username } = useParams();

  const pagesForUser = favorites ? state.favoritedArticles[username!] : state.userArticles[username!];

  const getArticlesForPage = async (page: number) => {
    if (pagesForUser?.articles[page]) return;

    const params: GetArticlesParams = { limit: environment.pageSizes.profile, offset: page };
    if (favorites) {
      params.favorited = username;
    } else {
      params.author = username;
    }

    return getArticles(params).then(({ articles, articlesCount }) => {
      if (favorites) {
        return dispatch({
          type: 'setFavoritedArticles',
          setAction: stateSlice => ({
            ...stateSlice,
            [username!]: { articles: { ...stateSlice[username!]?.articles, [page]: articles }, articlesCount },
          }),
        });
      }

      return dispatch({
        type: 'setUserArticles',
        setAction: stateSlice => ({
          ...stateSlice,
          [username!]: { articles: { ...stateSlice[username!]?.articles, [page]: articles }, articlesCount },
        }),
      });
    });
  };

  useEffect(() => {
    if (!pagesForUser) {
      getArticlesForPage(0);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (!pagesForUser) return <div className="article-preview">Loading articles...</div>;

  return (
    <ArticlePreviewList
      articlePages={pagesForUser}
      pageSize={environment.pageSizes.profile}
      getArticlesForPage={getArticlesForPage}
    />
  );
};

export default ProfileArticlesRoute;
