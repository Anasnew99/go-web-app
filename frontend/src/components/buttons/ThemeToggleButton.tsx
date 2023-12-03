import { ActionIcon } from "@mantine/core";
import { IconSun, IconMoon } from "@tabler/icons-react";
import cx from "clsx";
import classes from "./ThemeToggleButton.module.css";
import useTheme from "../../hooks/useTheme";
interface ThemeToggleProps {}
const ThemeToggleButton: React.FC<ThemeToggleProps> = () => {
  const { toggleColorScheme } = useTheme();
  return (
    <ActionIcon
      onClick={() => toggleColorScheme()}
      variant="default"
      size="xl"
      aria-label="Toggle color scheme"
    >
      <IconSun className={cx(classes.icon, classes.light)} stroke={1.5} />
      <IconMoon className={cx(classes.icon, classes.dark)} stroke={1.5} />
    </ActionIcon>
  );
};

export default ThemeToggleButton;
