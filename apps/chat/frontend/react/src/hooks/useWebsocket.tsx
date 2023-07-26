import { environment } from 'environments/environment';
import { useEffect, useState } from 'react';
import { Client } from '@stomp/stompjs';

const useWebsocket = () => {
  const [socket, setSocket] = useState<Client | null>(null);
  const [isConnected, setIsConnected] = useState<boolean>(false);

  useEffect(() => {
    const client = new Client({
      // corresponds to STOMP endpoint defined in the backend WebSocketConfig.registerStompEndpoints
      brokerURL: environment.ws.url,
      reconnectDelay: 5000,
      heartbeatIncoming: 4000,
      heartbeatOutgoing: 4000,
    });

    client.onConnect = () => {
      console.log('Connected to websocket');
      setIsConnected(true);
    };

    client.onDisconnect = () => {
      console.log('Disconnected from websocket');
      setIsConnected(false);
    };

    client.activate();

    setSocket(client);

    return () => {
      client.deactivate();
    };
  }, []);

  useEffect(() => {
    if (socket && isConnected) {
      const sub = socket.subscribe('/topic/login', message => {
        console.log(`${message.body} has logged in`);
      });

      return () => {
        sub.unsubscribe();
      };
    }
  }, [isConnected, socket]);

  if (socket && isConnected) {
    return socket;
  }

  return null;
};

export default useWebsocket;
