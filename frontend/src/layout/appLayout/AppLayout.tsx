import { Outlet } from "react-router-dom";
import useAuth from "../../hooks/useAuth";
import { useEffect } from "react";
import router from "../../router";
import { AppShell } from "@mantine/core";
import Header from "./Header";
// import Footer from "./Footer";


interface AppLayoutProps {
    children?: React.ReactNode;
}
const AppLayout:React.FC<AppLayoutProps> = () => {
    const { isAuthLoading, user } = useAuth();
    useEffect(() => {
      if(!isAuthLoading && !user){
          router.navigate("/auth/login")
      }
    }, [isAuthLoading, user])
    return(
        <AppShell header={{height: 64}}>
            <Header />
            <AppShell.Main>
                <Outlet />
            </AppShell.Main>
        </AppShell>
    )
};

export default AppLayout;