import {
  Box,
  Center,
  Heading,
  Button,
  Text,
  VStack,
} from "@chakra-ui/react";
import { ethers } from "ethers";
import { useState } from "react";

const provider = new ethers.BrowserProvider(window.ethereum);

function ConnectWallet() {
  const [votingPower, setVotingPower] = useState("");
  const [connectedAddress, setAddress] = useState("");
  const handleConnect = async () => {
    const signer = await provider.getSigner();
    const connectedAddress = signer.address;
    setAddress(connectedAddress);
    setVotingPower("5");
  };

  return (
    <Box>
      <Box>
        <Center>
          <Heading size="xl" textTransform="uppercase" margin="6">
            AGYSO DAO Voting
          </Heading>
        </Center>
      </Box>
      <Box>
        <Center>
          <Button colorScheme="blue" onClick={handleConnect}>
            Connect Wallet
          </Button>
        </Center>
        <Center mt="4">
          <VStack>
            <Text>Connected Address: {connectedAddress}</Text>
            <Text>Voting Power: {votingPower}</Text>
          </VStack>
        </Center>
      </Box>
    </Box>
  );
}

export default ConnectWallet;
