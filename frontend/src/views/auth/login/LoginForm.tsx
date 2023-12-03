import { useForm } from "@mantine/form";
import { TextInput, PasswordInput, Button } from "@mantine/core";
import useAuth from "../../../hooks/useAuth";

const LoginForm = () => {
  const { logIn, logInLoading } = useAuth();
  const form = useForm({
    initialValues: {
      username: "",
      password: "",
    },

    validate: {},
  });

  return (
    <form
      onSubmit={form.onSubmit((v) => {
        logIn(v.username, v.password)
      })}
    >
      <TextInput
        label="Username"
        placeholder="john_doe"
        required
        {...form.getInputProps("username")}
      />
      <PasswordInput
        label="Password"
        placeholder="Your password"
        required
        autoComplete="current-password"
        mt="md"
        {...form.getInputProps("password")}
      />
      <Button fullWidth mt="xl" type="submit" loading={logInLoading}>
        Sign in
      </Button>
    </form>
  );
};

export default LoginForm;
