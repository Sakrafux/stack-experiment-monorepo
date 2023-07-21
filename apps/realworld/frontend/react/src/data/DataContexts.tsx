import { ArticleContextProvider, ArticleContextState, TagContextProvider } from './state';

export type DataContextsProps = {
  defaultState?: { article?: ArticleContextState; tag?: string[] };
  children: React.ReactNode;
};

const DataContexts = ({ defaultState, children }: DataContextsProps) => {
  return (
    <ArticleContextProvider defaultState={defaultState?.article}>
      <TagContextProvider defaultState={defaultState?.tag}>{children}</TagContextProvider>
    </ArticleContextProvider>
  );
};

export default DataContexts;
