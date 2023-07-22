import { useLoginContext } from 'context';
import { createArticleFavorite, deleteArticleFavorite, useArticle } from 'data';
import { ArticleDto } from 'models';
import { Link, useNavigate } from 'react-router-dom';

export type ArticlePreviewProps = {
  article: ArticleDto;
};

const ArticlePreview = ({ article }: ArticlePreviewProps) => {
  const { user } = useLoginContext().state;

  const { replaceArticle } = useArticle();
  const navigate = useNavigate();

  const onClickFavorite = () => {
    if (!user) {
      navigate('/login');
      return;
    }

    if (article.favorited) {
      deleteArticleFavorite(article.slug).then(result => replaceArticle(article.slug, result.article));
    } else {
      createArticleFavorite(article.slug).then(result => replaceArticle(article.slug, result.article));
    }
  };

  return (
    <div className="article-preview">
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
        <button
          className={`btn btn-outline-primary btn-sm pull-xs-right ${article.favorited ? 'active' : ''}`}
          onClick={onClickFavorite}
        >
          <i className="ion-heart"></i> {article.favoritesCount}
        </button>
      </div>
      <Link to={`/article/${article.slug}`} className="preview-link">
        <h1>{article.title}</h1>
        <p>{article.description}</p>
        <span>Read more...</span>
        <ul className="tag-list">
          {article.tagList.map(tag => (
            <li key={tag} className="tag-default tag-pill tag-outline">
              {tag}
            </li>
          ))}
        </ul>
      </Link>
    </div>
  );
};

export default ArticlePreview;
