import Home from 'components/Home';
import Login from 'components/Login';
import SigninCallback from 'components/SigninCallback';
import { useAuth } from 'context/AuthContext';
import { Navigate, Route, Routes as RouteSwitch } from 'react-router-dom';

const Routes = () => {
  const { isLoggedIn } = useAuth();

  return (
    <RouteSwitch>
      <Route path="/" element={isLoggedIn ? <Home /> : <Login />} />
      <Route path="/signin-callback" element={<SigninCallback />} />
      <Route path="*" element={<Navigate to="/" />} />
    </RouteSwitch>
  );
};

export default Routes;
