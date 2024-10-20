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
import { useState } from "react";
function BulletVoting() {
  const votingPower = 5;
  const [value1, setValue1] = useState("");
  const [value2, setValue2] = useState("");
  const [value3, setValue3] = useState("");
  const [value4, setValue4] = useState("");
  const [loading, setLoading] = useState(false);
  const [loadingText, setLoadingText] = useState("");

  const handleVoting = () => {
    setLoading(true);
    setLoadingText("Submitting");
  };

  return (
    <Box margin="12">
      <Center>
        <VStack>
          <HStack spacing="24px">
            <Card boxShadow="xl">
              <CardBody>
                <Image src={scar} borderRadius="lg" />
              </CardBody>
              <Divider />
              <Center>
                <CardFooter>
                  <VStack>
                    <Button size="lg">AHK</Button>
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={value1}
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

            <Card boxShadow="xl">
              <CardBody>
                <Image src={scar} borderRadius="lg" />
              </CardBody>
              <Divider />
              <Center>
                <CardFooter>
                  <VStack>
                    <Button size="lg">Ali</Button>
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={value2}
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

            <Card boxShadow="xl">
              <CardBody>
                <Image src={scar} borderRadius="lg" />
              </CardBody>
              <Divider />
              <Center>
                <CardFooter>
                  <VStack>
                    <Button size="lg">Burak</Button>
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={value3}
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

            <Card boxShadow="xl">
              <CardBody>
                <Image src={scar} borderRadius="lg" />
              </CardBody>
              <Divider />
              <Center>
                <CardFooter>
                  <VStack>
                    <Button size="lg">Nasuh</Button>
                    <FormControl>
                      <FormLabel>Voting Weight</FormLabel>
                      <NumberInput
                        value={value4}
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
            mt="6"
            size="md"
            height="48px"
            width="200px"
            border="2px"
            onClick={handleVoting}
            isLoading={loading}
            loadingText={loadingText}
            colorScheme="teal"
          >
            Vote
          </Button>
        </VStack>
      </Center>
    </Box>
  );
}

export default BulletVoting;
