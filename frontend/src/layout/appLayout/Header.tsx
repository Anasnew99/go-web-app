import { ActionIcon, AppShell, Group, Text } from "@mantine/core";
import { IconLogout } from "@tabler/icons-react";
import useAuth from "../../hooks/useAuth";
import classes from "./Header.module.css";

const Header = () => {
  const { logOut, user } = useAuth();
  return (
    <AppShell.Header>
      <Group
        className={classes.headerItemsContainer}
        align="center"
        justify="space-between"
      >
        <Group></Group>
        <Text className={classes.headerTitle} size="lg">
          Chat App
        </Text>
        <Group>
          <Text className={classes.accountName} size="sm">
            {user?.username}
          </Text>
          <ActionIcon variant="default" size={"xl"} onClick={logOut}>
            <IconLogout />
          </ActionIcon>
        </Group>
      </Group>
    </AppShell.Header>
  );
};

export default Header;
