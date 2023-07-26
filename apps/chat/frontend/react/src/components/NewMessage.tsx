import { Client } from '@stomp/stompjs';
import { useAuthContext } from 'context/AuthContext';
import { NewMessageDto } from 'model/NewMessageDto';
import { useState } from 'react';

type NewMessageProps = {
  chatId: number;
  socket: Client | null;
};

const NewMessage = ({ chatId, socket }: NewMessageProps) => {
  const [message, setMessage] = useState('');

  const { auth } = useAuthContext();

  const sendMessage = () => {
    if (socket && chatId && message) {
      socket.publish({
        destination: `/ws/chat/${chatId}`,
        body: JSON.stringify({ chatId, text: message } as NewMessageDto),
        headers: { Authorization: `Bearer ${auth?.id_token}` },
      });
      setMessage('');
    } else {
      console.log('No socket or chatId or message');
    }
  };

  return (
    <div className="w-full sticky bottom-0 z-10 mt-2 flex flex-row">
      <textarea
        className="flex-1 h-24 p-2 rounded-l-xl bg-white bg-opacity-90 text-gray-900 shadow-md resize-none"
        value={message}
        onChange={e => setMessage(e.target.value)}
        onKeyDown={e => {
          if (e.key === 'Enter' && !e.shiftKey) {
            e.preventDefault();
            sendMessage();
          }
        }}
        maxLength={255}
      />
      <button
        className="flex-shrink-0 w-32 rounded-r-xl bg-gradient-to-tr from-blue-600 to-blue-400 text-white shadow-md shadow-blue-500/20 hover:shadow-lg hover:shadow-blue-500/40 active:opacity-[0.75]"
        onClick={sendMessage}
      >
        Send
      </button>
    </div>
  );
};

export default NewMessage;
