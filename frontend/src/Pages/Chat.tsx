import React from "react";
import MessageList from "../components/Chat/MessageList";
import ChatInput from "../components/Chat/ChatInput";
import "../Styles/Chat.css";

const Chat: React.FC = () => {
  console.log("Chat component rendered");
  return (
    <div className="chat-container">
      <h1 className="chat-title">Chat Room</h1>
      <MessageList />
      <ChatInput />
    </div>
  );
};

export default Chat;
