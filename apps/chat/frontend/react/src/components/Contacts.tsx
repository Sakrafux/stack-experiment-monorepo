import { api } from 'api/axios';
import { useContactContext } from 'context/ContactContext';
import { useEffect } from 'react';

const Contacts = () => {
  const { contacts, setContacts, activeContact, setActiveContact } = useContactContext();

  useEffect(() => {
    api.get('/user').then(res => setContacts(res.data));
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <div className="grid max-md:grid-cols-2 content-baseline gap-3 w-full rounded-xl border border-white/80 bg-white bg-opacity-80 py-2 px-4 text-white shadow-md backdrop-blur-2xl backdrop-saturate-200 lg:px-6 lg:py-4 overflow-auto">
      {contacts.map((contact, index) => {
        const color =
          activeContact?.id === contact.id
            ? 'from-red-600 to-red-400 shadow-red-500/20 hover:shadow-red-500/40'
            : 'from-orange-500/80 to-orange-300 shadow-orange-500/20 hover:shadow-orange-500/40';

        return (
          <button
            key={index}
            className={`flex items-center gap-2 p-2 relative w-full rounded-xl bg-gradient-to-tr bg-clip-border text-white shadow-md cursor-pointer transition-all hover:shadow-lg active:opacity-[0.75] ${color}`}
            onClick={() => setActiveContact(cur => (cur === contact ? null : contact))}
          >
            <img src={contact.picture} alt="Profile" className="w-8 h-8 rounded-full" />
            <span className="ml-2 text-sm font-semibold">{contact.name}</span>
          </button>
        );
      })}
    </div>
  );
};

export default Contacts;
