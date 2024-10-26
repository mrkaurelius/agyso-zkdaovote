import { Divider, GridItem, Flex, Heading } from "@chakra-ui/react";

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
        justifyContent="space-between"
      >
        <Flex
          p="5%"
          flexDir="column"
          w="100%"
          alignItems={"flex-start"}
          as="nav"
        ></Flex>

        <Flex p="5%" flexDir="column" w="100%" alignItems={"flex-start"} mb={4}>
          <Divider display={"flex"} />
          <Flex mt={4} align="center">
            <Flex flexDir="column" ml={4} display={"flex"}>
              <Heading as="h3" size="sm">
                AGYSO
              </Heading>
            </Flex>
          </Flex>
        </Flex>
      </Flex>
    </GridItem>
  );
}

export default Menu;
