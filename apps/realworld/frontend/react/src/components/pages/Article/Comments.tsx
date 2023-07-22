import { deleteArticleComment } from 'api';
import { useLoginContext } from 'context';
import { CommentDto } from 'models';
import { Link } from 'react-router-dom';

export type CommentsProps = {
  slug: string;
  comments: CommentDto[];
  setComments: React.Dispatch<React.SetStateAction<CommentDto[]>>;
};

const Comments = ({ slug, comments, setComments }: CommentsProps) => {
  const { username } = useLoginContext().state?.user ?? {};

  const onDelete = (comment: CommentDto) => {
    deleteArticleComment(slug, comment.id).then(() => {
      setComments(comments => comments.filter(c => c.id !== comment.id));
    });
  };

  return (
    <>
      {comments.map(comment => (
        <div className="card" key={comment.id}>
          <div className="card-block">
            <p className="card-text">{comment.body}</p>
          </div>
          <div className="card-footer">
            <Link to={`/profile/${comment.author.username}`} className="comment-author">
              <img src={comment.author.image} className="comment-author-img" alt="profile" />
            </Link>
            &nbsp;
            <Link to={`/profile/${comment.author.username}`} className="comment-author">
              {comment.author.username}
            </Link>
            <span className="date-posted">{comment.createdAt.toLocaleDateString()}</span>
            {username === comment.author.username && (
              <span className="mod-options">
                <i className="ion-trash-a" onClick={() => onDelete(comment)}></i>
              </span>
            )}
          </div>
        </div>
      ))}
    </>
  );
};

export default Comments;
