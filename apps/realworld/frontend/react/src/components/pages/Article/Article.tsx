import Meta from './Meta';
import { ArticleDto, CommentDto } from 'models';
import NewComment from './NewComment';
import { useLoginContext } from 'context';
import { Link, useLocation } from 'react-router-dom';
import Comments from './Comments';
import { useEffect, useState } from 'react';
import { getArticleComments } from 'api';

export type ArticleProps = {
  article: ArticleDto;
  setArticle: React.Dispatch<React.SetStateAction<ArticleDto | undefined>>;
};

const Article = ({ article, setArticle }: ArticleProps) => {
  const [comments, setComments] = useState<CommentDto[]>([]);

  const { user } = useLoginContext().state;

  const location = useLocation().pathname;

  useEffect(() => {
    if (article.slug) {
      getArticleComments(article.slug).then(result => setComments(result.comments));
    }
  }, [article.slug]);

  return (
    <div className="article-page">
      <div className="banner">
        <div className="container">
          <h1>{article.title}</h1>

          <Meta article={article} setArticle={setArticle} />
        </div>
      </div>

      <div className="container page">
        <div className="row article-content">
          <div className="col-md-12">
            <p>{article.body}</p>
            <ul className="tag-list">
              {article.tagList.map(tag => (
                <li key={tag} className="tag-default tag-pill tag-outline">
                  {tag}
                </li>
              ))}
            </ul>
          </div>
        </div>

        <hr />

        <div className="article-actions">
          <Meta article={article} setArticle={setArticle} />
        </div>

        <div className="row">
          <div className="col-xs-12 col-md-8 offset-md-2">
            {user ? (
              <NewComment author={article.author} slug={article.slug} setComments={setComments} />
            ) : (
              <>
                <div>
                  <Link to="/login" state={{ location }}>
                    Sign in
                  </Link>{' '}
                  or{' '}
                  <Link to="/register" state={{ location }}>
                    sign up
                  </Link>{' '}
                  to add comments on this article.
                </div>
                <br />
              </>
            )}

            <Comments slug={article.slug} comments={comments} setComments={setComments} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Article;
