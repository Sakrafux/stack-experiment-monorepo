import jwtDecode from 'jwt-decode';
import { UserDto } from 'models';
import { createContext, useContext, useEffect, useMemo, useState } from 'react';

export type LoginContextState = {
  user?: UserDto;
  isLoading: boolean;
};

export type LoginContextType = {
  state: LoginContextState;
  dispatch: React.Dispatch<React.SetStateAction<LoginContextState>>;
};

const LoginContext = createContext<LoginContextType>({
  state: { isLoading: true },
  dispatch: () => {
    throw new Error('dispatch function must be overridden');
  },
});

export const useLoginContext = () => {
  const context = useContext(LoginContext);
  if (!context) {
    throw new Error('useLoginContext must be used within a LoginContextProvider');
  }
  return context;
};

export type LoginContextProviderProps = {
  defaultState?: LoginContextState;
  children: React.ReactNode;
};

export const LoginContextProvider = ({ defaultState, children }: LoginContextProviderProps) => {
  const [state, dispatch] = useState<LoginContextState>(defaultState ?? { isLoading: true });

  const value = useMemo(() => ({ state, dispatch }), [state]);

  useEffect(() => {
    if (state.user) {
      localStorage.setItem('user', JSON.stringify(state.user));
    }
  }, [state.user]);

  useEffect(() => {
    const user = localStorage.getItem('user');
    if (user) {
      const parsedUser = JSON.parse(user) as UserDto;
      const decodedToken = jwtDecode<{ exp: number; rol: string[]; sub: string }>(parsedUser.token);

      if (new Date(decodedToken.exp * 1000) > new Date()) {
        dispatch({ user: parsedUser, isLoading: false });
      } else {
        localStorage.removeItem('user');
      }
    } else {
      dispatch({ isLoading: false });
    }
  }, []);

  return <LoginContext.Provider value={value}>{children}</LoginContext.Provider>;
};
