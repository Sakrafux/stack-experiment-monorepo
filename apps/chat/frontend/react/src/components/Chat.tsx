import { useAuthContext } from 'context/AuthContext';
import NewMessage from './NewMessage';
import { useEffect, useState } from 'react';
import { useContactContext } from 'context/ContactContext';
import { api } from 'api/axios';

type ChatMessage = {
  isMe: boolean;
  profile:
    | {
        name?: string;
        picture?: string;
      }
    | null
    | undefined;
  text: string;
  createdAt: Date;
};

const Chat = () => {
  const [chatId, setChatId] = useState<number>(0);
  const [messages, setMessages] = useState<ChatMessage[]>([]);

  const { auth } = useAuthContext();
  const { activeContact } = useContactContext();

  const myProfile = auth?.profile;
  const otherProfile = activeContact;

  useEffect(() => {
    api.post(`/chat/${otherProfile?.id}`).then(res => setChatId(res.data));
  }, [otherProfile?.id]);

  useEffect(() => {
    setMessages(
      new Array(20).fill(null).map((_, index) => {
        const isMe = index % 2 === 0;
        return {
          isMe,
          profile: isMe ? myProfile : otherProfile,
          text: `This is a message ${index}`,
          createdAt: new Date(new Date().getTime() + index * 1000),
        };
      })
    ); // TODO: replace with real messages
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // useEffect(() => {
  //   const chatEnd = document.getElementById('chat-end')!;

  //   const observer = new IntersectionObserver(
  //     entries => {
  //       if (entries[0].isIntersecting) {
  //         setMessages(cur => [
  //           ...cur,
  //           ...new Array(20).fill(null).map((_, index) => {
  //             const isMe = index % 2 === 0;
  //             return {
  //               isMe,
  //               profile: isMe ? myProfile : otherProfile,
  //               text: `This is a message loaded afterwareds ${index}`,
  //               createdAt: new Date(new Date().getTime() + index * 1000),
  //             };
  //           }),
  //         ]); // TODO: replace with real new loaded messages
  //       }
  //     },
  //     {
  //       root: document.getElementById('chat'),
  //       rootMargin: '0px',
  //       threshold: 0,
  //     }
  //   );
  //   observer.observe(chatEnd);

  //   return () => observer.disconnect();
  //   // eslint-disable-next-line react-hooks/exhaustive-deps
  // }, []);

  return (
    <div
      id="chat"
      className="flex flex-col-reverse gap-3 w-full rounded-xl border border-white/80 bg-white bg-opacity-80 py-2 px-4 text-white shadow-md backdrop-blur-2xl backdrop-saturate-200 lg:px-6 lg:py-4 overflow-auto"
    >
      <NewMessage chatId={chatId} />
      {messages.map((message, index) => {
        const color = message.isMe
          ? 'from-green-600 to-green-400 shadow-green-500/20 hover:shadow-green-500/40'
          : 'from-gray-600/80 to-gray-400 shadow-gray-500/20 hover:shadow-gray-500/40';

        const side = message.isMe ? 'self-end' : 'self-start';

        return (
          <div
            key={index}
            className={`w-max max-w-lg min-w-[150px] p-2 mb-2 relative rounded-xl bg-gradient-to-tr bg-clip-border text-white shadow-md transition-all ${color} ${side}`}
          >
            <div className="flex items-center gap-2">
              <img src={message.profile?.picture} alt="Profile" className="w-8 h-8 rounded-full" />
              <span className="ml-2 text-xs font-semibold">{message.profile?.name}</span>
            </div>
            <span className="text-sm">{message.text}</span>
            <div className="absolute m-2 top-100 right-0 text-xs text-gray-400">
              {message.createdAt.toLocaleDateString()} {message.createdAt.toLocaleTimeString()}
            </div>
          </div>
        );
      })}
      <div id="chat-end" />
    </div>
  );
};

export default Chat;
