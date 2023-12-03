import { RouteObject } from "react-router-dom";
import AppLayout from "../layout/appLayout/AppLayout";
import Home from "../views/app/Home";
import Room from "../views/app/Room";
const appRouter:RouteObject = {
    element: <AppLayout />,
    children: [
        {
            index: true,
            element: <Home />
        },
        {
            path: "/:roomId",
            element: <Room />
        }
    ]
}

export default appRouter;
