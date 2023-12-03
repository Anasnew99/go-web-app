// import classes from "./Login.module.css";
import AuthCard from "../../../layout/authLayout/AuthCard";
import ClientLink from "../../../components/buttons/ClientLink";
import RegisterForm from "./RegisterForm";

export function Register() {
  
  return (
    <AuthCard
      title={"Welcome!"}
      subtitle={
        <>
          Do not already have an account?{" "}
          <ClientLink size="sm" to="../login">
            Login
          </ClientLink>
        </>
      }
    >
      <RegisterForm />
    </AuthCard>
  );
}
