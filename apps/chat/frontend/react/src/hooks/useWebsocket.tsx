import { environment } from 'environments/environment';
import { useEffect, useState } from 'react';
import { Client } from '@stomp/stompjs';

const useWebsocket = (chatId: number) => {
  const [socket, setSocket] = useState<Client | null>(null);
  const [isConnected, setIsConnected] = useState<boolean>(false);

  useEffect(() => {
    const client = new Client({
      brokerURL: environment.ws.url,
      reconnectDelay: 5000,
      heartbeatIncoming: 4000,
      heartbeatOutgoing: 4000,
    });

    client.onConnect = () => {
      console.log('Connected to websocket');
      setIsConnected(true);
    };

    client.onStompError = frame => {
      console.log('Error', frame);
    };

    client.onDisconnect = () => {
      console.log('Disconnected from websocket');
    };

    client.activate();

    setSocket(client);

    return () => {
      client.deactivate();
    };
  }, []);

  if (socket && isConnected) {
    return socket;
  }

  return null;
};

export default useWebsocket;
