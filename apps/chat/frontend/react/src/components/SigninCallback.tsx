import { useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

const SigninCallback = () => {
  const completionPromise = useRef<Promise<void>>();

  const { completeLogin } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (!completionPromise.current) {
      completionPromise.current = completeLogin().then(() => {
        navigate('/', { replace: true });
      });
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return <div></div>;
};

export default SigninCallback;
