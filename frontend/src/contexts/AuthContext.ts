import { createContext } from "react";
import { UserResponse } from "../api/user";
import { RegisterRequest } from "../api/auth";



export interface AuthContextData {
  user: null|UserResponse;
  token: string | null;
  isAuthLoading: boolean;
  logIn: (username: string, password: string) => Promise<void>;
  logInLoading: boolean;
  register: (data: RegisterRequest) => Promise<void>;
  registerLoading: boolean;
  logOut: () => Promise<void>;
}
const AuthContext = createContext({} as AuthContextData);

export default AuthContext;
