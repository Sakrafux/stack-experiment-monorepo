import { api } from 'api/axios';
import { AxiosResponse } from 'axios';
import { LoginUserRequest, UserDto, UserResponse } from 'models';

export const login = async (email: string, password: string): Promise<UserDto> => {
  const result = await api.post<UserResponse, AxiosResponse<UserResponse>, LoginUserRequest>('/users/login', {
    user: {
      email,
      password,
    },
  });

  return result.data.user;
};
