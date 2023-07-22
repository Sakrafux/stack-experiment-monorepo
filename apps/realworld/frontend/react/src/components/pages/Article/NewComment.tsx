import { createArticleComment } from 'api';
import { CommentDto, ProfileDto } from 'models';
import { useState } from 'react';

export type NewCommentProps = {
  slug: string;
  author: ProfileDto;
  setComments: React.Dispatch<React.SetStateAction<CommentDto[]>>;
};

const NewComment = ({ slug, author, setComments }: NewCommentProps) => {
  const [comment, setComment] = useState('');
  const [loading, setLoading] = useState(false);

  const onClick = () => {
    setLoading(true);
    createArticleComment(slug, { body: comment }).then(result => {
      setComment('');
      setLoading(false);
      setComments(comments => [result.comment, ...comments]);
    });
  };

  return (
    <form className="card comment-form" onSubmit={e => e.preventDefault()}>
      <div className="card-block">
        <textarea
          className="form-control"
          placeholder="Write a comment..."
          rows={3}
          value={comment}
          onChange={e => setComment(e.target.value)}
          disabled={loading}
        ></textarea>
      </div>
      <div className="card-footer">
        <img src={author.image} alt="profile" className="comment-author-img" />
        <button className="btn btn-sm btn-primary" onClick={onClick} disabled={loading}>
          Post Comment
        </button>
      </div>
    </form>
  );
};

export default NewComment;
