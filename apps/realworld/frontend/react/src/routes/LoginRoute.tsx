import { Login } from 'components';
import { useLoginContext } from 'context';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const LoginRoute = () => {
  const { user } = useLoginContext().state;

  const navigate = useNavigate();

  useEffect(() => {
    if (user) {
      console.log('user', user);
      navigate('/');
    }
  }, [navigate, user]);

  return <Login />;
};

export default LoginRoute;
