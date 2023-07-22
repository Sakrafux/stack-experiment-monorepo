import { Login } from 'components';
import { useLoginContext } from 'context';
import { useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

const LoginRoute = () => {
  const { user } = useLoginContext().state;

  const location = useLocation();
  const navigate = useNavigate();

  const onSuccessfulLogin = async () => {
    const comesFromLogin = location.state?.location === '/login';
    navigate(comesFromLogin ? '/' : location.state?.location || '/');
  };

  useEffect(() => {
    if (user) {
      navigate('/');
    }
  }, [navigate, user]);

  return <Login onSuccessfulLogin={onSuccessfulLogin} />;
};

export default LoginRoute;
