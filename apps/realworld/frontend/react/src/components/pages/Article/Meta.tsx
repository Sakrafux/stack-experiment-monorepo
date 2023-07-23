import { followUserByUsername, unfollowUserByUsername } from 'api';
import { useLoginContext } from 'context';
import { createArticleFavorite, deleteArticleFavorite, useArticle } from 'data';
import { ArticleDto } from 'models';
import { Link, useLocation, useNavigate } from 'react-router-dom';

export type ArticleMetaProps = {
  article: ArticleDto;
  setArticle: React.Dispatch<React.SetStateAction<ArticleDto | undefined>>;
};

const ArticleMeta = ({ article, setArticle }: ArticleMetaProps) => {
  const { user } = useLoginContext().state;

  const { replaceArticle } = useArticle();
  const navigate = useNavigate();
  const location = useLocation().pathname;

  const isCurrentUser = user?.username === article.author.username;

  const onClickFavorite = () => {
    if (!user) {
      navigate('/login', { state: { location } });
      return;
    }

    if (article.favorited) {
      deleteArticleFavorite(article.slug).then(result => {
        replaceArticle(article.slug, result.article);
        setArticle(result.article);
      });
    } else {
      createArticleFavorite(article.slug).then(result => {
        replaceArticle(article.slug, result.article);
        setArticle(result.article);
      });
    }
  };

  const onClickFollow = () => {
    if (!user) {
      navigate('/login', { state: { location } });
      return;
    }

    if (article.author.following) {
      unfollowUserByUsername(article.author.username).then(result => {
        replaceArticle(article.slug, { ...article, author: result });
        setArticle({ ...article, author: result });
      });
    } else {
      followUserByUsername(article.author.username).then(result => {
        replaceArticle(article.slug, { ...article, author: result });
        setArticle({ ...article, author: result });
      });
    }
  };

  return (
    <div className="article-meta">
      <Link to={`/profile/${article.author.username}`}>
        <img src={article.author.image} alt="profile" />
      </Link>
      <div className="info">
        <Link to={`/profile/${article.author.username}`} className="author">
          {article.author.username}
        </Link>
        <span className="date">{article.createdAt.toLocaleDateString()}</span>
      </div>
      {!isCurrentUser && (
        <button className="btn btn-sm btn-outline-secondary action-btn" onClick={onClickFollow}>
          {article.author.following ? (
            <>
              <i className="ion-minus-round"></i>
              &nbsp; Unfollow {article.author.username}
            </>
          ) : (
            <>
              <i className="ion-plus-round"></i>
              &nbsp; Follow {article.author.username}
            </>
          )}
        </button>
      )}
      &nbsp;&nbsp;
      <button
        className={`btn btn-sm btn-outline-primary ${article.favorited ? 'active' : ''}`}
        onClick={onClickFavorite}
      >
        <i className="ion-heart"></i>
        &nbsp; {article.favorited ? 'Unfavorite' : 'Favorite'} Post{' '}
        <span className="counter">({article.favoritesCount})</span>
      </button>
      &nbsp;&nbsp;
      {isCurrentUser && (
        <Link className="btn btn-sm btn-outline-secondary action-btn" to={`/editor/${article.slug}`}>
          <i className="ion-edit"></i>
          &nbsp; Edit Article
        </Link>
      )}
    </div>
  );
};

export default ArticleMeta;
