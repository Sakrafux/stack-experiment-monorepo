import { createContext, useContext, useMemo, useState } from 'react';

export type TagContextState = string[] | undefined;

export type TagContextType = {
  state: TagContextState;
  dispatch: React.Dispatch<React.SetStateAction<TagContextState>>;
};

const TagContext = createContext<TagContextType>({
  state: undefined,
  dispatch: () => {
    throw new Error('dispatch function must be overridden');
  },
});

export const useTagContext = () => {
  const context = useContext(TagContext);
  if (!context) {
    throw new Error('useTagContext must be used within a TagContextProvider');
  }
  return context;
};

export type TagContextProviderProps = {
  defaultState?: TagContextState;
  children: React.ReactNode;
};

export const TagContextProvider = ({ defaultState, children }: TagContextProviderProps) => {
  const [state, dispatch] = useState<TagContextState>(defaultState);

  const value = useMemo(() => ({ state, dispatch }), [state]);

  return <TagContext.Provider value={value}>{children}</TagContext.Provider>;
};
