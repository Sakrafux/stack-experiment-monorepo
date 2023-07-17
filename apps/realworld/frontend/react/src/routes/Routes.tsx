import { Login } from 'components';
import { Route, Routes as RouteSwitch } from 'react-router-dom';

const Routes = () => {
  return (
    <RouteSwitch>
      <Route path="/" element={<div>Home</div>} />
      <Route path="/login" element={<Login />} />
    </RouteSwitch>
  );
};

export default Routes;
