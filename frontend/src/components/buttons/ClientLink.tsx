import { Anchor } from "@mantine/core";
import { Link } from "react-router-dom";


type AnchorProps = React.ComponentProps<typeof Anchor<typeof Link>>;

const ClientLink = (props: AnchorProps) => {
    return (
        <Anchor component={Link} {...props} />
    );
}

export default ClientLink;