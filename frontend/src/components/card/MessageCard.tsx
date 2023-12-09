import {  Text, Group, Stack } from "@mantine/core";
import { BaseMessage } from "../../@types/api";
import { fromDate } from "../../utils/dates";
import UserAvatar from "../avatar/UserAvatar";

// import router from "../../router";

interface MessageCardProps {
  message: BaseMessage;
}

export function MessageCard(props: MessageCardProps) {
  return (
    <Stack>
      <Group>
        <UserAvatar username={props.message?.username?? ""} radius="xl" />
        <div>
          <Text size="sm">{props.message.username}</Text>
          <Text size="xs" c="dimmed">
            {fromDate(props.message.timestamp * 1000)}
          </Text>
        </div>
      </Group>
      <Text pl={54} size="sm">
        {props.message.message}
      </Text>
    </Stack>
  );
}
