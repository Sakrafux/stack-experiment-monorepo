import { ReactNode } from 'react';
import { useNavigate } from 'react-router-dom';
import { convertIsoDatesToDates, getToken } from 'utils';
import { api } from './axios';
import { useLoginContext } from 'context';

export type AxiosInterceptorsProps = { children: ReactNode };

const AxiosInterceptors = ({ children }: AxiosInterceptorsProps) => {
  const { user } = useLoginContext().state;

  const navigate = useNavigate();

  api.interceptors.request.use(config => {
    const token = !user ? getToken() : user.token;
    if (token) {
      config.headers.Authorization = token;
    }
    return config;
  });

  api.interceptors.response.use(
    response => {
      convertIsoDatesToDates(response.data);
      return response;
    },
    error => {
      if (error.response?.status === 401) {
        navigate('/login');
      }

      return Promise.reject(error);
    }
  );

  // eslint-disable-next-line react/jsx-no-useless-fragment
  return <>{children}</>;
};

export default AxiosInterceptors;
