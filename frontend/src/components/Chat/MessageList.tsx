import React from "react";
import ChatMessage from "./ChatMessage";

const MessageList: React.FC = () => {
  console.log("MessageList component rendered");
  const messages = [
    { sender: "Alice", message: "Hello!", timestamp: "10:00 AM" },
    { sender: "Bob", message: "Hi there!", timestamp: "10:01 AM" },
    // Add more messages for demonstration
  ];

  return (
    <div className="message-list">
      {messages.map((msg, index) => (
        <ChatMessage
          key={index}
          sender={msg.sender}
          message={msg.message}
          timestamp={msg.timestamp}
        />
      ))}
    </div>
  );
};

export default MessageList;
