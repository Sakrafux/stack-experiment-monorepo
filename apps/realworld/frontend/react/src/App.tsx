import AxiosInterceptors from 'api/AxiosInterceptors';
import { DataContexts } from 'data';
import { BrowserRouter } from 'react-router-dom';
import { Footer, Header } from './components';
import { Contexts } from './context';
import Routes from './routes';

const App = () => {
  return (
    <BrowserRouter>
      <Contexts>
        <DataContexts>
          <AxiosInterceptors>
            <Header />
            <Routes />
            <Footer />
          </AxiosInterceptors>
        </DataContexts>
      </Contexts>
    </BrowserRouter>
  );
};

export default App;
