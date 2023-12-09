import {
  TextInput,
  TextInputProps,
  ActionIcon,
  useMantineTheme,
  rem,
} from "@mantine/core";
import {  IconArrowRight } from "@tabler/icons-react";
import useAuth from "../../hooks/useAuth";
import UserAvatar from "../avatar/UserAvatar";
import { useEffect, useState } from "react";

interface SendInputProps extends TextInputProps {
  onSend: (newValue: string) => void;
  initialValue?: string;
}

export function SendInput(props: SendInputProps) {
  const theme = useMantineTheme();
  const { user } = useAuth();
  const [value, setValue] = useState("");
  useEffect(() => {
    setValue(props.initialValue ?? "");
  }, [props.initialValue]);
  return (
    <TextInput
      radius="xl"
      size="md"
      placeholder="Send a message"
      rightSectionWidth={42}
      value={value}
      onChange={(event) => setValue(event.target.value)}
      leftSection={<UserAvatar username={""} />}
      onKeyDown={(event) => {
        if (event.key === "Enter") {
          event.preventDefault();
          event.stopPropagation();
          props.onSend(value);
        }
      }}
      rightSection={
        <ActionIcon
          size={32}
          radius="xl"
          color={theme.primaryColor}
          variant="filled"
          onClick={() => props.onSend(value)}
        >
          <IconArrowRight
            style={{ width: rem(18), height: rem(18) }}
            stroke={1.5}
          />
        </ActionIcon>
      }
      {...props}
    />
  );
}
