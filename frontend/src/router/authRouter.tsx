import { RouteObject } from "react-router-dom";
import AuthLayout from "../layout/authLayout/AuthLayout";
import { Login } from "../views/auth/login/Login";
import { Register } from "../views/auth/register/Register";

const authRouter: RouteObject = {
    path: "auth",
    element: <AuthLayout />,
    children: [
        {
            path: "login",
            element: <Login />
        },
        {
            path: "register",
            element: <Register />
        }
    ]
};

export default authRouter;
