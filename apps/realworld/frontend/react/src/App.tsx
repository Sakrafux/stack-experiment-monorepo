import { BrowserRouter } from 'react-router-dom';
import { Header } from './components';
import { Contexts } from './context';
import Routes from './routes';

const App = () => {
  return (
    <BrowserRouter>
      <Contexts>
        <Header />
        <Routes />
      </Contexts>
    </BrowserRouter>
  );
};

export default App;
