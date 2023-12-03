import { createBrowserRouter } from "react-router-dom";
import appRouter from "./appRouter";
import authRouter from "./authRouter";
import RootLayout from "../layout/Root";


const router = createBrowserRouter([
    {
        path: "/",
        element: <RootLayout />,
        children: [
            appRouter,
            authRouter
        ]
    }
]);

// eslint-disable-next-line react-refresh/only-export-components
export default router;
