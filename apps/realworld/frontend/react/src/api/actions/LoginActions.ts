import { api } from 'api/axios';
import { AxiosResponse } from 'axios';
import { LoginUserRequest, UserResponse } from 'models';

export const login = async (email: string, password: string) => {
  const result = await api.post<UserResponse, AxiosResponse<UserResponse>, LoginUserRequest>('/users/login', {
    user: {
      email,
      password,
    },
  });

  return result.data.user;
};
