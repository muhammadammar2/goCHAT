import { useEffect, useRef, useState } from "react";

const useWebSocket = (url: string) => {
  const [messages, setMessages] = useState<string[]>([]);
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    ws.current = new WebSocket(url);

    ws.current.onopen = () => {
      console.log("Connected to WebSocket");
    };

    ws.current.onmessage = (event) => {
      setMessages((prevMessages) => [...prevMessages, event.data]);
    };

    ws.current.onclose = () => {
      console.log("Disconnected from WebSocket");
    };

    return () => {
      ws.current?.close();
    };
  }, [url]);

  const sendMessage = (message: string) => {
    if (ws.current) {
      ws.current.send(message);
    }
  };

  return { messages, sendMessage };
};

export default useWebSocket;
