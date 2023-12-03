import { Container, Stack, Text } from "@mantine/core";
import { useUser } from "../../hooks/apis/user";
import { useMemo } from "react";
import classes from "./Home.module.css";
import { RoomCard } from "../../components/card/RoomCard";
const Home = () => {
  const { data: user } = useUser();

  const ownedRooms = useMemo(() => {
    return user?.rooms.filter(
      (room) => room.room_owner.username === user.username
    );
  }, [user]);

  const joinedRooms = useMemo(() => {
    return user?.rooms.filter(
      (room) => room.room_owner.username !== user.username
    );
  }, [user]);
  return (
    <Container p={"sm"}>
      <Stack>
        <Stack>
          <Text size={"md"}> Welcome, {user?.username}</Text>
        </Stack>
        <Stack className={classes.sections}>
          <Text size="xl">Owned Rooms</Text>
          <Stack>
            {ownedRooms?.map((room) => {
              return <RoomCard key={room.id} room={room} />;
            })}
          </Stack>
        </Stack>
        <Stack className={classes.sections}>
          <Text size="xl">Joined Rooms</Text>
          <Stack>
            {joinedRooms?.map((room) => {
              return <RoomCard key={room.id} room={room} />;
            })}
          </Stack>
        </Stack>
      </Stack>
    </Container>
  );
};

export default Home;
