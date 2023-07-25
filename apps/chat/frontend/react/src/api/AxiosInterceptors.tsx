import { ReactNode } from 'react';
import { api } from './axios';
import { useAuthContext } from 'context/AuthContext';

export type AxiosInterceptorsProps = { children: ReactNode };

const AxiosInterceptors = ({ children }: AxiosInterceptorsProps) => {
  const { auth } = useAuthContext();

  api.interceptors.request.clear();

  api.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${auth?.id_token}`;
    return config;
  });

  // eslint-disable-next-line react/jsx-no-useless-fragment
  return <>{children}</>;
};

export default AxiosInterceptors;
