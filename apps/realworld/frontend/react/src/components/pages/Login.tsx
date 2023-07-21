import { login } from 'api';
import { useLoginContext } from 'context';
import { useState } from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isError, setIsError] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const { dispatch } = useLoginContext();

  const location = useLocation();
  const navigate = useNavigate();

  const onSuccessfulLogin = async () => {
    const comesFromLogin = location.state?.location === '/login';
    navigate(comesFromLogin ? '/' : location.state.location || '/');
  };

  const onClickSignIn = async () => {
    setIsLoading(true);
    try {
      const result = await login(email, password);

      dispatch(result);

      onSuccessfulLogin();
    } catch (error) {
      setIsError(true);
      setIsLoading(false);
    }
  };

  return (
    <div className="auth-page">
      <div className="container page">
        <div className="row">
          <div className="col-md-6 offset-md-3 col-xs-12">
            <h1 className="text-xs-center">Sign in</h1>
            <p className="text-xs-center">
              <Link to="/register">Need an account?</Link>
            </p>

            {isError && (
              <ul className="error-messages">
                <li>email or password is invalid</li>
              </ul>
            )}

            <form onSubmit={e => e.preventDefault()}>
              <fieldset className="form-group">
                <input
                  className="form-control form-control-lg"
                  type="text"
                  placeholder="Email"
                  value={email}
                  onChange={event => setEmail(event.target.value)}
                  disabled={isLoading}
                />
              </fieldset>
              <fieldset className="form-group">
                <input
                  className="form-control form-control-lg"
                  type="password"
                  placeholder="Password"
                  value={password}
                  onChange={event => setPassword(event.target.value)}
                  disabled={isLoading}
                />
              </fieldset>
              <button className="btn btn-lg btn-primary pull-xs-right" onClick={onClickSignIn} disabled={isLoading}>
                Sign in
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
