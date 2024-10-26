import {
  Box,
  Button,
  Card,
  CardBody,
  CardFooter,
  Center,
  Divider,
  FormControl,
  FormLabel,
  HStack,
  Image,
  NumberDecrementStepper,
  NumberIncrementStepper,
  NumberInput,
  NumberInputField,
  NumberInputStepper,
  VStack,
} from "@chakra-ui/react";

import scar from "../assets/scar.jpeg";
import apuSmart from "../assets/apu_smart.jpg";
import apuHacker from "../assets/apu_hacker.jpg";
import apuRich from "../assets/apu_rich.jpg";

import { useEffect, useState } from "react";
import { getAGYSOContract } from "../lib";

function BulletVoting() {
  const votingPower = 5;
  const [vote1, setValue1] = useState("");
  const [vote2, setValue2] = useState("");
  const [vote3, setValue3] = useState("");
  const [vote4, setValue4] = useState("");
  const [loading, setLoading] = useState(false);
  const [loadingText, setLoadingText] = useState("");

  useEffect(() => {});

  // TODO get ballots from chain
  const fetchGenerateProof = async () => {
    const encryptedVotes = await getAGYSOContract().getBallotBox();
    console.log(encryptedVotes);
  };

  const handleVoting = async () => {
    console.log(vote1, vote2, vote3, vote4);
    await fetchGenerateProof();
    setLoading(true);
    setLoadingText("Generating proof for vote");
  };

  return (
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
                  <VStack>
                    <Button size="lg">Programmer, Hacker A.</Button>
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={vote1}
                        onChange={(valueString) => setValue1(valueString)}
                        max={votingPower}
                        min={0}
                      >
                        <NumberInputField />
                        <NumberInputStepper>
                          <NumberIncrementStepper />
                          <NumberDecrementStepper />
                        </NumberInputStepper>
                      </NumberInput>
                    </FormControl>
                  </VStack>
                </CardFooter>
              </Center>
            </Card>

            <Card boxShadow="xl" bg="whiteAlpha.600">
              <CardBody>
                <Image src={apuSmart} w={300} h={200} borderRadius="lg" />
              </CardBody>
              <Divider />
              <Center>
                <CardFooter>
                  <VStack>
                    <Button size="lg">Applied Cryptographer A.</Button>
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={vote2}
                        onChange={(valueString) => setValue2(valueString)}
                        max={votingPower}
                        min={0}
                      >
                        <NumberInputField />
                        <NumberInputStepper>
                          <NumberIncrementStepper />
                          <NumberDecrementStepper />
                        </NumberInputStepper>
                      </NumberInput>
                    </FormControl>
                  </VStack>
                </CardFooter>
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
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={vote3}
                        onChange={(valueString) => setValue3(valueString)}
                        max={votingPower}
                        min={0}
                      >
                        <NumberInputField />
                        <NumberInputStepper>
                          <NumberIncrementStepper />
                          <NumberDecrementStepper />
                        </NumberInputStepper>
                      </NumberInput>
                    </FormControl>
                  </VStack>
                </CardFooter>
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
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={vote4}
                        onChange={(valueString) => setValue4(valueString)}
                        max={votingPower}
                        min={0}
                      >
                        <NumberInputField />
                        <NumberInputStepper>
                          <NumberIncrementStepper />
                          <NumberDecrementStepper />
                        </NumberInputStepper>
                      </NumberInput>
                    </FormControl>
                  </VStack>
                </CardFooter>
              </Center>
            </Card>
          </HStack>
          <Button
            mt="20"
            size="md"
            height="70px"
            width="300px"
            border="2px"
            onClick={handleVoting}
            isLoading={loading}
            loadingText={loadingText}
            colorScheme="blue"
          >
            Generate ZK Vote
          </Button>
          {/* <Button
            mt="5"
            size="md"
            height="70px"
            width="300px"
            border="2px"
            onClick={handleVoting}
            isLoading={loading}
            loadingText={loadingText}
            colorScheme="blue"
          >
            Submit Vote On-chain
          </Button> */}
        </VStack>
      </Center>
    </Box>
  );
}

export default BulletVoting;
