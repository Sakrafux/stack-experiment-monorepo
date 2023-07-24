import { useAuth, useAuthContext } from '../context/AuthContext';

const Home = () => {
  const { state } = useAuthContext();

  const { login, logout, isLoggedIn } = useAuth();

  console.log(state, isLoggedIn);

  return (
    <div>
      <h1 className="text-2xl font-bold">Home</h1>
      <button onClick={() => login()} className="block p-2 bg-slate-300 hover:bg-slate-400 active:bg-slate-500">
        Login
      </button>
      <button onClick={() => logout()} className="block p-2 bg-slate-300 hover:bg-slate-400 active:bg-slate-500">
        Logout
      </button>
    </div>
  );
};

export default Home;
