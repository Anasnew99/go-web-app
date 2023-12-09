import { useCallback, useEffect, useRef, useState } from "react";
import useAuth from "./useAuth";
import { useRoomMessages } from "./apis/room";
import { BaseMessage } from "../@types/api";


const useRoomSocket = (roomId: string) => {
    const [client, setClient] = useState<WebSocket|null>(null);
    const {token} = useAuth();
    const [messages, setMessages] = useState<BaseMessage[]>([]);
    const [connected, setConnected] = useState(false);
    const {data: initialMessages, isLoading: isMessagesLoading, refetch} = useRoomMessages(roomId);
    const messagesSetRef = useRef<Set<string>>(new Set());


    const addMessage = useCallback((addToFront: boolean, ...messages: BaseMessage[]) => {
        messages = messages.filter((message) => !messagesSetRef.current.has(message.id));
        if(messages.length === 0) {
            return;
        }
        messages.forEach((message) => {
            messagesSetRef.current.add(message.id);
        });
     
        setMessages((prevMessages) => addToFront?[...messages, ...prevMessages]:[...prevMessages, ...messages]);
    }, []);

    useEffect(() => {
        if (initialMessages) {

            addMessage(true, ...initialMessages);
        }
    }, [initialMessages, addMessage]);

    const sendMessage = useCallback((message: string) => {
        if(!connected){
            return;
        }
        client?.send(message);
        
    }, [client, connected]);

    useEffect(() => {
        if(!token) {
            return;
        }
        const client = new WebSocket(`${import.meta.env.VITE_WEBSOCKET_URL}/room/${roomId}/ws?token=${encodeURIComponent(token)}`);
        client.onopen = () => {
            console.log("Connected to websocket");
            setConnected(true);
            
        };

        client.onmessage = (message) => {
            console.log("Message received from server: ", message);
            try {
                const data = JSON.parse(message.data) as BaseMessage;
                addMessage(false, data);
            } catch (error) {
                console.log("Error parsing message: ", error);
            
            }
            
        }

        client.onclose = () => {
            console.log("Disconnected from websocket");
            setConnected(false);
        }


        client.onerror = (error) => {
            console.log("Error: ", error);
        }
        setClient(client);

        return () => {
            client.close();
        };
    }, [roomId, token, addMessage]);
    return {
        raw: client,
        messages,
        sendMessage,
        connected,
        isMessagesLoading,
        refetch
    };

    
};


export default useRoomSocket;