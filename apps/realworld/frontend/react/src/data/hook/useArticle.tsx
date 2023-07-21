import { useArticleContext } from '../state';
import { getArticles, getArticlesFeed } from '../action';
import { useCallback, useRef } from 'react';
import { environment } from 'environments/environment';
import { useLoginContext } from 'context';

const useArticle = () => {
  const ongoingRequests = useRef<Record<string, Record<number, boolean>>>({ globalFeed: {}, myFeed: {} });

  const { state: articles, dispatch: setArticles } = useArticleContext();
  const { user } = useLoginContext().state;

  const getGlobalFeed = useCallback(
    async (page: number) => {
      if (articles.globalFeed?.articles[page] || ongoingRequests.current['globalFeed'][page]) return;

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
    [articles?.globalFeed, setArticles]
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

  return { getGlobalFeed, getMyFeed, getTaggedFeed };
};

export default useArticle;
