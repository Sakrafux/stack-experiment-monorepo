import { BrowserRouter } from 'react-router-dom';
import { AuthContextProvider } from './context/AuthContext';
import Routes from 'routes/Routes';
import AxiosInterceptors from 'api/AxiosInterceptors';
import { ContactContextProvider } from 'context/ContactContext';

export const App = () => {
  return (
    <AuthContextProvider>
      <ContactContextProvider>
        <BrowserRouter>
          <AxiosInterceptors>
            <Routes />
          </AxiosInterceptors>
        </BrowserRouter>
      </ContactContextProvider>
    </AuthContextProvider>
  );
};

export default App;
