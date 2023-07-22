import { ArticlePages, useArticleContext } from '../state';
import { getArticles, getArticlesFeed } from '../action';
import { useCallback, useRef } from 'react';
import { environment } from 'environments/environment';
import { useLoginContext } from 'context';
import { ArticleDto } from 'models';

const useArticle = () => {
  const ongoingRequests = useRef<Record<string, Record<number, boolean>>>({ globalFeed: {}, myFeed: {} });

  const { state: articles, dispatch: setArticles } = useArticleContext();
  const { user, isLoading } = useLoginContext().state;

  const getGlobalFeed = useCallback(
    async (page: number) => {
      if (isLoading || articles.globalFeed?.articles[page] || ongoingRequests.current['globalFeed'][page]) return;

      ongoingRequests.current['globalFeed'][page] = true;
      return getArticles({ limit: environment.pageSizes.home, offset: page }).then(({ articles, articlesCount }) => {
        ongoingRequests.current['globalFeed'][page] = false;

        return setArticles({
          type: 'setGlobalFeed',
          setAction: stateSlice => ({
            articles: { ...stateSlice?.articles, [page]: articles },
            articlesCount,
          }),
        });
      });
    },
    [articles.globalFeed?.articles, isLoading, setArticles]
  );

  const getMyFeed = useCallback(
    async (page: number) => {
      if (!user || articles.myFeed?.articles[page] || ongoingRequests.current['myFeed'][page]) return;

      ongoingRequests.current['myFeed'][page] = true;
      return getArticlesFeed({ limit: environment.pageSizes.home, offset: page }).then(
        ({ articles, articlesCount }) => {
          ongoingRequests.current['myFeed'][page] = false;

          return setArticles({
            type: 'setMyFeed',
            setAction: stateSlice => ({
              articles: { ...stateSlice?.articles, [page]: articles },
              articlesCount,
            }),
          });
        }
      );
    },
    [articles.myFeed?.articles, setArticles, user]
  );

  const getTaggedFeed = useCallback(
    (tag: string) => async (page: number) => {
      ongoingRequests.current[tag] = ongoingRequests.current[tag] || {};
      if (articles.taggedArticles[tag]?.articles[page] || ongoingRequests.current[tag][page]) return;

      ongoingRequests.current[tag][page] = true;
      return getArticles({ limit: environment.pageSizes.home, offset: page, tag }).then(
        ({ articles, articlesCount }) => {
          ongoingRequests.current[tag][page] = false;

          return setArticles({
            type: 'setTaggedArticles',
            setAction: stateSlice => ({
              ...stateSlice,
              [tag]: { articles: { ...stateSlice[tag]?.articles, [page]: articles }, articlesCount },
            }),
          });
        }
      );
    },
    [articles.taggedArticles, setArticles]
  );

  const replaceArticle = useCallback((slug: string, article: ArticleDto) => {
    setArticles({
      type: 'setAll',
      setAction: stateSlice => ({
        myFeed: stateSlice.myFeed && replaceArticleInSlice(stateSlice.myFeed, slug, article),
        globalFeed: stateSlice.globalFeed && replaceArticleInSlice(stateSlice.globalFeed, slug, article),
        taggedArticles: Object.fromEntries(
          Object.entries(stateSlice.taggedArticles).map(([tag, slice]) => [
            tag,
            replaceArticleInSlice(slice, slug, article),
          ])
        ),
        userArticles: Object.fromEntries(
          Object.entries(stateSlice.userArticles).map(([username, slice]) => [
            username,
            replaceArticleInSlice(slice, slug, article),
          ])
        ),
        favoritedArticles: Object.fromEntries(
          Object.entries(stateSlice.favoritedArticles).map(([username, slice]) => [
            username,
            replaceArticleInSlice(slice, slug, article),
          ])
        ),
      }),
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return { getGlobalFeed, getMyFeed, getTaggedFeed, replaceArticle };
};

const replaceArticleInSlice = (slice: ArticlePages, slug: string, article: ArticleDto): ArticlePages => {
  return {
    articles: Object.fromEntries(
      Object.entries(slice.articles).map(([page, articles]) => {
        const index = articles.findIndex(a => a.slug === slug);
        if (index > -1) {
          articles[index] = article;
        }
        return [page, [...articles]];
      })
    ),
    articlesCount: slice.articlesCount,
  };
};

export default useArticle;
