import { Navigate, Route, Routes as RouteSwitch } from 'react-router-dom';
import LoginRoute from './LoginRoute';
import ProfileArticlesRoute from './Profile/ProfileArticlesRoute';
import ProfileRoute from './Profile/ProfileRoute';

const Routes = () => {
  return (
    <RouteSwitch>
      <Route path="/" element={<div>Home</div>} />
      <Route path="/login" element={<LoginRoute />} />
      <Route path="/profile/:username" element={<ProfileRoute />}>
        <Route index element={<ProfileArticlesRoute key="articles" />} />
        <Route path="favorites" element={<ProfileArticlesRoute favorites key="favorites" />} />
      </Route>
      <Route path="*" element={<Navigate to="/" />} />
    </RouteSwitch>
  );
};

export default Routes;
