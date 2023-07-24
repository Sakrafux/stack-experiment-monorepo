import { useState } from 'react';

const NewMessage = () => {
  const [message, setMessage] = useState('');

  const sendMessage = () => {
    console.log(message); // TODO: send message
    setMessage('');
  };

  return (
    <div className="w-full sticky bottom-0 z-10 mt-2 flex flex-row">
      <textarea
        className="flex-1 h-20 p-2 rounded-l-xl bg-white bg-opacity-90 text-gray-900 shadow-md resize-none"
        value={message}
        onChange={e => setMessage(e.target.value)}
        onKeyDown={e => {
          if (e.key === 'Enter' && !e.shiftKey) {
            e.preventDefault();
            sendMessage();
          }
        }}
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
