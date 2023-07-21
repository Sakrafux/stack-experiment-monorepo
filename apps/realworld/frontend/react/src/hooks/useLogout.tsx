import { useLoginContext } from 'context';
import { useCallback } from 'react';
import { useNavigate } from 'react-router-dom';

const useLogout = () => {
  const { dispatch } = useLoginContext();

  const navigate = useNavigate();

  const logout = useCallback(() => {
    localStorage.removeItem('user');
    dispatch({ user: undefined, isLoading: false });
    navigate('/');
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return logout;
};

export default useLogout;
