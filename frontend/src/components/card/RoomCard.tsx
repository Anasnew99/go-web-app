import { Card, Avatar, Text, Group, ActionIcon, Tooltip } from "@mantine/core";
import { IconCircleMinus, IconDoorExit } from "@tabler/icons-react";
import { BaseRoom } from "../../@types/api";
import useAuth from "../../hooks/useAuth";

const avatars = [
  "https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/avatars/avatar-2.png",
  "https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/avatars/avatar-4.png",
  "https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/avatars/avatar-7.png",
];

interface TaskCardProps {
  room: BaseRoom;
}
export function RoomCard(props: TaskCardProps) {
  const { user } = useAuth();
  return (
    <Card withBorder padding="lg" radius="md">
      <Text fz="lg" fw={500} mt="md">
        {props.room.id}
      </Text>
      <Text fz="sm" c="dimmed" mt={5}>
        {props.room.description}
      </Text>

      <Text c="dimmed" fz="sm" mt="md">
        Created at:{" "}
        <Text span fw={500} c="bright">
          {props.room.created_at}
        </Text>
      </Text>
      <Group justify="space-between" mt="md">
        <Avatar.Group spacing="sm">
          <Avatar src={avatars[0]} radius="xl" />
          <Avatar src={avatars[1]} radius="xl" />
          <Avatar src={avatars[2]} radius="xl" />
          <Avatar radius="xl">+5</Avatar>
        </Avatar.Group>
        {user?.username === props.room.room_owner.username ? (
          <Tooltip label="Delete Room">
            <ActionIcon variant="default" size="lg" radius="md">
              <IconCircleMinus />
            </ActionIcon>
          </Tooltip>
        ) : (
          <Tooltip label="Leave Room">
            <ActionIcon variant="default" size="lg" radius="md">
              <IconDoorExit />
            </ActionIcon>
          </Tooltip>
        )}
      </Group>
    </Card>
  );
}
