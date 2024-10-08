import React, { useState } from "react";
import useWebSocket from "../../hooks/useWebSocket";

const ChatInput: React.FC = () => {
  const [inputValue, setInputValue] = useState("");
  const socket = useWebSocket("ws://localhost:8080/ws");

  const handleSendMessage = () => {
    if (inputValue.trim() && socket) {
      const message = {
        sender: "YourUsername",
        content: inputValue,
        timestamp: new Date().toLocaleTimeString(),
      };

      socket.send(JSON.stringify(message));
      setInputValue("");
    }
  };

  return (
    <div className="chat-input-container">
      <input
        type="text"
        value={inputValue}
        onChange={(e) => setInputValue(e.target.value)}
        placeholder="Type a message..."
        className="chat-input"
      />
      <button onClick={handleSendMessage} className="send-button">
        Send
      </button>
    </div>
  );
};

export default ChatInput;
