import { ReactNode, useCallback, useEffect, useState } from "react";
import AuthContext, { AuthContextData } from "../contexts/AuthContext";
import {  useQueryClient } from "react-query";
import { UserResponse } from "../api/user";
import { login, register as registerApi } from "../api/auth";
import router from "../router";
import { useUser } from "../hooks/apis/user";

interface AuthProviderProps {
  children: ReactNode;
}

const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [isAuthLoading, setIsAuthLoading] = useState(true);
  const [token, setToken] = useState<string | null>(null);
  const [logInLoading, setLogInLoading] = useState(false);
  const [registerLoading, setRegisterLoading] = useState(false);
  const queryClient = useQueryClient();
  const {data: user} = useUser({
      enabled: !!token
  });

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setToken(token);
    } else {
      setIsAuthLoading(false);
    }
  }, []);

  useEffect(() => {
    if (user) {
      setIsAuthLoading(false);
    }
  }, [user]);

  const logInFunc: AuthContextData["logIn"] = useCallback(async (username, password) => {
    try {
      setLogInLoading(true);
      const loginRes = await login(username, password);
      const token = loginRes.data.data.token;
      setToken(token);
      localStorage.setItem("token", token);
      setIsAuthLoading(true);
    //   nav("/");
    router.navigate("/");
    } catch (e) {
      setLogInLoading(false);
      throw e;
    } finally {
      setLogInLoading(false);
    }
  }, []);

  const registerFunc: AuthContextData["register"] = useCallback(async (user) => {
    try {
      setRegisterLoading(true);
      await registerApi(user);
    router.navigate("/auth/login");
      return;
    } catch (error) {
      setRegisterLoading(false);
      throw error;
    } finally {
      setRegisterLoading(false);
    }
  }, []);

  const logOutFunc: AuthContextData["logOut"] = useCallback(async () => {
    localStorage.removeItem("token");
    setToken(null);
    queryClient.removeQueries();
    router.navigate("/auth/login");
  }, [queryClient]);

  return (
    <AuthContext.Provider
      value={{
        isAuthLoading,
        user: user as UserResponse,
        token,
        logIn: logInFunc,
        logInLoading,
        register: registerFunc,
        registerLoading,
        logOut: logOutFunc,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
