import { Divider, GridItem, Flex, Heading, Button, Box } from "@chakra-ui/react";
import { Link } from "react-router-dom";

function Menu() {
  return (
    <GridItem area={"nav"}>
      <Flex
        pos="sticky"
        left="5"
        h="95vh"
        marginTop="2.5vh"
        boxShadow="2px 4px 12px 15px rgba(0, 0, 0, 0.05)"
        borderRadius={"30px"}
        shadow="dark-lg"
        w={"200px"}
        flexDir="column"
      >
        <Link to="/">
          <Button m="3" mt="5" size="md" height="60px" border="2px" colorScheme="blue">
            Cast Vote
          </Button>
        </Link>

        <Link to="/results">
          <Button m="3" mt="1" size="md" height="60px" border="2px" colorScheme="red">
            Reveal Results
          </Button>
        </Link>
      </Flex>
    </GridItem>
  );
}

export default Menu;
