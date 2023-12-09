import { Container, Stack, Text } from "@mantine/core";
import { useParams } from "react-router-dom";
import useRoomSocket from "../../hooks/useRoomSocket";
import { MessageCard } from "../../components/card/MessageCard";
import { SendInput } from "../../components/inputs/SendInput";

const Room = () => {
  const { roomId } = useParams();
  const { messages, isMessagesLoading, sendMessage, connected, raw, refetch } =
    useRoomSocket(roomId as string);

  return (
    <Container>
      <Text size="xl">Room {roomId}</Text>
      <Stack gap={'sm'}>
        {messages.map((message) => (
          <MessageCard key={message.id} message={message} />
        ))}
      </Stack>
      <SendInput 
        onSend={sendMessage}
        disabled={!connected}
        initialValue={""}
        placeholder={connected ? "Send a message" : "Connecting..."}
      />
    </Container>
  );
};

export default Room;
