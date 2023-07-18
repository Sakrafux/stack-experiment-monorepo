import { ArticleDto } from 'models';
import { createContext, useContext, useMemo, useReducer } from 'react';

export type ArticlePages = {
  articles: Record<number, ArticleDto[]>;
  articlesCount: number;
};

export type ArticleContextState = {
  myFeed?: ArticlePages;
  globalFeed?: ArticlePages;
  taggedArticles: Record<string, ArticlePages>;
  userArticles: Record<string, ArticlePages>;
  favoritedArticles: Record<string, ArticlePages>;
};

const createDefaultState = (): ArticleContextState => ({
  myFeed: undefined,
  globalFeed: undefined,
  taggedArticles: {},
  userArticles: {},
  favoritedArticles: {},
});

export type ArticleContextAction =
  | {
      type: 'setMyFeed' | 'setGlobalFeed';
      setAction: (stateSlice: ArticlePages | undefined) => ArticlePages;
    }
  | {
      type: 'setTaggedArticles' | 'setUserArticles' | 'setFavoritedArticles';
      setAction: (stateSlice: Record<string, ArticlePages>) => Record<string, ArticlePages>;
    };

const articleReducer = (state: ArticleContextState, action: ArticleContextAction): ArticleContextState => {
  const { type, setAction } = action;
  switch (type) {
    case 'setMyFeed':
      return { ...state, myFeed: setAction(state.myFeed) };
    case 'setGlobalFeed':
      return { ...state, globalFeed: setAction(state.globalFeed) };
    case 'setTaggedArticles':
      return { ...state, taggedArticles: setAction(state.taggedArticles) };
    case 'setUserArticles':
      return { ...state, userArticles: setAction(state.userArticles) };
    case 'setFavoritedArticles':
      return { ...state, favoritedArticles: setAction(state.favoritedArticles) };
    default:
      throw new Error(`Unhandled action type: ${type}`);
  }
};

export type ArticleContextType = {
  state: ArticleContextState;
  dispatch: React.Dispatch<ArticleContextAction>;
};

const ArticleContext = createContext<ArticleContextType>({
  state: createDefaultState(),
  dispatch: () => {
    throw new Error('dispatch function must be overridden');
  },
});

export const useArticleContext = () => {
  const context = useContext(ArticleContext);
  if (!context) {
    throw new Error('useArticleContext must be used within a ArticleContextProvider');
  }
  return context;
};

export type ArticleContextProviderProps = {
  defaultState?: ArticleContextState;
  children: React.ReactNode;
};

export const ArticleContextProvider = ({ defaultState, children }: ArticleContextProviderProps) => {
  const [state, dispatch] = useReducer(articleReducer, defaultState || createDefaultState());

  const value = useMemo(() => ({ state, dispatch }), [state]);

  return <ArticleContext.Provider value={value}>{children}</ArticleContext.Provider>;
};
