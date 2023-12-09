import { Avatar, AvatarProps } from "@mantine/core";
import React from "react";
const colorMaps: Record<string, string> = {
    A: "blue",
    B: "red",
    C: "green",
    D: "yellow",
    E: "orange",
    F: "pink",
    G: "purple",
    H: "cyan",
    I: "teal",
    J: "lime",
    K: "gray",
    L: "indigo",
    M: "brown",
    N: "pink",
    O: "purple",
    P: "cyan",
    Q: "teal",
    R: "lime",
    S: "gray",
    T: "indigo",
    U: "brown",
    V: "pink",
    W: "purple",
    X: "cyan",
    Y: "teal",
    Z: "lime",
  };

interface UserAvatarProps extends AvatarProps {
    username: string;
}

export default function UserAvatar({username, ...restProps}: UserAvatarProps){
    console.log("Username", username);
    const initials = username.slice(0).toUpperCase();
    const twoLetters = username.slice(0, 2).toUpperCase();
    return (
        <Avatar
            alt={username}
            radius="xl"
            color={colorMaps[initials] as string}
            {...restProps}
        >
            {twoLetters}
        </Avatar>
    )
}