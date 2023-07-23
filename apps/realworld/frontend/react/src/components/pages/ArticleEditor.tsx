import { createArticle, updateArticle } from 'data';
import { ArticleDto, NewArticleDto } from 'models';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

export type ArticleEditorProps = {
  article: ArticleDto;
  slug: string | undefined;
};

const ArticleEditor = ({ article, slug }: ArticleEditorProps) => {
  const [newArticle, setNewArticle] = useState<NewArticleDto>({
    title: '',
    description: '',
    body: '',
    tagList: [],
  });
  const [tagString, setTagString] = useState('');
  const [error, setError] = useState<string>();

  const navigate = useNavigate();

  const onClickPublish = () => {
    if (slug) {
      updateArticle(slug, {
        body: newArticle.body || undefined,
        description: newArticle.description || undefined,
        title: newArticle.title || '',
      }).then(result => {
        navigate(`/article/${result.article.slug}`, { state: { article: result.article } });
      });
    } else {
      createArticle({ ...newArticle, tagList: tagString.split(',').map(tag => tag.trim()) })
        .then(result => {
          navigate(`/article/${result.article.slug}`, { state: { article: result.article } });
        })
        .catch(err => {
          if (err.response?.status === 422) {
            setError('Title is already taken');
          } else {
            setError("Couldn't publish article");
          }
        });
    }
  };

  useEffect(() => {
    setNewArticle({
      title: article.title,
      description: article.description,
      body: article.body,
      tagList: article.tagList,
    });
    setTagString(article.tagList.join(', '));
  }, [article]);

  return (
    <div className="editor-page">
      <div className="container page">
        <div className="row">
          <div className="col-md-10 offset-md-1 col-xs-12">
            <form onSubmit={e => e.preventDefault()}>
              {error && (
                <ul className="error-messages">
                  <li>{error}</li>
                </ul>
              )}

              <fieldset>
                <fieldset className="form-group">
                  <input
                    type="text"
                    className="form-control form-control-lg"
                    placeholder="Article Title"
                    value={newArticle?.title ?? ''}
                    onChange={e => setNewArticle({ ...newArticle, title: e.target.value })}
                  />
                </fieldset>
                <fieldset className="form-group">
                  <input
                    type="text"
                    className="form-control"
                    placeholder="What's this article about?"
                    value={newArticle?.description ?? ''}
                    onChange={e => setNewArticle({ ...newArticle, description: e.target.value })}
                  />
                </fieldset>
                <fieldset className="form-group">
                  <textarea
                    className="form-control"
                    rows={8}
                    placeholder="Write your article (in markdown)"
                    value={newArticle?.body ?? ''}
                    onChange={e => setNewArticle({ ...newArticle, body: e.target.value })}
                  ></textarea>
                </fieldset>
                {!slug && (
                  <fieldset className="form-group">
                    <input
                      type="text"
                      className="form-control"
                      placeholder="Enter tags"
                      value={tagString}
                      onChange={e => setTagString(e.target.value)}
                    />
                  </fieldset>
                )}
                <button className="btn btn-lg pull-xs-right btn-primary" type="button" onClick={onClickPublish}>
                  {slug ? 'Update' : 'Publish'} Article
                </button>
              </fieldset>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ArticleEditor;
