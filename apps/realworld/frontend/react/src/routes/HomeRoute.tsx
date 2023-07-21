import { Home } from 'components';
import { getTags, useArticle, useTagContext } from 'data';
import { useEffect } from 'react';

const HomeRoute = () => {
  const { state: tags, dispatch: setTags } = useTagContext();

  const { getGlobalFeed, getMyFeed } = useArticle();

  useEffect(() => {
    if (!tags) {
      getTags().then(setTags);
    }
  }, [setTags, tags]);

  useEffect(() => {
    getGlobalFeed(0);
  }, [getGlobalFeed]);

  useEffect(() => {
    getMyFeed(0);
  }, [getMyFeed]);

  return <Home />;
};

export default HomeRoute;
