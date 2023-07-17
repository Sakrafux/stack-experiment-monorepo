import { Link, Route, Routes as RouteSwitch } from 'react-router-dom';

const Routes = () => {
  return (
    <RouteSwitch>
      <Route
        path="/"
        element={
          <div>
            This is the generated root route. <Link to="/page-2">Click here for page 2.</Link>
          </div>
        }
      />
      <Route
        path="/page-2"
        element={
          <div>
            <Link to="/">Click here to go back to root page.</Link>
          </div>
        }
      />
    </RouteSwitch>
  );
};

export default Routes;
