import { Container } from "@mantine/core";
import { Outlet } from "react-router-dom";
import ThemeToggleButton from "../components/buttons/ThemeToggleButton";
import classes from "./Root.module.css";
const RootLayout = () => {


  return (
    <Container className={classes.container}>
      <Outlet />
      <Container className={classes.themeToggleButton}>
        <ThemeToggleButton />
      </Container>
    </Container>
  );
};

export default RootLayout;
