import Chat from './Chat';
import Contacts from './Contacts';
import Header from './Header';

const Home = () => {
  // TODO: hide chat if no contact is selected
  return (
    <div className="h-screen flex flex-col p-4 gap-4">
      <Header />
      <div className="flex-1 mx-auto grid max-md:grid-rows-[1fr_3fr] md:grid-cols-[1fr_3fr] w-full max-w-screen-xl gap-4 min-h-0">
        <Contacts />
        <Chat />
      </div>
    </div>
  );
};

export default Home;
