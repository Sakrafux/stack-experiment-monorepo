import { login } from 'api';
import { useLoginContext } from 'context';
import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';

const Register = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isError, setIsError] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const { dispatch } = useLoginContext();

  const navigate = useNavigate();

  const onClickSignIn = async () => {
    setIsLoading(true);
    try {
      const result = await login(email, password);

      dispatch({ user: result, isLoading: false });

      navigate('/');
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
            <h1 className="text-xs-center">Sign up</h1>
            <p className="text-xs-center">
              <Link to="/login">Have an account?</Link>
            </p>

            {isError && (
              <ul className="error-messages">
                <li>That email or username is already taken</li>
              </ul>
            )}

            <form onSubmit={e => e.preventDefault()}>
              <fieldset className="form-group">
                <input
                  className="form-control form-control-lg"
                  type="text"
                  placeholder="Username"
                  value={username}
                  onChange={event => setUsername(event.target.value)}
                  disabled={isLoading}
                />
              </fieldset>
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
                Sign up
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Register;
