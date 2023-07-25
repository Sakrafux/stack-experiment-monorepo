import { BrowserRouter } from 'react-router-dom';
import { AuthContextProvider } from './context/AuthContext';
import Routes from 'routes/Routes';
import AxiosInterceptors from 'api/AxiosInterceptors';

export const App = () => {
  return (
    <AuthContextProvider>
      <BrowserRouter>
        <AxiosInterceptors>
          <Routes />
        </AxiosInterceptors>
      </BrowserRouter>
    </AuthContextProvider>
  );
};

export default App;
