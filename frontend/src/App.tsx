import "./App.css";
import "@mantine/core/styles.css";
import {  RouterProvider } from "react-router-dom";
import router from "./router";
import { QueryClientProvider } from "react-query";
import queryClient from "./utils/queryClient";
import { MantineProvider } from "@mantine/core";
import theme from "./theme";
import { Notifications } from "@mantine/notifications";
import AuthProvider from "./providers/AuthProvider";
function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <MantineProvider theme={theme}>
        <Notifications />

        <AuthProvider>
        <RouterProvider router={router} />

        </AuthProvider>
      </MantineProvider>
    </QueryClientProvider>
  );
}

export default App;
