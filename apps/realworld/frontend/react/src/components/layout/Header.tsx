import { Link, useLocation } from 'react-router-dom';
import { useLoginContext } from 'context';

const Header = () => {
  const {
    state: { user },
  } = useLoginContext();

  const location = useLocation().pathname;

  return (
    <nav className="navbar navbar-light">
      <div className="container">
        <a className="navbar-brand" href="index.html">
          conduit
        </a>
        <ul className="nav navbar-nav pull-xs-right">
          <li className="nav-item">
            <Link to="/" className={`nav-link ${location === '/' ? 'active' : ''}`}>
              Home
            </Link>
          </li>
          {user ? (
            <>
              <li className="nav-item">
                <Link to="/editor" className="nav-link">
                  {' '}
                  <i className="ion-compose"></i>&nbsp;New Article{' '}
                </Link>
              </li>
              <li className="nav-item">
                <Link to="/settings" className="nav-link">
                  {' '}
                  <i className="ion-gear-a"></i>&nbsp;Settings{' '}
                </Link>
              </li>
              <li className="nav-item">
                <Link to={`/profile/${user.username}`} className="nav-link">
                  {user.image && <img src={user.image} alt="profile" className="user-pic" />}
                  {user.username}
                </Link>
              </li>
            </>
          ) : (
            <>
              <li className="nav-item">
                <Link to="/login" state={{ location }} className={`nav-link ${location === '/login' ? 'active' : ''}`}>
                  Sign in
                </Link>
              </li>
              <li className="nav-item">
                <Link to="/register" className={`nav-link ${location === '/register' ? 'active' : ''}`}>
                  Sign up
                </Link>
              </li>
            </>
          )}
        </ul>
      </div>
    </nav>
  );
};

export default Header;
