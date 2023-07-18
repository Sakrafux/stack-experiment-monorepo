import { Login } from 'components';
import { useLocation, useNavigate } from 'react-router-dom';

const LoginRoute = () => {
  const location = useLocation();
  const navigate = useNavigate();

  const onSuccessfulLogin = async () => {
    const comesFromLogin = location.state?.location === '/login';
    navigate(comesFromLogin ? '/' : location.state.location || '/');
  };

  return <Login onSuccessfulLogin={onSuccessfulLogin} />;
};

export default LoginRoute;
