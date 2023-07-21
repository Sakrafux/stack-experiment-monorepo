import { api } from 'api';

export const getTags = async (): Promise<string[]> => {
  const result = await api.get('/tags');

  return result.data.tags;
};
