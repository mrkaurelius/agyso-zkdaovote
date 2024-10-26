import {
  Box,
  Flex,
  Button,
  Spacer,
  GridItem,
  Text,
  VStack,
} from "@chakra-ui/react";
import { useState } from "react";
import { ethers } from "ethers";

const provider = new ethers.BrowserProvider(window.ethereum);

function Wallet() {
  const [votingPower, setVotingPower] = useState("");
  const [connectedAddress, setAddress] = useState("");

  const [myBool, setmyBool] = useState(true);

  function toggleBool() {
    setmyBool(!myBool);
  }

  const handleConnect = async () => {
    const signer = await provider.getSigner();
    const connectedAddress = signer.address;
    toggleBool();
    setAddress(connectedAddress);
    setVotingPower("5");
  };

  return myBool ? (
    <GridItem area={"header"}>
      <Flex>
        <Spacer />
        <Box
          p="4"
          paddingRight={4}
          alignSelf="center"
          display="flex"
          justifyContent="flex-end"
        >
          <Button
            colorScheme="purple"
            size="lg"
            variant="outline"
            onClick={handleConnect}
          >
            Connect Wallet
          </Button>
        </Box>
      </Flex>
    </GridItem>
  ) : (
    <Box>
      <Flex>
        <Spacer />
        <Box
          p="4"
          paddingRight={4}
          alignSelf="center"
          display="flex"
          justifyContent="flex-end"
        >
          <Flex
            pos="sticky"
            h="7vh"
            boxShadow="2px 4px 12px 15px rgba(0, 0, 0, 0.05)"
            borderRadius={"30px"}
            shadow="dark-lg"
            w="40vw"
            flexDir="column"
            justifyContent="center"
            alignItems="center"
          >
            <VStack fontSize="lg" color="black">
              <Text >Wallet Address: {connectedAddress}</Text>
              <Text>Voting Power: {votingPower}</Text>
            </VStack>
          </Flex>
        </Box>
      </Flex>
    </Box>
  );
}

export default Wallet;
