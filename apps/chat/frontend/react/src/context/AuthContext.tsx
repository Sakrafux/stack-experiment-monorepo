import { User, UserManager } from 'oidc-client-ts';
import { createContext, useCallback, useContext, useEffect, useMemo, useState } from 'react';

const UserManagerContext = createContext(
  new UserManager({
    authority: 'https://accounts.google.com/o/oauth2/auth',
    client_id: process.env.NX_GOOGLE_CLIENT_ID!,
    client_secret: process.env.NX_GOOGLE_CLIENT_SECRET!,
    redirect_uri: 'http://localhost:4200/signin-callback',
    scope: 'openid profile',
    response_type: 'code',
    post_logout_redirect_uri: 'http://localhost:4200/signout-callback',
    automaticSilentRenew: true,
    silent_redirect_uri: 'http://localhost:4200/assets/silent-refresh.html',
    metadata: {
      issuer: 'https://accounts.google.com',
      authorization_endpoint: 'https://accounts.google.com/o/oauth2/auth',
      token_endpoint: 'https://oauth2.googleapis.com/token',
    },
  })
);

export type AuthContextState = User | null;

export type AuthContextType = {
  auth: AuthContextState;
  setAuth: React.Dispatch<React.SetStateAction<AuthContextState>>;
};

const AuthContext = createContext<AuthContextType>({
  auth: null,
  setAuth: () => {
    throw new Error('setAuth function must be overridden');
  },
});

export const useAuthContext = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuthContext must be used within a AuthContextProvider');
  }
  return context;
};

export type AuthContextProviderProps = {
  defaultState?: AuthContextState;
  children: React.ReactNode;
};

export const AuthContextProvider = ({ defaultState, children }: AuthContextProviderProps) => {
  const [auth, setAuth] = useState<AuthContextState>(defaultState || null);

  const userManager = useContext(UserManagerContext);

  const value = useMemo(() => ({ auth, setAuth }), [auth]);

  useEffect(() => {
    userManager.getUser().then(setAuth);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = () => {
  const userManager = useContext(UserManagerContext);
  const { auth, setAuth } = useAuthContext();

  const login = useCallback(() => {
    return userManager.signinRedirect();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const isLoggedIn = useMemo(() => {
    return !!auth && !auth.expired;
  }, [auth]);

  const completeLogin = useCallback(async () => {
    const user = await userManager.signinRedirectCallback();
    setAuth(user);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const logout = useCallback(() => {
    setAuth(null);
    return userManager.removeUser();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return { login, isLoggedIn, completeLogin, logout };
};
