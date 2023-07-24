import { BrowserRouter } from 'react-router-dom';
import { AuthContextProvider } from './context/AuthContext';
import Routes from 'routes/Routes';

export const App = () => {
  return (
    <AuthContextProvider>
      <BrowserRouter>
        <Routes />
      </BrowserRouter>
    </AuthContextProvider>
  );
};

export default App;
