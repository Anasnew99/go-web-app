
import React from "react";
import classes from "./AuthCard.module.css";
import {
    Paper,
    Title,
    Text,
    Container,
  } from '@mantine/core';

interface IAuthCardProps {
    children: React.ReactNode;
    title: React.ReactNode;
    subtitle: React.ReactNode;
}

const AuthCard: React.FC<IAuthCardProps> = ({ children, title, subtitle }) => {
    return (
        <Container size={420} my={40}>
          <Title ta="center" className={classes.title}>
            {title}
          </Title>
          <Text c="dimmed" size="sm" ta="center" mt={5}>
            {subtitle}
          </Text>
    
          <Paper withBorder shadow="md" p={30} mt={30} radius="md">
            {children}
          </Paper>
        </Container>
      );
};

export default AuthCard;