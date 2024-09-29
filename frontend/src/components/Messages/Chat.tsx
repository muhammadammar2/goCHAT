import React, { useEffect, useState, useRef } from "react";
import "./styles.css"; // Importing the CSS file

const Chat = () => {
  const [messages, setMessages] = useState<string[]>([]);
  const [input, setInput] = useState("");
  const socket = useRef<WebSocket | null>(null);

  useEffect(() => {
    socket.current = new WebSocket("ws://localhost:6969/ws");

    socket.current.onmessage = (event) => {
      const newMessage = event.data;
      setMessages((prevMessages) => [...prevMessages, newMessage]);
    };

    return () => {
      if (socket.current) {
        socket.current.close();
      }
    };
  }, []);

  const sendMessage = () => {
    if (input && socket.current) {
      socket.current.send(input);
      setInput("");
    }
  };

  return (
    <div className="chat-container">
      <div className="messages">
        {messages.map((msg, index) => (
          <div key={index} className="message">
            {msg}
          </div>
        ))}
      </div>
      <input
        type="text"
        className="chat-input"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        onKeyPress={(e) => {
          if (e.key === "Enter") sendMessage();
        }}
      />
    </div>
  );
};

export default Chat;
