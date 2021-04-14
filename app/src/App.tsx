import "./index.css";
import React, { useCallback, useState } from "react";
import { Launcher } from "react-chat-window";

interface Message {
	author: "me" | "them";
	type: "text";
	data: {
		text: string;
	};
}

function App() {
	const [sessionId, setSessionId] = useState<string>("");
	const [messageList, setMessageList] = useState<Message[]>([
		{
			author: "them",
			type: "text",
			data: {
				text: "Hi 👋 How can I help you today?",
			},
		},
	]);

	const sendMessage = useCallback(
		(message: Message) => {
			fetch("https://m6mecdmxol.execute-api.us-west-2.amazonaws.com/prod", {
				method: "POST",
				body: JSON.stringify({ text: message.data.text, sessionId: sessionId }),
			})
				.then((response) => {
					if (!response.ok) {
						throw new Error("Cannot add message");
					}

					return response.json();
				})
				.then((data: { message: string; sessionId: string }) => {
					setSessionId(data.sessionId);
					console.log("data", data);
					setMessageList((messages) => [
						...messages,
						{
							author: "them",
							type: "text",
							data: {
								text: data.message,
							},
						} as Message,
					]);
				});
		},
		[sessionId]
	);

	return (
		<Launcher
			agentProfile={{
				teamName: "Streem Boi",
				imageUrl: "/1.png",
			}}
			onMessageWasSent={(d: Message) => {
				setMessageList((messages) => [...messages, d]);
				sendMessage(d);
			}}
			messageList={messageList}
		/>
	);
}

export default App;
