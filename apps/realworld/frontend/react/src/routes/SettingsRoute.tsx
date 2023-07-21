import { Settings } from 'components';
import { useLoginContext } from 'context';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const SettingsRoute = () => {
  const user = useLoginContext().state;

  const navigate = useNavigate();

  useEffect(() => {
    if (!user.isLoading && !user.user) {
      navigate('/login');
    }
  }, [navigate, user]);

  return <Settings />;
};

export default SettingsRoute;
