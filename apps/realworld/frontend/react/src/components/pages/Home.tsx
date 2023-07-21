import { useLoginContext } from 'context';
import { ArticlePreviewList } from '../domain';
import { useArticle, useArticleContext, useTagContext } from 'data';
import { environment } from 'environments/environment';
import { useState } from 'react';

const Home = () => {
  const [selectedTab, setSelectedTab] = useState<string>('global');

  const { user } = useLoginContext().state;
  const tags = useTagContext().state;
  const { state: articles } = useArticleContext();

  const { getGlobalFeed, getMyFeed, getTaggedFeed } = useArticle();

  const isTagSelected = selectedTab !== 'global' && selectedTab !== 'myFeed';

  const getSelectedArticles = () => {
    switch (selectedTab) {
      case 'global':
        return articles.globalFeed;
      case 'myFeed':
        return articles.myFeed;
      default:
        return articles.taggedArticles[selectedTab];
    }
  };

  const getSelectedArticlesFunction = () => {
    switch (selectedTab) {
      case 'global':
        return getGlobalFeed;
      case 'myFeed':
        return getMyFeed;
      default:
        return getTaggedFeed(selectedTab);
    }
  };

  return (
    <div className="home-page">
      <div className="banner">
        <div className="container">
          <h1 className="logo-font">conduit</h1>
          <p>A place to share your knowledge.</p>
        </div>
      </div>

      <div className="container page">
        <div className="row">
          <div className="col-md-9">
            <div className="feed-toggle">
              <ul className="nav nav-pills outline-active">
                {user && (
                  <li className="nav-item" onClick={() => setSelectedTab('myFeed')}>
                    <span className={`nav-link ${selectedTab === 'myFeed' ? 'active' : ''}`}>Your Feed</span>
                  </li>
                )}
                <li className="nav-item" onClick={() => setSelectedTab('global')}>
                  <span className={`nav-link ${selectedTab === 'global' ? 'active' : ''}`}>Global Feed</span>
                </li>
                {isTagSelected && (
                  <li className="nav-item">
                    <span className="nav-link active">#{selectedTab}</span>
                  </li>
                )}
              </ul>
            </div>

            <ArticlePreviewList
              key={selectedTab}
              articlePages={getSelectedArticles()}
              pageSize={environment.pageSizes.home}
              getArticlesForPage={getSelectedArticlesFunction()}
            />
          </div>

          <div className="col-md-3">
            <div className="sidebar">
              <p>Popular Tags</p>

              {tags ? (
                <div className="tag-list">
                  {tags.map(tag => (
                    <button className="tag-pill tag-default" key={tag} onClick={() => setSelectedTab(tag)}>
                      {tag}
                    </button>
                  ))}
                </div>
              ) : (
                <div>Loading Tags...</div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;
