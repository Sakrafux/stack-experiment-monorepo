import { Register } from 'components';
import { useLoginContext } from 'context';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const RegisterRoute = () => {
  const { user } = useLoginContext().state;

  const navigate = useNavigate();

  useEffect(() => {
    if (user) {
      navigate('/');
    }
  }, [navigate, user]);

  useEffect(() => {
    document.title = 'Sign up — Conduit';
  }, []);

  return <Register />;
};

export default RegisterRoute;
