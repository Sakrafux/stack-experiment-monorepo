import { Link, BrowserRouter } from 'react-router-dom';
import Routes from './routes';

const App = () => {
  return (
    <BrowserRouter>
      <div>
        {/* START: routes */}
        {/* These routes and navigation have been generated for you */}
        {/* Feel free to move and update them to fit your needs */}
        <br />
        <hr />
        <br />
        <div role="navigation">
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/page-2">Page 2</Link>
            </li>
          </ul>
        </div>
        <Routes />
        {/* END: routes */}
      </div>
    </BrowserRouter>
  );
};

export default App;
