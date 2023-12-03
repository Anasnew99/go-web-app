import { Outlet } from "react-router-dom";
import useAuth from "../../hooks/useAuth";
import { useEffect } from "react";
import router from "../../router";

// create empty component for now
const AuthLayout = () => {
   const { isAuthLoading, user } = useAuth();
   useEffect(() => {
     if(!isAuthLoading && user){
         router.navigate("/")
     }
   }, [isAuthLoading, user])
  
 return(
    isAuthLoading? null: <Outlet />
 );
};

export default AuthLayout;