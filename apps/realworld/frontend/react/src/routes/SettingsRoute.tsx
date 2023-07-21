import { Settings } from 'components';
import { useLoginContext } from 'context';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const SettingsRoute = () => {
  const user = useLoginContext();

  const navigate = useNavigate();

  useEffect(() => {
    if (!user) {
      navigate('/login');
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return <Settings />;
};

export default SettingsRoute;
