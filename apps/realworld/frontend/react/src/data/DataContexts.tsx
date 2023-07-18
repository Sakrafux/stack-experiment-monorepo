import { ArticleContextProvider, ArticleContextState } from './state';

export type DataContextsProps = {
  defaultState?: { article?: ArticleContextState };
  children: React.ReactNode;
};

const DataContexts = ({ defaultState, children }: DataContextsProps) => {
  return <ArticleContextProvider defaultState={defaultState?.article}>{children}</ArticleContextProvider>;
};

export default DataContexts;
