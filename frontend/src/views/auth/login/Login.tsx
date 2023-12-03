
// import classes from "./Login.module.css";
import AuthCard from "../../../layout/authLayout/AuthCard";
import ClientLink from "../../../components/buttons/ClientLink";
import LoginForm from "./LoginForm";

export function Login() {
  return (
    <AuthCard
      title={"Welcome Back"}
      subtitle={
        <>
          Do not have an account yet?{" "}
          <ClientLink size="sm"  to="../register">
            Create account
          </ClientLink>
        </>
      }
    >
      <LoginForm />
    </AuthCard>
  );
}
