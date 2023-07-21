import { UserDto } from 'models';

export const getToken = () => {
  const user = localStorage.getItem('user');
  if (user) {
    const parsedUser = JSON.parse(user) as UserDto;
    return parsedUser.token;
  }
  return undefined;
};
