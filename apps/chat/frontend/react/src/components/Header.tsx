import { useAuth, useAuthContext } from 'context/AuthContext';

const Header = () => {
  const { state } = useAuthContext();

  const { logout } = useAuth();

  return (
    <nav className="mx-auto block w-full max-w-screen-xl rounded-xl border border-white/80 bg-white bg-opacity-80 py-2 px-4 text-white shadow-md backdrop-blur-2xl backdrop-saturate-200 lg:px-8 lg:py-4">
      <div className="mx-auto flex items-center justify-between text-gray-900">
        <span className="flex items-center gap-2">
          <img src={state?.profile.picture} alt="Profile" className="w-8 h-8 rounded-full" />
          <span className="ml-2 text-sm font-semibold">{state?.profile.name}</span>
        </span>
        <button
          className="middle center inline-block rounded-lg bg-gradient-to-tr from-blue-600 to-blue-400 py-2 px-4 font-sans text-xs font-bold uppercase text-white shadow-md shadow-blue-500/20 transition-all hover:shadow-lg hover:shadow-blue-500/40 active:opacity-[0.85] disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none"
          type="button"
          onClick={logout}
        >
          <span>Sign out</span>
        </button>
      </div>
    </nav>
  );
};

export default Header;
