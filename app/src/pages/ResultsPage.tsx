import {
  Box,
  Button,
  Card,
  CardBody,
  CardFooter,
  Center,
  Divider,
  Grid,
  GridItem,
  HStack,
  Image,
  Text,
  useToast,
  VStack,
} from "@chakra-ui/react";

import { useState } from "react";

import Wallet from "../components/Wallet";
import Menu from "../components/Menu";

import scar from "../assets/scar.jpeg";
import apuSmart from "../assets/apu_smart.jpg";
import apuHacker from "../assets/apu_hacker.jpg";
import apuRich from "../assets/apu_rich.jpg";

import { getAGYSOContract } from "../lib";

function ResultsPage() {
  const [vote0, setVote0] = useState("0");
  const [vote1, setVote1] = useState("0");
  const [vote2, setVote2] = useState("0");
  const [vote3, setVote3] = useState("0");

  const toast = useToast();

  const handleRevealResults = async () => {
    let encryptedBallots = (await getAGYSOContract().getBallotBox()) as string;
    if (encryptedBallots == "0x") encryptedBallots = "0";

    let encryptedVotesNoPrefix = encryptedBallots.slice(2, encryptedBallots.length);

    console.log(encryptedBallots)

    const payload = {
      encryptedVotes: encryptedVotesNoPrefix,
    };

    const options = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    };

    let res = await fetch("http://localhost:2929/decrypt", options);
    let resJson = await res.json();
    let decryptedVoteArr = resJson.decryptedVotes as number[];

    if (res.status < 300) {
      toast({
        position: "top",
        title: "Total votes revealed!",
        status: "success",
        duration: 10000,
        isClosable: true,
      });
    } else if (res.status >= 300) {
      toast({
        position: "top",
        title: "Vote reveal error",
        status: "error",
        isClosable: true,
      });
    }

    console.log(decryptedVoteArr)

    setVote0(decryptedVoteArr[0].toString());
    setVote1(decryptedVoteArr[1].toString());
    setVote2(decryptedVoteArr[2].toString());
    setVote3(decryptedVoteArr[3].toString());
  };

  return (
    <Grid
      templateAreas={`
      "nav header"
      "nav main"
      `}
      gridTemplateRows={"auto auto"}
      gridTemplateColumns={"250px auto"}
      bgGradient="linear-gradient(120deg, #a6c0fe 0%, #f68084 100%)"
      w="100vw"
      h="100vh"
      color="blackAlpha.700"
      fontWeight="bold"
    >
      <Wallet></Wallet>
      <Menu></Menu>
      <GridItem area={"main"}>
        <Box margin="12">
          <Center>
            <VStack>
              <HStack spacing="48px">
                <Card boxShadow="xl" bg="whiteAlpha.600">
                  <CardBody>
                    <Image src={apuHacker} w={300} h={200} borderRadius="lg" />
                  </CardBody>
                  <Divider />
                  <Center>
                    <CardFooter>
                      <Button size="lg">Programmer, Hacker A.</Button>
                    </CardFooter>
                  </Center>
                  <Center>
                    <Text fontSize="2xl" mb="4">
                      Vote: {vote0}
                    </Text>
                  </Center>
                </Card>

                <Card boxShadow="xl" bg="whiteAlpha.600">
                  <CardBody>
                    <Image src={apuSmart} w={300} h={200} borderRadius="lg" />
                  </CardBody>
                  <Divider />
                  <Center>
                    <CardFooter>
                      <Button size="lg">Applied Cryptographer A.</Button>
                    </CardFooter>
                  </Center>
                  <Center>
                    <Text fontSize="2xl" mb="4">
                      Vote: {vote1}
                    </Text>
                  </Center>
                </Card>

                <Card boxShadow="xl" bg="whiteAlpha.600">
                  <CardBody>
                    <Image src={scar} borderRadius="lg" />
                  </CardBody>
                  <Divider />
                  <Center>
                    <CardFooter>
                      <VStack>
                        <Button size="lg">Programmer, UI/UX 6-6 B.</Button>
                      </VStack>
                    </CardFooter>
                  </Center>
                  <Center>
                    <Text fontSize="2xl" mb="4">
                      Vote: {vote2}
                    </Text>
                  </Center>
                </Card>

                <Card boxShadow="xl" bg="whiteAlpha.600">
                  <CardBody>
                    <Image src={apuRich} w={300} h={200} borderRadius="lg" />
                  </CardBody>
                  <Divider />
                  <Center>
                    <CardFooter>
                      <VStack>
                        <Button size="lg">Dapp, DeFi, Web3 Expert N.</Button>
                      </VStack>
                    </CardFooter>
                  </Center>
                  <Center>
                    <Text fontSize="2xl" mb="5">
                      Vote: {vote3}
                    </Text>
                  </Center>
                </Card>
              </HStack>

              <Button
                mt="20"
                size="md"
                height="70px"
                width="300px"
                border="2px"
                onClick={handleRevealResults}
                colorScheme="red"
              >
                Reveal Results
              </Button>
            </VStack>
          </Center>
        </Box>
      </GridItem>
    </Grid>
  );
}

export default ResultsPage;
