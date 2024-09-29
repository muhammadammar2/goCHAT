import React from "react";
import "../../Styles/ChatMessage.css";
interface ChatMessageProps {
  sender: string;
  message: string;
  timestamp: string;
}

const ChatMessage: React.FC<ChatMessageProps> = ({
  sender,
  message,
  timestamp,
}) => {
  return (
    <div className="chat-message">
      <div className="message-header">
        <span className="message-sender">{sender}</span>
        <span className="message-timestamp">{timestamp}</span>
      </div>
      <div className="message-content">{message}</div>
    </div>
  );
};

export default ChatMessage;
