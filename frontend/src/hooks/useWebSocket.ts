import { useEffect, useRef } from "react";

const useWebSocket = (url: string) => {
  const socketRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    socketRef.current = new WebSocket("ws://localhost:8080/ws");

    socketRef.current.onopen = () => {
      console.log("WebSocket connection established");
    };

    socketRef.current.onmessage = (event) => {
      const message = JSON.parse(event.data);
      console.log("Message received:", message);
      // Handle incoming messages (you might want to update state here)
    };

    socketRef.current.onerror = (error) => {
      console.error("WebSocket error:", error);
      console.log("WebSocket state:", socketRef.current?.readyState);
    };

    socketRef.current.onclose = () => {
      console.log("WebSocket connection closed");
    };

    return () => {
      socketRef.current?.close();
    };
  }, [url]);

  return socketRef.current;
};

export default useWebSocket;
