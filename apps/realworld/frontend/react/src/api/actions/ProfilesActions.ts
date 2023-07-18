import { api } from 'api/axios';
import { AxiosResponse } from 'axios';
import { ProfileDto, ProfileResponse } from 'models';

export const getProfileByUsername = async (username: string): Promise<ProfileDto> => {
  const result = await api.get<ProfileResponse, AxiosResponse<ProfileResponse>>(`/profiles/${username}`);
  const { profile } = result.data;

  profile.image = profile.image || 'https://api.realworld.io/images/demo-avatar.png';

  return profile;
};
