import { LoginContextProvider, LoginContextState } from './LoginContext';

export type ContextsProps = {
  defaultState?: { login?: LoginContextState };
  children: React.ReactNode;
};

const Contexts = ({ defaultState, children }: ContextsProps) => {
  return <LoginContextProvider defaultState={defaultState?.login}>{children}</LoginContextProvider>;
};

export default Contexts;
