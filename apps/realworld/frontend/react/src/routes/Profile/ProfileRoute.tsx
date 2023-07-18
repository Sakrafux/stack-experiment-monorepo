import { getProfileByUsername } from 'api';
import { Profile } from 'components';
import { ProfileDto } from 'models';
import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

const ProfileRoute = () => {
  const [profile, setProfile] = useState<ProfileDto | undefined>();

  const { username } = useParams();
  const navigate = useNavigate();

  useEffect(() => {
    if (!profile || profile.username !== username) {
      getProfileByUsername(username!)
        .then(profile => setProfile(profile))
        .catch(() => navigate('/'));
    }
  }, [navigate, profile, username]);

  if (!profile) return <div>Loading...</div>;

  return <Profile profile={profile} />;
};

export default ProfileRoute;
