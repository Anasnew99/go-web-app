import { Container, Text } from "@mantine/core";
import { useParams } from "react-router-dom";



const Room = () => {
    const {roomId} = useParams();

    return(
        <Container>
            {/* <Room /> */}
            <Text size="xl">
                Room {roomId}
            </Text>
        </Container>
    )
};

export default Room;