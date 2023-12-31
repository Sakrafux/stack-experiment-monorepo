import { followUserByUsername, unfollowUserByUsername } from 'api';
import { useLoginContext } from 'context';
import { ProfileDto } from 'models';
import { Link, Outlet, useLocation, useNavigate } from 'react-router-dom';

export type ProfileProps = {
  profile: ProfileDto;
  setProfile: React.Dispatch<React.SetStateAction<ProfileDto | undefined>>;
};

const Profile = ({ profile, setProfile }: ProfileProps) => {
  const { user } = useLoginContext().state;
  const currentUsername = user?.username;

  const navigate = useNavigate();
  const location = useLocation().pathname;
  const subPath = location.split('/').at(-1);

  const isCurrentUser = currentUsername === profile.username;

  const onClickFollow = () => {
    if (!user) {
      navigate('/login', { state: { location } });
      return;
    }

    if (profile.following) {
      unfollowUserByUsername(profile.username).then(profile => setProfile(profile));
    } else {
      followUserByUsername(profile.username).then(profile => setProfile(profile));
    }
  };

  return (
    <div className="profile-page">
      <div className="user-info">
        <div className="container">
          <div className="row">
            <div className="col-xs-12 col-md-10 offset-md-1">
              {profile.image && <img src={profile.image} alt="profile" className="user-img" />}
              <h4>{profile.username}</h4>
              {profile.bio && <p>{profile.bio}</p>}
              {!isCurrentUser && (
                <button className="btn btn-sm btn-outline-secondary action-btn" onClick={onClickFollow}>
                  {profile.following ? (
                    <>
                      <i className="ion-minus-round"></i>
                      &nbsp; Unfollow {profile.username}
                    </>
                  ) : (
                    <>
                      <i className="ion-plus-round"></i>
                      &nbsp; Follow {profile.username}
                    </>
                  )}
                </button>
              )}
            </div>
          </div>
        </div>
      </div>

      <div className="container">
        <div className="row">
          <div className="col-xs-12 col-md-10 offset-md-1">
            <div className="articles-toggle">
              <ul className="nav nav-pills outline-active">
                <li className="nav-item">
                  <Link className={`nav-link ${subPath !== 'favorites' ? 'active' : ''}`} to="">
                    My Articles
                  </Link>
                </li>
                <li className="nav-item">
                  <Link className={`nav-link ${subPath === 'favorites' ? 'active' : ''}`} to="favorites">
                    Favorited Articles
                  </Link>
                </li>
              </ul>
            </div>

            <Outlet />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Profile;
