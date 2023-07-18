import { ReactNode, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { convertIsoDatesToDates } from 'utils';
import { api } from './axios';

export type AxiosInterceptorsProps = { children: ReactNode };

const AxiosInterceptors = ({ children }: AxiosInterceptorsProps) => {
  const navigate = useNavigate();

  useEffect(() => {
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
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // eslint-disable-next-line react/jsx-no-useless-fragment
  return <>{children}</>;
};

export default AxiosInterceptors;
