import { Navigate, Route, Routes as RouteSwitch } from 'react-router-dom';
import LoginRoute from './LoginRoute';
import ProfileArticlesRoute from './Profile/ProfileArticlesRoute';
import ProfileRoute from './Profile/ProfileRoute';
import SettingsRoute from './SettingsRoute';
import RegisterRoute from './RegisterRoute';
import HomeRoute from './HomeRoute';
import ArticleRoute from './ArticleRoute';

const Routes = () => {
  return (
    <RouteSwitch>
      <Route path="/" element={<HomeRoute />} />
      <Route path="/login" element={<LoginRoute />} />
      <Route path="/register" element={<RegisterRoute />} />
      <Route path="/settings" element={<SettingsRoute />} />
      <Route path="/profile/:username" element={<ProfileRoute />}>
        <Route index element={<ProfileArticlesRoute key="articles" />} />
        <Route path="favorites" element={<ProfileArticlesRoute favorites key="favorites" />} />
      </Route>
      <Route path="/article/:slug" element={<ArticleRoute />} />
      <Route path="*" element={<Navigate to="/" />} />
    </RouteSwitch>
  );
};

export default Routes;
