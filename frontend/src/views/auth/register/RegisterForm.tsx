import { useForm } from "@mantine/form";
import { TextInput, PasswordInput, Button } from "@mantine/core";
import useAuth from "../../../hooks/useAuth";

const RegisterForm = () => {
  const { register, registerLoading } = useAuth();
  const form = useForm({
    initialValues: {
      username: "",
      password: "",
      resetPassword: "",
      email: "",
    },

    validate: {
      username: (value) =>
        /^[a-zA-Z0-9]{3,}$/.test(value) ? null : "Invalid username",
      password: (value) =>
        value.length < 6
          ? "Password should be at least 6 characters long"
          : null,
      resetPassword: (value, values) =>
        value !== values.password ? "Passwords don't match" : null,
      email: (value) =>
        /|^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[A-Za-z]+$/.test(value)
          ? null
          : "Invalid email",
    },
    validateInputOnChange: true,
  });

  return (
    <form
      onSubmit={form.onSubmit((v) => {
        register({
          username: v.username,
          password: v.password,
          email: v.email,
        })
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
        mt="md"
        {...form.getInputProps("password")}
      />

      <PasswordInput
        label="Repeat Password"
        placeholder="Confirm your password"
        required
        mt="md"
        {...form.getInputProps("resetPassword")}
      />

      <TextInput
        label="Email (optional)"
        placeholder="a@b.com"
        mt="md"
        {...form.getInputProps("email")}
      />

      <Button fullWidth mt="xl" type="submit" loading={registerLoading}>
        Register
      </Button>
    </form>
  );
};

export default RegisterForm;
