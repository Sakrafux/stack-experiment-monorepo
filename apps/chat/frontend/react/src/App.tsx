import { BrowserRouter, Route, Routes } from 'react-router-dom';
import SigninCallback from './components/SigninCallback';
import Home from './components/Home';
import { AuthContextProvider } from './context/AuthContext';

export const App = () => {
  return (
    <AuthContextProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/signin-callback" element={<SigninCallback />} />
        </Routes>
      </BrowserRouter>
    </AuthContextProvider>
  );
};

export default App;
