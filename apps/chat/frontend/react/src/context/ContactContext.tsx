import { UserDto } from 'model/UserDto';
import { createContext, useContext, useMemo, useState } from 'react';

export type ContactContextType = {
  contacts: UserDto[];
  setContacts: React.Dispatch<React.SetStateAction<UserDto[]>>;
  activeContact: UserDto | null;
  setActiveContact: React.Dispatch<React.SetStateAction<UserDto | null>>;
};

const ContactContext = createContext<ContactContextType>({
  contacts: [],
  setContacts: () => {},
  activeContact: null,
  setActiveContact: () => {},
});

export const useContactContext = () => {
  const context = useContext(ContactContext);
  if (!context) {
    throw new Error('useContactContext must be used within a ContactContextProvider');
  }
  return context;
};

export type ContactContextProviderProps = {
  children: React.ReactNode;
};

export const ContactContextProvider = ({ children }: ContactContextProviderProps) => {
  const [contacts, setContacts] = useState<UserDto[]>([]);
  const [activeContact, setActiveContact] = useState<UserDto | null>(null);

  const value = useMemo(() => ({ contacts, setContacts, activeContact, setActiveContact }), [contacts, activeContact]);

  return <ContactContext.Provider value={value}>{children}</ContactContext.Provider>;
};
